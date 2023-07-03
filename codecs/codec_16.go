package codecs

import (
	"encoding/binary"

	models "github.com/rojack96/gotlk/models/codec_16"
)

func c16AVLData(dataPacket []byte) models.AVLDataArray {
	var avlDataArray models.AVLDataArray

	avlDataPacket := dataParser(&dataPacket)

	avlDataArray.Preamble = binary.BigEndian.Uint32(avlDataPacket.Preamble)
	avlDataArray.CodecID = avlDataPacket.CodecID
	avlDataArray.NumberOfData1 = avlDataPacket.NumberOfData1
	avlDataArray.NumberOfData2 = avlDataPacket.NumberOfData2
	avlDataArray.CRC16 = binary.BigEndian.Uint32(avlDataPacket.CRC16)

	body := avlDataPacket.AVLdata

	startIndex := 0

	if avlDataArray.NumberOfData1 == avlDataArray.NumberOfData2 {
		for i := 0; i < int(avlDataArray.NumberOfData1); i++ {
			avlData := models.AVLData{}

			var timestampEndIndex int
			avlData.Timestamp, timestampEndIndex = parseTimestamp(startIndex, body)

			var priorityIndex int
			avlData.Priority, priorityIndex = parsePriority(timestampEndIndex, body)

			var gpsEndIndex int
			avlData.GPSElement, gpsEndIndex = parseGPSElement(priorityIndex, body)

			var eventIOIDIndexEnd int
			avlData.EventIOID, eventIOIDIndexEnd = c16ParseEventIO(gpsEndIndex, body)

			var generationTypeIndexEnd int
			avlData.GenerationType, generationTypeIndexEnd = c16ParseGenerationType(eventIOIDIndexEnd, body)

			var noOfTotalIOIndexEnd int
			avlData.NoOfTotalIO, noOfTotalIOIndexEnd = c16ParseTotalNumberOfIO(generationTypeIndexEnd, body)

			var oneByteIOEndIndex int
			avlData.NoOfOneByte, avlData.OneByteIO, oneByteIOEndIndex = c16ParseOneByteIO(noOfTotalIOIndexEnd, body)

			var twoByteIOEndIndex int
			avlData.NoOfTwoByte, avlData.TwoByteIO, twoByteIOEndIndex = c16ParseTwoByteIO(oneByteIOEndIndex, body)

			var fourByteIOEndIndex int
			avlData.NoOfFourByte, avlData.FourByteIO, fourByteIOEndIndex = c16ParseFourByteIO(twoByteIOEndIndex, body)

			var eightByteIOEndIndex int
			avlData.NoOfEightByte, avlData.EightByteIO, eightByteIOEndIndex = c16ParseEightByteIO(fourByteIOEndIndex, body)

			startIndex = eightByteIOEndIndex

			avlDataArray.AVLData = append(avlDataArray.AVLData, avlData)
		}
	}

	return avlDataArray
}
