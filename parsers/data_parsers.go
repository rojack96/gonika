package parsers

import (
	"encoding/binary"
	"fmt"

	"github.com/rojack96/gonika/models"
)

type DataParser struct {
	Body []byte
}

// ParseEventIO This function parse the Event IO data from AVL data.
//
// Event IO changed
// if data is acquired on events,
// this field defines which IO property has changed and generated an event.
func (p *DataParser) ParseEventIO(startIndex int, extended bool) (uint16, int) {
	endIndex := startIndex + 1
	eventIoId := []byte{p.Body[startIndex]}

	if extended {
		endIndex = startIndex + 2
		eventIoId = p.Body[startIndex:endIndex]
	}

	result := binary.BigEndian.Uint16(eventIoId)
	return result, endIndex
}

// ParseTimestamp This function parse the timestamp from AVL data.
func (p *DataParser) ParseTimestamp(startIndex int) (uint64, int) {
	endIndex := startIndex + 8
	ts := p.Body[startIndex:endIndex]

	unixTs := binary.BigEndian.Uint64(ts)
	return unixTs, endIndex
}

// ParsePriority This function parses the priority from AVL data.
func (p *DataParser) ParsePriority(index int) (byte, int) {
	priority := p.Body[index]

	return priority, index + 1
}

// ParseGpsElement This function parse the GPS data from AVL data.
func (p *DataParser) ParseGpsElement(startIndex int) (models.GpsElement, int) {

	endIndex := startIndex + 15
	data := p.Body[startIndex:endIndex]
	var gps models.GpsElement

	gps.Longitude = positionAnalyzer(data[0:4])
	gps.Latitude = positionAnalyzer(data[4:8])
	gps.Altitude = binary.BigEndian.Uint16(data[8:10])
	gps.Angle = binary.BigEndian.Uint16(data[10:12])
	gps.Satellites = data[12]
	gps.Speed = binary.BigEndian.Uint16(data[13:15])

	return gps, endIndex
}

// positionAnalyzer This function return a position negative.
//
// To determine if the coordinate is negative, convert it to binary format and check the very first bit.
// If it is 0, coordinate is positive, if it is 1, coordinate is negative.
func positionAnalyzer(position []byte) float64 {
	const LatLongConverter = 0.0000001

	var (
		bits      string
		resInUint uint32
		resInInt  int32
		resFloat  float64
	)
	firstByte := position[0]
	bits = fmt.Sprintf("%08b", firstByte)

	if bits[0:1] == "1" {
		resInUint = binary.BigEndian.Uint32(position)
		resInInt = -(int32(resInUint))
		resFloat = float64(resInInt) * LatLongConverter
	} else {
		resInUint = binary.BigEndian.Uint32(position)
		resFloat = float64(resInUint) * LatLongConverter
	}

	return resFloat
}
