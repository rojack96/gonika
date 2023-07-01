package codec14

import "github.com/rojack96/teltonika-parser/models"

type CommandMessage struct {
	// Premble (4 bytes) the packet starts with four zero bytes.
	Preamble string `json:"preamble"`
	// Data Size (4 bytes)  size is calculated from Codec ID field to the second command or response quantity field.
	DataSize string `json:"data_size"`
	// Codec ID (1 byte)
	CodecID string `json:"codec_id"`
	// Command Quantity 1 (1 byte) it is ignored when parsing the message.
	CommandQuantity1 string `json:"command_quantity_1"`
	// Type (1 byte) it can be 0x05 to denote command or 0x06 to denote response.
	Type string `json:"type"`
	// Command Size (4 bytes) command or response length.
	CommandAndImeiSize string `json:"command_and_imei_size"`
	// IMEI
	IMEI string `json:"imei"`
	// Command (X bytes) command or response in HEX.
	Command string `json:"command"`
	// Command Quantity 2 (1 byte) a byte which defines how many records (commands or responses) is in the packet.
	// This byte will not be parsed but it’s recommended that it should contain same value as Command/Response Quantity 1.
	CommandQuantity2 string `json:"command_quantity_2"`
	// CRC-16 (4 bytes) calculated from Codec ID to the Second Number of Data.
	// CRC (Cyclic Redundancy Check) is an error-detecting code using for detect accidental changes to RAW data.
	// For calculation we are using CRC-16/IBM.
	CRC16 string `json:"crc_16"`
}

type ResponseMessage struct {
	models.CodecMessage
	models.CodecResponse
	// Response Size (4 bytes) command or response length.
	ResponseAndImeiSize []byte `json:"response_and_imei_size"`
	// IMEI
	IMEI []byte `json:"imei"`
}
