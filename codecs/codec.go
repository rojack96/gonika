package codecs

import (
	"encoding/hex"
	"fmt"

	codec8 "github.com/rojack96/gonika/codecs/codec_8"
)

type Decoder interface {
	Decode() any
}

type Encoder interface {
	Encode(avlDataArray any) []byte
	EncodeString(avlDataArray any) string
}

func DecoderFactory(avlDataPacket any) (Decoder, error) {
	const codecIdIndex = 8

	var (
		data []byte
		err  error
	)

	// Convert string to []byte if necessary
	switch v := avlDataPacket.(type) {
	case []byte:
		data = v
	case string:
		if data, err = hex.DecodeString(v); err != nil {
			return nil, fmt.Errorf("invalid hex string: %v", err)
		}
	default:
		return nil, fmt.Errorf("invalid packet type: expected []byte or string, got %T", avlDataPacket)
	}

	if len(data) <= codecIdIndex {
		return nil, fmt.Errorf("invalid packet length: %d", len(data))
	}

	codecID := data[codecIdIndex]

	var decoders = map[byte]func([]byte) Decoder{
		0x08: func(b []byte) Decoder { return codec8.New(b) },
		0x8E: func(b []byte) Decoder { return codec8.New(b) },
	}

	if ctor, ok := decoders[codecID]; ok {
		return ctor(data), nil
	}
	return nil, fmt.Errorf("unsupported codec: 0x%X", codecID)

}
