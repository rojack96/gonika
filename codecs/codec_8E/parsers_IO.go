package codec8e

import (
	"encoding/binary"
)

// This function parse the Event IO data from AVL data.
//
// Event IO changed
// if data is acquired on event
// this field defines which IO property has changed and generated an event.
func parseEventIO(startIndex int, body []byte) (uint16, int) {
	endIndex := startIndex + 2
	eventIOID := binary.BigEndian.Uint16(body[startIndex:endIndex])

	return eventIOID, endIndex
}

// Total Number of IO.
//
// A total number of properties coming with record (N = N1 + N2 + N4 + N8).
func parseTotalNumberOfIO(startIndex int, body []byte) (uint16, int) {
	endIndex := startIndex + 2
	noOfTotalIO := binary.BigEndian.Uint16(body[startIndex:endIndex])

	return noOfTotalIO, endIndex
}

// This function parse one byte IO.
//
// N1
// number of properties, which length is 1 byte.
func parseOneByteIO(startIndex int, body []byte) (uint16, map[uint16]uint8, int) {
	oneByteIO := map[uint16]uint8{}

	// One Byte IO Number
	noOfByteIOStartIndex := startIndex
	noOfByteIOEndIndex := noOfByteIOStartIndex + 2
	noOfOneByteIO := binary.BigEndian.Uint16(body[noOfByteIOStartIndex:noOfByteIOEndIndex])
	// One Byte IO Data
	IOstartIndex := noOfByteIOEndIndex
	IOendIndex := IOstartIndex + int(noOfOneByteIO)*3
	data := body[IOstartIndex:IOendIndex]

	for i := 0; i < len(data); i += 3 {
		id := binary.BigEndian.Uint16(data[i : i+2])
		value := data[i+2]

		oneByteIO[id] = value
	}

	return noOfOneByteIO, oneByteIO, IOendIndex
}

// This function parse two byte IO.
//
// N2
// number of properties, which length is 2 byte.
func parseTwoByteIO(startIndex int, body []byte) (uint16, map[uint16]uint16, int) {
	twoByteIO := map[uint16]uint16{}

	// Two Byte IO Number
	noOfTwoByteIOIndexStart := startIndex
	noOfTwoByteIOIndexEnd := noOfTwoByteIOIndexStart + 2
	noOfTwoByteIO := binary.BigEndian.Uint16(body[noOfTwoByteIOIndexStart:noOfTwoByteIOIndexEnd])
	// Two Byte IO Data
	twoByteIOStartIndex := noOfTwoByteIOIndexEnd
	twoByteIOEndIndex := twoByteIOStartIndex + int(noOfTwoByteIO)*4
	data := body[twoByteIOStartIndex:twoByteIOEndIndex]

	for i := 0; i < len(data); i += 4 {
		id := binary.BigEndian.Uint16(data[i : i+2])
		value := binary.BigEndian.Uint16(data[i+2 : i+4])

		twoByteIO[id] = value
	}

	return noOfTwoByteIO, twoByteIO, twoByteIOEndIndex
}

// This function parse four byte IO.
//
// N4
// number of properties, which length is 4 byte.
func parseFourByteIO(startIndex int, body []byte) (uint16, map[uint16]uint32, int) {
	fourByteIO := map[uint16]uint32{}

	// Four Byte IO Number
	noOfFourByteIOIndexStart := startIndex
	noOfFourByteIOIndexEnd := noOfFourByteIOIndexStart + 2
	noOfFourByteIO := binary.BigEndian.Uint16(body[noOfFourByteIOIndexStart:noOfFourByteIOIndexEnd])
	// Four Byte IO Number
	fourByteIOStartIndex := noOfFourByteIOIndexEnd
	fourByteIOEndIndex := fourByteIOStartIndex + int(noOfFourByteIO)*6
	data := body[fourByteIOStartIndex:fourByteIOEndIndex]

	for i := 0; i < len(data); i += 6 {
		id := binary.BigEndian.Uint16(data[i : i+2])
		value := binary.BigEndian.Uint32(data[i+2 : i+6])

		fourByteIO[id] = value
	}

	return noOfFourByteIO, fourByteIO, fourByteIOEndIndex
}

// This function parse eight byte IO.
//
// N8
// number of properties, which length is 8 byte.
// Eight Byte IO Number
func parseEightByteIO(startIndex int, body []byte) (uint16, map[uint16]uint64, int) {
	eightByteIO := map[uint16]uint64{}

	// Eight Byte IO Number
	noOfEightByteIOIndexStart := startIndex
	noOfEightByteIOIndexEnd := noOfEightByteIOIndexStart + 2
	noOfEightByteIO := binary.BigEndian.Uint16(body[noOfEightByteIOIndexStart:noOfEightByteIOIndexEnd])
	// Eight Byte IO Data
	eightByteIOStartIndex := noOfEightByteIOIndexEnd
	eightByteIOEndIndex := eightByteIOStartIndex + int(noOfEightByteIO)*10
	data := body[eightByteIOStartIndex:eightByteIOEndIndex]

	for i := 0; i < len(data); i += 10 {
		id := binary.BigEndian.Uint16(data[i : i+2])
		value := binary.BigEndian.Uint64(data[i+2 : i+10])

		eightByteIO[id] = value

	}

	return noOfEightByteIO, eightByteIO, eightByteIOEndIndex
}

// This function parse X byte IO.
//
// NX
// a number of properties, which length is defined by length element.
// X Byte IO Number
func parseXByteIO(startIndex int, body []byte) (uint16, map[uint16]uint, int) {
	xByteIO := map[uint16]uint{}

	// Eight Byte IO Number
	noOfXByteIOIndexStart := startIndex
	noOfXByteIOIndexEnd := noOfXByteIOIndexStart + 2
	noOfXByteIO := binary.BigEndian.Uint16(body[noOfXByteIOIndexStart:noOfXByteIOIndexEnd])
	// Eight Byte IO Data
	if noOfXByteIO != 0 {
		xByteIOStartIndex := noOfXByteIOIndexEnd

		var i uint16 = 0
		j := 0

		for i < noOfXByteIO {
			xIOIDStartIndex := j + xByteIOStartIndex
			xIOIDEndIndex := xIOIDStartIndex + 2

			id := binary.BigEndian.Uint16(body[xIOIDStartIndex:xIOIDEndIndex])

			xValueLengthStartIndex := xIOIDEndIndex
			xValueLengthEndIndex := xValueLengthStartIndex + 2

			valueLength := binary.BigEndian.Uint16(body[xValueLengthStartIndex:xValueLengthEndIndex])

			if valueLength == 0 {
				return noOfXByteIO, xByteIO, xValueLengthEndIndex
			}

			xValueStartIndex := xValueLengthEndIndex
			xValueEndIndex := xValueStartIndex + int(valueLength)

			value := binary.BigEndian.Uint64(body[xValueStartIndex:xValueEndIndex])

			xByteIO[id] = uint(value)

			i = i + 1
			j = j + xValueEndIndex
		}

		xByteIOEndIndex := j
		return noOfXByteIO, xByteIO, xByteIOEndIndex
	}

	return noOfXByteIO, xByteIO, startIndex + 2
}
