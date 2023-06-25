package codec8

import (
	parser "github.com/rojack96/teltonika-parser/common"
	"github.com/rojack96/teltonika-parser/helpers"
	"github.com/rojack96/teltonika-parser/models"
)

func Codec8Bytes(dataPacket []byte) models.AVLDataArray {
	var avlDataArray models.AVLDataArray

	avlDataPacket := helpers.DataBytesParser(&dataPacket)

	// avlDataArray.Preamble =
	avlDataArray.CodecID = avlDataPacket.CodecID
	avlDataArray.NumberOfData1 = avlDataPacket.NumberOfData1
	avlDataArray.NumberOfData2 = avlDataPacket.NumberOfData2
	// avlDataArray.CRC16 = avlDataArray

	body := avlDataPacket.AVLdata

	startIndex := 0

	if avlDataArray.NumberOfData1 == avlDataArray.NumberOfData2 {
		for i := 0; i < int(avlDataArray.NumberOfData1); i++ {
			avlData := models.AVLDataCodec8{}

			avlData.IMEI = avlDataPacket.IMEI

			var timestampEndIndex int
			avlData.Timestamp, avlData.RawData.Timestamp, timestampEndIndex = parser.ParseTimestamp(startIndex, body)

			var priorityIndex int
			avlData.Priority, priorityIndex = parser.ParsePriority(timestampEndIndex, body)
			avlData.RawData.Priority = avlData.Priority

			var gpsEndIndex int
			avlData.GPSElement, avlData.RawData.Gps, gpsEndIndex = parser.ParseGPSElement(priorityIndex, body)

			var eventIOIDIndexEnd int
			avlData.EventIOID, eventIOIDIndexEnd = parseEventIO(gpsEndIndex, body)
			avlData.RawData.EventIOID = uint(avlData.EventIOID)

			var noOfTotalIOIndexEnd int
			avlData.NoOfTotalIO, noOfTotalIOIndexEnd = parseTotalNumberOfIO(eventIOIDIndexEnd, body)

			var oneByteIOEndIndex int
			avlData.NoOfOneByte, avlData.OneByteIO, IORawData, oneByteIOEndIndex = parseOneByteIO(noOfTotalIOIndexEnd, body, IORawData)

			var twoByteIOEndIndex int
			avlData.NoOfTwoByte, avlData.TwoByteIO, IORawData, twoByteIOEndIndex = parseTwoByteIO(oneByteIOEndIndex, body, IORawData)

			var fourByteIOEndIndex int
			avlData.NoOfFourByte, avlData.FourByteIO, IORawData, fourByteIOEndIndex = parseFourByteIO(twoByteIOEndIndex, body, IORawData)

			var eightByteIOEndIndex int
			avlData.NoOfEightByte, avlData.EightByteIO, IORawData, eightByteIOEndIndex = parseEightByteIO(fourByteIOEndIndex, body, IORawData)

			startIndex = eightByteIOEndIndex

			avlDataArray.AVLDataCodec8 = append(avlDataArray.AVLDataCodec8, avlData)
		}

	}

	return avlDataArray
}
