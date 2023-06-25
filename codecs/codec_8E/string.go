package codec8e

import (
	"encoding/hex"

	models "github.com/rojack96/teltonika-parser/models/codec_8E"
)

func Codec8EString(dataPacket string) models.AVLDataArray {
	dataPacketBytes, err := hex.DecodeString(dataPacket)

	if err != nil {
		panic(err)
	}

	return Codec8EBytes(dataPacketBytes)
}
