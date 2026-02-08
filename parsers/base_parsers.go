package parsers

import (
	"encoding/binary"
	"fmt"

	"github.com/rojack96/gonika/models"
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

// CodecId This function parse the codec id from AVL data.
func CodecId(data byte) models.CodecID {
	return models.CodecID(data)
}

// NumberOfData This function parse the number of data from AVL data.
func NumberOfData(data byte) models.NumberOfData {
	return models.NumberOfData(data)
}

// Crc16 This function parse the crc16 from AVL data.
func Crc16(data []byte) models.Crc16 {
	crc16 := binary.BigEndian.Uint32(data)
	return models.Crc16(crc16)
}

// Timestamp This function parse the timestamp from AVL data.
func Timestamp(startIndex int, body []byte) (models.Timestamp, int) {
	endIndex := startIndex + 8
	ts := body[startIndex:endIndex]

	unixTs := binary.BigEndian.Uint64(ts)
	return models.Timestamp(unixTs), endIndex
}

// Priority This function parse the priority from AVL data.
func Priority(index int, body []byte) (models.Priority, int) {
	priority := body[index]

	return models.Priority(priority), index + 1
}

// GpsElement This function parse the GPS data from AVL data.
func GpsElement(startIndex int, body []byte) (models.GpsElement, int) {

	endIndex := startIndex + 15
	data := body[startIndex:endIndex]
	var gps models.GpsElement

	gps.Longitude = models.Longitude(PositionAnalyzer(data[0:4]))
	gps.Latitude = models.Latitude(PositionAnalyzer(data[4:8]))
	gps.Altitude = models.Altitude(binary.BigEndian.Uint16(data[8:10]))
	gps.Angle = models.Angle(binary.BigEndian.Uint16(data[10:12]))
	gps.Satellites = models.Satellites(data[12])
	gps.Speed = models.Speed(binary.BigEndian.Uint16(data[13:15]))

	return gps, endIndex
}

// PositionAnalyzer This function return a position negative.
//
// To determine if the coordinate is negative, convert it to binary format and check the very first bit.
// If it is 0, coordinate is positive, if it is 1, coordinate is negative.
func PositionAnalyzer(position []byte) int32 {
	var bits string
	firstByte := position[0]
	bits = fmt.Sprintf("%08b", byte(firstByte))

	if bits[0:1] == "1" {
		r := binary.BigEndian.Uint32(position)
		res := -(int32(r))
		return res
	}
	res := binary.BigEndian.Uint32(position)
	return int32(res)
}
