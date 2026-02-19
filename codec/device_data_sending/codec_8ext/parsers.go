package codec8ext

import (
	"encoding/binary"
	"encoding/hex"
)

// ParseEventIO This function parse the Event IO data from AVL data.
//
// Event IO changed
// if data is acquired on event
// this field defines which IO property has changed and generated an event.
func (c *codec8ext) parseEventIO(startIndex int, body []byte) (uint16, int) {
	endIndex := startIndex + 2
	eventIOID := binary.BigEndian.Uint16(body[startIndex:endIndex])

	return eventIOID, endIndex
}

// ParseTotalNumberOfIO Total Number of IO.
//
// A total number of properties coming with record (N = N1 + N2 + N4 + N8).
func (c *codec8ext) parseTotalNumberOfIO(startIndex int, body []byte) (uint16, int) {
	endIndex := startIndex + 2
	noOfTotalIO := binary.BigEndian.Uint16(body[startIndex:endIndex])

	return noOfTotalIO, endIndex
}

// parseIo This function parse byte IO.
func (c *codec8ext) parseIo(valueLength int, startIndex int, body []byte) (uint16, map[uint16]string, int) {
	const idLength = 2
	splitByte := idLength + valueLength
	result := map[uint16]string{}

	if startIndex >= len(body) {
		return 0, result, len(body)
	}

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
		value := hex.EncodeToString(data[i+2 : i+splitByte])

		result[id] = value
	}

	return nOfIOelements, result, IOelementsEndIndex
}

// ParseXByteIO This function parse X byte IO.
//
// NX
// a number of properties, which length is defined by length element.
// X Byte IO Number
func (c *codec8ext) parseXByteIO(startIndex int, body []byte) (uint16, map[uint16]string, int) {
	result := map[uint16]string{}

	// Eight Byte IO Number
	nOfIOstartIndex := startIndex
	nOfIOendIndex := nOfIOstartIndex + 2
	noOfXByteIO := binary.BigEndian.Uint16(body[nOfIOstartIndex:nOfIOendIndex])
	// X Byte IO Data

	if noOfXByteIO == 0 { // if there is no X byte IO, return the index after reading number of X byte IO
		return noOfXByteIO, result, startIndex + 2
	}

	xByteIOStartIndex := nOfIOendIndex

	var i uint16 = 0
	j := 0

	for i < noOfXByteIO {
		xIOIDStartIndex := j + xByteIOStartIndex
		xIOIDEndIndex := xIOIDStartIndex + 2

		xValueLengthStartIndex := xIOIDEndIndex
		xValueLengthEndIndex := xValueLengthStartIndex + 2

		valueLength := binary.BigEndian.Uint16(body[xValueLengthStartIndex:xValueLengthEndIndex])

		if valueLength == 0 {
			return noOfXByteIO, result, xValueLengthEndIndex
		}

		xValueStartIndex := xValueLengthEndIndex
		xValueEndIndex := xValueStartIndex + int(valueLength)

		id := binary.BigEndian.Uint16(body[xIOIDStartIndex:xIOIDEndIndex])
		value := hex.EncodeToString(body[xValueStartIndex:xValueEndIndex])

		result[id] = value

		i = i + 1
		xByteIOStartIndex = xValueEndIndex
	}

	xByteIOEndIndex := xByteIOStartIndex
	return noOfXByteIO, result, xByteIOEndIndex

}
