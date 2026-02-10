package codec

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/getrak/crc16"
	"github.com/rojack96/gonika/codec/constant"
)

// ImeiChecker This function checks if the data buffer contains a valid IMEI number and returns it if found.
func ImeiChecker(dataBuffer any) ([]byte, bool, error) {
	var (
		data []byte
		err  error
	)

	if data, err = transformData(dataBuffer); err != nil {
		return nil, false, fmt.Errorf("failed to transform data: %v", err)
	}

	if len(data) == 0 {
		return nil, false, nil
	}

	if bytes.HasPrefix(data, []byte(constant.ImeiPrefix)) && len(data) == 17 {
		return data[2:], true, nil
	}

	return nil, false, nil
}

func Crc16Checker(dataBuffer any) error {
	var (
		data []byte
		err  error
	)

	if data, err = transformData(dataBuffer); err != nil {
		return fmt.Errorf("failed to transform data: %v", err)
	}

	if len(data) < 4 {
		return fmt.Errorf("invalid packet length: %d", len(data))
	}

	crc := data[len(data)-2:] //take last 2 bytes as CRC16
	crcProvided := binary.BigEndian.Uint16(crc)

	sourceCrc := data[8 : len(data)-4]
	crcTable := crc16.MakeTable(crc16.CRC16_ARC)
	crcCalculated := crc16.Checksum(sourceCrc, crcTable)

	if crcProvided != crcCalculated {
		return fmt.Errorf("CRC16 mismatch: provided 0x%X, calculated 0x%X", crcProvided, crcCalculated)
	}

	return nil
}

// CodecRecognizer This function identifies the codec type from the AVL data packet and returns the codec ID if recognized.
func CodecRecognizer(avlDataPacket any) (byte, error) {
	var (
		data []byte
		err  error
	)

	if data, err = transformData(avlDataPacket); err != nil {
		return 0x00, fmt.Errorf("failed to transform data: %v", err)
	}

	if len(data) <= 8 {
		return 0x00, fmt.Errorf("invalid packet length: %d", len(data))
	}

	codecID := data[8]

	switch codecID {
	case constant.Codec8, constant.Codec8ext, constant.Codec16, constant.Codec12, constant.Codec13, constant.Codec14:
		return codecID, nil
	default:
		return 0x00, fmt.Errorf("unsupported codec: 0x%X", codecID)
	}

}
