package n41types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalN41SMReqFlags(t *testing.T) {
	testData := N41SMReqFlags{
		Qaurr: true,
		Sndem: false,
		Drobu: true,
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{5}, buf)
}

func TestUnmarshalN41SMReqFlags(t *testing.T) {
	buf := []byte{5}
	var testData N41SMReqFlags
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := N41SMReqFlags{
		Qaurr: true,
		Sndem: false,
		Drobu: true,
	}
	assert.Equal(t, expectData, testData)
}
