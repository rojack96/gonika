package codec8e

import (
	"encoding/binary"
	"github.com/rojack96/gonika/parsers"
)

type Codec8e struct{}

func (c *Codec8e) Decode(dataPacket []byte) AVLDataArray {
	var avlDataArray AVLDataArray

	avlDataPacket := parsers.DataParser(&dataPacket)

	avlDataArray.Preamble = binary.BigEndian.Uint32(avlDataPacket.Preamble)
	avlDataArray.CodecID = avlDataPacket.CodecID
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
			avlData.GPSElement, gpsEndIndex = parsers.ParseGPSElement(priorityIndex, body)

			var eventIOIDIndexEnd int
			avlData.EventIOID, eventIOIDIndexEnd = c.parseEventIO(gpsEndIndex, body)

			var noOfTotalIOIndexEnd int
			avlData.NoOfTotalIO, noOfTotalIOIndexEnd = c.parseTotalNumberOfIO(eventIOIDIndexEnd, body)

			var oneByteIOEndIndex int
			avlData.NoOfOneByte, avlData.OneByteIO, oneByteIOEndIndex = c.parseOneByteIO(noOfTotalIOIndexEnd, body)

			var twoByteIOEndIndex int
			avlData.NoOfTwoByte, avlData.TwoByteIO, twoByteIOEndIndex = c.parseTwoByteIO(oneByteIOEndIndex, body)

			var fourByteIOEndIndex int
			avlData.NoOfFourByte, avlData.FourByteIO, fourByteIOEndIndex = c.parseFourByteIO(twoByteIOEndIndex, body)

			var eightByteIOEndIndex int
			avlData.NoOfEightByte, avlData.EightByteIO, eightByteIOEndIndex = c.parseEightByteIO(fourByteIOEndIndex, body)

			var xByteIOEndIndex int
			avlData.NoOfXByte, avlData.XByteIO, xByteIOEndIndex = c.parseXByteIO(eightByteIOEndIndex, body)

			startIndex = xByteIOEndIndex

			avlDataArray.AVLData = append(avlDataArray.AVLData, avlData)
		}

	}

	return avlDataArray
}
