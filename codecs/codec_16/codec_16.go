package codec16

import (
	"github.com/rojack96/gonika/codecs/utils"
	"github.com/rojack96/gonika/models"
	"github.com/rojack96/gonika/parsers"
)

type AvlDataArray struct {
	models.Preamble        `json:"preamble"`
	models.DataFieldLength `json:"dataFieldLength"`
	models.CodecID         `json:"codecId"`
	NumberOfData1          models.NumberOfData `json:"numberOfData1"`
	AvlData                []AvlData           `json:"avlData"`
	// Number of Data 2 (1 byte)
	//
	// number which defines how many records is in the packet.
	// This number must be the same as “Number of Data 1”.
	NumberOfData2 models.NumberOfData `json:"numberOfData2"`
	models.Crc16  `json:"crc16"`
}

type AvlData struct {
	models.Timestamp  `json:"timestamp"`
	models.Priority   `json:"priority"`
	models.GpsElement `json:"gpsElement"`
	// Event IO ID (1 byte) this field defines which IO property has changed and generated an event.
	EventIOID uint16 `json:"eventIoID"`
	// Data event generation type
	GenerationType uint8 `json:"generationType"`
	// Number of Total IO (1 byte) a total number of properties coming with record (N = N1 + N2 + N4 + N8).
	NoOfTotalIO uint8 `json:"numberOfTotalIO"`
	// Number of One Byte IO (1 byte) number of properties which length is 1 byte.
	NoOfOneByte uint8 `json:"numberOfOneByte"`
	// Map id:value with properties which length is 1 byte.
	OneByteIO map[uint16]uint8 `json:"oneByteIO"`
	// Number of Two Byte IO (1 byte) number of properties which length is 2 bytes.
	NoOfTwoByte uint8 `json:"numberOfTwoByte"`
	// Map id:value with properties which length is 2 bytes.
	TwoByteIO map[uint16]uint16 `json:"twoByteIO"`
	// Number of Four Byte IO (1 byte) number of properties which length is 4 bytes.
	NoOfFourByte uint8 `json:"numberOfFourByte"`
	// Map id:value with properties which length is 4 bytes.
	FourByteIO map[uint16]uint32 `json:"fourByteIO"`
	// Number of Eight Byte IO (1 byte) number of properties which length is 8 bytes.
	NoOfEightByte uint8 `json:"numberOfEightByte"`
	// Map id:value with properties which length is 8 bytes.
	EightByteIO map[uint16]uint64 `json:"eightByteIO"`
}

type Codec16 struct{ avlDataPacket []byte }

func New(avlDataPacket []byte) *Codec16 {
	return &Codec16{avlDataPacket: avlDataPacket}
}

func (c *Codec16) Decode() any {
	var result AvlDataArray

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
		avlData := AvlData{}

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
		//index += 2

		var generationTypeIndexEnd int
		avlData.GenerationType, generationTypeIndexEnd = c.parseGenerationType(eventIOIDIndexEnd, body)
		//index += 1

		var noOfTotalIOIndexEnd int
		avlData.NoOfTotalIO, noOfTotalIOIndexEnd = c.parseTotalNumberOfIO(generationTypeIndexEnd, body)
		//index += 1

		var oneByteIOEndIndex int
		avlData.NoOfOneByte, avlData.OneByteIO, oneByteIOEndIndex = c.parseOneByteIO(noOfTotalIOIndexEnd, body)

		var twoByteIOEndIndex int
		avlData.NoOfTwoByte, avlData.TwoByteIO, twoByteIOEndIndex = c.parseTwoByteIO(oneByteIOEndIndex, body)

		var fourByteIOEndIndex int
		avlData.NoOfFourByte, avlData.FourByteIO, fourByteIOEndIndex = c.parseFourByteIO(twoByteIOEndIndex, body)

		var eightByteIOEndIndex int
		avlData.NoOfEightByte, avlData.EightByteIO, eightByteIOEndIndex = c.parseEightByteIO(fourByteIOEndIndex, body)

		startIndex = eightByteIOEndIndex

		result.AvlData = append(result.AvlData, avlData)
	}

	return result
}
