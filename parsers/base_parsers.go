package parsers

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/rojack96/gonika/models"
)

// ParseTimestamp This function parse the timestamp from AVL data.
func ParseTimestamp(startIndex int, body []byte) (time.Time, int) {
	endIndex := startIndex + 8
	ts := body[startIndex:endIndex]

	unixTs := binary.BigEndian.Uint64(ts)
	return time.UnixMilli(int64(unixTs)), endIndex
}

// ParsePriority This function parse the priority from AVL data.
func ParsePriority(index int, body []byte) (byte, int) {
	priority := body[index]

	return priority, index + 1
}

// ParseGPSElement This function parse the GPS data from AVL data.
func ParseGPSElement(startIndex int, body []byte) (models.GPSElement, int) {

	endIndex := startIndex + 15
	data := body[startIndex:endIndex]
	var gps models.GPSElement

	gps.Longitude = PositionAnalyzer(data[0:4])
	gps.Latitude = PositionAnalyzer(data[4:8])
	gps.Altitude = binary.BigEndian.Uint16(data[8:10])
	gps.Angle = binary.BigEndian.Uint16(data[10:12])
	gps.Satellites = data[12]
	gps.Speed = binary.BigEndian.Uint16(data[13:15])

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
