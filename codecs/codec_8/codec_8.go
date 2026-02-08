package codec8

import (
	"github.com/rojack96/gonika/codecs/utils"
	"github.com/rojack96/gonika/models"
	"github.com/rojack96/gonika/parsers"
)

type Codec8 struct{ avlDataPacket []byte }

func New(avlDataPacket []byte) *Codec8 {
	return &Codec8{avlDataPacket: avlDataPacket}
}

func (c *Codec8) Decode() *models.AvlDataArray {
	var result models.AvlDataArray

	data := utils.DataMapping(c.avlDataPacket)

	result.Preamble = parsers.Preamble(data.Preamble)
	result.CodecID = parsers.CodecId(data.CodecID)
	result.DataFieldLength = parsers.DataFieldLength(data.DataFieldLength)
	result.NumberOfData1 = parsers.NumberOfData(data.NumberOfData1)
	result.NumberOfData2 = parsers.NumberOfData(data.NumberOfData2)
	result.Crc16 = parsers.Crc16(data.Crc16)

	body := data.Avldata

	startIndex := 0
	//index := 0

	if result.NumberOfData1 != result.NumberOfData2 {
		return nil
	}

	for i := 0; i < int(result.NumberOfData1); i++ {
		avlData := models.AvlData8{}

		var timestampEndIndex int
		avlData.Timestamp, timestampEndIndex = parsers.Timestamp(startIndex, body)
		//index += 8

		var priorityIndex int
		avlData.Priority, priorityIndex = parsers.Priority(timestampEndIndex, body)
		//index += 1

		var gpsEndIndex int
		avlData.GpsElement, gpsEndIndex = parsers.GpsElement(priorityIndex, body)
		// index += 15

		var eventIOIDIndexEnd int
		avlData.EventIOID, eventIOIDIndexEnd = c.parseEventIO(gpsEndIndex, body)
		//index += 1

		var noOfTotalIOIndexEnd int
		avlData.NoOfTotalIO, noOfTotalIOIndexEnd = c.parseTotalNumberOfIO(eventIOIDIndexEnd, body)
		//index += 1

		var oneByteIOEndIndex int
		avlData.NoOfOneByte, avlData.OneByteIO, oneByteIOEndIndex = c.parseIo(1, noOfTotalIOIndexEnd, body)

		var twoByteIOEndIndex int
		avlData.NoOfTwoByte, avlData.TwoByteIO, twoByteIOEndIndex = c.parseIo(2, oneByteIOEndIndex, body)

		var fourByteIOEndIndex int
		avlData.NoOfFourByte, avlData.FourByteIO, fourByteIOEndIndex = c.parseIo(4, twoByteIOEndIndex, body)

		var eightByteIOEndIndex int
		avlData.NoOfEightByte, avlData.EightByteIO, eightByteIOEndIndex = c.parseIo(8, fourByteIOEndIndex, body)

		startIndex = eightByteIOEndIndex

		result.AvlData = append(result.AvlData, avlData)
	}

	return &result
}
