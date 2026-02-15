package models

type CodeSpecificMapperParam interface {
	isSpecificMapperParam()
}

func (Codec13) isSpecificMapperParam() {}
func (Codec14) isSpecificMapperParam() {}

type CommandMessage struct {
	// Preamble (4 bytes) the packet starts with four zero bytes.
	Preamble []byte
	// Data Size (4 bytes)  size is calculated from Codec ID field to the second command or response quantity field.
	DataSize []byte
	// Codec ID (1 byte)
	CodecID byte
	// Command Quantity 1 (1 byte) it is ignored when parsing the message.
	CommandQuantity1 byte
	// Type (1 byte) it can be 0x05 to denote command or 0x06 to denote response.
	Type byte
	// Command Size (4 bytes) command or response length.
	CommandSize []byte
	CodeSpecificMapperParam
	// Command (X bytes) command or response in HEX.
	Command []byte
	// Command Quantity 2 (1 byte) a byte which defines how many records (commands or responses) is in the packet.
	// This byte will not be parsed but it’s recommended that it should contain same value as Command/Response Quantity 1.
	CommandQuantity2 byte
	// CRC-16 (4 bytes) calculated from Codec ID to the Second Number of Data.
	// CRC (Cyclic Redundancy Check) is an error-detecting code using for detect accidental changes to RAW data.
	// For calculation, we are using CRC-16/IBM.
	Crc16 []byte
}

type ResponseMessage struct {
	// Preamble (4 bytes) the packet starts with four zero bytes.
	Preamble []byte
	// Data Size (4 bytes)  size is calculated from Codec ID field to the second command or response quantity field.
	DataSize []byte
	// Codec ID (1 byte)
	CodecID byte
	// Command Quantity 1 (1 byte) it is ignored when parsing the message.
	ResponseQuantity1 byte
	// Type (1 byte) it can be 0x05 to denote command or 0x06 to denote response.
	Type byte
	// Command Size (4 bytes) command or response length.
	ResponseSize []byte
	CodeSpecificMapperParam
	// Command (X bytes) command or response in HEX.
	Response []byte
	// Command Quantity 2 (1 byte) a byte which defines how many records (commands or responses) is in the packet.
	// This byte will not be parsed but it’s recommended that it should contain same value as Command/Response Quantity 1.
	ResponseQuantity2 byte
	// CRC-16 (4 bytes) calculated from Codec ID to the Second Number of Data.
	// CRC (Cyclic Redundancy Check) is an error-detecting code using for detect accidental changes to RAW data.
	// For calculation, we are using CRC-16/IBM.
	Crc16 []byte
}

type Codec13 struct {
	// Timestamp (4 bytes) a difference, in seconds, between the current time and midnight, January, 1970 UTC (UNIX time).
	Timestamp [4]byte
}

type Codec14 struct {
	// Imei (8 bytes) it is send as HEX value. Example if device IMEI is 123456789123456 then IMEI data field will contain 0x0123456789123456 value.
	Imei [8]byte
}
