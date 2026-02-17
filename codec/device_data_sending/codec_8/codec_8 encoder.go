package codec8

import (
	"errors"
	"fmt"

	"github.com/rojack96/gonika/codec/constant"
	"github.com/rojack96/gonika/codec/device_data_sending/models"
)

func (c *codec8) EncoderTCP(numberOfData uint8, gpsList []string) (*models.AvlDataPacketByte, error) {
	var (
		result     models.AvlDataPacketByte
		gpsElement [15]byte = [15]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	)

	if len(gpsList) < int(numberOfData) {
		return nil, errors.New("the number of data should be greather then or equals to numer of gps element and viceversa")
	}

	result.AvlDataPacketHeader.Preamble = [4]byte{0x00, 0x00, 0x00, 0x00}
	result.AvlDataArray.CodecID = constant.Codec8
	result.AvlDataArray.NumberOfData1 = numberOfData
	result.AvlDataArray.NumberOfData2 = numberOfData

	for n := range numberOfData {
		fmt.Println(n)
		timestamp := c.builder.Timestamp()
		priority := c.builder.Priority()
		result.AvlDataArray.AvlData = append(result.AvlDataArray.AvlData, timestamp[:]...)
		result.AvlDataArray.AvlData = append(result.AvlDataArray.AvlData, priority)
		result.AvlDataArray.AvlData = append(result.AvlDataArray.AvlData, gpsElement[:]...)
		eventIo := c.builder.EventIo1Byte()
		result.AvlDataArray.AvlData = append(result.AvlDataArray.AvlData, eventIo)
		// TODO qui fare prima tutta la parte id value e  poi fare total io

		//totalIo := 0
	}

	result.AvlDataPacketHeader.DataFieldLength = c.builder.DataFieldLength(result)

	crcSource := c.builder.MergeData(result)
	result.Crc16 = c.builder.Crc16Builder(crcSource)

	return &result, nil
}
