package common

import (
	// golang import
	"encoding/binary"
	"encoding/hex"
	"fmt"
	// project import
	// external import
)

// This function parse the Event IO data from AVL data.
//
// Event IO changed
// if data is acquired on event
// this field defines which IO property has changed and generated an event.
func parseEventIO(startIndex int, body []byte) (uint8, int) {
	eventIOID := body[startIndex]

	return uint8(eventIOID), startIndex + 1
}

// Total Number of IO.
//
// A total number of properties coming with record (N = N1 + N2 + N4 + N8).
func parseTotalNumberOfIO(startIndex int, body []byte) (uint8, int) {
	noOfTotalIO := body[startIndex]

	return uint8(noOfTotalIO), startIndex + 1
}

// This function parse one byte IO.
//
// N1
// number of properties, which length is 1 byte.
func parseOneByteIO(startIndex int, body []byte, IORawData []string) (uint8, map[uint8]uint8, []string, int) {
	oneByteIO := map[uint8]uint8{}
	splitByte := 2 // ID 1 byte + VALUE 1 byte
	if startIndex < len(body) {
		// One Byte IO Number
		noOfOneByteIO := body[startIndex]
		// One Byte IO Data
		oneByteIOStartIndex := startIndex + 1
		oneByteIOEndIndex := oneByteIOStartIndex + int(noOfOneByteIO)*splitByte
		data := body[oneByteIOStartIndex:oneByteIOEndIndex]

		for i := 0; i < len(data); i += splitByte {
			id := data[i]
			value := data[i+1]

			oneByteIO[id] = value

			IORawData = append(IORawData, fmt.Sprintf("%s=%s", hex.EncodeToString([]byte{id}), hex.EncodeToString([]byte{value})))
		}

		return noOfOneByteIO, oneByteIO, IORawData, oneByteIOEndIndex
	}
	return 0, oneByteIO, []string{}, len(body)
}

// This function parse two byte IO.
//
// N2
// number of properties, which length is 2 byte.
func parseTwoByteIO(startIndex int, body []byte, IORawData []string) (uint8, map[uint8]uint16, []string, int) {
	twoByteIO := map[uint8]uint16{}
	splitByte := 3 // ID 1 byte + VALUE 2 bytes
	if startIndex < len(body) {
		// Two Byte IO Number
		noOfTwoByteIO := body[startIndex]
		// Two Byte IO Data
		twoByteIOStartIndex := startIndex + 1
		twoByteIOEndIndex := twoByteIOStartIndex + int(noOfTwoByteIO)*splitByte
		data := body[twoByteIOStartIndex:twoByteIOEndIndex]

		for i := 0; i < len(data); i += splitByte {
			id := data[i]
			value := binary.BigEndian.Uint16(data[i+1 : i+splitByte])

			twoByteIO[id] = value

			IORawData = append(IORawData, fmt.Sprintf("%s=%s", hex.EncodeToString([]byte{id}), hex.EncodeToString(data[i+1:i+splitByte])))
		}

		return noOfTwoByteIO, twoByteIO, IORawData, twoByteIOEndIndex
	}
	return 0, twoByteIO, []string{}, len(body)
}

// This function parse four byte IO.
//
// N4
// number of properties, which length is 4 byte.
func parseFourByteIO(startIndex int, body []byte, IORawData []string) (uint8, map[uint8]uint32, []string, int) {
	fourByteIO := map[uint8]uint32{}
	splitByte := 5 // ID 1 byte + VALUE 4 bytes

	if startIndex < len(body) {
		// Four Byte IO Number
		noOfFourByteIO := body[startIndex]
		// Four Byte IO Number
		fourByteIOStartIndex := startIndex + 1
		fourByteIOEndIndex := fourByteIOStartIndex + int(noOfFourByteIO)*splitByte
		data := body[fourByteIOStartIndex:fourByteIOEndIndex]

		for i := 0; i < len(data); i += splitByte {
			id := data[i]
			value := binary.BigEndian.Uint32(data[i+1 : i+splitByte])

			fourByteIO[id] = value

			IORawData = append(IORawData, fmt.Sprintf("%s=%s", hex.EncodeToString([]byte{id}), hex.EncodeToString(data[i+1:i+splitByte])))
		}

		return noOfFourByteIO, fourByteIO, IORawData, fourByteIOEndIndex
	}
	return 0, fourByteIO, []string{}, len(body)
}

// This function parse eight byte IO.
//
// N8
// number of properties, which length is 8 byte.
// Eight Byte IO Number
func parseEightByteIO(startIndex int, body []byte, IORawData []string) (uint8, map[uint8]uint64, []string, int) {
	eightByteIO := map[uint8]uint64{}
	splitByte := 9 // ID 1 byte + VALUE 8 bytes

	if startIndex < len(body) {
		// Eight Byte IO Number
		noOfEightByteIO := body[startIndex]
		// Eight Byte IO Data
		eightByteIOStartIndex := startIndex + 1
		eightByteIOEndIndex := eightByteIOStartIndex + int(noOfEightByteIO)*splitByte
		data := body[eightByteIOStartIndex:eightByteIOEndIndex]

		for i := 0; i < len(data); i += splitByte {
			id := data[i]
			value := binary.BigEndian.Uint64(data[i+1 : i+splitByte])

			eightByteIO[id] = value

			IORawData = append(IORawData, fmt.Sprintf("%s=%s", hex.EncodeToString([]byte{id}), hex.EncodeToString(data[i+1:i+splitByte])))
		}

		return noOfEightByteIO, eightByteIO, IORawData, eightByteIOEndIndex
	}

	return 0, eightByteIO, []string{}, len(body)
}
