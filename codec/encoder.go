package codec

import (
	"encoding/hex"
	"fmt"

	"github.com/rojack96/gonika/codec/constant"
	codec16 "github.com/rojack96/gonika/codec/device_data_sending/codec_16"
	codec8 "github.com/rojack96/gonika/codec/device_data_sending/codec_8"
	codec8ext "github.com/rojack96/gonika/codec/device_data_sending/codec_8ext"
	codec12 "github.com/rojack96/gonika/codec/gprs_message/codec_12"
	codec13 "github.com/rojack96/gonika/codec/gprs_message/codec_13"
	codec14 "github.com/rojack96/gonika/codec/gprs_message/codec_14"
	m "github.com/rojack96/gonika/codec/models"
)

type AvlEncoder interface {
	EncodeTCP(avlDataArray []m.AvlDataArrayEncoder) ([]byte, error)
	EncodeUDP(imei string, avlDataArray []m.AvlDataArrayEncoder) ([]byte, error)
}

type GprsEncoder interface {
	Encode() []byte
}

func DeviceDataSendingEncoderFactory(codecID byte) (AvlEncoder, error) {
	switch codecID {
	case constant.Codec8:
		return codec8.NewEncoder(), nil
	case constant.Codec8ext:
		return codec8ext.NewEncoder(), nil
	case constant.Codec16:
		return codec16.NewEncoder(), nil
	default:
		return nil, fmt.Errorf("unsupported codec: 0x%X", codecID)
	}
}

func GprsMessageEncoderFactory(codecID byte) (GprsEncoder, error) {
	switch codecID {
	case constant.Codec12:
		return codec12.NewEncoder(), nil
	case constant.Codec13:
		return codec13.NewEncoder(), nil
	case constant.Codec14:
		return codec14.NewEncoder(), nil
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
