package codec8

import (
	"maps"

	"github.com/rojack96/gonika/codec/constant"
	"github.com/rojack96/gonika/codec/device_data_sending/utils"
	"github.com/rojack96/gonika/codec/models"
	"github.com/rojack96/gonika/codec/parsers"
)

type codec8 struct {
	avlDataPacket []byte
	parser        parsers.BaseParser
}

func New(avlDataPacket []byte) *codec8 {
	return &codec8{
		avlDataPacket: avlDataPacket,
		parser:        parsers.NewBaseParser(),
	}
}

func (c *codec8) Decode() *models.AvlDataPacket {
	var result models.AvlDataPacket

	data := utils.DataMapping(c.avlDataPacket)

	result.Preamble = c.parser.Preamble(data.AvlDataPacketHeader.Preamble)
	result.DataFieldLength = c.parser.DataFieldLength(data.AvlDataPacketHeader.DataFieldLength)
	result.CodecID = c.parser.CodecId(data.AvlDataArray.CodecID)
	result.NumberOfData1 = c.parser.NumberOfData(data.AvlDataArray.NumberOfData1)
	result.NumberOfData2 = c.parser.NumberOfData(data.AvlDataArray.NumberOfData2)
	result.Crc16 = c.parser.Crc16(data.Crc16)

	body := data.AvlDataArray.AvlData

	index := 0

	if result.NumberOfData1 != result.NumberOfData2 {
		return nil
	}

	for i := 0; i < int(result.NumberOfData1); i++ {
		avl := models.AvlData8{}

		avl.Timestamp, index = c.parser.Timestamp(index, body)
		avl.Priority, index = c.parser.Priority(index, body)
		avl.GpsElement, index = c.parser.GpsElement(index, body)
		avl.EventIOID, index = c.parseEventIO(index, body)
		avl.NoOfTotalIO, index = c.parseTotalNumberOfIO(index, body)
		avl.NoOfOneByte, avl.OneByteIO, index = c.parseIo(constant.OneByteIo, index, body)
		avl.NoOfTwoByte, avl.TwoByteIO, index = c.parseIo(constant.TwoByteIo, index, body)
		avl.NoOfFourByte, avl.FourByteIO, index = c.parseIo(constant.FourByteIo, index, body)
		avl.NoOfEightByte, avl.EightByteIO, index = c.parseIo(constant.EightByteIo, index, body)

		result.AvlData = append(result.AvlData, avl)
	}

	return &result
}

func (c *codec8) DecodeFlat() *models.AvlDataPacketFlat {
	var result models.AvlDataPacketFlat

	data := utils.DataMapping(c.avlDataPacket)

	result.Preamble = c.parser.Preamble(data.AvlDataPacketHeader.Preamble)
	result.DataFieldLength = c.parser.DataFieldLength(data.AvlDataPacketHeader.DataFieldLength)
	result.CodecID = c.parser.CodecId(data.AvlDataArray.CodecID)
	result.NumberOfData1 = c.parser.NumberOfData(data.AvlDataArray.NumberOfData1)
	result.NumberOfData2 = c.parser.NumberOfData(data.AvlDataArray.NumberOfData2)
	result.Crc16 = c.parser.Crc16(data.Crc16)

	body := data.AvlDataArray.AvlData

	index := 0

	if result.NumberOfData1 != result.NumberOfData2 {
		return nil
	}

	for i := 0; i < int(result.NumberOfData1); i++ {
		avl := models.AvlData8Flat{}

		avl.Timestamp, index = c.parser.Timestamp(index, body)
		avl.Priority, index = c.parser.Priority(index, body)

		var gps models.GpsElement
		gps, index = c.parser.GpsElement(index, body)
		avl.Latitude = gps.Latitude
		avl.Longitude = gps.Longitude
		avl.Altitude = gps.Altitude
		avl.Angle = gps.Angle
		avl.Satellites = gps.Satellites
		avl.Speed = gps.Speed

		_, index = c.parseEventIO(index, body)
		_, index = c.parseTotalNumberOfIO(index, body)

		var ioData map[uint8]string

		avl.IO = make(map[uint8]string)

		_, ioData, index = c.parseIo(constant.OneByteIo, index, body)
		maps.Copy(avl.IO, ioData)
		_, ioData, index = c.parseIo(constant.TwoByteIo, index, body)
		maps.Copy(avl.IO, ioData)
		_, ioData, index = c.parseIo(constant.FourByteIo, index, body)
		maps.Copy(avl.IO, ioData)
		_, ioData, index = c.parseIo(constant.EightByteIo, index, body)
		maps.Copy(avl.IO, ioData)

		result.AvlData = append(result.AvlData, avl)
	}

	return &result
}
