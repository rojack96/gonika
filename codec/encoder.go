package codec

import (
	"encoding/hex"
	"fmt"

	m "github.com/rojack96/gonika/codec/models"
)

type AvlEncoder interface {
	EncodeTCP(avlDataArray []m.AvlDataArrayEncoder) ([]byte, error)
	EncodeUDP(imei string, avlDataArray []m.AvlDataArrayEncoder) ([]byte, error)
}

type GprsEncoder interface {
	Encode() []byte
}

func transformData(dataBuffer any) ([]byte, error) {
	switch v := dataBuffer.(type) {

	case []byte:
		return v, nil

	case string:
		data, err := hex.DecodeString(v)
		if err != nil {
			return nil, fmt.Errorf("invalid hex string: %w", err)
		}
		return data, nil

	default:
		return nil, fmt.Errorf(
			"invalid packet type: expected []byte or hex string, got %T",
			dataBuffer,
		)
	}
}
