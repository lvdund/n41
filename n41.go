package n41

import (
	"fmt"
	"n41/n41msg"
	"n41/n41types"
	"sync"
	"sync/atomic"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	N41_MSG_RECV_QUEUE_SIZE int = 1024 //size of the queue to receive N41 message from forwwarder
	N41_MSG_SEND_QUEUE_SIZE int = 1024 //size of the queue to receive N41 message for sending to forwwarder
)

type Endpoint interface {
	// UdpAddr() *net.UDPAddr
	Addr() *n41types.Sbi
}

type N41Session interface {
	Endpoint
	RemoteSeid() uint64
	LocalSeid() uint64
	FillDeletionRequest(*n41msg.N41SessionDeletionRequest)
	FillEstablishmentRequest(*n41msg.N41SessionEstablishmentRequest)
	FillModificationRequest(*n41msg.N41SessionModificationRequest)
}

type NodeContext interface {
	NodeId() *n41types.NodeID
	Port() int
	SessionHandler() SessionProducer
	AssociationHandler() AssociationProducer
}

type N41 struct {
	ctx      NodeContext
	fwd      *Forwarder
	seq      uint32 //for generating request sequence number
	wg       sync.WaitGroup
	shandler SessionProducer     //upper handler to response messages
	ahandler AssociationProducer //upper handler to response messages
	done     chan bool           //trigger termination in child loops
	sending  chan *ReqSendingInfo
	queue    ExpiringList
}

func NewN41(ctx NodeContext) *N41 {
	id := ctx.NodeId()
	// addr := net.UDPAddr{
	// 	IP:   id.ResolveNodeIdToIp(),
	// 	Port: ctx.Port(),
	// }
	addr := n41types.Sbi{
		IP:   id.ResolveNodeIdToIp(),
		Port: ctx.Port(),
	}
	ret := &N41{
		ctx:      ctx,
		fwd:      newForwarder(addr),
		sending:  make(chan *ReqSendingInfo, N41_MSG_SEND_QUEUE_SIZE),
		done:     make(chan bool),
		queue:    newExpiringList(),
		shandler: ctx.SessionHandler(),
		ahandler: ctx.AssociationHandler(),
	}
	return ret
}

func (proto *N41) Start() (err error) {
	recv := make(chan RecvInfo, N41_MSG_RECV_QUEUE_SIZE)
	if err = proto.fwd.start(recv); err == nil {
		go proto.receivingloop(recv)
		go proto.sendingloop()
	}
	return
}

func (proto *N41) Stop() {
	proto.fwd.stop()
	close(proto.done)
	proto.wg.Wait()
}

// waits for request messages to send
func (proto *N41) sendingloop() {
	proto.wg.Add(1)
	defer proto.wg.Done()
	ticker := time.NewTicker(EXPIRING_CHECK_INTERVAL)
	for {
		select {
		case <-proto.done:
			return
		case <-ticker.C:
			proto.queue.flush()

		case info := <-proto.sending:
			//sending the message
			if info.err = proto.fwd.WriteTo(info.msg, info.remote); info.err != nil {
				//terminate sending
				close(info.done)
			} else {
				//cache the request for resending (in case of timeout)  and
				//searching (in case of receiving a response)
				proto.queue.add(info)
			}
		}
	}
}

// receiving messages from forwarder
func (proto *N41) receivingloop(recv chan RecvInfo) {
	proto.wg.Add(1)
	defer proto.wg.Done()
	for info := range recv {
		proto.handle(info.remote, info.msg)
	}

}

// func (proto *N41) handle(remote *net.UDPAddr, msg *n41msg.Message) {
func (proto *N41) handle(remote *n41types.Sbi, msg *n41msg.Message) {
	if msg.IsRequest() {
		proto.handleReq(remote, msg)
	} else {
		proto.handleRsp(remote, msg)
	}
}

// func (proto *N41) handleRsp(remote *net.UDPAddr, msg *n41msg.Message) {
func (proto *N41) handleRsp(remote *n41types.Sbi, msg *n41msg.Message) {
	logrus.Debugf("receive a response of type %d from %s", msg.Header.MessageType, remote)
	if infoinf := proto.queue.pop(remote.String(), msg.Header.SequenceNumber); infoinf != nil {
		if info, ok := infoinf.(*ReqSendingInfo); ok {
			match := false
			switch msg.Header.MessageType {
			case n41msg.N41_HEARTBEAT_RESPONSE:
				match = info.msg.Header.MessageType == n41msg.N41_HEARTBEAT_REQUEST
			case n41msg.N41_PFD_MANAGEMENT_RESPONSE:
				match = info.msg.Header.MessageType == n41msg.N41_PFD_MANAGEMENT_REQUEST
			case n41msg.N41_ASSOCIATION_SETUP_RESPONSE:
				match = info.msg.Header.MessageType == n41msg.N41_ASSOCIATION_SETUP_REQUEST
			case n41msg.N41_ASSOCIATION_UPDATE_RESPONSE:
				match = info.msg.Header.MessageType == n41msg.N41_ASSOCIATION_UPDATE_REQUEST
			case n41msg.N41_ASSOCIATION_RELEASE_RESPONSE:
				match = info.msg.Header.MessageType == n41msg.N41_ASSOCIATION_RELEASE_REQUEST
			case n41msg.N41_NODE_REPORT_RESPONSE:
				match = info.msg.Header.MessageType == n41msg.N41_NODE_REPORT_REQUEST
			case n41msg.N41_SESSION_SET_DELETION_RESPONSE:
				match = info.msg.Header.MessageType == n41msg.N41_SESSION_SET_DELETION_REQUEST
			case n41msg.N41_SESSION_ESTABLISHMENT_RESPONSE:
				match = info.msg.Header.MessageType == n41msg.N41_SESSION_ESTABLISHMENT_REQUEST
			case n41msg.N41_SESSION_MODIFICATION_RESPONSE:
				match = info.msg.Header.MessageType == n41msg.N41_SESSION_MODIFICATION_REQUEST
			case n41msg.N41_SESSION_DELETION_RESPONSE:
				match = info.msg.Header.MessageType == n41msg.N41_SESSION_DELETION_REQUEST
			case n41msg.N41_SESSION_REPORT_RESPONSE:
				match = info.msg.Header.MessageType == n41msg.N41_SESSION_REPORT_REQUEST
			default:
			}

			if !match {
				info.err = fmt.Errorf("Mismatched response")
			} else {
				info.rsp = msg
			}
			//terminate the sending task
			close(info.done)
		}
	}
}

// send a request then wait for a response.
// func (proto *N41) sendReq(msg *n41msg.Message, remote *net.UDPAddr) (rsp *n41msg.Message, err error) {
func (proto *N41) sendReq(msg *n41msg.Message, remote *n41types.Sbi) (rsp *n41msg.Message, err error) {
	info := newReqSendingInfo(msg, remote, proto.scheduleReqSending)
	//schedule for sending
	proto.scheduleReqSending(info)
	//wait for a response
	<-info.done
	rsp = info.rsp
	err = info.err
	return
}

// func (proto *N41) sendRsp(msg *n41msg.Message, remote *net.UDPAddr) (err error) {
func (proto *N41) sendRsp(msg *n41msg.Message, remote *n41types.Sbi) (err error) {
	if err = proto.fwd.WriteTo(msg, remote); err == nil {
		//cache the message for a certain time duration to resend in cases
		//where duplicated requests arrive
		proto.queue.add(newRspSendingInfo(msg, remote))
	}

	return
}

func (proto *N41) scheduleReqSending(info *ReqSendingInfo) {
	//increate sending retry counter for the request
	info.retry++
	//logrus.Infof("schedule for sending %d", info.retry)
	//push to sending chanel
	proto.sending <- info
}

// generate sequence number for N41 sending request
func (proto *N41) sequenceNumber() uint32 {
	return atomic.AddUint32(&proto.seq, 1)
}
