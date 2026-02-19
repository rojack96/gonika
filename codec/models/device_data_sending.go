package models

type AvlData interface {
	isAvlData()
}

func (AvlData8) isAvlData()    {}
func (AvlData8ext) isAvlData() {}
func (AvlData16) isAvlData()   {}

type AvlDataPacketTCP struct {
	Preamble Preamble `json:"preamble"`
	// Data Field Length (4 bytes)
	//
	// size is calculated starting from Codec ID to Number of Data 2.
	DataFieldLength uint32       `json:"dataFieldLength"`
	CodecID         CodecID      `json:"codecID"`
	NumberOfData1   NumberOfData `json:"numberOfData1"`
	AvlData         []AvlData    `json:"avlData"`
	// Number of Data 2 (1 byte)
	//
	// number which defines how many records is in the packet.
	// This number must be the same as “Number of Data 1”.
	NumberOfData2 NumberOfData `json:"numberOfData2"`
	Crc16         Crc16        `json:"crc16"`
}

type AvlDataPacketUDP struct {
	// Length
	//
	// packet length (excluding this field) in big ending byte order.
	Length uint16
	// PacketID
	//
	// packet ID unique for this channel
	PacketID uint16
	// NotUsableByte
	//
	// not usable byte.
	NotUsableByte uint8
	// AvlPacketID
	//
	// ID identifying this AVL packet.
	AvlPacketID uint8
	// ImeiLength
	//
	// always will be 0x000F
	ImeiLength    uint16
	Imei          string
	CodecID       CodecID      `json:"codecID"`
	NumberOfData1 NumberOfData `json:"numberOfData1"`
	AvlData       []AvlData    `json:"avlData"`
	// Number of Data 2 (1 byte)
	//
	// number which defines how many records is in the packet.
	// This number must be the same as “Number of Data 1”.
	NumberOfData2 NumberOfData `json:"numberOfData2"`
}

type AvlData8 struct {
	Timestamp  Timestamp  `json:"timestamp"`
	Priority   Priority   `json:"priority"`
	GpsElement GpsElement `json:"gpsElement"`
	// Event IO ID (1 byte) this field defines which IO property has changed and generated an event.
	EventIOID uint8 `json:"eventIoID"`
	// Number of Total IO (1 byte) a total number of properties coming with record (N = N1 + N2 + N4 + N8).
	NoOfTotalIO uint8 `json:"numberOfTotalIO"`
	// Number of One Byte IO (1 byte) number of properties which length is 1 byte.
	NoOfOneByte uint8 `json:"numberOfOneByte"`
	// Map id:value with properties which length is 1 byte.
	OneByteIO map[uint8]string `json:"oneByteIO"`
	// Number of Two Byte IO (1 byte) number of properties which length is 2 bytes.
	NoOfTwoByte uint8 `json:"numberOfTwoByte"`
	// Map id:value with properties which length is 2 bytes.
	TwoByteIO map[uint8]string `json:"twoByteIO"`
	// Number of Four Byte IO (1 byte) number of properties which length is 4 bytes.
	NoOfFourByte uint8 `json:"numberOfFourByte"`
	// Map id:value with properties which length is 4 bytes.
	FourByteIO map[uint8]string `json:"fourByteIO"`
	// Number of Eight Byte IO (1 byte) number of properties which length is 8 bytes.
	NoOfEightByte uint8 `json:"numberOfEightByte"`
	// Map id:value with properties which length is 8 bytes.
	EightByteIO map[uint8]string `json:"eightByteIO"`
}

type AvlData8ext struct {
	Timestamp  Timestamp  `json:"timestamp"`
	Priority   Priority   `json:"priority"`
	GpsElement GpsElement `json:"gpsElement"`
	// Event IO ID (2 bytes) this field defines which IO property has changed and generated an event.
	EventIOID uint16 `json:"eventIoID"`
	// Number of Total IO (2 bytes) a total number of properties coming with record (N = N1 + N2 + N4 + N8).
	NoOfTotalIO uint16 `json:"numberOfTotalIO"`
	// Number of One Byte IO (2 bytes) number of properties which length is 1 byte.
	NoOfOneByte uint16 `json:"numberOfOneByte"`
	// Map id:value with properties which length is 1 byte.
	OneByteIO map[uint16]string `json:"oneByteIO"`
	// Number of Two Byte IO (2 bytes) number of properties which length is 2 bytes.
	NoOfTwoByte uint16 `json:"numberOfTwoByte"`
	// Map id:value with properties which length is 2 bytes.
	TwoByteIO map[uint16]string `json:"twoByteIO"`
	// Number of Four Byte IO (2 bytes) number of properties which length is 4 bytes.
	NoOfFourByte uint16 `json:"numberOfFourByte"`
	// Map id:value with properties which length is 4 bytes.
	FourByteIO map[uint16]string `json:"fourByteIO"`
	// Number of Eight Byte IO (2 bytes) number of properties which length is 8 bytes.
	NoOfEightByte uint16 `json:"numberOfEightByte"`
	// Map id:value with properties which length is 8 bytes.
	EightByteIO map[uint16]string `json:"eightByteIO"`
	// Number of X Byte IO (2 bytes)  a number of properties which length is defined by length element.
	NoOfXByte uint16 `json:"numberOfXByte"`
	// Map id:value with properties which length is defined by length element.
	XByteIO map[uint16]string `json:"xByteIO"`
}

type AvlData16 struct {
	Timestamp  Timestamp  `json:"timestamp"`
	Priority   Priority   `json:"priority"`
	GpsElement GpsElement `json:"gpsElement"`
	// Event IO ID (1 byte) this field defines which IO property has changed and generated an event.
	EventIOID uint16 `json:"eventIoID"`
	// Data event generation type
	GenerationType uint8 `json:"generationType"`
	// Number of Total IO (1 byte) a total number of properties coming with record (N = N1 + N2 + N4 + N8).
	NoOfTotalIO uint8 `json:"numberOfTotalIO"`
	// Number of One Byte IO (1 byte) number of properties which length is 1 byte.
	NoOfOneByte uint8 `json:"numberOfOneByte"`
	// Map id:value with properties which length is 1 byte.
	OneByteIO map[uint16]string `json:"oneByteIO"`
	// Number of Two Byte IO (1 byte) number of properties which length is 2 bytes.
	NoOfTwoByte uint8 `json:"numberOfTwoByte"`
	// Map id:value with properties which length is 2 bytes.
	TwoByteIO map[uint16]string `json:"twoByteIO"`
	// Number of Four Byte IO (1 byte) number of properties which length is 4 bytes.
	NoOfFourByte uint8 `json:"numberOfFourByte"`
	// Map id:value with properties which length is 4 bytes.
	FourByteIO map[uint16]string `json:"fourByteIO"`
	// Number of Eight Byte IO (1 byte) number of properties which length is 8 bytes.
	NoOfEightByte uint8 `json:"numberOfEightByte"`
	// Map id:value with properties which length is 8 bytes.
	EightByteIO map[uint16]string `json:"eightByteIO"`
}
