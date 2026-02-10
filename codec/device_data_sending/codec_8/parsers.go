package codec8

import (
	"encoding/hex"
)

// ParseEventIO  This function parse the Event IO data from AVL data.
//
// Event IO changed
// if data is acquired on event
// this field defines which IO property has changed and generated an event.
func (c *codec8) parseEventIO(startIndex int, body []byte) (uint8, int) {
	eventIOID := body[startIndex]

	return eventIOID, startIndex + 1
}

// parseTotalNumberOfIO Total Number of IO.
//
// A total number of properties coming with record (N = N1 + N2 + N4 + N8).
func (c *codec8) parseTotalNumberOfIO(startIndex int, body []byte) (uint8, int) {
	noOfTotalIO := body[startIndex]

	return noOfTotalIO, startIndex + 1
}

// parseIo This function parse byte IO.
func (c *codec8) parseIo(valueLength int, startIndex int, body []byte) (uint8, map[uint8]string, int) {
	const idLength = 1
	splitByte := idLength + valueLength
	result := map[uint8]string{}

	if startIndex >= len(body) {
		return 0, result, len(body)
	}

	// One Byte IO Number
	nOfIOelements := body[startIndex]
	// One Byte IO Data
	IOelementsStartIndex := startIndex + 1
	IOelementsEndIndex := IOelementsStartIndex + int(nOfIOelements)*splitByte
	data := body[IOelementsStartIndex:IOelementsEndIndex]

	for i := 0; i < len(data); i += splitByte {
		id := data[i]
		value := hex.EncodeToString(data[i+1 : i+splitByte])

		result[id] = value
	}

	return nOfIOelements, result, IOelementsEndIndex
}
