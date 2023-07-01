package codecs

import (
	"encoding/binary"
)

/* CODEC 8 EXTENDED */

// This function parse the Event IO data from AVL data.
//
// Event IO changed
// if data is acquired on event
// this field defines which IO property has changed and generated an event.
func codec8ParseEventIO(startIndex int, body []byte) (uint8, int) {
	eventIOID := body[startIndex]

	return eventIOID, startIndex + 1
}

// Total Number of IO.
//
// A total number of properties coming with record (N = N1 + N2 + N4 + N8).
func codec8ParseTotalNumberOfIO(startIndex int, body []byte) (uint8, int) {
	noOfTotalIO := body[startIndex]

	return noOfTotalIO, startIndex + 1
}

// This function parse one byte IO.
//
// N1
// number of properties, which length is 1 byte.
func codec8ParseOneByteIO(startIndex int, body []byte) (uint8, map[uint8]uint8, int) {
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
func codec8ParseTwoByteIO(startIndex int, body []byte) (uint8, map[uint8]uint16, int) {
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
func codec8ParseFourByteIO(startIndex int, body []byte) (uint8, map[uint8]uint32, int) {
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
func codec8ParseEightByteIO(startIndex int, body []byte) (uint8, map[uint8]uint64, int) {
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

/* CODEC 8 EXTENDED */

// This function parse the Event IO data from AVL data.
//
// Event IO changed
// if data is acquired on event
// this field defines which IO property has changed and generated an event.
func codec8eParseEventIO(startIndex int, body []byte) (uint16, int) {
	endIndex := startIndex + 2
	eventIOID := binary.BigEndian.Uint16(body[startIndex:endIndex])

	return eventIOID, endIndex
}

// Total Number of IO.
//
// A total number of properties coming with record (N = N1 + N2 + N4 + N8).
func codec8eParseTotalNumberOfIO(startIndex int, body []byte) (uint16, int) {
	endIndex := startIndex + 2
	noOfTotalIO := binary.BigEndian.Uint16(body[startIndex:endIndex])

	return noOfTotalIO, endIndex
}

// This function parse one byte IO.
//
// N1
// number of properties, which length is 1 byte.
func codec8eParseOneByteIO(startIndex int, body []byte) (uint16, map[uint16]uint8, int) {
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
func codec8eParseTwoByteIO(startIndex int, body []byte) (uint16, map[uint16]uint16, int) {
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
func codec8eParseFourByteIO(startIndex int, body []byte) (uint16, map[uint16]uint32, int) {
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
func codec8eParseEightByteIO(startIndex int, body []byte) (uint16, map[uint16]uint64, int) {
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
func codec8eParseXByteIO(startIndex int, body []byte) (uint16, map[uint16]uint, int) {
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

/* CODEC 16 */

// This function parse the Event IO data from AVL data.
//
// Event IO changed
// if data is acquired on event
// this field defines which IO property has changed and generated an event.
func codec16ParseEventIO(startIndex int, body []byte) (uint16, int) {
	endIndex := startIndex + 2
	eventIOID := binary.BigEndian.Uint16(body[startIndex:endIndex])

	return eventIOID, endIndex
}

func codec16ParseGenerationType(startIndex int, body []byte) (uint8, int) {
	generationType := body[startIndex]

	return generationType, startIndex + 1
}

// Total Number of IO.
//
// A total number of properties coming with record (N = N1 + N2 + N4 + N8).
func codec16ParseTotalNumberOfIO(startIndex int, body []byte) (uint8, int) {
	noOfTotalIO := body[startIndex]

	return noOfTotalIO, startIndex + 1
}

// This function parse one byte IO.
//
// N1
// number of properties, which length is 1 byte.
func codec16ParseOneByteIO(startIndex int, body []byte) (uint8, map[uint16]uint8, int) {
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
func codec16ParseTwoByteIO(startIndex int, body []byte) (uint8, map[uint16]uint16, int) {
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
func codec16ParseFourByteIO(startIndex int, body []byte) (uint8, map[uint16]uint32, int) {
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
func codec16ParseEightByteIO(startIndex int, body []byte) (uint8, map[uint16]uint64, int) {
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
