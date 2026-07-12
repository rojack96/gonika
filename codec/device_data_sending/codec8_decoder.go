package dds

import (
	"encoding/hex"
	"fmt"
	"maps"

	"github.com/rojack96/gonika/codec/constant"
	"github.com/rojack96/gonika/codec/device_data_sending/helpers"
	utils "github.com/rojack96/gonika/codec/device_data_sending/helpers"
	"github.com/rojack96/gonika/codec/models"
	"github.com/rojack96/gonika/codec/parsers"
)

type codec8 struct {
	parser  parsers.BaseParser
	builder *utils.Builders
}

func NewCodec8Decoder() *codec8 {
	return &codec8{
		parser:  parsers.NewBaseParser(),
		builder: utils.NewBuilders(),
	}
}

// ----------------------------------------
// TCP Decoder for Codec 8
// ----------------------------------------

// DecodeTCP - This function decodes the AVL data packet and returns all information based on Teltonika documentation.
func (c *codec8) DecodeTCP(avlDataPacket any) (*models.AvlDataPacketTCP, error) {
	var result models.AvlDataPacketTCP

	transformedData, err := helpers.TransformData(avlDataPacket)
	if err != nil {
		return nil, fmt.Errorf("failed to transform data: %v", err)
	}

	data := utils.DataMapping(transformedData)

	result.Preamble = c.parser.Preamble(data.AvlDataPacketHeader.Preamble)
	result.DataFieldLength = c.parser.Parse4bytes(data.AvlDataPacketHeader.DataFieldLength)
	result.CodecID = c.parser.CodecID(data.AvlDataArray.CodecID)
	result.NumberOfData1 = c.parser.NumberOfData(data.AvlDataArray.NumberOfData1)
	result.NumberOfData2 = c.parser.NumberOfData(data.AvlDataArray.NumberOfData2)
	result.Crc16 = c.parser.Crc16(data.Crc16)

	body := data.AvlDataArray.AvlData

	index := 0

	if result.NumberOfData1 != result.NumberOfData2 {
		return nil, fmt.Errorf("inconsistent data count")
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

	return &result, nil
}

// DecodeTCPflat - This function decodes the AVL data packet and returns a flat structured representation.
func (c *codec8) DecodeTCPflat(avlDataPacket any) (*models.AvlDataPacketFlat, error) {
	var result models.AvlDataPacketFlat

	transformedData, err := helpers.TransformData(avlDataPacket)
	if err != nil {
		return nil, fmt.Errorf("failed to transform data: %v", err)
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
// UDP Decoder for Codec 8
// ----------------------------------------

// DecodeUDP - This function decodes an AVL packet received over UDP.
func (c *codec8) DecodeUDP(avlDataPacket any) (*models.AvlDataPacketUDP, error) {
	var result models.AvlDataPacketUDP

	transformedData, err := helpers.TransformData(avlDataPacket)
	if err != nil {
		return nil, fmt.Errorf("failed to transform data: %v", err)
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

	return &result, nil
}

// DecodeUDPflat decodes an AVL packet received over UDP into a flat representation.
func (c *codec8) DecodeUDPflat(avlDataPacket any) (*models.AvlDataPacketFlat, error) {
	var result models.AvlDataPacketFlat

	transformedData, err := helpers.TransformData(avlDataPacket)
	if err != nil {
		return nil, fmt.Errorf("failed to transform data: %v", err)
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
// Parsers for Codec 8
// ----------------------------------------

// ParseEventIO  This function parse the Event IO data from AVL data.
//
// Event IO changed
// if data is acquired on event
// this field defines which IO property has changed and generated an event.
func (c *codec8) parseEventIO(startIndex int, body []byte) (uint8, int) {
	eventIOID := body[startIndex]

	return eventIOID, startIndex + 1
}

// parseTotalNumberOfIO Total Number of IO.
//
// A total number of properties coming with record (N = N1 + N2 + N4 + N8).
func (c *codec8) parseTotalNumberOfIO(startIndex int, body []byte) (uint8, int) {
	noOfTotalIO := body[startIndex]

	return noOfTotalIO, startIndex + 1
}

// parseIo This function parse byte IO.
func (c *codec8) parseIo(valueLength int, startIndex int, body []byte) (uint8, map[uint8]string, int) {
	const idLength = 1
	splitByte := idLength + valueLength
	result := map[uint8]string{}

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
		id := data[i]
		value := hex.EncodeToString(data[i+1 : i+splitByte])

		result[id] = value
	}

	return nOfIOelements, result, IOelementsEndIndex
}
