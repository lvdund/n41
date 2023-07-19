package n41msg

import (
	"encoding/json"
	"fmt"
)

func (m *Message) N41Unmarshal(data []byte) (err error) {

	// load header
	var header struct {
		Header Header `json:"header"`
	}
	if err = json.Unmarshal(data, &header); err != nil {
		return
	}

	// load Body
	switch m.Header.MessageType {
	case N41_HEARTBEAT_REQUEST:
		var loadBody struct {
			Body HeartbeatRequest
		}
		if err := json.Unmarshal(data, &loadBody); err != nil {
			return err
		}
		m.Body = loadBody.Body
	case N41_HEARTBEAT_RESPONSE:
		var loadBody struct {
			Body HeartbeatResponse
		}
		if err := json.Unmarshal(data, &loadBody); err != nil {
			return err
		}
		m.Body = loadBody
	case N41_PFD_MANAGEMENT_REQUEST:
		var loadBody struct {
			Body N41PFDManagementRequest
		}
		if err := json.Unmarshal(data, &loadBody); err != nil {
			return err
		}
		m.Body = loadBody
	case N41_PFD_MANAGEMENT_RESPONSE:
		var loadBody struct {
			Body N41PFDManagementResponse
		}
		if err := json.Unmarshal(data, &loadBody); err != nil {
			return err
		}
		m.Body = loadBody
	case N41_ASSOCIATION_SETUP_REQUEST:
		var loadBody struct {
			Body N41AssociationSetupRequest
		}
		if err := json.Unmarshal(data, &loadBody); err != nil {
			return err
		}
		m.Body = loadBody
	case N41_ASSOCIATION_SETUP_RESPONSE:
		var loadBody struct {
			Body N41AssociationSetupRequest
		}
		if err := json.Unmarshal(data, &loadBody); err != nil {
			return err
		}
		m.Body = loadBody
	case N41_ASSOCIATION_UPDATE_REQUEST:
		var loadBody struct {
			Body N41AssociationUpdateRequest
		}
		if err := json.Unmarshal(data, &loadBody); err != nil {
			return err
		}
		m.Body = loadBody
	case N41_ASSOCIATION_UPDATE_RESPONSE:
		var loadBody struct {
			Body N41AssociationUpdateResponse
		}
		if err := json.Unmarshal(data, &loadBody); err != nil {
			return err
		}
		m.Body = loadBody
	case N41_ASSOCIATION_RELEASE_REQUEST:
		var loadBody struct {
			Body N41AssociationReleaseRequest
		}
		if err := json.Unmarshal(data, &loadBody); err != nil {
			return err
		}
		m.Body = loadBody
	case N41_ASSOCIATION_RELEASE_RESPONSE:
		var loadBody struct {
			Body N41AssociationReleaseResponse
		}
		if err := json.Unmarshal(data, &loadBody); err != nil {
			return err
		}
		m.Body = loadBody
	case N41_NODE_REPORT_REQUEST:
		var loadBody struct {
			Body N41NodeReportRequest
		}
		if err := json.Unmarshal(data, &loadBody); err != nil {
			return err
		}
		m.Body = loadBody
	case N41_NODE_REPORT_RESPONSE:
		var loadBody struct {
			Body N41NodeReportResponse
		}
		if err := json.Unmarshal(data, &loadBody); err != nil {
			return err
		}
		m.Body = loadBody
	case N41_SESSION_SET_DELETION_REQUEST:
		var loadBody struct {
			Body N41SessionSetDeletionRequest
		}
		if err := json.Unmarshal(data, &loadBody); err != nil {
			return err
		}
		m.Body = loadBody
	case N41_SESSION_SET_DELETION_RESPONSE:
		var loadBody struct {
			Body N41SessionSetDeletionResponse
		}
		if err := json.Unmarshal(data, &loadBody); err != nil {
			return err
		}
		m.Body = loadBody
	case N41_SESSION_ESTABLISHMENT_REQUEST:
		var loadBody struct {
			Body N41SessionEstablishmentRequest
		}
		if err := json.Unmarshal(data, &loadBody); err != nil {
			return err
		}
		m.Body = loadBody
	case N41_SESSION_ESTABLISHMENT_RESPONSE:
		var loadBody struct {
			Body N41SessionEstablishmentResponse
		}
		if err := json.Unmarshal(data, &loadBody); err != nil {
			return err
		}
		m.Body = loadBody
	case N41_SESSION_MODIFICATION_REQUEST:
		var loadBody struct {
			Body N41SessionModificationRequest
		}
		if err := json.Unmarshal(data, &loadBody); err != nil {
			return err
		}
		m.Body = loadBody
	case N41_SESSION_MODIFICATION_RESPONSE:
		var loadBody struct {
			Body N41SessionModificationResponse
		}
		if err := json.Unmarshal(data, &loadBody); err != nil {
			return err
		}
		m.Body = loadBody
	case N41_SESSION_DELETION_REQUEST:
		var loadBody struct {
			Body N41SessionDeletionRequest
		}
		if err := json.Unmarshal(data, &loadBody); err != nil {
			return err
		}
		m.Body = loadBody
	case N41_SESSION_DELETION_RESPONSE:
		var loadBody struct {
			Body N41SessionDeletionResponse
		}
		if err := json.Unmarshal(data, &loadBody); err != nil {
			return err
		}
		m.Body = loadBody
	case N41_SESSION_REPORT_REQUEST:
		var loadBody struct {
			Body N41SessionReportRequest
		}
		if err := json.Unmarshal(data, &loadBody); err != nil {
			return err
		}
		m.Body = loadBody
	case N41_SESSION_REPORT_RESPONSE:
		var loadBody struct {
			Body N41SessionReportResponse
		}
		if err := json.Unmarshal(data, &loadBody); err != nil {
			return err
		}
		m.Body = loadBody
	default:
		return fmt.Errorf("n41: unmarshal msg type %d not supported", m.Header.MessageType)
	}

	return nil
}
