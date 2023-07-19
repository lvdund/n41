package n41types

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/sirupsen/logrus"
)

// PFDContents - describe in TS 29.244 Figure 8.2.39-1: PFD Contents
type PFDContents struct {
	FlowDescription  string
	URL              string
	DomainName       string
	CustomPFDContent []byte
}

func (p *PFDContents) MarshalBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(nil)
	var presenceByte, spareByte byte
	// set presence header
	if p.FlowDescription != "" {
		presenceByte |= BitMask1
	}
	if p.URL != "" {
		presenceByte |= BitMask2
	}
	if p.DomainName != "" {
		presenceByte |= BitMask3
	}
	if p.CustomPFDContent != nil {
		presenceByte |= BitMask4
	}

	if err := binary.Write(buf, binary.BigEndian, presenceByte); err != nil {
		logrus.Warnf("presenceByte write failed: %v", err)
	}
	if err := binary.Write(buf, binary.BigEndian, spareByte); err != nil {
		logrus.Warnf("spareByte write failed: %v", err)
	}

	if p.FlowDescription != "" {
		if err := binary.Write(buf, binary.BigEndian, uint16(len(p.FlowDescription))); err != nil {
			logrus.Warnf("FlowDescription write failed: %v", err)
		}
		buf.WriteString(p.FlowDescription)
	}

	if p.URL != "" {
		if err := binary.Write(buf, binary.BigEndian, uint16(len(p.URL))); err != nil {
			logrus.Warnf("URL write failed: %v", err)
		}
		buf.WriteString(p.URL)
	}

	if p.DomainName != "" {
		if err := binary.Write(buf, binary.BigEndian, uint16(len(p.DomainName))); err != nil {
			logrus.Warnf("DomainName write failed: %v", err)
		}
		buf.WriteString(p.DomainName)
	}

	if p.CustomPFDContent != nil {
		if err := binary.Write(buf, binary.BigEndian, uint16(len(p.CustomPFDContent))); err != nil {
			logrus.Warnf("CustomPFDContent length write failed: %v", err)
		}
		if err := binary.Write(buf, binary.BigEndian, p.CustomPFDContent); err != nil {
			logrus.Warnf("CustomPFDContent write failed: %v", err)
		}
	}

	return buf.Bytes(), nil
}

func (p *PFDContents) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0

	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}

	presenceByte := data[idx]
	// presenceByte & spareByte
	idx = idx + 2

	// flow Description presence
	if utob(presenceByte & BitMask1) {
		if length < idx+2 {
			return fmt.Errorf("Inadequate TLV length: %d", length)
		}
		flowDescriptionLen := binary.BigEndian.Uint16(data[idx:])
		idx = idx + 2

		if length < idx+flowDescriptionLen {
			return fmt.Errorf("Inadequate TLV length: %d", length)
		}
		p.FlowDescription = string(data[idx : idx+flowDescriptionLen])
		idx = idx + flowDescriptionLen
	}

	// URL presence
	if utob(presenceByte & BitMask2) {
		if length < idx+2 {
			return fmt.Errorf("Inadequate TLV length: %d", length)
		}
		urlLen := binary.BigEndian.Uint16(data[idx:])
		idx = idx + 2

		if length < idx+urlLen {
			return fmt.Errorf("Inadequate TLV length: %d", length)
		}
		p.URL = string(data[idx : idx+urlLen])
		idx = idx + urlLen
	}

	// domain name presence
	if utob(presenceByte & BitMask3) {
		if length < idx+2 {
			return fmt.Errorf("Inadequate TLV length: %d", length)
		}
		domainNameLen := binary.BigEndian.Uint16(data[idx:])
		idx = idx + 2

		if length < idx+domainNameLen {
			return fmt.Errorf("Inadequate TLV length: %d", length)
		}
		p.DomainName = string(data[idx : idx+domainNameLen])
		idx = idx + domainNameLen
	}

	// custom PFD content presence
	if utob(presenceByte & BitMask4) {
		if length < idx+2 {
			return fmt.Errorf("Inadequate TLV length: %d", length)
		}
		custemPFDContentLen := binary.BigEndian.Uint16(data[idx:])
		idx += 2

		if length < idx+custemPFDContentLen {
			return fmt.Errorf("Inadequate TLV length: %d", length)
		}
		p.CustomPFDContent = data[idx : idx+custemPFDContentLen]
		// idx += custemPFDContentLen
	}

	return nil
}
