package codec

import (
	"fmt"

	"github.com/rojack96/gonika/codec/constant"
	codec16 "github.com/rojack96/gonika/codec/device_data_sending/codec_16"
	codec8 "github.com/rojack96/gonika/codec/device_data_sending/codec_8"
	codec8ext "github.com/rojack96/gonika/codec/device_data_sending/codec_8e"
	codec12 "github.com/rojack96/gonika/codec/gprs_message/codec_12"
	"github.com/rojack96/gonika/codec/models"
	"github.com/rojack96/gonika/codec/utils"
)

type AvlDecoder interface {
	Decode() *models.AvlDataPacket
}

type GprsDecoder interface {
	Decode() *models.ResponseMessage
}

type Encoder interface {
	Encode() []byte
	EncodeString() string
}

func DecoderFactory(avlDataPacket any) (any, error) {
	const codecIdIndex = 8

	var (
		data []byte
		err  error
	)

	if data, err = utils.TransformData(avlDataPacket); err != nil {
		return nil, fmt.Errorf("failed to transform data: %v", err)
	}

	if len(data) <= codecIdIndex {
		return nil, fmt.Errorf("invalid packet length: %d", len(data))
	}

	codecID := data[codecIdIndex]

	var decoders = map[byte]func([]byte) any{
		constant.Codec8:    func(b []byte) any { return codec8.New(b) },
		constant.Codec8ext: func(b []byte) any { return codec8ext.New(b) },
		constant.Codec16:   func(b []byte) any { return codec16.New(b) },
		constant.Codec12:   func(b []byte) any { return codec12.New(b) },
		constant.Codec13:   func(b []byte) any { return nil },
		constant.Codec14:   func(b []byte) any { return nil },
	}

	if ctor, ok := decoders[codecID]; ok {
		return ctor(data), nil
	}
	return nil, fmt.Errorf("unsupported codec: 0x%X", codecID)
}
