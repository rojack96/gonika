package models

import (
	"time"
)

type AVLDataPacket struct {
	// Premble (4 bytes)
	//
	// the packet starts with four zero bytes.
	Preamble []byte `json:"preamble"`
	// Data Field Length (4 bytes)
	//
	// size is calculated starting from Codec ID to Number of Data 2.
	DataFieldLength []byte `json:"data_field_length"`
	// Codec ID (1 byte)
	CodecID byte `json:"codec_id"`
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

type GPSElement struct {
	// Longitude (4 bytes) east – west position.
	Longitude int32 `json:"lon"`
	// Latitude (4 bytes) north – south position.
	Latitude int32 `json:"lat"`
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

type AVLDataArray struct {
	// Premble (4 bytes)
	//
	// the packet starts with four zero bytes.
	Preamble uint32 `json:"preamble"`
	// Data Field Length (4 bytes)
	//
	// size is calculated starting from Codec ID to Number of Data 2.
	DataFieldLength uint32 `json:"data_field_length"`
	// Codec ID (1 byte)
	CodecID uint8 `json:"codec_id"`
	// Number of Data 1 (1 byte)
	//
	// a number which defines how many records is in the packet.
	NumberOfData1 uint8 `json:"no_of_data_1"`
	// Number of Data 12 (1 byte)
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

type AVLDataTsToGps struct {
	// Timestamp (8 bytes)
	//
	// A difference, in milliseconds,
	// between the current time and midnight, January, 1970 UTC (UNIX time).
	Timestamp time.Time `json:"timestamp"`
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
	GPSElement GPSElement `json:"gps"`
}
