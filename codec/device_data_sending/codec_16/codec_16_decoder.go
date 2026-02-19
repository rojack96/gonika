package codec16

import (
	"encoding/hex"
	"maps"

	"github.com/rojack96/gonika/codec/constant"
	"github.com/rojack96/gonika/codec/device_data_sending/utils"
	"github.com/rojack96/gonika/codec/models"
	"github.com/rojack96/gonika/codec/parsers"
)

type codec16 struct {
	avlDataPacket []byte
	parser        parsers.BaseParser
	builder       *utils.Builders
}

func New(avlDataPacket []byte) *codec16 {
	return &codec16{
		avlDataPacket: avlDataPacket,
		parser:        parsers.NewBaseParser(),
		builder:       utils.NewBuilders(),
	}
}

func (c *codec16) DecodeTCP() *models.AvlDataPacketTCP {
	var result models.AvlDataPacketTCP

	data := utils.DataMapping(c.avlDataPacket)

	result.Preamble = c.parser.Preamble(data.AvlDataPacketHeader.Preamble)
	result.CodecID = c.parser.CodecID(data.AvlDataArray.CodecID)
	result.DataFieldLength = c.parser.Parse4bytes(data.AvlDataPacketHeader.DataFieldLength)
	result.NumberOfData1 = c.parser.NumberOfData(data.AvlDataArray.NumberOfData1)
	result.NumberOfData2 = c.parser.NumberOfData(data.AvlDataArray.NumberOfData2)
	result.Crc16 = c.parser.Crc16(data.Crc16)

	body := data.AvlDataArray.AvlData

	index := 0

	if result.NumberOfData1 != result.NumberOfData2 {
		return nil
	}

	for i := 0; i < int(result.NumberOfData1); i++ {
		avl := models.AvlData16{}

		avl.Timestamp, index = c.parser.Timestamp(index, body)
		avl.Priority, index = c.parser.Priority(index, body)
		avl.GpsElement, index = c.parser.GpsElement(index, body)
		avl.EventIOID, index = c.parseEventIO(index, body)
		avl.GenerationType, index = c.parseGenerationType(index, body)
		avl.NoOfTotalIO, index = c.parseTotalNumberOfIO(index, body)
		avl.NoOfOneByte, avl.OneByteIO, index = c.parseIo(constant.OneByteIo, index, body)
		avl.NoOfTwoByte, avl.TwoByteIO, index = c.parseIo(constant.TwoByteIo, index, body)
		avl.NoOfFourByte, avl.FourByteIO, index = c.parseIo(constant.FourByteIo, index, body)
		avl.NoOfEightByte, avl.EightByteIO, index = c.parseIo(constant.EightByteIo, index, body)

		result.AvlData = append(result.AvlData, avl)
	}

	return &result
}

func (c *codec16) DecodeTCPflat() *models.AvlDataPacketFlat {
	var result models.AvlDataPacketFlat

	data := utils.DataMapping(c.avlDataPacket)

	result.CodecID = c.parser.CodecID(data.AvlDataArray.CodecID)
	result.NumberOfData1 = c.parser.NumberOfData(data.AvlDataArray.NumberOfData1)
	result.NumberOfData2 = c.parser.NumberOfData(data.AvlDataArray.NumberOfData2)

	body := data.AvlDataArray.AvlData

	index := 0

	if result.NumberOfData1 != result.NumberOfData2 {
		return nil
	}

	for i := 0; i < int(result.NumberOfData1); i++ {
		avl := models.AvlData8extFlat{}

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
		_, index = c.parseGenerationType(index, body)
		_, index = c.parseTotalNumberOfIO(index, body)

		var ioData map[uint16]string

		avl.IO = make(map[uint16]string)

		ioByteLen := []int{constant.OneByteIo, constant.TwoByteIo, constant.FourByteIo, constant.EightByteIo}
		for _, t := range ioByteLen {
			_, ioData, index = c.parseIo(t, index, body)
			maps.Copy(avl.IO, ioData)
		}

		result.AvlData = append(result.AvlData, avl)
	}

	return &result
}

func (c *codec16) DecodeUDP() *models.AvlDataPacketUDP {
	var result models.AvlDataPacketUDP

	data := utils.UdpDataMapping(c.avlDataPacket)

	result.Length = c.parser.Parse2bytes(data.UdpChannelHeader.Length)
	result.PacketID = c.parser.Parse2bytes(data.UdpChannelHeader.PacketID)
	result.NotUsableByte = data.UdpChannelHeader.NotUsableByte
	result.AvlPacketID = data.UdpAvlPacketHeader.AvlPacketID
	result.ImeiLength = c.parser.Parse2bytes(data.UdpAvlPacketHeader.ImeiLength)
	result.Imei = hex.EncodeToString(data.UdpAvlPacketHeader.Imei[:])
	result.CodecID = c.parser.CodecID(data.AvlDataArray.CodecID)
	result.NumberOfData1 = c.parser.NumberOfData(data.AvlDataArray.NumberOfData1)
	result.NumberOfData2 = c.parser.NumberOfData(data.AvlDataArray.NumberOfData2)

	body := data.AvlDataArray.AvlData

	index := 0

	if result.NumberOfData1 != result.NumberOfData2 {
		return nil
	}

	for i := 0; i < int(result.NumberOfData1); i++ {
		avl := models.AvlData16{}

		avl.Timestamp, index = c.parser.Timestamp(index, body)
		avl.Priority, index = c.parser.Priority(index, body)
		avl.GpsElement, index = c.parser.GpsElement(index, body)
		avl.EventIOID, index = c.parseEventIO(index, body)
		avl.GenerationType, index = c.parseGenerationType(index, body)
		avl.NoOfTotalIO, index = c.parseTotalNumberOfIO(index, body)
		avl.NoOfOneByte, avl.OneByteIO, index = c.parseIo(constant.OneByteIo, index, body)
		avl.NoOfTwoByte, avl.TwoByteIO, index = c.parseIo(constant.TwoByteIo, index, body)
		avl.NoOfFourByte, avl.FourByteIO, index = c.parseIo(constant.FourByteIo, index, body)
		avl.NoOfEightByte, avl.EightByteIO, index = c.parseIo(constant.EightByteIo, index, body)

		result.AvlData = append(result.AvlData, avl)
	}

	return &result
}

func (c *codec16) DecodeUDPflat() *models.AvlDataPacketFlat {
	var result models.AvlDataPacketFlat

	data := utils.UdpDataMapping(c.avlDataPacket)

	result.CodecID = c.parser.CodecID(data.AvlDataArray.CodecID)
	result.NumberOfData1 = c.parser.NumberOfData(data.AvlDataArray.NumberOfData1)
	result.NumberOfData2 = c.parser.NumberOfData(data.AvlDataArray.NumberOfData2)

	body := data.AvlDataArray.AvlData

	index := 0

	if result.NumberOfData1 != result.NumberOfData2 {
		return nil
	}

	for i := 0; i < int(result.NumberOfData1); i++ {
		avl := models.AvlData16Flat{}

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
		_, index = c.parseGenerationType(index, body)
		_, index = c.parseTotalNumberOfIO(index, body)

		var ioData map[uint16]string

		avl.IO = make(map[uint16]string)

		ioByteLen := []int{constant.OneByteIo, constant.TwoByteIo, constant.FourByteIo, constant.EightByteIo}
		for _, t := range ioByteLen {
			_, ioData, index = c.parseIo(t, index, body)
			maps.Copy(avl.IO, ioData)
		}

		result.AvlData = append(result.AvlData, avl)
	}

	return &result
}
