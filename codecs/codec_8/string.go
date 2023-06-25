package codec8

import (
	"encoding/hex"

	models "github.com/rojack96/teltonika-parser/models/codec_8"
)

func Codec8String(dataPacket string) models.AVLDataArray {
	dataPacketBytes, err := hex.DecodeString(dataPacket)

	if err != nil {
		panic(err)
	}

	return Codec8Bytes(dataPacketBytes)
}
