package codec8

import (
	"github.com/rojack96/gonika/codec/models"
	"github.com/rojack96/gonika/codec/parsers"
	"github.com/rojack96/gonika/codec/utils"
)

type Codec8 struct{ avlDataPacket []byte }

func New(avlDataPacket []byte) *Codec8 {
	return &Codec8{avlDataPacket: avlDataPacket}
}

func (c *Codec8) Decode() *models.AvlDataPacket {
	var result models.AvlDataPacket

	data := utils.DataMapping(c.avlDataPacket)

	result.Preamble = parsers.Preamble(data.Preamble)
	result.DataFieldLength = parsers.DataFieldLength(data.DataFieldLength)
	result.CodecID = parsers.CodecId(data.CodecID)
	result.NumberOfData1 = parsers.NumberOfData(data.NumberOfData1)
	result.NumberOfData2 = parsers.NumberOfData(data.NumberOfData2)
	result.Crc16 = parsers.Crc16(data.Crc16)

	body := data.Avldata

	index := 0

	if result.NumberOfData1 != result.NumberOfData2 {
		return nil
	}

	for i := 0; i < int(result.NumberOfData1); i++ {
		avl := models.AvlData8{}

		avl.Timestamp, index = parsers.Timestamp(index, body)
		avl.Priority, index = parsers.Priority(index, body)
		avl.GpsElement, index = parsers.GpsElement(index, body)
		avl.EventIOID, index = c.parseEventIO(index, body)
		avl.NoOfTotalIO, index = c.parseTotalNumberOfIO(index, body)
		avl.NoOfOneByte, avl.OneByteIO, index = c.parseIo(1, index, body)
		avl.NoOfTwoByte, avl.TwoByteIO, index = c.parseIo(2, index, body)
		avl.NoOfFourByte, avl.FourByteIO, index = c.parseIo(4, index, body)
		avl.NoOfEightByte, avl.EightByteIO, index = c.parseIo(8, index, body)

		result.AvlData = append(result.AvlData, avl)
	}

	return &result
}
