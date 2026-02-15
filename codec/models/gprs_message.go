package models

type CodecSpecificParam interface {
	isSpecificParam()
}

func (Codec13) isSpecificParam() {}
func (Codec14) isSpecificParam() {}

/* ---------- Codec 12 ---------- */

type CommandMessageByte struct {
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

type CommandMessage struct {
	Preamble Preamble `json:"preamble"`
	// Data Size (4 bytes)
	//
	// size is calculated from Codec ID field to the second command or response quantity field.
	DataSize uint32  `json:"dataSize"`
	CodecID  CodecID `json:"codecId"`
	// Command Quantity 1 (1 byte) it is ignored when parsing the message.
	CommandQuantity1 Quantity `json:"commandQuantity1"`
	Type             Type     `json:"type"`
	// Command Size (4 bytes)
	//
	// command or response length.
	CommandSize uint32 `json:"commandSize"`
	CodecSpecificParam
	// Command (X bytes) command or response in HEX.
	Command string `json:"command"`
	// Command Quantity 2 (1 byte) a byte which defines how many records (commands or responses) is in the packet.
	// This byte will not be parsed but it’s recommended that it should contain same value as Command/Response Quantity 1.
	CommandQuantity2 Quantity `json:"commandQuantity2"`
	Crc16            Crc16    `json:"crc16"`
}

type ResponseMessageByte struct {
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

type ResponseMessage struct {
	Preamble Preamble `json:"preamble"`
	// Data Size (4 bytes)
	//
	// size is calculated from Codec ID field to the second command or response quantity field.
	DataSize uint32  `json:"dataSize"`
	CodecID  CodecID `json:"codecId"`
	// Command Quantity 1 (1 byte) it is ignored when parsing the message.
	ResponseQuantity1 Quantity `json:"responseQuantity1"`
	Type              Type     `json:"type"`
	// Response Size (4 bytes)
	//
	// command or response length.
	ResponseSize uint32 `json:"responseSize"`
	CodecSpecificParam
	// Command (X bytes) command or response in HEX.
	Response string `json:"response"`
	// Command Quantity 2 (1 byte) a byte which defines how many records (commands or responses) is in the packet.
	// This byte will not be parsed but it’s recommended that it should contain same value as Command/Response Quantity 1.
	ResponseQuantity2 Quantity `json:"responseQuantity2"`
	Crc16             Crc16    `json:"crc16"`
}

/* ---------- Codec 13 ---------- */

type ResponseMessageCodec13Byte struct {
	// Preamble (4 bytes) the packet starts with four zero bytes.
	Preamble []byte
	// Data Size (4 bytes)  size is calculated from Codec ID field to the second command or response quantity field.
	DataSize []byte
	// Codec ID (1 byte)
	CodecID byte
	// Response Quantity 1 (1 byte) it is ignored when parsing the message.
	ResponseQuantity1 byte
	// Type (1 byte) it can be 0x05 to denote command or 0x06 to denote response.
	Type byte
	// Response Size (4 bytes) command or response length.
	ResponseSize []byte
	// Timestamp (4 bytes) a difference, in seconds, between the current time and midnight, January, 1970 UTC (UNIX time).
	Timestamp []byte
	// Response (X bytes) command or response in HEX.
	Response []byte
	// Response Quantity 2 (1 byte) a byte which defines how many records (commands or responses) is in the packet.
	// This byte will not be parsed but it’s recommended that it should contain same value as Command/Response Quantity 1.
	ResponseQuantity2 byte
	// CRC-16 (4 bytes) calculated from Codec ID to the Second Number of Data.
	// CRC (Cyclic Redundancy Check) is an error-detecting code using for detect accidental changes to RAW data.
	// For calculation, we are using CRC-16/IBM.
	Crc16 []byte
}

type Codec13 struct {
	// Timestamp (4 bytes) a difference, in seconds, between the current time and midnight, January, 1970 UTC (UNIX time).
	Timestamp uint32 `json:"timestamp"`
}

/* ---------- Codec 14 ---------- */

type CommandMessage14Byte struct {
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
	// Note: make sure that size is IMEI size 8 + actual command size.
	// Minimal value is 8 because Codec14 always contain IMEI and it’s 8 bytes.
	CommandSize []byte
	// Imei (8 bytes) it is send as HEX value. Example if device IMEI is 123456789123456 then IMEI data field will contain 0x0123456789123456 value.
	Imei []byte
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

type ResponseMessageCodec14Byte struct {
	// Preamble (4 bytes) the packet starts with four zero bytes.
	Preamble []byte
	// Data Size (4 bytes)  size is calculated from Codec ID field to the second command or response quantity field.
	DataSize []byte
	// Codec ID (1 byte)
	CodecID byte
	// Response Quantity 1 (1 byte) it is ignored when parsing the message.
	ResponseQuantity1 byte
	// Type (1 byte) it can be 0x05 to denote command or 0x06 to denote response.
	Type byte
	// Response Size (4 bytes) command or response length.
	// Note: make sure that size is IMEI size 8 + actual command size.
	// Minimal value is 8 because Codec14 always contain IMEI and it’s 8 bytes.
	ResponseSize []byte
	// Imei (8 bytes) it is send as HEX value. Example if device IMEI is 123456789123456 then IMEI data field will contain 0x0123456789123456 value.
	Imei []byte
	// Response (X bytes) command or response in HEX.
	Response []byte
	// Response Quantity 2 (1 byte) a byte which defines how many records (commands or responses) is in the packet.
	// This byte will not be parsed but it’s recommended that it should contain same value as Command/Response Quantity 1.
	ResponseQuantity2 byte
	// CRC-16 (4 bytes) calculated from Codec ID to the Second Number of Data.
	// CRC (Cyclic Redundancy Check) is an error-detecting code using for detect accidental changes to RAW data.
	// For calculation, we are using CRC-16/IBM.
	Crc16 []byte
}

type Codec14 struct {
	// Imei (8 bytes) it is send as HEX value. Example if device IMEI is 123456789123456 then IMEI data field will contain 0x0123456789123456 value.
	Imei string `json:"imei"`
}
