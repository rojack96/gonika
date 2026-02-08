package codec8e

import (
	"github.com/rojack96/gonika/codecs/utils"
	"github.com/rojack96/gonika/models"
	"github.com/rojack96/gonika/parsers"
)

type Codec8e struct{ avlDataPacket []byte }

func New(avlDataPacket []byte) *Codec8e {
	return &Codec8e{avlDataPacket: avlDataPacket}
}

func (c *Codec8e) Decode() *models.AvlDataArray {
	var result models.AvlDataArray

	data := utils.DataMapping(c.avlDataPacket)

	result.Preamble = parsers.Preamble(data.Preamble)
	result.CodecID = parsers.CodecId(data.CodecID)
	result.DataFieldLength = parsers.DataFieldLength(data.DataFieldLength)
	result.NumberOfData1 = parsers.NumberOfData(data.NumberOfData1)
	result.NumberOfData2 = parsers.NumberOfData(data.NumberOfData2)
	result.Crc16 = parsers.Crc16(data.Crc16)

	body := data.Avldata

	index := 0

	if result.NumberOfData1 != result.NumberOfData2 {
		return nil
	}

	for i := 0; i < int(result.NumberOfData1); i++ {
		avlData := models.AvlData8ext{}

		avlData.Timestamp, index = parsers.Timestamp(index, body)
		avlData.Priority, index = parsers.Priority(index, body)
		avlData.GpsElement, index = parsers.GpsElement(index, body)

		avlData.EventIOID = c.parseEventIO(index, body)
		index += 2

		avlData.NoOfTotalIO = c.parseTotalNumberOfIO(index, body)
		index += 2

		avlData.NoOfOneByte, avlData.OneByteIO, index = c.parseIo(1, index, body)
		avlData.NoOfTwoByte, avlData.TwoByteIO, index = c.parseIo(2, index, body)
		avlData.NoOfFourByte, avlData.FourByteIO, index = c.parseIo(4, index, body)
		avlData.NoOfFourByte, avlData.FourByteIO, index = c.parseIo(4, index, body)
		avlData.NoOfEightByte, avlData.EightByteIO, index = c.parseIo(8, index, body)
		avlData.NoOfXByte, avlData.XByteIO, index = c.parseXByteIO(index, body)

		result.AvlData = append(result.AvlData, avlData)
	}

	return &result
}
