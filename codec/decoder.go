package codec

import (
	"fmt"

	"github.com/rojack96/gonika/codec/constant"
	codec16 "github.com/rojack96/gonika/codec/device_data_sending/codec_16"
	codec8 "github.com/rojack96/gonika/codec/device_data_sending/codec_8"
	codec8ext "github.com/rojack96/gonika/codec/device_data_sending/codec_8e"
	codec12 "github.com/rojack96/gonika/codec/gprs_message/codec_12"
	"github.com/rojack96/gonika/codec/models"
)

// configuration for the decoder
type Config struct {
	timestampLayout string
}

type Option func(*Config)

// WithTimestampLayout sets the timestamp layout for the decoder.
func WithTimestampLayout(timestampLayout string) Option {
	return func(c *Config) {
		c.timestampLayout = timestampLayout
	}
}

// Decoder is the main struct responsible for decoding AVL data packets and GPRS messages.

type AvlDecoder interface {
	// Decode decodes the AVL data packet and returns a structured AvlDataPacket.
	Decode() *models.AvlDataPacket
	// DecodeFlat decodes the AVL data packet and returns a structured AvlDataPacketFlat.
	DecodeFlat() *models.AvlDataPacketFlat
}

type GprsDecoder interface {
	Decode() *models.ResponseMessage
}

func AvlDataDecoderFactory(avlDataPacket any) (AvlDecoder, error) {
	const codecIdIndex = 8

	var (
		data []byte
		err  error
	)

	if data, err = transformData(avlDataPacket); err != nil {
		return nil, fmt.Errorf("failed to transform data: %v", err)
	}

	if len(data) <= codecIdIndex {
		return nil, fmt.Errorf("invalid packet length: %d", len(data))
	}

	codecID := data[codecIdIndex]

	var decoders = map[byte]func([]byte) AvlDecoder{
		constant.Codec8:    func(b []byte) AvlDecoder { return codec8.New(b) },
		constant.Codec8ext: func(b []byte) AvlDecoder { return codec8ext.New(b) },
		constant.Codec16:   func(b []byte) AvlDecoder { return codec16.New(b) },
	}

	if ctor, ok := decoders[codecID]; ok {
		return ctor(data), nil
	}

	return nil, fmt.Errorf("unsupported codec: 0x%X", codecID)
}

func GprsMessageDecoderFactory(gprsMessagePacket any) (GprsDecoder, error) {
	const codecIdIndex = 8

	var (
		data []byte
		err  error
	)

	if data, err = transformData(gprsMessagePacket); err != nil {
		return nil, fmt.Errorf("failed to transform data: %v", err)
	}

	if len(data) <= codecIdIndex {
		return nil, fmt.Errorf("invalid packet length: %d", len(data))
	}

	codecID := data[codecIdIndex]

	var decoders = map[byte]func([]byte) GprsDecoder{
		constant.Codec12: func(b []byte) GprsDecoder { return codec12.New(b) },
		constant.Codec13: func(b []byte) GprsDecoder { return nil },
		constant.Codec14: func(b []byte) GprsDecoder { return nil },
	}

	if ctor, ok := decoders[codecID]; ok {
		return ctor(data), nil
	}
	return nil, fmt.Errorf("unsupported codec: 0x%X", codecID)
}
