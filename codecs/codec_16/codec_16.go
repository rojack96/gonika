package codec16

import (
	"encoding/binary"

	"github.com/rojack96/gonika/parsers"
)

type Codec16 struct{}

func (c *Codec16) Decode(dataPacket []byte) interface{} {
	var avlDataArray AvlDataArray[uint8]

	avlDataPacket := parsers.DataParser(&dataPacket)

	avlDataArray.Preamble = binary.BigEndian.Uint32(avlDataPacket.Preamble)
	avlDataArray.CodecId = avlDataPacket.CodecId
	avlDataArray.NumberOfData1 = avlDataPacket.NumberOfData1
	avlDataArray.NumberOfData2 = avlDataPacket.NumberOfData2
	avlDataArray.CRC16 = binary.BigEndian.Uint32(avlDataPacket.CRC16)

	body := avlDataPacket.AVLdata

	startIndex := 0

	if avlDataArray.NumberOfData1 == avlDataArray.NumberOfData2 {
		for i := 0; i < int(avlDataArray.NumberOfData1); i++ {
			avlData := AVLData{}

			var timestampEndIndex int
			avlData.Timestamp, timestampEndIndex = parsers.ParseTimestamp(startIndex, body)

			var priorityIndex int
			avlData.Priority, priorityIndex = parsers.ParsePriority(timestampEndIndex, body)

			var gpsEndIndex int
			avlData.GPSElement, gpsEndIndex = parsers.ParseGpsElement(priorityIndex, body)

			var eventIOIDIndexEnd int
			avlData.EventIOID, eventIOIDIndexEnd = c.parseEventIO(gpsEndIndex, body)

			var generationTypeIndexEnd int
			avlData.GenerationType, generationTypeIndexEnd = c.parseGenerationType(eventIOIDIndexEnd, body)

			var noOfTotalIOIndexEnd int
			avlData.NoOfTotalIO, noOfTotalIOIndexEnd = c.parseTotalNumberOfIO(generationTypeIndexEnd, body)

			var oneByteIOEndIndex int
			avlData.NoOfOneByte, avlData.OneByteIO, oneByteIOEndIndex = c.parseOneByteIO(noOfTotalIOIndexEnd, body)

			var twoByteIOEndIndex int
			avlData.NoOfTwoByte, avlData.TwoByteIO, twoByteIOEndIndex = c.parseTwoByteIO(oneByteIOEndIndex, body)

			var fourByteIOEndIndex int
			avlData.NoOfFourByte, avlData.FourByteIO, fourByteIOEndIndex = c.parseFourByteIO(twoByteIOEndIndex, body)

			var eightByteIOEndIndex int
			avlData.NoOfEightByte, avlData.EightByteIO, eightByteIOEndIndex = c.parseEightByteIO(fourByteIOEndIndex, body)

			startIndex = eightByteIOEndIndex

			avlDataArray.AVLData = append(avlDataArray.AVLData, avlData)
		}
	}

	return avlDataArray
}
