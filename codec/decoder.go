package codec

import (
	"fmt"

	"github.com/rojack96/gonika/codec/constant"
	dds "github.com/rojack96/gonika/codec/device_data_sending"
	codec12 "github.com/rojack96/gonika/codec/gprs_message/codec_12"
	codec13 "github.com/rojack96/gonika/codec/gprs_message/codec_13"
	codec14 "github.com/rojack96/gonika/codec/gprs_message/codec_14"
	"github.com/rojack96/gonika/codec/models"
)

// Decoder is the main struct responsible for decoding AVL data packets and GPRS messages.

type AvlDecoder interface {
	// DecodeTCP decodes the AVL data packet and returns all information based on Teltonika documentation.
	//
	// Example JSON result for Codec 8:
	// {
	//	 "preamble": 0,
	//	 "dataFieldLength": 35,
	//	 "codecID": 8,
	//	 "numberOfData1": 1,
	//	 "avlData": [
	//		 {
	//			 "timestamp": 1710000000000,
	//			 "priority": 0,
	//			 "gpsElement": {
	//				 "longitude": 12.4964,
	//				 "latitude": 41.9028,
	//				 "altitude": 21,
	//				 "angle": 180,
	//				 "satellites": 9,
	//				 "speed": 42
	//			 },
	//			 "eventIoID": 1,
	//			 "numberOfTotalIO": 1,
	//			 "numberOfOneByte": 1,
	//			 "oneByteIO": {"1": "01"},
	//			 "numberOfTwoByte": 0,
	//			 "twoByteIO": {},
	//			 "numberOfFourByte": 0,
	//			 "fourByteIO": {},
	//			 "numberOfEightByte": 0,
	//			 "eightByteIO": {}
	//		 }
	//	 ],
	//	 "numberOfData2": 1,
	//	 "crc16": 39184
	// }
	DecodeTCP(avlDataPacket any) (*models.AvlDataPacketTCP, error)

	// DecodeTCPflat decodes the AVL data packet and returns a flat structured representation.
	//
	// Example JSON result for Codec 8:
	// {
	//	 "codecID": 8,
	//	 "numberOfData1": 1,
	//	 "avlData": [
	//		 {
	//			 "timestamp": 1710000000000,
	//			 "priority": 0,
	//			 "longitude": 12.4964,
	//			 "latitude": 41.9028,
	//			 "altitude": 21,
	//			 "angle": 180,
	//			 "satellites": 9,
	//			 "speed": 42,
	//			 "ioData": {"1": "01"}
	//		 }
	//	 ],
	//	 "numberOfData2": 1
	// }
	DecodeTCPflat(avlDataPacket any) (*models.AvlDataPacketFlat, error)

	// DecodeUDP decodes an AVL packet received over UDP.
	//
	// Example JSON result:
	// {
	//	 "length": 45,
	//	 "packetID": 1,
	//	 "notUsableByte": 1,
	//	 "avlPacketID": 1,
	//	 "imeiLength": 15,
	//	 "imei": "123456789012345",
	//	 "codecID": 8,
	//	 "numberOfData1": 0,
	//	 "avlData": [],
	//	 "numberOfData2": 0
	// }
	DecodeUDP(avlDataPacket any) (*models.AvlDataPacketUDP, error)

	// DecodeUDPflat decodes an AVL packet received over UDP into a flat representation.
	//
	// Example JSON result:
	// {
	//	 "codecID": 8,
	//	 "numberOfData1": 0,
	//	 "avlData": [],
	//	 "numberOfData2": 0
	// }(
	DecodeUDPflat(avlDataPacket any) (*models.AvlDataPacketFlat, error)
}

type GprsDecoder interface {
	// DecodeCommand decodes a Codec 12 GPRS command.
	//
	// Example JSON result:
	// {
	//	 "preamble": 0,
	//	 "dataSize": 15,
	//	 "codecID": 12,
	//	 "commandQuantity1": 1,
	//	 "type": 5,
	//	 "commandSize": 7,
	//	 "CodecSpecificParam": null,
	//	 "command": "getinfo",
	//	 "commandQuantity2": 1,
	//	 "crc16": 43137
	// }
	// Codec 13 and Codec 14 return an unsupported-operation error.
	DecodeCommand() (*models.CommandMessage, error)

	// DecodeResponse decodes a Codec 12, Codec 13, or Codec 14 GPRS response.
	//
	// Example JSON result for Codec 12:
	// {
	//	 "preamble": 0,
	//	 "dataSize": 19,
	//	 "codecID": 12,
	//	 "responseQuantity1": 1,
	//	 "type": 6,
	//	 "responseSize": 11,
	//	 "response": "Hello World",
	//	 "responseQuantity2": 1,
	//	 "crc16": 49926
	// }
	// Codec 13 additionally includes "timestamp"; Codec 14 includes "imei".
	DecodeResponse() *models.ResponseMessage
}

// Deprecated: Use NewDecoder instead.
func DeviceDataSendingDecoderFactory(avlDataPacket any) (AvlDecoder, error)

func NewDecoder(codecID byte) (AvlDecoder, error) {
	switch codecID {
	case constant.Codec8:
		return dds.NewCodec8Decoder(), nil
	case constant.Codec8ext:
		return dds.NewCodec8ExtDecoder(), nil
	case constant.Codec16:
		return dds.NewCodec16Decoder(), nil
	default:
		return nil, fmt.Errorf("unsupported codec: 0x%X", codecID)
	}
}

// Deprecated: Use GprsMessageDecoder instead.
func GprsMessageDecoderFactory(gprsMessagePacket any) (GprsDecoder, error)

func GprsMessageDecoder(gprsMessagePacket any) (GprsDecoder, error) {
	const codecIDIndex = 8

	data, err := transformData(gprsMessagePacket)
	if err != nil {
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
