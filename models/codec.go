package models

type AvlDataPacket struct {
	// Preamble (4 bytes)
	//
	// the packet starts with four zero bytes.
	Preamble []byte `json:"preamble"`
	// Data Field Length (4 bytes)
	//
	// size is calculated starting from Codec ID to Number of Data 2.
	DataFieldLength []byte `json:"data_field_length"`
	// Codec ID (1 byte)
	CodecId byte `json:"codec_id"`
	// Number of Data 1 (1 byte)
	//
	// a number which defines how many records is in the packet.
	NumberOfData1 byte `json:"no_of_data_1"`
	// AVL Data (X bytes)
	//
	// actual data in the packet.
	AVLdata []byte `json:"avl_data"`
	// Number of Data 12 (1 byte)
	//
	// number which defines how many records is in the packet.
	// This number must be the same as “Number of Data 1”.
	NumberOfData2 byte `json:"no_of_data_2"`
	// CRC-16 (4 bytes)
	//
	// calculated from Codec ID to the Second Number of Data.
	// CRC (Cyclic Redundancy Check) is an error-detecting code using for detect accidental changes to RAW data.
	// For calculation we are using CRC-16/IBM.
	CRC16 []byte `json:"crc_16"`
}

type AvlDataArray struct {
	// Preamble (4 bytes)
	//
	// the packet starts with four zero bytes.
	Preamble uint32 `json:"preamble"`
	// Data Field Length (4 bytes)
	//
	// size is calculated starting from Codec ID to Number of Data 2.
	DataFieldLength uint32 `json:"data_field_length"`
	// Codec ID (1 byte)
	CodecId uint8 `json:"codec_id"`
	// Number of Data 1 (1 byte)
	//
	// a number which defines how many records is in the packet.
	NumberOfData1 uint8     `json:"no_of_data_1"`
	AvlData       []AvlData `json:"avl_data"`
	// Number of Data 2 (1 byte)
	//
	// number which defines how many records is in the packet.
	// This number must be the same as “Number of Data 1”.
	NumberOfData2 uint8 `json:"no_of_data_2"`
	// CRC-16 (4 bytes)
	//
	// calculated from Codec ID to the Second Number of Data.
	// CRC (Cyclic Redundancy Check) is an error-detecting code using for detect accidental changes to RAW data.
	// For calculation we are using CRC-16/IBM.
	CRC16 uint32 `json:"crc_16"`
}

type AvlData struct {
	AVLDataTsToGps
	// Event IO ID (2 bytes) this field defines which IO property has changed and generated an event.
	EventIOId uint16 `json:"event_IO_id"`
	// Data event generation type
	GenerationType uint8 `json:"generation_type"`
	// Number of Total IO (2 bytes) a total number of properties coming with record (N = N1 + N2 + N4 + N8).
	NoOfTotalIO uint16 `json:"-"`
	// Number of One Byte IO (2 bytes) number of properties which length is 1 byte.
	NoOfOneByte uint16 `json:"-"`
	// Number of Two Byte IO (2 bytes) number of properties which length is 2 bytes.
	NoOfTwoByte uint16 `json:"-"`
	// Number of Four Byte IO (2 bytes) number of properties which length is 4 bytes.
	NoOfFourByte uint16 `json:"-"`
	// Number of Eight Byte IO (2 bytes) number of properties which length is 8 bytes.
	NoOfEightByte uint16 `json:"-"`
	// Number of X Byte IO (2 bytes)  a number of properties which length is defined by length element.
	NoOfXByte uint16 `json:"-"`
	// Map id:value with properties which length is 1 byte.
	OneByteIO map[uint16]interface{} `json:"-"`
	// Map id:value with properties which length is 2 bytes.
	TwoByteIO map[uint16]interface{} `json:"-"`
	// Map id:value with properties which length is 4 bytes.
	FourByteIO map[uint16]interface{} `json:"-"`
	// Map id:value with properties which length is 8 bytes.
	EightByteIO map[uint16]interface{} `json:"-"`
	// Map id:value with properties which length is defined by length element.
	XByteIO map[uint16]interface{} `json:"-"`
}

type AVLDataTsToGps struct {
	// Timestamp (8 bytes)
	//
	// A difference, in milliseconds,
	// between the current time and midnight, January, 1970 UTC (UNIX time).
	Timestamp uint64 `json:"timestamp"`
	// Priority (1 byte)
	//
	// Field which define AVL data priority.
	// 0 -> Low
	// 1 -> High
	// 2 -> Panic
	Priority uint8 `json:"priority"`
	// GPS (15 bytes)
	//
	// Location information of the AVL data.
	GpsElement GpsElement `json:"gps"`
}

type GpsElement struct {
	// Longitude (4 bytes) east – west position.
	Longitude float64 `json:"lon"`
	// Latitude (4 bytes) north – south position.
	Latitude float64 `json:"lat"`
	// Altitude (2 bytes) meters above sea level.
	Altitude uint16 `json:"alt"`
	// Angle (2 bytes) degrees from north pole.
	Angle uint16 `json:"angle"`
	// Satellites (1 byte)  number of visible satellites.
	Satellites uint8 `json:"sat"`
	// Speed (2 bytes) speed calculated from satellites.
	// Speed will be 0x0000 if GPS data is invalid.
	Speed uint16 `json:"speed"`
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
