package n41types

import (
	"fmt"
)

type N41SRRspFlags struct {
	Drobu bool
}

func (p *N41SRRspFlags) MarshalBinary() (data []byte, err error) {
	// Octet 5
	tmpUint8 := btou(p.Drobu)
	data = append([]byte(""), tmpUint8)

	return data, nil
}

func (p *N41SRRspFlags) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}
	p.Drobu = utob(data[idx] & BitMask1)
	idx = idx + 1

	if length != idx {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}

	return nil
}
