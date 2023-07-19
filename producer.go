package n41

import (
	"n41/n41msg"
	"n41/n41types"

	"github.com/sirupsen/logrus"
)

type AssociationProducer interface {
	HandleAssociationSetupRequest(string, *n41msg.N41AssociationSetupRequest) (*n41msg.N41AssociationSetupResponse, error)
	HandleAssociationReleaseRequest(string, *n41msg.N41AssociationReleaseRequest) (*n41msg.N41AssociationReleaseResponse, error)
	HandleHeartbeatRequest(string, *n41msg.HeartbeatRequest) (*n41msg.HeartbeatResponse, error)
}
type SessionProducer interface {
	HandleSessionReportRequest(string, uint64, *n41msg.N41SessionReportRequest) (*n41msg.N41SessionReportResponse, uint64, error)
}

type Producer interface {
	AssociationProducer
	SessionProducer
}

// func (proto *N41) handleReq(remote *net.UDPAddr, msg *n41msg.Message) {
func (proto *N41) handleReq(remote *n41types.Sbi, msg *n41msg.Message) {
	logrus.Debugf("receive a request of type %d from %s", msg.Header.MessageType, remote)
	if infoinf := proto.queue.find(remote.String(), msg.Header.SequenceNumber); infoinf != nil {
		//duplicated request
		if info, ok := infoinf.(*RspSendingInfo); ok {
			//resend the response
			logrus.Warnf("re-send the response to %s", info.remote.String())
			proto.fwd.WriteTo(info.msg, remote)
		}
	}
	//a new request
	switch msg.Header.MessageType {

	case n41msg.N41_HEARTBEAT_REQUEST:
		proto.handleHeartbeatReq(remote, msg)

	case n41msg.N41_PFD_MANAGEMENT_REQUEST:

	case n41msg.N41_ASSOCIATION_SETUP_REQUEST:
		proto.handleAssSetReq(remote, msg)

	case n41msg.N41_ASSOCIATION_UPDATE_REQUEST:

	case n41msg.N41_ASSOCIATION_RELEASE_REQUEST:
		proto.handleAssRelReq(remote, msg)

	case n41msg.N41_NODE_REPORT_REQUEST:
		//proto.handleNodeRepReq(remote, msg)

	case n41msg.N41_SESSION_SET_DELETION_REQUEST:

	case n41msg.N41_SESSION_ESTABLISHMENT_REQUEST:

	case n41msg.N41_SESSION_MODIFICATION_REQUEST:

	case n41msg.N41_SESSION_DELETION_REQUEST:

	case n41msg.N41_SESSION_REPORT_REQUEST:
		proto.handleSessRepReq(remote, msg)

	default:
	}
}

// func (proto *N41) handleAssSetReq(remote *net.UDPAddr, msg *n41msg.Message) {
func (proto *N41) handleAssSetReq(remote *n41types.Sbi, msg *n41msg.Message) {
	req := msg.Body.(n41msg.N41AssociationSetupRequest)
	if body, err := proto.ahandler.HandleAssociationSetupRequest(remote.String(), &req); err == nil {
		body.RecoveryTimeStamp = &n41types.RecoveryTimeStamp{
			RecoveryTimeStamp: proto.fwd.When(),
		}
		body.NodeID = proto.ctx.NodeId()

		rsp := &n41msg.Message{
			Header: n41msg.Header{
				Version:        n41msg.N41Version,
				MP:             0,
				S:              n41msg.SEID_NOT_PRESENT,
				MessageType:    n41msg.N41_ASSOCIATION_SETUP_RESPONSE,
				// SequenceNumber: msg.Header.SequenceNumber,
			},
			Body: body,
		}
		proto.sendRsp(rsp, remote)
	}
}

// func (proto *N41) handleAssRelReq(remote *net.UDPAddr, msg *n41msg.Message) {
func (proto *N41) handleAssRelReq(remote *n41types.Sbi, msg *n41msg.Message) {
	req := msg.Body.(n41msg.N41AssociationReleaseRequest)
	if body, err := proto.ahandler.HandleAssociationReleaseRequest(remote.String(), &req); err == nil {
		body.NodeID = proto.ctx.NodeId()
		rsp := &n41msg.Message{
			Header: n41msg.Header{
				Version:        n41msg.N41Version,
				MP:             0,
				S:              n41msg.SEID_NOT_PRESENT,
				MessageType:    n41msg.N41_ASSOCIATION_RELEASE_RESPONSE,
				// SequenceNumber: msg.Header.SequenceNumber,
			},
			Body: body,
		}
		proto.sendRsp(rsp, remote)
	}
}

// func (proto *N41) handleSessRepReq(remote *net.UDPAddr, msg *n41msg.Message) {
func (proto *N41) handleSessRepReq(remote *n41types.Sbi, msg *n41msg.Message) {
	req := msg.Body.(n41msg.N41SessionReportRequest)
	if body, seid, err := proto.shandler.HandleSessionReportRequest(remote.String(), msg.Header.SEID, &req); err == nil {
		rsp := &n41msg.Message{
			Header: n41msg.Header{
				Version:        n41msg.N41Version,
				MP:             0,
				S:              n41msg.SEID_PRESENT,
				MessageType:    n41msg.N41_SESSION_REPORT_RESPONSE,
				// SequenceNumber: msg.Header.SequenceNumber,
				SEID:           seid,
			},
			Body: body,
		}
		proto.sendRsp(rsp, remote)
	}

}

// func (proto *N41) handleHeartbeatReq(remote *net.UDPAddr, msg *n41msg.Message) {
func (proto *N41) handleHeartbeatReq(remote *n41types.Sbi, msg *n41msg.Message) {
	req := msg.Body.(n41msg.HeartbeatRequest)
	if body, err := proto.ahandler.HandleHeartbeatRequest(remote.String(), &req); err == nil {
		body.RecoveryTimeStamp = &n41types.RecoveryTimeStamp{
			RecoveryTimeStamp: proto.fwd.When(),
		}

		rsp := &n41msg.Message{
			Header: n41msg.Header{
				Version:        n41msg.N41Version,
				MP:             0,
				S:              n41msg.SEID_NOT_PRESENT,
				MessageType:    n41msg.N41_HEARTBEAT_RESPONSE,
				// SequenceNumber: msg.Header.SequenceNumber,
			},
			Body: body,
		}
		proto.sendRsp(rsp, remote)
	}
}
