package n41types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalN41SRRspFlags(t *testing.T) {
	testData := N41SRRspFlags{
		Drobu: true,
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{1}, buf)
}

func TestUnmarshalN41SRRspFlags(t *testing.T) {
	buf := []byte{1}
	var testData N41SRRspFlags
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := N41SRRspFlags{
		Drobu: true,
	}
	assert.Equal(t, expectData, testData)
}
