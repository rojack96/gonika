package codec8e

import (
	"encoding/binary"
	"encoding/hex"
)

// ParseEventIO This function parse the Event IO data from AVL data.
//
// Event IO changed
// if data is acquired on event
// this field defines which IO property has changed and generated an event.
func (c *Codec8e) parseEventIO(startIndex int, body []byte) (uint16, int) {
	endIndex := startIndex + 2
	eventIOID := binary.BigEndian.Uint16(body[startIndex:endIndex])

	return eventIOID, endIndex
}

// ParseTotalNumberOfIO Total Number of IO.
//
// A total number of properties coming with record (N = N1 + N2 + N4 + N8).
func (c *Codec8e) parseTotalNumberOfIO(startIndex int, body []byte) (uint16, int) {
	endIndex := startIndex + 2
	noOfTotalIO := binary.BigEndian.Uint16(body[startIndex:endIndex])

	return noOfTotalIO, endIndex
}

// parseIo This function parse byte IO.
func (c *Codec8e) parseIo(valueLength int, startIndex int, body []byte) (uint16, map[uint16]string, int) {
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
func (c *Codec8e) parseXByteIO(startIndex int, body []byte) (uint16, map[uint16]string, int) {
	// TODO questo codice va cambiato perché ora cé quello funzionante
	IOelements := map[uint16]string{}

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

			value := hex.EncodeToString(body[valueStartIndex:valueEndIndex])

			IOelements[id] = value

			i = i + 1
			j = j + valueEndIndex
		}

		IOelementsEndIndex := j
		return nOfIOelements, IOelements, IOelementsEndIndex
	}

	return nOfIOelements, IOelements, startIndex + 2
}
