package dds

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"maps"

	"github.com/rojack96/gonika/codec/constant"
	utils "github.com/rojack96/gonika/codec/device_data_sending/helpers"
	"github.com/rojack96/gonika/codec/helpers"
	"github.com/rojack96/gonika/codec/models"
	"github.com/rojack96/gonika/codec/parsers"
)

type codec16 struct {
	parser  *parsers.BaseParser
	builder *utils.Builders
}

func NewCodec16Decoder() *codec16 {
	return &codec16{
		parser:  parsers.NewBaseParser(),
		builder: utils.NewBuilders(),
	}
}

// ----------------------------------------
// TCP Decoder for Codec 16
// ----------------------------------------

func (c *codec16) DecodeTCP(avlDataPacket any) (*models.AvlDataPacketTCP, error) {
	var result models.AvlDataPacketTCP

	transformedData, err := helpers.TransformData(avlDataPacket)
	if err != nil {
		return nil, err
	}

	data := utils.DataMapping(transformedData)

	result.Preamble = c.parser.Preamble(data.AvlDataPacketHeader.Preamble)
	result.CodecID = c.parser.CodecID(data.AvlDataArray.CodecID)
	result.DataFieldLength = c.parser.Parse4bytes(data.AvlDataPacketHeader.DataFieldLength)
	result.NumberOfData1 = c.parser.NumberOfData(data.AvlDataArray.NumberOfData1)
	result.NumberOfData2 = c.parser.NumberOfData(data.AvlDataArray.NumberOfData2)
	result.Crc16 = c.parser.Crc16(data.Crc16)

	body := data.AvlDataArray.AvlData

	index := 0

	if result.NumberOfData1 != result.NumberOfData2 {
		return nil, fmt.Errorf("inconsistent data count")
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

	return &result, nil
}

func (c *codec16) DecodeTCPflat(avlDataPacket any) (*models.AvlDataPacketFlat, error) {
	var result models.AvlDataPacketFlat

	transformedData, err := helpers.TransformData(avlDataPacket)
	if err != nil {
		return nil, err
	}

	data := utils.DataMapping(transformedData)

	result.CodecID = c.parser.CodecID(data.AvlDataArray.CodecID)
	result.NumberOfData1 = c.parser.NumberOfData(data.AvlDataArray.NumberOfData1)
	result.NumberOfData2 = c.parser.NumberOfData(data.AvlDataArray.NumberOfData2)

	body := data.AvlDataArray.AvlData

	index := 0

	if result.NumberOfData1 != result.NumberOfData2 {
		return nil, fmt.Errorf("inconsistent data count")
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

	return &result, nil
}

// ----------------------------------------
// UDP Decoder for Codec 16
// ----------------------------------------

func (c *codec16) DecodeUDP(avlDataPacket any) (*models.AvlDataPacketUDP, error) {
	var result models.AvlDataPacketUDP

	transformedData, err := helpers.TransformData(avlDataPacket)
	if err != nil {
		return nil, err
	}

	data := utils.UdpDataMapping(transformedData)

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
		return nil, fmt.Errorf("inconsistent data count")
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

	return &result, nil
}

func (c *codec16) DecodeUDPflat(avlDataPacket any) (*models.AvlDataPacketFlat, error) {
	var result models.AvlDataPacketFlat

	transformedData, err := helpers.TransformData(avlDataPacket)
	if err != nil {
		return nil, err
	}

	data := utils.UdpDataMapping(transformedData)

	result.CodecID = c.parser.CodecID(data.AvlDataArray.CodecID)
	result.NumberOfData1 = c.parser.NumberOfData(data.AvlDataArray.NumberOfData1)
	result.NumberOfData2 = c.parser.NumberOfData(data.AvlDataArray.NumberOfData2)

	body := data.AvlDataArray.AvlData

	index := 0

	if result.NumberOfData1 != result.NumberOfData2 {
		return nil, fmt.Errorf("inconsistent data count")
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

	return &result, nil
}

// ----------------------------------------
// Parsers for Codec 16
// ----------------------------------------

// This function parse the Event IO data from AVL data.
//
// Event IO changed
// if data is acquired on event
// this field defines which IO property has changed and generated an event.
func (c *codec16) parseEventIO(startIndex int, body []byte) (uint16, int) {
	endIndex := startIndex + 2
	eventIOID := binary.BigEndian.Uint16(body[startIndex:endIndex])

	return eventIOID, endIndex
}

func (c *codec16) parseGenerationType(startIndex int, body []byte) (uint8, int) {
	generationType := body[startIndex]

	return generationType, startIndex + 1
}

// Total Number of IO.
//
// A total number of properties coming with record (N = N1 + N2 + N4 + N8).
func (c *codec16) parseTotalNumberOfIO(startIndex int, body []byte) (uint8, int) {
	noOfTotalIO := body[startIndex]

	return noOfTotalIO, startIndex + 1
}

// parseIo This function parse byte IO.
func (c *codec16) parseIo(valueLength int, startIndex int, body []byte) (uint8, map[uint16]string, int) {
	const idLength = 2
	splitByte := idLength + valueLength
	result := map[uint16]string{}

	if startIndex >= len(body) {
		return 0, result, len(body)
	}

	// One Byte IO Number
	nOfIOelements := body[startIndex]
	// One Byte IO Data
	IOelementsStartIndex := startIndex + 1
	IOelementsEndIndex := IOelementsStartIndex + int(nOfIOelements)*splitByte
	data := body[IOelementsStartIndex:IOelementsEndIndex]

	for i := 0; i < len(data); i += splitByte {
		id := binary.BigEndian.Uint16(data[i : i+2])
		value := hex.EncodeToString(data[i+2 : i+splitByte])

		result[id] = value
	}

	return nOfIOelements, result, IOelementsEndIndex
}
