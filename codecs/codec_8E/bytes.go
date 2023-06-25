package codec8e

import (
	"encoding/binary"

	"github.com/rojack96/teltonika-parser/helpers"
	models "github.com/rojack96/teltonika-parser/models/codec_8E"
)

func Codec8EBytes(dataPacket []byte) models.AVLDataArray {
	var avlDataArray models.AVLDataArray

	avlDataPacket := helpers.DataBytesParser(&dataPacket)

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
			avlData.EventIOID, eventIOIDIndexEnd = parseEventIO(gpsEndIndex, body)

			var noOfTotalIOIndexEnd int
			avlData.NoOfTotalIO, noOfTotalIOIndexEnd = parseTotalNumberOfIO(eventIOIDIndexEnd, body)

			var oneByteIOEndIndex int
			avlData.NoOfOneByte, avlData.OneByteIO, oneByteIOEndIndex = parseOneByteIO(noOfTotalIOIndexEnd, body)

			var twoByteIOEndIndex int
			avlData.NoOfTwoByte, avlData.TwoByteIO, twoByteIOEndIndex = parseTwoByteIO(oneByteIOEndIndex, body)

			var fourByteIOEndIndex int
			avlData.NoOfFourByte, avlData.FourByteIO, fourByteIOEndIndex = parseFourByteIO(twoByteIOEndIndex, body)

			var eightByteIOEndIndex int
			avlData.NoOfEightByte, avlData.EightByteIO, eightByteIOEndIndex = parseEightByteIO(fourByteIOEndIndex, body)

			var xByteIOEndIndex int
			avlData.NoOfXByte, avlData.XByteIO, xByteIOEndIndex = parseXByteIO(eightByteIOEndIndex, body)

			startIndex = xByteIOEndIndex

			avlDataArray.AVLData = append(avlDataArray.AVLData, avlData)
		}

	}

	return avlDataArray
}
