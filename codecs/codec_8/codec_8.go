package codec8

import (
	"encoding/binary"

	"github.com/rojack96/gonika/models"
	"github.com/rojack96/gonika/parsers"
)

type Codec8 struct {
	Body []byte
}

func (c *Codec8) Decode(dataPacket []byte) models.AvlDataArray {
	var avlDataArray models.AvlDataArray

	mapper := models.Mapper{Data: dataPacket}
	avlDataPacket := mapper.CodecDataSending()

	avlDataArray.Preamble = binary.BigEndian.Uint32(avlDataPacket.Preamble)
	avlDataArray.CodecId = avlDataPacket.CodecId
	avlDataArray.NumberOfData1 = avlDataPacket.NumberOfData1
	avlDataArray.NumberOfData2 = avlDataPacket.NumberOfData2
	avlDataArray.CRC16 = binary.BigEndian.Uint32(avlDataPacket.CRC16)

	if avlDataArray.NumberOfData1 != avlDataArray.NumberOfData2 {
		return avlDataArray // Return empty array if the counts do not match
	}

	c.Body = avlDataPacket.AVLdata
	parser := parsers.DataParser{Body: c.Body}

	startIndex := 0

	for i := 0; i < int(avlDataArray.NumberOfData1); i++ {
		avlData := models.AvlData{}

		var timestampEndIndex int
		avlData.Timestamp, timestampEndIndex = parser.ParseTimestamp(startIndex)

		var priorityIndex int
		avlData.Priority, priorityIndex = parser.ParsePriority(timestampEndIndex)

		var gpsEndIndex int
		avlData.GpsElement, gpsEndIndex = parser.ParseGpsElement(priorityIndex)

		var eventIOIDIndexEnd int
		avlData.EventIOId, eventIOIDIndexEnd = parser.ParseEventIO(gpsEndIndex, false)

		var noOfTotalIOIndexEnd int
		avlData.NoOfTotalIO, noOfTotalIOIndexEnd = c.parseTotalNumberOfIO(eventIOIDIndexEnd)

		var oneByteIOEndIndex int
		avlData.NoOfOneByte, avlData.OneByteIO, oneByteIOEndIndex = c.parseOneByteIO(noOfTotalIOIndexEnd)

		var twoByteIOEndIndex int
		avlData.NoOfTwoByte, avlData.TwoByteIO, twoByteIOEndIndex = c.parseTwoByteIO(oneByteIOEndIndex)

		var fourByteIOEndIndex int
		avlData.NoOfFourByte, avlData.FourByteIO, fourByteIOEndIndex = c.parseFourByteIO(twoByteIOEndIndex)

		var eightByteIOEndIndex int
		avlData.NoOfEightByte, avlData.EightByteIO, eightByteIOEndIndex = c.parseEightByteIO(fourByteIOEndIndex)

		startIndex = eightByteIOEndIndex

		avlDataArray.AVLData = append(avlDataArray.AVLData, avlData)
	}

	return avlDataArray
}

/*

-------------------- Functions and utility types  --------------------
Here are the structures, functions useful for the operations

*/

// parseTotalNumberOfIO Total Number of IO.
//
// A total number of properties coming with record (N = N1 + N2 + N4 + N8).
func (c *Codec8) parseTotalNumberOfIO(startIndex int) (uint16, int) {
	noOfTotalIO := c.Body[startIndex]

	return uint16(noOfTotalIO), startIndex + 1
}

// generic parseIo function
func (c *Codec8) parseIo(startIndex int, splitByte int) (uint8, map[uint8]interface{}, int) {
	IoElements := map[uint8]interface{}{}

	if startIndex >= len(c.Body) {
		// No IO elements found
		return 0, IoElements, len(c.Body)
	}

	// One Byte IO Number
	nOfIoElements := c.Body[startIndex]
	// One Byte IO Data
	IoElementsStartIndex := startIndex + 1
	IoElementsEndIndex := IoElementsStartIndex + int(nOfIoElements)*splitByte
	data := c.Body[IoElementsStartIndex:IoElementsEndIndex]

	for i := 0; i < len(data); i += splitByte {
		id := data[i]
		value := data[i+1]

		IoElements[id] = value
	}
	return nOfIoElements, IoElements, IoElementsEndIndex

}

// ParseOneByteIO This function parse one byte IO.
//
// N1
// number of properties, which length is 1 byte.
func (c *Codec8) parseOneByteIO(startIndex int) (uint8, map[uint8]interface{}, int) {
	splitByte := 2 // ID 1 byte + VALUE 1 byte
	return c.parseIo(startIndex, splitByte)
}

// ParseTwoByteIO This function parse two byte IO.
//
// N2
// number of properties, which length is 2 byte.
func (c *Codec8) parseTwoByteIO(startIndex int) (uint8, map[uint8]interface{}, int) {
	splitByte := 3 // ID 1 byte + VALUE 2 bytes
	return c.parseIo(startIndex, splitByte)
}

// ParseFourByteIO This function parse four byte IO.
//
// N4
// number of properties, which length is 4 byte.
func (c *Codec8) parseFourByteIO(startIndex int) (uint8, map[uint8]interface{}, int) {
	splitByte := 5 // ID 1 byte + VALUE 4 bytes
	return c.parseIo(startIndex, splitByte)
}

// ParseEightByteIO This function parse eight byte IO.
//
// N8
// number of properties, which length is 8 byte.
// Eight Byte IO Number
func (c *Codec8) parseEightByteIO(startIndex int) (uint8, map[uint8]interface{}, int) {
	splitByte := 9 // ID 1 byte + VALUE 8 bytes
	return c.parseIo(startIndex, splitByte)
}
