package n41msg

import "n41/n41types"

type Message struct {
	Header Header
	Body   interface{}
}

func (message *Message) IsRequest() (IsRequest bool) {
	switch message.Header.MessageType {
	case N41_HEARTBEAT_REQUEST:
		IsRequest = true
	case N41_PFD_MANAGEMENT_REQUEST:
		IsRequest = true
	case N41_ASSOCIATION_SETUP_REQUEST:
		IsRequest = true
	case N41_ASSOCIATION_UPDATE_REQUEST:
		IsRequest = true
	case N41_ASSOCIATION_RELEASE_REQUEST:
		IsRequest = true
	case N41_NODE_REPORT_REQUEST:
		IsRequest = true
	case N41_SESSION_SET_DELETION_REQUEST:
		IsRequest = true
	case N41_SESSION_ESTABLISHMENT_REQUEST:
		IsRequest = true
	case N41_SESSION_MODIFICATION_REQUEST:
		IsRequest = true
	case N41_SESSION_DELETION_REQUEST:
		IsRequest = true
	case N41_SESSION_REPORT_REQUEST:
		IsRequest = true
	default:
		IsRequest = false
	}

	return
}

func (message *Message) IsResponse() (IsResponse bool) {
	IsResponse = false
	switch message.Header.MessageType {
	case N41_HEARTBEAT_RESPONSE:
		IsResponse = true
	case N41_PFD_MANAGEMENT_RESPONSE:
		IsResponse = true
	case N41_ASSOCIATION_SETUP_RESPONSE:
		IsResponse = true
	case N41_ASSOCIATION_UPDATE_RESPONSE:
		IsResponse = true
	case N41_ASSOCIATION_RELEASE_RESPONSE:
		IsResponse = true
	case N41_NODE_REPORT_RESPONSE:
		IsResponse = true
	case N41_SESSION_SET_DELETION_RESPONSE:
		IsResponse = true
	case N41_SESSION_ESTABLISHMENT_RESPONSE:
		IsResponse = true
	case N41_SESSION_MODIFICATION_RESPONSE:
		IsResponse = true
	case N41_SESSION_DELETION_RESPONSE:
		IsResponse = true
	case N41_SESSION_REPORT_RESPONSE:
		IsResponse = true
	default:
		IsResponse = false
	}

	return
}

type HeartbeatRequest struct {
	RecoveryTimeStamp *n41types.RecoveryTimeStamp `tlv:"96"`
}

type HeartbeatResponse struct {
	RecoveryTimeStamp *n41types.RecoveryTimeStamp `tlv:"96"`
}

type N41PFDManagementRequest struct {
	ApplicationIDsPFDs []ApplicationIDsPFDs `tlv:"58"`
}

type ApplicationIDsPFDs struct {
	ApplicationID n41types.ApplicationID `tlv:"24"`
	PFD           *PFD                   `tlv:"59"`
}

type PFD struct {
	PFDContents []n41types.PFDContents `tlv:"61"`
}

type N41PFDManagementResponse struct {
	Cause       *n41types.Cause       `tlv:"19"`
	OffendingIE *n41types.OffendingIE `tlv:"40"`
}

type N41AssociationSetupRequest struct {
	NodeID                         *n41types.NodeID                         `tlv:"60"`
	RecoveryTimeStamp              *n41types.RecoveryTimeStamp              `tlv:"96"`
	UPFunctionFeatures             *n41types.UPFunctionFeatures             `tlv:"43"`
	CPFunctionFeatures             *n41types.CPFunctionFeatures             `tlv:"89"`
	UserPlaneIPResourceInformation *n41types.UserPlaneIPResourceInformation `tlv:"116"`
}

type N41AssociationSetupResponse struct {
	NodeID                         *n41types.NodeID                         `tlv:"60"`
	Cause                          *n41types.Cause                          `tlv:"19"`
	RecoveryTimeStamp              *n41types.RecoveryTimeStamp              `tlv:"96"`
	UPFunctionFeatures             *n41types.UPFunctionFeatures             `tlv:"43"`
	CPFunctionFeatures             *n41types.CPFunctionFeatures             `tlv:"89"`
	UserPlaneIPResourceInformation *n41types.UserPlaneIPResourceInformation `tlv:"116"`
}

type N41AssociationUpdateRequest struct {
	NodeID                         *n41types.NodeID                         `tlv:"60"`
	UPFunctionFeatures             *n41types.UPFunctionFeatures             `tlv:"43"`
	CPFunctionFeatures             *n41types.CPFunctionFeatures             `tlv:"89"`
	N41AssociationReleaseRequest   *N41AssociationReleaseRequest            `tlv:"111"`
	GracefulReleasePeriod          *n41types.GracefulReleasePeriod          `tlv:"112"`
	UserPlaneIPResourceInformation *n41types.UserPlaneIPResourceInformation `tlv:"116"`
}

type N41AssociationUpdateResponse struct {
	NodeID             *n41types.NodeID             `tlv:"60"`
	Cause              *n41types.Cause              `tlv:"19"`
	UPFunctionFeatures *n41types.UPFunctionFeatures `tlv:"43"`
	CPFunctionFeatures *n41types.CPFunctionFeatures `tlv:"89"`
}

type N41AssociationReleaseRequest struct {
	NodeID *n41types.NodeID `tlv:"60"`
}

type N41AssociationReleaseResponse struct {
	NodeID *n41types.NodeID `tlv:"60"`
	Cause  *n41types.Cause  `tlv:"19"`
}

type N41NodeReportRequest struct {
	NodeID                     *n41types.NodeID                     `tlv:"60"`
	NodeReportType             *n41types.NodeReportType             `tlv:"101"`
	UserPlanePathFailureReport *n41types.UserPlanePathFailureReport `tlv:"102"`
}

type UserPlanePathFailure struct {
	RemoteGTPUPeer *n41types.RemoteGTPUPeer `tlv:"103"`
}

type N41NodeReportResponse struct {
	NodeID      *n41types.NodeID      `tlv:"60"`
	Cause       *n41types.Cause       `tlv:"19"`
	OffendingIE *n41types.OffendingIE `tlv:"40"`
}

type N41SessionSetDeletionRequest struct {
	NodeID     *n41types.NodeID `tlv:"60"`
	SGWCFQCSID *n41types.FQCSID `tlv:"65"`
	PGWCFQCSID *n41types.FQCSID `tlv:"65"`
	SGWUFQCSID *n41types.FQCSID `tlv:"65"`
	PGWUFQCSID *n41types.FQCSID `tlv:"65"`
	TWANFQCSID *n41types.FQCSID `tlv:"65"`
	EPDGFQCSID *n41types.FQCSID `tlv:"65"`
	MMEFQCSID  *n41types.FQCSID `tlv:"65"`
}

type N41SessionSetDeletionResponse struct {
	NodeID      *n41types.NodeID      `tlv:"60"`
	Cause       *n41types.Cause       `tlv:"19"`
	OffendingIE *n41types.OffendingIE `tlv:"40"`
}

type N41SessionEstablishmentRequest struct {
	NodeID                   *n41types.NodeID                   `tlv:"60"`
	CPFSEID                  *n41types.FSEID                    `tlv:"57"`
	CreatePDR                []*CreatePDR                       `tlv:"1"`
	CreateFAR                []*CreateFAR                       `tlv:"3"`
	CreateURR                []*CreateURR                       `tlv:"6"`
	CreateQER                []*CreateQER                       `tlv:"7"`
	CreateBAR                []*CreateBAR                       `tlv:"85"`
	CreateTrafficEndpoint    *CreateTrafficEndpoint             `tlv:"127"`
	PDNType                  *n41types.PDNType                  `tlv:"113"`
	SGWCFQCSID               *n41types.FQCSID                   `tlv:"65"`
	MMEFQCSID                *n41types.FQCSID                   `tlv:"65"`
	PGWCFQCSID               *n41types.FQCSID                   `tlv:"65"`
	EPDGFQCSID               *n41types.FQCSID                   `tlv:"65"`
	TWANFQCSID               *n41types.FQCSID                   `tlv:"65"`
	UserPlaneInactivityTimer *n41types.UserPlaneInactivityTimer `tlv:"117"`
	UserID                   *n41types.UserID                   `tlv:"141"`
	TraceInformation         *n41types.TraceInformation         `tlv:"152"`
}

type CreatePDR struct {
	PDRID                   *n41types.PacketDetectionRuleID   `tlv:"56"`
	Precedence              *n41types.Precedence              `tlv:"29"`
	PDI                     *PDI                              `tlv:"2"`
	OuterHeaderRemoval      *n41types.OuterHeaderRemoval      `tlv:"95"`
	FARID                   *n41types.FARID                   `tlv:"108"`
	URRID                   []*n41types.URRID                 `tlv:"81"`
	QERID                   []*n41types.QERID                 `tlv:"109"`
	ActivatePredefinedRules *n41types.ActivatePredefinedRules `tlv:"106"`
}

type PDI struct {
	SourceInterface               *n41types.SourceInterface               `tlv:"20"`
	LocalFTEID                    *n41types.FTEID                         `tlv:"21"`
	NetworkInstance               *n41types.NetworkInstance               `tlv:"22"`
	UEIPAddress                   *n41types.UEIPAddress                   `tlv:"93"`
	TrafficEndpointID             *n41types.TrafficEndpointID             `tlv:"131"`
	SDFFilter                     *n41types.SDFFilter                     `tlv:"23"`
	ApplicationID                 *n41types.ApplicationID                 `tlv:"24"`
	EthernetPDUSessionInformation *n41types.EthernetPDUSessionInformation `tlv:"142"`
	EthernetPacketFilter          *EthernetPacketFilter                   `tlv:"132"`
	QFI                           []*n41types.QFI                         `tlv:"124"`
	FramedRoute                   *n41types.FramedRoute                   `tlv:"153"`
	FramedRouting                 *n41types.FramedRouting                 `tlv:"154"`
	FramedIPv6Route               *n41types.FramedIPv6Route               `tlv:"155"`
}

type EthernetPacketFilter struct {
	EthernetFilterID         *n41types.EthernetFilterID         `tlv:"138"`
	EthernetFilterProperties *n41types.EthernetFilterProperties `tlv:"139"`
	MACAddress               *n41types.MACAddress               `tlv:"133"`
	Ethertype                *n41types.Ethertype                `tlv:"136"`
	CTAG                     *n41types.CTAG                     `tlv:"134"`
	STAG                     *n41types.STAG                     `tlv:"135"`
	SDFFilter                *n41types.SDFFilter                `tlv:"23"`
}

type CreateFAR struct {
	FARID                 *n41types.FARID                 `tlv:"108"`
	ApplyAction           *n41types.ApplyAction           `tlv:"44"`
	ForwardingParameters  *ForwardingParametersIEInFAR    `tlv:"4"`
	DuplicatingParameters *n41types.DuplicatingParameters `tlv:"5"`
	BARID                 *n41types.BARID                 `tlv:"88"`
}

type ForwardingParametersIEInFAR struct {
	DestinationInterface    *n41types.DestinationInterface  `tlv:"42"`
	NetworkInstance         *n41types.NetworkInstance       `tlv:"22"`
	RedirectInformation     *n41types.RedirectInformation   `tlv:"38"`
	OuterHeaderCreation     *n41types.OuterHeaderCreation   `tlv:"84"`
	TransportLevelMarking   *n41types.TransportLevelMarking `tlv:"30"`
	ForwardingPolicy        *n41types.ForwardingPolicy      `tlv:"41"`
	HeaderEnrichment        *n41types.HeaderEnrichment      `tlv:"98"`
	LinkedTrafficEndpointID *n41types.TrafficEndpointID     `tlv:"131"`
	Proxying                *n41types.Proxying              `tlv:"137"`
}

type DuplicatingParametersIEInFAR struct {
	DestinationInterface  *n41types.DestinationInterface  `tlv:"42"`
	OuterHeaderCreation   *n41types.OuterHeaderCreation   `tlv:"84"`
	TransportLevelMarking *n41types.TransportLevelMarking `tlv:"30"`
	ForwardingPolicy      *n41types.ForwardingPolicy      `tlv:"41"`
}

type CreateURR struct {
	URRID                     *n41types.URRID                     `tlv:"81"`
	MeasurementMethod         *n41types.MeasurementMethod         `tlv:"62"`
	ReportingTriggers         *n41types.ReportingTriggers         `tlv:"37"`
	MeasurementPeriod         *n41types.MeasurementPeriod         `tlv:"64"`
	VolumeThreshold           *n41types.VolumeThreshold           `tlv:"31"`
	VolumeQuota               *n41types.VolumeQuota               `tlv:"73"`
	TimeThreshold             *n41types.TimeThreshold             `tlv:"32"`
	TimeQuota                 *n41types.TimeQuota                 `tlv:"74"`
	QuotaHoldingTime          *n41types.QuotaHoldingTime          `tlv:"71"`
	DroppedDLTrafficThreshold *n41types.DroppedDLTrafficThreshold `tlv:"72"`
	MonitoringTime            *n41types.MonitoringTime            `tlv:"33"`
	EventInformation          *EventInformation                   `tlv:"148"`
	SubsequentVolumeThreshold *n41types.SubsequentVolumeThreshold `tlv:"34"`
	SubsequentTimeThreshold   *n41types.SubsequentTimeThreshold   `tlv:"35"`
	SubsequentVolumeQuota     *n41types.SubsequentVolumeQuota     `tlv:"121"`
	SubsequentTimeQuota       *n41types.SubsequentTimeQuota       `tlv:"122"`
	InactivityDetectionTime   *n41types.InactivityDetectionTime   `tlv:"36"`
	LinkedURRID               *n41types.LinkedURRID               `tlv:"82"`
	MeasurementInformation    *n41types.MeasurementInformation    `tlv:"100"`
	TimeQuotaMechanism        *n41types.TimeQuotaMechanism        `tlv:"115"`
	AggregatedURRs            []*AggregatedURRs                   `tlv:"118"`
	FARIDForQuotaAction       *n41types.FARID                     `tlv:"108"`
	EthernetInactivityTimer   *n41types.EthernetInactivityTimer   `tlv:"146"`
	AdditionalMonitoringTime  *AdditionalMonitoringTime           `tlv:"147"`
	QuotaValidityTime         *n41types.QuotaValidityTime         `tlv:"181"`
}

type AggregatedURRs struct {
	AggregatedURRID *n41types.AggregatedURRID `tlv:"120"`
	Multiplier      *n41types.Multiplier      `tlv:"119"`
}

type AdditionalMonitoringTime struct {
	MonitoringTime            *n41types.MonitoringTime            `tlv:"33"`
	SubsequentVolumeThreshold *n41types.SubsequentVolumeThreshold `tlv:"34"`
	SubsequentTimeThreshold   *n41types.SubsequentTimeThreshold   `tlv:"35"`
	SubsequentVolumeQuota     *n41types.SubsequentVolumeQuota     `tlv:"121"`
	SubsequentTimeQuota       *n41types.SubsequentTimeQuota       `tlv:"122"`
}

type EventInformation struct {
	EventID        *n41types.EventID        `tlv:"150"`
	EventThreshold *n41types.EventThreshold `tlv:"151"`
}

type CreateQER struct {
	QERID              *n41types.QERID              `tlv:"109"`
	QERCorrelationID   *n41types.QERCorrelationID   `tlv:"28"`
	GateStatus         *n41types.GateStatus         `tlv:"25"`
	MaximumBitrate     *n41types.MBR                `tlv:"26"`
	GuaranteedBitrate  *n41types.GBR                `tlv:"27"`
	PacketRate         *n41types.PacketRate         `tlv:"94"`
	DLFlowLevelMarking *n41types.DLFlowLevelMarking `tlv:"97"`
	QoSFlowIdentifier  *n41types.QFI                `tlv:"124"`
	ReflectiveQoS      *n41types.RQI                `tlv:"123"`
}

type CreateBAR struct {
	BARID                          *n41types.BARID                          `tlv:"88"`
	DownlinkDataNotificationDelay  *n41types.DownlinkDataNotificationDelay  `tlv:"46"`
	SuggestedBufferingPacketsCount *n41types.SuggestedBufferingPacketsCount `tlv:"140"`
}

type CreateTrafficEndpoint struct {
	TrafficEndpointID             *n41types.TrafficEndpointID             `tlv:"131"`
	LocalFTEID                    *n41types.FTEID                         `tlv:"21"`
	NetworkInstance               *n41types.NetworkInstance               `tlv:"22"`
	UEIPAddress                   *n41types.UEIPAddress                   `tlv:"93"`
	EthernetPDUSessionInformation *n41types.EthernetPDUSessionInformation `tlv:"142"`
	FramedRoute                   *n41types.FramedRoute                   `tlv:"153"`
	FramedRouting                 *n41types.FramedRouting                 `tlv:"154"`
	FramedIPv6Route               *n41types.FramedIPv6Route               `tlv:"155"`
}

type N41SessionEstablishmentResponse struct {
	NodeID                     *n41types.NodeID            `tlv:"60"`
	Cause                      *n41types.Cause             `tlv:"19"`
	OffendingIE                *n41types.OffendingIE       `tlv:"40"`
	UPFSEID                    *n41types.FSEID             `tlv:"57"`
	CreatedPDR                 *CreatedPDR                 `tlv:"8"`
	LoadControlInformation     *LoadControlInformation     `tlv:"51"`
	OverloadControlInformation *OverloadControlInformation `tlv:"54"`
	SGWUFQCSID                 *n41types.FQCSID            `tlv:"65"`
	PGWUFQCSID                 *n41types.FQCSID            `tlv:"65"`
	FailedRuleID               *n41types.FailedRuleID      `tlv:"114"`
	CreatedTrafficEndpoint     *CreatedTrafficEndpoint     `tlv:"128"`
}

type CreatedPDR struct {
	PDRID      *n41types.PacketDetectionRuleID `tlv:"56"`
	LocalFTEID *n41types.FTEID                 `tlv:"21"`
}

type LoadControlInformation struct {
	// LoadControlSequenceNumber *n41types.SequenceNumber `tlv:"52"`
	LoadMetric                *n41types.Metric         `tlv:"53"`
}

type OverloadControlInformation struct {
	// OverloadControlSequenceNumber   *n41types.SequenceNumber `tlv:"52"`
	OverloadReductionMetric         *n41types.Metric         `tlv:"53"`
	PeriodOfValidity                *n41types.Timer          `tlv:"55"`
	OverloadControlInformationFlags *n41types.OCIFlags       `tlv:"110"`
}

type CreatedTrafficEndpoint struct {
	TrafficEndpointID *n41types.TrafficEndpointID `tlv:"131"`
	LocalFTEID        *n41types.FTEID             `tlv:"21"`
}

type N41SessionModificationRequest struct {
	CPFSEID                  *n41types.FSEID                         `tlv:"57"`
	RemovePDR                []*RemovePDR                            `tlv:"15"`
	RemoveFAR                []*RemoveFAR                            `tlv:"16"`
	RemoveURR                []*RemoveURR                            `tlv:"17"`
	RemoveQER                []*RemoveQER                            `tlv:"18"`
	RemoveBAR                []*RemoveBAR                            `tlv:"87"`
	RemoveTrafficEndpoint    *RemoveTrafficEndpoint                  `tlv:"130"`
	CreatePDR                []*CreatePDR                            `tlv:"1"`
	CreateFAR                []*CreateFAR                            `tlv:"3"`
	CreateURR                []*CreateURR                            `tlv:"6"`
	CreateQER                []*CreateQER                            `tlv:"7"`
	CreateBAR                []*CreateBAR                            `tlv:"85"`
	CreateTrafficEndpoint    *CreateTrafficEndpoint                  `tlv:"127"`
	UpdatePDR                []*UpdatePDR                            `tlv:"9"`
	UpdateFAR                []*UpdateFAR                            `tlv:"10"`
	UpdateURR                []*UpdateURR                            `tlv:"13"`
	UpdateQER                []*UpdateQER                            `tlv:"14"`
	UpdateBAR                *UpdateBARN41SessionModificationRequest `tlv:"86"`
	UpdateTrafficEndpoint    *UpdateTrafficEndpoint                  `tlv:"129"`
	N41SMReqFlags            *n41types.N41SMReqFlags                 `tlv:"49"`
	QueryURR                 *QueryURR                               `tlv:"77"`
	PGWCFQCSID               *n41types.FQCSID                        `tlv:"65"`
	SGWCFQCSID               *n41types.FQCSID                        `tlv:"65"`
	MMEFQCSID                *n41types.FQCSID                        `tlv:"65"`
	EPDGFQCSID               *n41types.FQCSID                        `tlv:"65"`
	TWANFQCSID               *n41types.FQCSID                        `tlv:"65"`
	UserPlaneInactivityTimer *n41types.UserPlaneInactivityTimer      `tlv:"117"`
	QueryURRReference        *n41types.QueryURRReference             `tlv:"125"`
	TraceInformation         *n41types.TraceInformation              `tlv:"152"`
}

type UpdatePDR struct {
	PDRID                     *n41types.PacketDetectionRuleID     `tlv:"56"`
	OuterHeaderRemoval        *n41types.OuterHeaderRemoval        `tlv:"95"`
	Precedence                *n41types.Precedence                `tlv:"29"`
	PDI                       *PDI                                `tlv:"2"`
	FARID                     *n41types.FARID                     `tlv:"108"`
	URRID                     []*n41types.URRID                   `tlv:"81"`
	QERID                     []*n41types.QERID                   `tlv:"109"`
	ActivatePredefinedRules   *n41types.ActivatePredefinedRules   `tlv:"106"`
	DeactivatePredefinedRules *n41types.DeactivatePredefinedRules `tlv:"107"`
}

type UpdateFAR struct {
	FARID                       *n41types.FARID                       `tlv:"108"`
	ApplyAction                 *n41types.ApplyAction                 `tlv:"44"`
	UpdateForwardingParameters  *UpdateForwardingParametersIEInFAR    `tlv:"11"`
	UpdateDuplicatingParameters *n41types.UpdateDuplicatingParameters `tlv:"105"`
	BARID                       *n41types.BARID                       `tlv:"88"`
}

type UpdateForwardingParametersIEInFAR struct {
	DestinationInterface    *n41types.DestinationInterface  `tlv:"42"`
	NetworkInstance         *n41types.NetworkInstance       `tlv:"22"`
	RedirectInformation     *n41types.RedirectInformation   `tlv:"38"`
	OuterHeaderCreation     *n41types.OuterHeaderCreation   `tlv:"84"`
	TransportLevelMarking   *n41types.TransportLevelMarking `tlv:"30"`
	ForwardingPolicy        *n41types.ForwardingPolicy      `tlv:"41"`
	HeaderEnrichment        *n41types.HeaderEnrichment      `tlv:"98"`
	N41SMReqFlags           *n41types.N41SMReqFlags         `tlv:"49"`
	LinkedTrafficEndpointID *n41types.TrafficEndpointID     `tlv:"131"`
}

type UpdateDuplicatingParametersIEInFAR struct {
	DestinationInterface  *n41types.DestinationInterface  `tlv:"42"`
	OuterHeaderCreation   *n41types.OuterHeaderCreation   `tlv:"84"`
	TransportLevelMarking *n41types.TransportLevelMarking `tlv:"30"`
	ForwardingPolicy      *n41types.ForwardingPolicy      `tlv:"41"`
}

type UpdateURR struct {
	URRID                     *n41types.URRID                     `tlv:"81"`
	MeasurementMethod         *n41types.MeasurementMethod         `tlv:"62"`
	ReportingTriggers         *n41types.ReportingTriggers         `tlv:"37"`
	MeasurementPeriod         *n41types.MeasurementPeriod         `tlv:"64"`
	VolumeThreshold           *n41types.VolumeThreshold           `tlv:"31"`
	VolumeQuota               *n41types.VolumeQuota               `tlv:"73"`
	TimeThreshold             *n41types.TimeThreshold             `tlv:"32"`
	TimeQuota                 *n41types.TimeQuota                 `tlv:"74"`
	QuotaHoldingTime          *n41types.QuotaHoldingTime          `tlv:"71"`
	DroppedDLTrafficThreshold *n41types.DroppedDLTrafficThreshold `tlv:"72"`
	MonitoringTime            *n41types.MonitoringTime            `tlv:"33"`
	EventInformation          *EventInformation                   `tlv:"148"`
	SubsequentVolumeThreshold *n41types.SubsequentVolumeThreshold `tlv:"34"`
	SubsequentTimeThreshold   *n41types.SubsequentTimeThreshold   `tlv:"35"`
	SubsequentVolumeQuota     *n41types.SubsequentVolumeQuota     `tlv:"121"`
	SubsequentTimeQuota       *n41types.SubsequentTimeQuota       `tlv:"122"`
	InactivityDetectionTime   *n41types.InactivityDetectionTime   `tlv:"36"`
	LinkedURRID               *n41types.LinkedURRID               `tlv:"82"`
	MeasurementInformation    *n41types.MeasurementInformation    `tlv:"100"`
	TimeQuotaMechanism        *n41types.TimeQuotaMechanism        `tlv:"115"`
	AggregatedURRs            *AggregatedURRs                     `tlv:"118"`
	FARIDForQuotaAction       *n41types.FARID                     `tlv:"108"`
	EthernetInactivityTimer   *n41types.EthernetInactivityTimer   `tlv:"146"`
	AdditionalMonitoringTime  *AdditionalMonitoringTime           `tlv:"147"`
	QuotaValidityTime         *n41types.QuotaValidityTime         `tlv:"181"`
}

type UpdateQER struct {
	QERID              *n41types.QERID              `tlv:"109"`
	QERCorrelationID   *n41types.QERCorrelationID   `tlv:"28"`
	GateStatus         *n41types.GateStatus         `tlv:"25"`
	MaximumBitrate     *n41types.MBR                `tlv:"26"`
	GuaranteedBitrate  *n41types.GBR                `tlv:"27"`
	PacketRate         *n41types.PacketRate         `tlv:"94"`
	DLFlowLevelMarking *n41types.DLFlowLevelMarking `tlv:"97"`
	QoSFlowIdentifier  *n41types.QFI                `tlv:"124"`
	ReflectiveQoS      *n41types.RQI                `tlv:"123"`
}

type RemovePDR struct {
	PDRID *n41types.PacketDetectionRuleID `tlv:"56"`
}

type RemoveFAR struct {
	FARID *n41types.FARID `tlv:"108"`
}

type RemoveURR struct {
	URRID *n41types.URRID `tlv:"81"`
}

type RemoveQER struct {
	QERID *n41types.QERID `tlv:"109"`
}

type QueryURR struct {
	URRID *n41types.URRID `tlv:"81"`
}

type UpdateBARN41SessionModificationRequest struct {
	BARID                          *n41types.BARID                          `tlv:"88"`
	DownlinkDataNotificationDelay  *n41types.DownlinkDataNotificationDelay  `tlv:"46"`
	SuggestedBufferingPacketsCount *n41types.SuggestedBufferingPacketsCount `tlv:"140"`
}

type RemoveBAR struct {
	BARID *n41types.BARID `tlv:"88"`
}

type UpdateTrafficEndpoint struct {
	TrafficEndpointID *n41types.TrafficEndpointID `tlv:"131"`
	LocalFTEID        *n41types.FTEID             `tlv:"21"`
	NetworkInstance   *n41types.NetworkInstance   `tlv:"22"`
	UEIPAddress       *n41types.UEIPAddress       `tlv:"93"`
	FramedRoute       *n41types.FramedRoute       `tlv:"153"`
	FramedRouting     *n41types.FramedRouting     `tlv:"154"`
	FramedIPv6Route   *n41types.FramedIPv6Route   `tlv:"155"`
}

type RemoveTrafficEndpoint struct {
	TrafficEndpointID *n41types.TrafficEndpointID `tlv:"131"`
}

type N41SessionModificationResponse struct {
	Cause                             *n41types.Cause                              `tlv:"19"`
	OffendingIE                       *n41types.OffendingIE                        `tlv:"40"`
	CreatedPDR                        *CreatedPDR                                  `tlv:"8"`
	LoadControlInformation            *LoadControlInformation                      `tlv:"51"`
	OverloadControlInformation        *OverloadControlInformation                  `tlv:"54"`
	UsageReport                       []*UsageReportN41SessionModificationResponse `tlv:"78"`
	FailedRuleID                      *n41types.FailedRuleID                       `tlv:"114"`
	AdditionalUsageReportsInformation *n41types.AdditionalUsageReportsInformation  `tlv:"126"`
	CreatedUpdatedTrafficEndpoint     *CreatedTrafficEndpoint                      `tlv:"128"`
}

type UsageReportN41SessionModificationResponse struct {
	URRID                      *n41types.URRID               `tlv:"81"`
	URSEQN                     *n41types.URSEQN              `tlv:"104"`
	UsageReportTrigger         *n41types.UsageReportTrigger  `tlv:"63"`
	StartTime                  *n41types.StartTime           `tlv:"75"`
	EndTime                    *n41types.EndTime             `tlv:"76"`
	VolumeMeasurement          *n41types.VolumeMeasurement   `tlv:"66"`
	DurationMeasurement        *n41types.DurationMeasurement `tlv:"67"`
	TimeOfFirstPacket          *n41types.TimeOfFirstPacket   `tlv:"69"`
	TimeOfLastPacket           *n41types.TimeOfLastPacket    `tlv:"70"`
	UsageInformation           *n41types.UsageInformation    `tlv:"90"`
	QueryURRReference          *n41types.QueryURRReference   `tlv:"125"`
	EthernetTrafficInformation *EthernetTrafficInformation   `tlv:"143"`
}

type N41SessionDeletionRequest struct{}

type N41SessionDeletionResponse struct {
	Cause                      *n41types.Cause                          `tlv:"19"`
	OffendingIE                *n41types.OffendingIE                    `tlv:"40"`
	LoadControlInformation     *LoadControlInformation                  `tlv:"51"`
	OverloadControlInformation *OverloadControlInformation              `tlv:"54"`
	UsageReport                []*UsageReportN41SessionDeletionResponse `tlv:"79"`
}

type UsageReportN41SessionDeletionResponse struct {
	URRID                      *n41types.URRID               `tlv:"81"`
	URSEQN                     *n41types.URSEQN              `tlv:"104"`
	UsageReportTrigger         *n41types.UsageReportTrigger  `tlv:"63"`
	StartTime                  *n41types.StartTime           `tlv:"75"`
	EndTime                    *n41types.EndTime             `tlv:"76"`
	VolumeMeasurement          *n41types.VolumeMeasurement   `tlv:"66"`
	DurationMeasurement        *n41types.DurationMeasurement `tlv:"67"`
	TimeOfFirstPacket          *n41types.TimeOfFirstPacket   `tlv:"69"`
	TimeOfLastPacket           *n41types.TimeOfLastPacket    `tlv:"70"`
	UsageInformation           *n41types.UsageInformation    `tlv:"90"`
	EthernetTrafficInformation *EthernetTrafficInformation   `tlv:"143"`
}

type N41SessionReportRequest struct {
	ReportType                        *n41types.ReportType                        `tlv:"39"`
	DownlinkDataReport                *DownlinkDataReport                         `tlv:"83"`
	UsageReport                       []*UsageReportN41SessionReportRequest       `tlv:"80"`
	ErrorIndicationReport             *ErrorIndicationReport                      `tlv:"99"`
	LoadControlInformation            *LoadControlInformation                     `tlv:"51"`
	OverloadControlInformation        *OverloadControlInformation                 `tlv:"54"`
	AdditionalUsageReportsInformation *n41types.AdditionalUsageReportsInformation `tlv:"126"`
}

type DownlinkDataReport struct {
	PDRID                          *n41types.PacketDetectionRuleID          `tlv:"56"`
	DownlinkDataServiceInformation *n41types.DownlinkDataServiceInformation `tlv:"45"`
}

type UsageReportN41SessionReportRequest struct {
	URRID                           *n41types.URRID                  `tlv:"81"`
	URSEQN                          *n41types.URSEQN                 `tlv:"104"`
	UsageReportTrigger              *n41types.UsageReportTrigger     `tlv:"63"`
	StartTime                       *n41types.StartTime              `tlv:"75"`
	EndTime                         *n41types.EndTime                `tlv:"76"`
	VolumeMeasurement               *n41types.VolumeMeasurement      `tlv:"66"`
	DurationMeasurement             *n41types.DurationMeasurement    `tlv:"67"`
	ApplicationDetectionInformation *ApplicationDetectionInformation `tlv:"68"`
	UEIPAddress                     *n41types.UEIPAddress            `tlv:"93"`
	NetworkInstance                 *n41types.NetworkInstance        `tlv:"22"`
	TimeOfFirstPacket               *n41types.TimeOfFirstPacket      `tlv:"69"`
	TimeOfLastPacket                *n41types.TimeOfLastPacket       `tlv:"70"`
	UsageInformation                *n41types.UsageInformation       `tlv:"90"`
	QueryURRReference               *n41types.QueryURRReference      `tlv:"125"`
	EventReporting                  *EventReporting                  `tlv:"149"`
	EthernetTrafficInformation      *EthernetTrafficInformation      `tlv:"143"`
}

type ApplicationDetectionInformation struct {
	ApplicationID         *n41types.ApplicationID         `tlv:"24"`
	ApplicationInstanceID *n41types.ApplicationInstanceID `tlv:"91"`
	FlowInformation       *n41types.FlowInformation       `tlv:"92"`
}

type EventReporting struct {
	EventID *n41types.EventID `tlv:"150"`
}

type EthernetTrafficInformation struct {
	MACAddressesDetected *n41types.MACAddressesDetected `tlv:"144"`
	MACAddressesRemoved  *n41types.MACAddressesRemoved  `tlv:"145"`
}

type ErrorIndicationReport struct {
	RemoteFTEID *n41types.FTEID `tlv:"21"`
}

type N41SessionReportResponse struct {
	Cause        *n41types.Cause                             `tlv:"19"`
	OffendingIE  *n41types.OffendingIE                       `tlv:"40"`
	UpdateBAR    *n41types.UpdateBARN41SessionReportResponse `tlv:"12"`
	SxSRRspFlags *n41types.N41SRRspFlags                     `tlv:"50"`
}

type UpdateBARIEInN41SessionReportResponse struct {
	BARID                           *n41types.BARID                           `tlv:"88"`
	DownlinkDataNotificationDelay   *n41types.DownlinkDataNotificationDelay   `tlv:"46"`
	DLBufferingDuration             *n41types.DLBufferingDuration             `tlv:"47"`
	DLBufferingSuggestedPacketCount *n41types.DLBufferingSuggestedPacketCount `tlv:"48"`
	SuggestedBufferingPacketsCount  *n41types.SuggestedBufferingPacketsCount  `tlv:"140"`
}
