package models

type Result struct {
	AvlDataPacket     *AvlDataPacket
	AvlDataPacketFlat *AvlDataPacketFlat
}

type AvlData interface {
	isAvlData()
}

func (AvlData8) isAvlData()    {}
func (AvlData8ext) isAvlData() {}
func (AvlData16) isAvlData()   {}

type AvlDataPacketByte struct {
	// Preamble (4 bytes)
	//
	// the packet starts with four zero bytes.
	Preamble []byte `json:"preamble"`
	// Data Field Length (4 bytes)
	//
	// size is calculated starting from Codec ID to Number of Data 2.
	DataFieldLength []byte `json:"dataFieldLength"`
	// Codec ID (1 byte)
	CodecID byte `json:"codecId"`
	// Number of Data 1 (1 byte)
	//
	// a number which defines how many records is in the packet.
	NumberOfData1 byte `json:"numberOfData1"`
	// AVL Data (X bytes)
	//
	// actual data in the packet.
	Avldata []byte `json:"avlData"`
	// Number of Data 12 (1 byte)
	//
	// number which defines how many records is in the packet.
	// This number must be the same as “Number of Data 1”.
	NumberOfData2 byte `json:"numberOfData2"`
	// CRC-16 (4 bytes)
	//
	// calculated from Codec ID to the Second Number of Data.
	// CRC (Cyclic Redundancy Check) is an error-detecting code using for detect accidental changes to RAW data.
	// For calculation we are using CRC-16/IBM.
	Crc16 []byte `json:"crc16"`
}

type AvlDataPacket struct {
	Preamble        Preamble        `json:"preamble"`
	DataFieldLength DataFieldLength `json:"dataFieldLength"`
	CodecID         CodecID         `json:"codecId"`
	NumberOfData1   NumberOfData    `json:"numberOfData1"`
	AvlData         []AvlData       `json:"avlData"`
	// Number of Data 2 (1 byte)
	//
	// number which defines how many records is in the packet.
	// This number must be the same as “Number of Data 1”.
	NumberOfData2 NumberOfData `json:"numberOfData2"`
	Crc16         Crc16        `json:"crc16"`
}

type AvlData8 struct {
	Timestamp  `json:"timestamp"`
	Priority   `json:"priority"`
	GpsElement `json:"gpsElement"`
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
	Timestamp  `json:"timestamp"`
	Priority   `json:"priority"`
	GpsElement `json:"gpsElement"`
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
	Timestamp  `json:"timestamp"`
	Priority   `json:"priority"`
	GpsElement `json:"gpsElement"`
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

type CommandMessage struct {
	// Preamble (4 bytes) the packet starts with four zero bytes.
	Preamble []byte `json:"preamble"`
	// Data Size (4 bytes)  size is calculated from Codec ID field to the second command or response quantity field.
	DataSize []byte `json:"dataSize"`
	// Codec ID (1 byte)
	CodecID byte `json:"codecId"`
	// Command Quantity 1 (1 byte) it is ignored when parsing the message.
	CommandQuantity1 byte `json:"commandQuantity1"`
	// Type (1 byte) it can be 0x05 to denote command or 0x06 to denote response.
	Type byte `json:"type"`
	// Command Size (4 bytes) command or response length.
	CommandSize []byte `json:"commandSize"`
	// Command (X bytes) command or response in HEX.
	Command []byte `json:"command"`
	// Command Quantity 2 (1 byte) a byte which defines how many records (commands or responses) is in the packet.
	// This byte will not be parsed but it’s recommended that it should contain same value as Command/Response Quantity 1.
	CommandQuantity2 byte `json:"commandQuantity2"`
	// CRC-16 (4 bytes) calculated from Codec ID to the Second Number of Data.
	// CRC (Cyclic Redundancy Check) is an error-detecting code using for detect accidental changes to RAW data.
	// For calculation, we are using CRC-16/IBM.
	Crc16 []byte `json:"crc16"`
}

type ResponseMessageByte struct {
	// Preamble (4 bytes) the packet starts with four zero bytes.
	Preamble []byte `json:"preamble"`
	// Data Size (4 bytes)  size is calculated from Codec ID field to the second command or response quantity field.
	DataSize []byte `json:"dataSize"`
	// Codec ID (1 byte)
	CodecID byte `json:"codecId"`
	// Command Quantity 1 (1 byte) it is ignored when parsing the message.
	ResponseQuantity1 byte `json:"responseQuantity1"`
	// Type (1 byte) it can be 0x05 to denote command or 0x06 to denote response.
	Type byte `json:"type"`
	// Command Size (4 bytes) command or response length.
	ResponseSize []byte `json:"responseSize"`
	// Command (X bytes) command or response in HEX.
	Response []byte `json:"response"`
	// Command Quantity 2 (1 byte) a byte which defines how many records (commands or responses) is in the packet.
	// This byte will not be parsed but it’s recommended that it should contain same value as Command/Response Quantity 1.
	ResponseQuantity2 byte `json:"responseQuantity2"`
	// CRC-16 (4 bytes) calculated from Codec ID to the Second Number of Data.
	// CRC (Cyclic Redundancy Check) is an error-detecting code using for detect accidental changes to RAW data.
	// For calculation, we are using CRC-16/IBM.
	Crc16 []byte `json:"crc16"`
}

type ResponseMessage struct {
	Preamble `json:"preamble"`
	DataSize `json:"dataSize"`
	CodecID  `json:"codecId"`
	// Command Quantity 1 (1 byte) it is ignored when parsing the message.
	ResponseQuantity1 Quantity     `json:"responseQuantity1"`
	Type              Type         `json:"type"`
	ResponseSize      ResponseSize `json:"responseSize"`
	// Command (X bytes) command or response in HEX.
	Response string `json:"response"`
	// Command Quantity 2 (1 byte) a byte which defines how many records (commands or responses) is in the packet.
	// This byte will not be parsed but it’s recommended that it should contain same value as Command/Response Quantity 1.
	ResponseQuantity2 Quantity `json:"responseQuantity2"`
	Crc16             `json:"crc16"`
}
