package codec12

import (
	// golang import
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"

	// project import
	models "sipli/device/teltonika/tcp_server/models"
	// external import
)

// This function parse the timestamp from AVL data.
func ParseTimestamp(startIndex int, body []byte) (time.Time, string, int) {

	timestampStartIndex := startIndex
	timestampEndIndex := startIndex + 8
	timestamp := body[timestampStartIndex:timestampEndIndex]
	rowTimestamp := hex.EncodeToString(timestamp)

	unixTimestamp := binary.BigEndian.Uint64(timestamp)
	return time.UnixMilli(int64(unixTimestamp)), rowTimestamp, timestampEndIndex
}

// This function parse the priority from AVL data.
func ParsePriority(index int, body []byte) (byte, int) {
	priority := body[index]

	return priority, index + 1
}

// This function parse the GPS data from AVL data.
func ParseGPSElement(startIndex int, body []byte) (models.GPSElement, string, int) {

	gpsStartIndex := startIndex
	gpsEndIndex := gpsStartIndex + 15
	data := body[gpsStartIndex:gpsEndIndex]
	var gps models.GPSElement

	gps.Longitude = positionAnalyzer(data[0:4])
	gps.Latitude = positionAnalyzer(data[4:8])
	gps.Altitude = binary.BigEndian.Uint16(data[8:10])
	gps.Angle = binary.BigEndian.Uint16(data[10:12])
	gps.Satellites = data[12]
	gps.Speed = binary.BigEndian.Uint16(data[13:15])

	gpsRaw := gpsRawData(data)

	return gps, gpsRaw, gpsEndIndex
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

func gpsRawData(data []byte) string {
	rawData := fmt.Sprintf("%s,%s,%s,%s,%s,%s",
		hex.EncodeToString(data[0:4]),
		hex.EncodeToString(data[4:8]),
		hex.EncodeToString(data[8:10]),
		hex.EncodeToString(data[10:12]),
		hex.EncodeToString([]byte{data[12]}),
		hex.EncodeToString(data[13:15]),
	)

	return rawData
}
