package codec8

import (
	"errors"
	"fmt"

	"github.com/rojack96/gonika/codec/constant"
	"github.com/rojack96/gonika/codec/device_data_sending/models"
	"github.com/rojack96/gonika/codec/device_data_sending/utils"
	m "github.com/rojack96/gonika/codec/models"
)

func (c *codec8) EncoderTCP(numberOfData uint8, gpsList []m.GpsElementEncoder, io m.Codec8Encoder) ([]byte, error) {
	var (
		packet     models.AvlDataPacketByte
		gpsElement [15]byte = [15]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	)

	if len(gpsList) < int(numberOfData) {
		return nil, errors.New("the number of data should be greather then or equals to numer of gps element and viceversa")
	}

	packet.AvlDataPacketHeader.Preamble = [4]byte{0x00, 0x00, 0x00, 0x00}
	packet.AvlDataArray.CodecID = constant.Codec8
	packet.AvlDataArray.NumberOfData1 = numberOfData
	packet.AvlDataArray.NumberOfData2 = numberOfData

	for n := range numberOfData {
		fmt.Println(n)
		timestamp := c.builder.Timestamp()
		priority := c.builder.Priority()
		packet.AvlDataArray.AvlData = append(packet.AvlDataArray.AvlData, timestamp[:]...)
		packet.AvlDataArray.AvlData = append(packet.AvlDataArray.AvlData, priority)
		// TODO take gps from list
		packet.AvlDataArray.AvlData = append(packet.AvlDataArray.AvlData, gpsElement[:]...)
		eventIo := c.builder.EventIo1Byte()
		packet.AvlDataArray.AvlData = append(packet.AvlDataArray.AvlData, eventIo)
		packet.AvlDataArray.AvlData = append(packet.AvlDataArray.AvlData, ioData(*c.builder, io)...)
	}

	packet.AvlDataPacketHeader.DataFieldLength = c.builder.DataFieldLength(packet)

	result := c.builder.MergeData(packet)
	packet.Crc16 = c.builder.Crc16Builder(result)

	result = append(result, packet.Crc16[:]...)

	return result, nil
}

func ioData(b utils.Builders, io m.Codec8Encoder) []byte {
	result := make([]byte, 0)

	nOfOneByte := uint8(len(io.OneByte))
	nOfTwoByte := uint8(len(io.TwoByte))
	nOfFourByte := uint8(len(io.FourByte))
	nOfEightByte := uint8(len(io.EightByte))
	nOfTotalIo := nOfOneByte + nOfTwoByte + nOfFourByte + nOfEightByte

	result = append(result, nOfTotalIo)

	result = append(result, nOfOneByte)
	for k, v := range io.OneByte {
		result = append(result, k)
		result = append(result, v)
	}

	result = append(result, nOfTwoByte)
	for k, v := range io.TwoByte {
		result = append(result, k)
		value := b.Uint16ToBytes(v)
		result = append(result, value[:]...)
	}

	result = append(result, nOfFourByte)
	for k, v := range io.FourByte {
		result = append(result, k)
		value := b.Uint32ToBytes(v)
		result = append(result, value[:]...)
	}

	result = append(result, nOfEightByte)
	for k, v := range io.EightByte {
		result = append(result, k)
		value := b.Uint64ToBytes(v)
		result = append(result, value[:]...)
	}

	return result
}
