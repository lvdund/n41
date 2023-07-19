package n41msg

import (
	"fmt"

	"github.com/free5gc/tlv"
)

func (m *Message) Unmarshal(data []byte) error {
	if err := m.Header.UnmarshalBinary(data); err != nil {
		return fmt.Errorf("n41: unmarshal msg failed: %s", err)
	}

	// Check Message Length field in header
	if int(m.Header.MessageLength) != len(data)-4 {
		return fmt.Errorf("Incorrect Message Length: Expected %d, got %d", m.Header.MessageLength, len(data)-4)
	}
	switch m.Header.MessageType {
	case N41_HEARTBEAT_REQUEST:
		Body := HeartbeatRequest{}
		if err := tlv.Unmarshal(data[m.Header.Len():], &Body); err != nil {
			return err
		}
		m.Body = Body
	case N41_HEARTBEAT_RESPONSE:
		Body := HeartbeatResponse{}
		if err := tlv.Unmarshal(data[m.Header.Len():], &Body); err != nil {
			return err
		}
		m.Body = Body
	case N41_PFD_MANAGEMENT_REQUEST:
		Body := N41PFDManagementRequest{}
		if err := tlv.Unmarshal(data[m.Header.Len():], &Body); err != nil {
			return err
		}
		m.Body = Body
	case N41_PFD_MANAGEMENT_RESPONSE:
		Body := N41PFDManagementResponse{}
		if err := tlv.Unmarshal(data[m.Header.Len():], &Body); err != nil {
			return err
		}
		m.Body = Body
	case N41_ASSOCIATION_SETUP_REQUEST:
		Body := N41AssociationSetupRequest{}
		if err := tlv.Unmarshal(data[m.Header.Len():], &Body); err != nil {
			return err
		}
		m.Body = Body
	case N41_ASSOCIATION_SETUP_RESPONSE:
		Body := N41AssociationSetupResponse{}
		if err := tlv.Unmarshal(data[m.Header.Len():], &Body); err != nil {
			return err
		}
		m.Body = Body
	case N41_ASSOCIATION_UPDATE_REQUEST:
		Body := N41AssociationUpdateRequest{}
		if err := tlv.Unmarshal(data[m.Header.Len():], &Body); err != nil {
			return err
		}
		m.Body = Body
	case N41_ASSOCIATION_UPDATE_RESPONSE:
		Body := N41AssociationUpdateResponse{}
		if err := tlv.Unmarshal(data[m.Header.Len():], &Body); err != nil {
			return err
		}
		m.Body = Body
	case N41_ASSOCIATION_RELEASE_REQUEST:
		Body := N41AssociationReleaseRequest{}
		if err := tlv.Unmarshal(data[m.Header.Len():], &Body); err != nil {
			return err
		}
		m.Body = Body
	case N41_ASSOCIATION_RELEASE_RESPONSE:
		Body := N41AssociationReleaseResponse{}
		if err := tlv.Unmarshal(data[m.Header.Len():], &Body); err != nil {
			return err
		}
		m.Body = Body
	case N41_NODE_REPORT_REQUEST:
		Body := N41NodeReportRequest{}
		if err := tlv.Unmarshal(data[m.Header.Len():], &Body); err != nil {
			return err
		}
		m.Body = Body
	case N41_NODE_REPORT_RESPONSE:
		Body := N41NodeReportResponse{}
		if err := tlv.Unmarshal(data[m.Header.Len():], &Body); err != nil {
			return err
		}
		m.Body = Body
	case N41_SESSION_SET_DELETION_REQUEST:
		Body := N41SessionSetDeletionRequest{}
		if err := tlv.Unmarshal(data[m.Header.Len():], &Body); err != nil {
			return err
		}
		m.Body = Body
	case N41_SESSION_SET_DELETION_RESPONSE:
		Body := N41SessionSetDeletionResponse{}
		if err := tlv.Unmarshal(data[m.Header.Len():], &Body); err != nil {
			return err
		}
		m.Body = Body
	case N41_SESSION_ESTABLISHMENT_REQUEST:
		Body := N41SessionEstablishmentRequest{}
		if err := tlv.Unmarshal(data[m.Header.Len():], &Body); err != nil {
			return err
		}
		m.Body = Body
	case N41_SESSION_ESTABLISHMENT_RESPONSE:
		Body := N41SessionEstablishmentResponse{}
		if err := tlv.Unmarshal(data[m.Header.Len():], &Body); err != nil {
			return err
		}
		m.Body = Body
	case N41_SESSION_MODIFICATION_REQUEST:
		Body := N41SessionModificationRequest{}
		if err := tlv.Unmarshal(data[m.Header.Len():], &Body); err != nil {
			return err
		}
		m.Body = Body
	case N41_SESSION_MODIFICATION_RESPONSE:
		Body := N41SessionModificationResponse{}
		if err := tlv.Unmarshal(data[m.Header.Len():], &Body); err != nil {
			return err
		}
		m.Body = Body
	case N41_SESSION_DELETION_REQUEST:
		Body := N41SessionDeletionRequest{}
		if err := tlv.Unmarshal(data[m.Header.Len():], &Body); err != nil {
			return err
		}
		m.Body = Body
	case N41_SESSION_DELETION_RESPONSE:
		Body := N41SessionDeletionResponse{}
		if err := tlv.Unmarshal(data[m.Header.Len():], &Body); err != nil {
			return err
		}
		m.Body = Body
	case N41_SESSION_REPORT_REQUEST:
		Body := N41SessionReportRequest{}
		if err := tlv.Unmarshal(data[m.Header.Len():], &Body); err != nil {
			return err
		}
		m.Body = Body
	case N41_SESSION_REPORT_RESPONSE:
		Body := N41SessionReportResponse{}
		if err := tlv.Unmarshal(data[m.Header.Len():], &Body); err != nil {
			return err
		}
		m.Body = Body
	default:
		return fmt.Errorf("n41: unmarshal msg type %d not supported", m.Header.MessageType)
	}
	return nil
}
