package codec8e

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/rojack96/teltonika-parser/models"
)

// This function parse the timestamp from AVL data.
func parseTimestamp(startIndex int, body []byte) (time.Time, int) {
	endIndex := startIndex + 8
	ts := body[startIndex:endIndex]

	unixTs := binary.BigEndian.Uint64(ts)
	return time.UnixMilli(int64(unixTs)), endIndex
}

// This function parse the priority from AVL data.
func parsePriority(index int, body []byte) (byte, int) {
	priority := body[index]

	return priority, index + 1
}

// This function parse the GPS data from AVL data.
func parseGPSElement(startIndex int, body []byte) (models.GPSElement, int) {

	endIndex := startIndex + 15
	data := body[startIndex:endIndex]
	var gps models.GPSElement

	gps.Longitude = positionAnalyzer(data[0:4])
	gps.Latitude = positionAnalyzer(data[4:8])
	gps.Altitude = binary.BigEndian.Uint16(data[8:10])
	gps.Angle = binary.BigEndian.Uint16(data[10:12])
	gps.Satellites = data[12]
	gps.Speed = binary.BigEndian.Uint16(data[13:15])

	return gps, endIndex
}

// This function return a position negative.
//
// To determine if the coordinate is negative, convert it to binary format and check the very first bit.
// If it is 0, coordinate is positive, if it is 1, coordinate is negative.
func positionAnalyzer(position []byte) int32 {
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
