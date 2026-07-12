package codec

import (
	"encoding/hex"
	"fmt"

	"github.com/rojack96/gonika/codec/constant"
	dds "github.com/rojack96/gonika/codec/device_data_sending"
	gprs "github.com/rojack96/gonika/codec/gprs_message"
	"github.com/rojack96/gonika/codec/models"
)

type AvlEncoder interface {
	EncodeTCP(avlDataArray []models.AvlDataArrayEncoder) ([]byte, error)
	EncodeUDP(imei string, avlDataArray []models.AvlDataArrayEncoder) ([]byte, error)
}

type GprsEncoder interface {
	Encode() []byte
}

func NewDataSendingEncoder(codecID byte) (AvlEncoder, error) {
	switch codecID {
	case constant.Codec12:
		return dds.NewCodec8Encoder(), nil
	case constant.Codec13:
		return dds.NewCodec8Encoder(), nil
	case constant.Codec14:
		return dds.NewCodec16Encoder(), nil
	default:
		return nil, fmt.Errorf("unsupported codec: 0x%X", codecID)
	}
}

func NewGprsMessageEncoder(codecID byte) (GprsEncoder, error) {
	switch codecID {
	case constant.Codec12:
		return gprs.NewCodec12Encoder(), nil
	case constant.Codec13:
		return gprs.NewCodec13Encoder(), nil
	case constant.Codec14:
		return gprs.NewCodec14Encoder(), nil
	default:
		return nil, fmt.Errorf("unsupported codec: 0x%X", codecID)
	}
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

// Deprecated: Use NewDataSendingEncoder instead.
func DeviceDataSendingEncoderFactory(codecID byte) (AvlEncoder, error)

// Deprecated: Use NewGprsMessageEncoder instead.
func GprsMessageEncoderFactory(codecID byte) (GprsEncoder, error)
