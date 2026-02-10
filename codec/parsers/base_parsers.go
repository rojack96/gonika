package parsers

import (
	"encoding/binary"

	"github.com/rojack96/gonika/codec/models"
)

// Preamble This function parse the preamble from AVL data.
func Preamble(data []byte) models.Preamble {
	preamble := binary.BigEndian.Uint32(data)
	return models.Preamble(preamble)
}

// DataFieldLength This function parse the data field length from AVL data.
func DataFieldLength(data []byte) models.DataFieldLength {
	dataFieldLength := binary.BigEndian.Uint32(data)
	return models.DataFieldLength(dataFieldLength)
}

// DataSize This function parse the data field length from AVL data.
func DataSize(data []byte) models.DataSize {
	dataFieldLength := binary.BigEndian.Uint32(data)
	return models.DataSize(dataFieldLength)
}

// CodecId This function parse the codec id from AVL data.
func CodecId(data byte) models.CodecID {
	return models.CodecID(data)
}

// NumberOfData This function parse the number of data from AVL data.
func NumberOfData(data byte) models.NumberOfData {
	return models.NumberOfData(data)
}

// Quantity This function parse the number of data from AVL data.
func Quantity(data byte) models.Quantity {
	return models.Quantity(data)
}

// Type This function parse the type from AVL data.
func Type(data byte) models.Type {
	return models.Type(data)
}

// ResponseSize This function parse the response size from AVL data.
func ResponseSize(data []byte) models.ResponseSize {
	responseSize := binary.BigEndian.Uint32(data)
	return models.ResponseSize(responseSize)
}

// Crc16 This function parse the crc16 from AVL data.
func Crc16(data []byte) models.Crc16 {
	crc16 := binary.BigEndian.Uint32(data)
	return models.Crc16(crc16)
}

// Timestamp This function parse the timestamp from AVL data.
func Timestamp(startIndex int, body []byte) (models.Timestamp, int) {
	const timestampLength = 8

	endIndex := startIndex + timestampLength
	ts := body[startIndex:endIndex]

	unixTs := binary.BigEndian.Uint64(ts)

	return models.Timestamp(unixTs), endIndex
}

// Priority This function parse the priority from AVL data.
func Priority(index int, body []byte) (models.Priority, int) {
	const priorityLength = 1

	priority := body[index]

	return models.Priority(priority), index + priorityLength
}

// GpsElement This function parse the GPS data from AVL data.
func GpsElement(startIndex int, body []byte) (models.GpsElement, int) {
	const gpsElementLength = 15

	endIndex := startIndex + gpsElementLength
	data := body[startIndex:endIndex]
	var gps models.GpsElement

	gps.Longitude = models.Longitude(decodeCoordinate(data[0:4]))
	gps.Latitude = models.Latitude(decodeCoordinate(data[4:8]))
	gps.Altitude = models.Altitude(binary.BigEndian.Uint16(data[8:10]))
	gps.Angle = models.Angle(binary.BigEndian.Uint16(data[10:12]))
	gps.Satellites = models.Satellites(data[12])
	gps.Speed = models.Speed(binary.BigEndian.Uint16(data[13:15]))

	return gps, endIndex
}

// decodeCoordinate This function return a position negative.
//
// To determine if the coordinate is negative, convert it to binary format and check the very first bit.
// If it is 0, coordinate is positive, if it is 1, coordinate is negative.
func decodeCoordinate(position []byte) float64 {
	raw := int32(binary.BigEndian.Uint32(position))
	return float64(raw) / 1e7
}
