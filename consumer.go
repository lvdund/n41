package n41

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"n41/n41msg"
	"n41/n41types"
	"net/http"
)

/*
func (proto *N41) SendN41AssociationSetupRequest(remote Endpoint) (rsp *n41msg.N41AssociationSetupResponse, err error) {
	req := &n41msg.N41AssociationSetupRequest{
		NodeID: proto.ctx.NodeId(),
		RecoveryTimeStamp: &n41types.RecoveryTimeStamp{
			RecoveryTimeStamp: proto.fwd.When(),
		},
		CPFunctionFeatures: &n41types.CPFunctionFeatures{
			SupportedFeatures: 0,
		},
	}

	reqmsg := &n41msg.Message{
		Header: n41msg.Header{
			Version:        n41msg.N41Version,
			MP:             0,
			S:              n41msg.SEID_NOT_PRESENT,
			MessageType:    n41msg.N41_ASSOCIATION_SETUP_REQUEST,
		},
		Body: req,
	}
	var rspmsg *n41msg.Message
	// if rspmsg, err = proto.sendReq(reqmsg, remote.UdpAddr()); err == nil {
	if rspmsg, err = proto.sendReq(reqmsg, remote.Addr()); err == nil {
		body := rspmsg.Body.(n41msg.N41AssociationSetupResponse)
		rsp = &body
	}
	return
}
*/

func (proto *N41) SendN41AssociationSetupRequest(remote Endpoint) (rsp *n41msg.N41AssociationSetupResponse, err error) {
	req := &n41msg.N41AssociationSetupRequest{
		NodeID: proto.ctx.NodeId(),
		RecoveryTimeStamp: &n41types.RecoveryTimeStamp{
			RecoveryTimeStamp: proto.fwd.When(),
		},
		CPFunctionFeatures: &n41types.CPFunctionFeatures{
			SupportedFeatures: 0,
		},
	}

	reqmsg := &n41msg.Message{
		Header: n41msg.Header{
			Version:     n41msg.N41Version,
			MP:          0,
			S:           n41msg.SEID_NOT_PRESENT,
			MessageType: n41msg.N41_ASSOCIATION_SETUP_REQUEST,
		},
		Body: req,
	}

	reqData, err := reqmsg.Marshal()
	if err != nil {
		return nil, err
	}

	// Gửi yêu cầu bằng HTTP POST request
	respData, err := proto.sendHTTPRequest(remote.Addr().String(), "POST", reqData)
	if err != nil {
		return nil, err
	}

	// Giải mã phản hồi từ dạng bytes thành cấu trúc N41AssociationSetupResponse
	rspmsg := &n41msg.Message{}
	if err := rspmsg.N41Unmarshal(respData); err != nil {
		return nil, err
	}

	body, ok := rspmsg.Body.(n41msg.N41AssociationSetupResponse)
	if !ok {
		return nil, fmt.Errorf("Invalid response body type")
	}

	return &body, nil
}

func (proto *N41) sendHTTPRequest(remoteAddr string, method string, data []byte) ([]byte, error) {
	url := "http://" + remoteAddr // URL của đích cần gửi yêu cầu HTTP
	resp, err := http.Post(url, "application/octet-stream", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respData, nil
}

func (proto *N41) SendN41AssociationReleaseRequest(remote Endpoint) (rsp *n41msg.N41AssociationReleaseResponse, err error) {
	req := &n41msg.N41AssociationReleaseRequest{
		NodeID: proto.ctx.NodeId(),
	}

	reqmsg := &n41msg.Message{
		Header: n41msg.Header{
			Version:     n41msg.N41Version,
			MP:          0,
			S:           n41msg.SEID_NOT_PRESENT,
			MessageType: n41msg.N41_ASSOCIATION_RELEASE_REQUEST,
		},
		Body: req,
	}

	var rspmsg *n41msg.Message
	// if rspmsg, err = proto.sendReq(reqmsg, remote.UdpAddr()); err == nil {
	if rspmsg, err = proto.sendReq(reqmsg, remote.Addr()); err == nil {
		body := rspmsg.Body.(n41msg.N41AssociationReleaseResponse)
		rsp = &body
	}
	return
}

func (proto *N41) SendN41SessionDeletionRequest(session N41Session) (rsp *n41msg.N41SessionDeletionResponse, err error) {
	reqbody := &n41msg.N41SessionDeletionRequest{}
	reqmsg := &n41msg.Message{
		Header: n41msg.Header{
			Version:         n41msg.N41Version,
			MP:              1,
			S:               n41msg.SEID_PRESENT,
			MessageType:     n41msg.N41_SESSION_DELETION_REQUEST,
			SEID:            session.RemoteSeid(),
			MessagePriority: 12,
		},
		Body: reqbody,
	}
	session.FillDeletionRequest(reqbody)
	var rspmsg *n41msg.Message
	// if rspmsg, err = proto.sendReq(reqmsg, session.UdpAddr()); err == nil {
	if rspmsg, err = proto.sendReq(reqmsg, session.Addr()); err == nil {
		if rspmsg.Header.SEID == session.LocalSeid() {
			body := rspmsg.Body.(n41msg.N41SessionDeletionResponse)
			rsp = &body
		} else {
			err = fmt.Errorf("mismatched SEID")
		}
	}
	return
}

func (proto *N41) SendN41HeartbeatRequest(remote Endpoint) (rsp *n41msg.HeartbeatResponse, err error) {
	req := &n41msg.HeartbeatRequest{
		RecoveryTimeStamp: &n41types.RecoveryTimeStamp{
			RecoveryTimeStamp: proto.fwd.When(),
		},
	}

	reqmsg := &n41msg.Message{
		Header: n41msg.Header{
			Version:     n41msg.N41Version,
			MP:          0,
			S:           n41msg.SEID_NOT_PRESENT,
			MessageType: n41msg.N41_HEARTBEAT_REQUEST,
		},
		Body: req,
	}
	var rspmsg *n41msg.Message
	if rspmsg, err = proto.sendReq(reqmsg, remote.Addr()); err == nil {
		body := rspmsg.Body.(n41msg.HeartbeatResponse)
		rsp = &body
	}

	return
}

func (proto *N41) SendN41SessionEstablishmentRequest(session N41Session) (rsp *n41msg.N41SessionEstablishmentResponse, err error) {
	reqbody := &n41msg.N41SessionEstablishmentRequest{
		NodeID: proto.ctx.NodeId(),
	}
	reqmsg := &n41msg.Message{
		Header: n41msg.Header{
			Version:         n41msg.N41Version,
			MP:              1,
			S:               n41msg.SEID_PRESENT,
			MessageType:     n41msg.N41_SESSION_ESTABLISHMENT_REQUEST,
			SEID:            0, /*session.RemoteSeid()*/
			MessagePriority: 0,
		},
		Body: reqbody,
	}

	session.FillEstablishmentRequest(reqbody)

	var rspmsg *n41msg.Message
	// if rspmsg, err = proto.sendReq(reqmsg, session.UdpAddr()); err == nil {
	if rspmsg, err = proto.sendReq(reqmsg, session.Addr()); err == nil {
		if rspmsg.Header.SEID == session.LocalSeid() {
			body := rspmsg.Body.(n41msg.N41SessionEstablishmentResponse)
			rsp = &body
		} else {
			err = fmt.Errorf("mismatched SEID")
		}
	}

	return
}

func (proto *N41) SendN41SessionModificationRequest(session N41Session) (rsp *n41msg.N41SessionModificationResponse, err error) {
	reqbody := &n41msg.N41SessionModificationRequest{}
	reqmsg := &n41msg.Message{
		Header: n41msg.Header{
			Version:         n41msg.N41Version,
			MP:              1,
			S:               n41msg.SEID_PRESENT,
			MessageType:     n41msg.N41_SESSION_MODIFICATION_REQUEST,
			SEID:            session.RemoteSeid(),
			MessagePriority: 12,
		},
		Body: reqbody,
	}

	session.FillModificationRequest(reqbody)

	var rspmsg *n41msg.Message
	if rspmsg, err = proto.sendReq(reqmsg, session.Addr()); err == nil {
		if rspmsg.Header.SEID == session.LocalSeid() {
			body := rspmsg.Body.(n41msg.N41SessionModificationResponse)
			rsp = &body
		} else {
			err = fmt.Errorf("mismatched SEID")
		}
	}

	return
}
