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
	IOelements := map[uint16]uint8{}
	splitByte := 3
	// One Byte IO Number
	nOfIOstartIndex := startIndex
	nOfIOendIndex := nOfIOstartIndex + 2
	nOfIOelements := binary.BigEndian.Uint16(body[nOfIOstartIndex:nOfIOendIndex])
	// One Byte IO Data
	IOelementsStartIndex := nOfIOendIndex
	IOelementsEndIndex := IOelementsStartIndex + int(nOfIOelements)*splitByte
	data := body[IOelementsStartIndex:IOelementsEndIndex]

	for i := 0; i < len(data); i += splitByte {
		id := binary.BigEndian.Uint16(data[i : i+2])
		value := data[i+2]

		IOelements[id] = value
	}

	return nOfIOelements, IOelements, IOelementsEndIndex
}

// This function parse two byte IO.
//
// N2
// number of properties, which length is 2 byte.
func parseTwoByteIO(startIndex int, body []byte) (uint16, map[uint16]uint16, int) {
	IOelements := map[uint16]uint16{}
	splitByte := 4
	// Two Byte IO Number
	nOfIOstartIndex := startIndex
	nOfIOendIndex := nOfIOstartIndex + 2
	nOfIOelements := binary.BigEndian.Uint16(body[nOfIOstartIndex:nOfIOendIndex])
	// Two Byte IO Data
	IOelementsStartIndex := nOfIOendIndex
	IOelementsEndIndex := IOelementsStartIndex + int(nOfIOelements)*splitByte
	data := body[IOelementsStartIndex:IOelementsEndIndex]

	for i := 0; i < len(data); i += splitByte {
		id := binary.BigEndian.Uint16(data[i : i+2])
		value := binary.BigEndian.Uint16(data[i+2 : i+splitByte])

		IOelements[id] = value
	}

	return nOfIOelements, IOelements, IOelementsEndIndex
}

// This function parse four byte IO.
//
// N4
// number of properties, which length is 4 byte.
func parseFourByteIO(startIndex int, body []byte) (uint16, map[uint16]uint32, int) {
	IOelements := map[uint16]uint32{}
	splitByte := 6
	// Four Byte IO Number
	nOfIOstartIndex := startIndex
	nOfIOendIndex := nOfIOstartIndex + 2
	nOfIOelements := binary.BigEndian.Uint16(body[nOfIOstartIndex:nOfIOendIndex])
	// Four Byte IO Number
	IOelementsStartIndex := nOfIOendIndex
	IOelementsEndIndex := IOelementsStartIndex + int(nOfIOelements)*splitByte
	data := body[IOelementsStartIndex:IOelementsEndIndex]

	for i := 0; i < len(data); i += splitByte {
		id := binary.BigEndian.Uint16(data[i : i+2])
		value := binary.BigEndian.Uint32(data[i+2 : i+splitByte])

		IOelements[id] = value
	}

	return nOfIOelements, IOelements, IOelementsEndIndex
}

// This function parse eight byte IO.
//
// N8
// number of properties, which length is 8 byte.
// Eight Byte IO Number
func parseEightByteIO(startIndex int, body []byte) (uint16, map[uint16]uint64, int) {
	IOelements := map[uint16]uint64{}
	splitByte := 10
	// Eight Byte IO Number
	nOfIOstartIndex := startIndex
	nOfIOendIndex := nOfIOstartIndex + 2
	nOfIOelements := binary.BigEndian.Uint16(body[nOfIOstartIndex:nOfIOendIndex])
	// Eight Byte IO Data
	IOelementsStartIndex := nOfIOendIndex
	IOelementsEndIndex := IOelementsStartIndex + int(nOfIOelements)*splitByte
	data := body[IOelementsStartIndex:IOelementsEndIndex]

	for i := 0; i < len(data); i += splitByte {
		id := binary.BigEndian.Uint16(data[i : i+2])
		value := binary.BigEndian.Uint64(data[i+2 : i+splitByte])

		IOelements[id] = value
	}

	return nOfIOelements, IOelements, IOelementsEndIndex
}

// This function parse X byte IO.
//
// NX
// a number of properties, which length is defined by length element.
// X Byte IO Number
func parseXByteIO(startIndex int, body []byte) (uint16, map[uint16]uint, int) {
	IOelements := map[uint16]uint{}

	// Eight Byte IO Number
	nOfIOstartIndex := startIndex
	nOfIOendIndex := nOfIOstartIndex + 2
	nOfIOelements := binary.BigEndian.Uint16(body[nOfIOstartIndex:nOfIOendIndex])
	// Eight Byte IO Data
	if nOfIOelements != 0 {
		IOelementsStartIndex := nOfIOendIndex

		var i uint16 = 0
		j := 0

		for i < nOfIOelements {
			IOelementsIDstartIndex := j + IOelementsStartIndex
			IOelementsIDendIndex := IOelementsIDstartIndex + 2

			id := binary.BigEndian.Uint16(body[IOelementsIDstartIndex:IOelementsIDendIndex])

			valueLengthStartIndex := IOelementsIDendIndex
			valueLengthEndIndex := valueLengthStartIndex + 2

			valueLength := binary.BigEndian.Uint16(body[valueLengthStartIndex:valueLengthEndIndex])

			if valueLength == 0 {
				return nOfIOelements, IOelements, valueLengthEndIndex
			}

			valueStartIndex := valueLengthEndIndex
			valueEndIndex := valueStartIndex + int(valueLength)

			value := binary.BigEndian.Uint64(body[valueStartIndex:valueEndIndex])

			IOelements[id] = uint(value)

			i = i + 1
			j = j + valueEndIndex
		}

		IOelementsEndIndex := j
		return nOfIOelements, IOelements, IOelementsEndIndex
	}

	return nOfIOelements, IOelements, startIndex + 2
}
