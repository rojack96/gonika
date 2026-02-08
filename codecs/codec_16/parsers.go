package codec16

import (
	"encoding/binary"
	"encoding/hex"
)

// This function parse the Event IO data from AVL data.
//
// Event IO changed
// if data is acquired on event
// this field defines which IO property has changed and generated an event.
func (c *Codec16) parseEventIO(startIndex int, body []byte) uint16 {
	endIndex := startIndex + 2
	eventIOID := binary.BigEndian.Uint16(body[startIndex:endIndex])

	return eventIOID
}

func (c *Codec16) parseGenerationType(startIndex int, body []byte) uint8 {
	generationType := body[startIndex]

	return generationType
}

// Total Number of IO.
//
// A total number of properties coming with record (N = N1 + N2 + N4 + N8).
func (c *Codec16) parseTotalNumberOfIO(startIndex int, body []byte) uint8 {
	noOfTotalIO := body[startIndex]

	return noOfTotalIO
}

// parseIo This function parse byte IO.
func (c *Codec16) parseIo(valueLength int, startIndex int, body []byte) (uint8, map[uint16]string, int) {
	const idLength = 2
	splitByte := idLength + valueLength
	result := map[uint16]string{}

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
		id := binary.BigEndian.Uint16(data[i : i+2])
		value := hex.EncodeToString(data[i+2 : i+splitByte])

		result[id] = value
	}

	return nOfIOelements, result, IOelementsEndIndex
}
