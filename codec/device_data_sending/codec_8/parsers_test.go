package codec8

import "testing"

func TestParseIoOneByte(t *testing.T) {
	c8 := &codec8{
		avlDataPacket: []byte{0x00},
	}

	c8.parseIo(1, 0, nil)

}
