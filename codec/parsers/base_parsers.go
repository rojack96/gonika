package parsers

import (
	"encoding/binary"

	"github.com/rojack96/gonika/codec/models"
)

type BaseParser struct{}

func NewBaseParser() *BaseParser {
	return &BaseParser{}
}

func (bp *BaseParser) Parse2bytes(data [2]byte) uint16 {
	return binary.BigEndian.Uint16(data[:])
}

func (bp *BaseParser) Parse4bytes(data [4]byte) uint32 {
	return binary.BigEndian.Uint32(data[:])
}

// Preamble This function parse the preamble from AVL data.
func (bp *BaseParser) Preamble(data [4]byte) models.Preamble {
	preamble := bp.Parse4bytes(data)
	return models.Preamble(preamble)
}

// CodecID This function parse the codec id from AVL data.
func (bp *BaseParser) CodecID(data byte) models.CodecID {
	return models.CodecID(data)
}

// NumberOfData This function parse the number of data from AVL data.
func (bp *BaseParser) NumberOfData(data byte) models.NumberOfData {
	return models.NumberOfData(data)
}

// Quantity This function parse the number of data from AVL data.
func (bp *BaseParser) Quantity(data byte) models.Quantity {
	return models.Quantity(data)
}

// Type This function parse the type from AVL data.
func (bp *BaseParser) Type(data byte) models.Type {
	return models.Type(data)
}

// Crc16 This function parse the crc16 from AVL data.
func (bp *BaseParser) Crc16(data [4]byte) models.Crc16 {
	crc16 := bp.Parse4bytes(data)
	return models.Crc16(crc16)
}

// Timestamp This function parse the timestamp from AVL data.
func (bp *BaseParser) Timestamp(startIndex int, body []byte) (models.Timestamp, int) {
	const timestampLength = 8

	endIndex := startIndex + timestampLength
	ts := body[startIndex:endIndex]

	unixTs := binary.BigEndian.Uint64(ts)

	return models.Timestamp(unixTs), endIndex
}

// Priority This function parse the priority from AVL data.
func (bp *BaseParser) Priority(index int, body []byte) (models.Priority, int) {
	const priorityLength = 1

	priority := body[index]

	return models.Priority(priority), index + priorityLength
}

// GpsElement This function parse the GPS data from AVL data.
func (bp *BaseParser) GpsElement(startIndex int, body []byte) (models.GpsElement, int) {
	const gpsElementLength = 15

	endIndex := startIndex + gpsElementLength
	data := body[startIndex:endIndex]
	var gps models.GpsElement

	gps.Longitude = models.Longitude(decodeCoordinate(data[0:4]))
	gps.Latitude = models.Latitude(decodeCoordinate(data[4:8]))
	gps.Altitude = models.Altitude(bp.Parse2bytes([2]byte{data[8], data[9]}))
	gps.Angle = models.Angle(bp.Parse2bytes([2]byte{data[10], data[11]}))
	gps.Satellites = models.Satellites(data[12])
	gps.Speed = models.Speed(bp.Parse2bytes([2]byte{data[13], data[14]}))

	return gps, endIndex
}

// decodeCoordinate This function return a position negative.
//
// To determine if the coordinate is negative,
// convert it to binary format and check the very first bit.
// If it is 0, coordinate is positive,
// if it is 1, coordinate is negative.
func decodeCoordinate(position []byte) float64 {
	raw := int32(binary.BigEndian.Uint32(position))
	return float64(raw) / 1e7
}
