package codec

import (
	"fmt"

	"github.com/rojack96/gonika/codec/constant"
	codec16 "github.com/rojack96/gonika/codec/device_data_sending/codec_16"
	codec8 "github.com/rojack96/gonika/codec/device_data_sending/codec_8"
	codec8ext "github.com/rojack96/gonika/codec/device_data_sending/codec_8e"
	codec12 "github.com/rojack96/gonika/codec/gprs_message/codec_12"
	codec13 "github.com/rojack96/gonika/codec/gprs_message/codec_13"
	codec14 "github.com/rojack96/gonika/codec/gprs_message/codec_14"
	"github.com/rojack96/gonika/codec/models"
)

// Decoder is the main struct responsible for decoding AVL data packets and GPRS messages.

type AvlDecoder interface {
	// DecodeTCP decodes the AVL data packet and returns all information based on Teltonika documentation.
	DecodeTCP() *models.AvlDataPacketTCP
	// DecodeTCPflat decodes the AVL data packet and returns a flat structured representation.
	DecodeTCPflat() *models.AvlDataPacketFlat
	DecodeUDP() *models.AvlDataPacketUDP
	DecodeUDPflat() *models.AvlDataPacketFlat
}

type GprsDecoder interface {
	DecodeCommand() (*models.CommandMessage, error)
	DecodeResponse() *models.ResponseMessage
}

func DeviceDataSendingDecoderFactory(avlDataPacket any) (AvlDecoder, error) {
	const codecIDIndex = 8

	var (
		data []byte
		err  error
	)

	if data, err = transformData(avlDataPacket); err != nil {
		return nil, fmt.Errorf("failed to transform data: %v", err)
	}

	if len(data) <= codecIDIndex {
		return nil, fmt.Errorf("invalid packet length: %d", len(data))
	}

	codecID := data[codecIDIndex]

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
	const codecIDIndex = 8

	var (
		data []byte
		err  error
	)

	if data, err = transformData(gprsMessagePacket); err != nil {
		return nil, fmt.Errorf("failed to transform data: %v", err)
	}

	if len(data) <= codecIDIndex {
		return nil, fmt.Errorf("invalid packet length: %d", len(data))
	}

	codecID := data[codecIDIndex]

	var decoders = map[byte]func([]byte) GprsDecoder{
		constant.Codec12: func(b []byte) GprsDecoder { return codec12.New(b) },
		constant.Codec13: func(b []byte) GprsDecoder { return codec13.New(b) },
		constant.Codec14: func(b []byte) GprsDecoder { return codec14.New(b) },
	}

	if ctor, ok := decoders[codecID]; ok {
		return ctor(data), nil
	}
	return nil, fmt.Errorf("unsupported codec: 0x%X", codecID)
}
