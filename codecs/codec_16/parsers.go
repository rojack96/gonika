package codec16

import "encoding/binary"

// This function parse the Event IO data from AVL data.
//
// Event IO changed
// if data is acquired on event
// this field defines which IO property has changed and generated an event.
func (c *Codec16) parseEventIO(startIndex int, body []byte) (uint16, int) {
	endIndex := startIndex + 2
	eventIOID := binary.BigEndian.Uint16(body[startIndex:endIndex])

	return eventIOID, endIndex
}

func (c *Codec16) parseGenerationType(startIndex int, body []byte) (uint8, int) {
	generationType := body[startIndex]

	return generationType, startIndex + 1
}

// Total Number of IO.
//
// A total number of properties coming with record (N = N1 + N2 + N4 + N8).
func (c *Codec16) parseTotalNumberOfIO(startIndex int, body []byte) (uint8, int) {
	noOfTotalIO := body[startIndex]

	return noOfTotalIO, startIndex + 1
}

// This function parse one byte IO.
//
// N1
// number of properties, which length is 1 byte.
func (c *Codec16) parseOneByteIO(startIndex int, body []byte) (uint8, map[uint16]uint8, int) {
	IOelements := map[uint16]uint8{}
	splitByte := 3
	// One Byte IO Number
	nOfIOelements := body[startIndex]
	// One Byte IO Data
	IOelementsStartIndex := startIndex + 1
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
func (c *Codec16) parseTwoByteIO(startIndex int, body []byte) (uint8, map[uint16]uint16, int) {
	IOelements := map[uint16]uint16{}
	splitByte := 4
	// Two Byte IO Number
	nOfIOelements := body[startIndex]
	// Two Byte IO Data
	IOelementsStartIndex := startIndex + 1
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
func (c *Codec16) parseFourByteIO(startIndex int, body []byte) (uint8, map[uint16]uint32, int) {
	IOelements := map[uint16]uint32{}
	splitByte := 6
	// Four Byte IO Number
	nOfIOelements := body[startIndex]
	// Four Byte IO Number
	IOelementsStartIndex := startIndex + 1
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
func (c *Codec16) parseEightByteIO(startIndex int, body []byte) (uint8, map[uint16]uint64, int) {
	IOelements := map[uint16]uint64{}
	splitByte := 10
	// Eight Byte IO Number
	nOfIOelements := body[startIndex]
	// Eight Byte IO Data
	IOelementsStartIndex := startIndex + 1
	IOelementsEndIndex := IOelementsStartIndex + int(nOfIOelements)*splitByte
	data := body[IOelementsStartIndex:IOelementsEndIndex]

	for i := 0; i < len(data); i += splitByte {
		id := binary.BigEndian.Uint16(data[i : i+2])
		value := binary.BigEndian.Uint64(data[i+2 : i+splitByte])

		IOelements[id] = value
	}

	return nOfIOelements, IOelements, IOelementsEndIndex
}
