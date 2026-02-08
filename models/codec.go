package models

type AvlDataPacket struct {
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

type CodecMessage struct {
	// Preamble (4 bytes) the packet starts with four zero bytes.
	Preamble []byte `json:"preamble"`
	// Data Size (4 bytes)  size is calculated from Codec ID field to the second command or response quantity field.
	DataSize []byte `json:"data_size"`
	// Codec ID (1 byte)
	CodecID byte `json:"codec_id"`
	// Type (1 byte) it can be 0x05 to denote command or 0x06 to denote response.
	Type byte `json:"type"`
	// CRC-16 (4 bytes) calculated from Codec ID to the Second Number of Data.
	// CRC (Cyclic Redundancy Check) is an error-detecting code using for detect accidental changes to RAW data.
	// For calculation we are using CRC-16/IBM.
	CRC16 []byte `json:"crc_16"`
}

type CodecResponse struct {
	// Command Quantity 1 (1 byte) it is ignored when parsing the message.
	ResponseQuantity1 byte `json:"response_quantity_1"`
	// Command (X bytes) command or response in HEX.
	Response []byte `json:"response"`
	// Command Quantity 2 (1 byte) a byte which defines how many records (commands or responses) is in the packet.
	// This byte will not be parsed but it’s recommended that it should contain same value as Command/Response Quantity 1.
	ResponseQuantity2 byte `json:"response_quantity_2"`
}
