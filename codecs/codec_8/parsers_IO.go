package codec8

import (
	"encoding/binary"
)

// This function parse the Event IO data from AVL data.
//
// Event IO changed
// if data is acquired on event
// this field defines which IO property has changed and generated an event.
func parseEventIO(startIndex int, body []byte) (uint8, int) {
	eventIOID := body[startIndex]

	return eventIOID, startIndex + 1
}

// Total Number of IO.
//
// A total number of properties coming with record (N = N1 + N2 + N4 + N8).
func parseTotalNumberOfIO(startIndex int, body []byte) (uint8, int) {
	noOfTotalIO := body[startIndex]

	return noOfTotalIO, startIndex + 1
}

// This function parse one byte IO.
//
// N1
// number of properties, which length is 1 byte.
func parseOneByteIO(startIndex int, body []byte) (uint8, map[uint8]uint8, int) {
	IOelements := map[uint8]uint8{}
	splitByte := 2 // ID 1 byte + VALUE 1 byte

	if startIndex < len(body) {
		// One Byte IO Number
		nOfIOelements := body[startIndex]
		// One Byte IO Data
		IOelementsStartIndex := startIndex + 1
		IOelementsEndIndex := IOelementsStartIndex + int(nOfIOelements)*splitByte
		data := body[IOelementsStartIndex:IOelementsEndIndex]

		for i := 0; i < len(data); i += splitByte {
			id := data[i]
			value := data[i+1]

			IOelements[id] = value
		}
		return nOfIOelements, IOelements, IOelementsEndIndex
	}
	return 0, IOelements, len(body)
}

// This function parse two byte IO.
//
// N2
// number of properties, which length is 2 byte.
func parseTwoByteIO(startIndex int, body []byte) (uint8, map[uint8]uint16, int) {
	IOelements := map[uint8]uint16{}
	splitByte := 3 // ID 1 byte + VALUE 2 bytes
	if startIndex < len(body) {
		// Two Byte IO Number
		nOfIOelements := body[startIndex]
		// Two Byte IO Data
		IOelementsStartIndex := startIndex + 1
		IOelementsEndIndex := IOelementsStartIndex + int(nOfIOelements)*splitByte
		data := body[IOelementsStartIndex:IOelementsEndIndex]

		for i := 0; i < len(data); i += splitByte {
			id := data[i]
			value := binary.BigEndian.Uint16(data[i+1 : i+splitByte])

			IOelements[id] = value
		}
		return nOfIOelements, IOelements, IOelementsEndIndex
	}
	return 0, IOelements, len(body)
}

// This function parse four byte IO.
//
// N4
// number of properties, which length is 4 byte.
func parseFourByteIO(startIndex int, body []byte) (uint8, map[uint8]uint32, int) {
	IOelements := map[uint8]uint32{}
	splitByte := 5 // ID 1 byte + VALUE 4 bytes

	if startIndex < len(body) {
		// Four Byte IO Number
		nOfIOelements := body[startIndex]
		// Four Byte IO Number
		IOelementsStartIndex := startIndex + 1
		IOelementsEndIndex := IOelementsStartIndex + int(nOfIOelements)*splitByte
		data := body[IOelementsStartIndex:IOelementsEndIndex]

		for i := 0; i < len(data); i += splitByte {
			id := data[i]
			value := binary.BigEndian.Uint32(data[i+1 : i+splitByte])

			IOelements[id] = value
		}

		return nOfIOelements, IOelements, IOelementsEndIndex
	}
	return 0, IOelements, len(body)
}

// This function parse eight byte IO.
//
// N8
// number of properties, which length is 8 byte.
// Eight Byte IO Number
func parseEightByteIO(startIndex int, body []byte) (uint8, map[uint8]uint64, int) {
	IOelements := map[uint8]uint64{}
	splitByte := 9 // ID 1 byte + VALUE 8 bytes

	if startIndex < len(body) {
		// Eight Byte IO Number
		nOfIOelements := body[startIndex]
		// Eight Byte IO Data
		IOelementsStartIndex := startIndex + 1
		IOelementsEndIndex := IOelementsStartIndex + int(nOfIOelements)*splitByte
		data := body[IOelementsStartIndex:IOelementsEndIndex]

		for i := 0; i < len(data); i += splitByte {
			id := data[i]
			value := binary.BigEndian.Uint64(data[i+1 : i+splitByte])

			IOelements[id] = value
		}

		return nOfIOelements, IOelements, IOelementsEndIndex
	}

	return 0, IOelements, len(body)
}
