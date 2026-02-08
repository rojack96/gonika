package models

// Preamble (4 bytes)
//
// the packet starts with four zero bytes.
type Preamble uint32

// Data Field Length (4 bytes)
//
// size is calculated starting from Codec ID to Number of Data 2.
type DataFieldLength uint32

// Codec ID (1 byte)
type CodecID uint8

// Number of Data
//
// a number which defines how many records is in the packet.
type NumberOfData uint8

// CRC-16 (4 bytes)
//
// calculated from Codec ID to the Second Number of Data.
// CRC (Cyclic Redundancy Check) is an error-detecting code using for detect accidental changes to RAW data.
// For calculation we are using CRC-16/IBM.
type Crc16 uint32

// Timestamp (8 bytes)
//
// A difference, in milliseconds,
// between the current time and midnight, January, 1970 UTC (UNIX time).
type Timestamp uint64

// Priority (1 byte)
//
// Field which define AVL data priority.
// 0 -> Low
// 1 -> High
// 2 -> Panic
type Priority uint8

// Longitude (4 bytes) east – west position.
type Longitude int32

// Latitude (4 bytes) north – south position.
type Latitude int32

// Altitude (2 bytes) meters above sea level.
type Altitude uint16

// Angle (2 bytes) degrees from north pole.
type Angle uint16

// Satellites (1 byte)  number of visible satellites.
type Satellites uint8

// Speed (2 bytes) speed calculated from satellites.
// Speed will be 0x0000 if GPS data is invalid.
type Speed uint16

// GPS (15 bytes)
//
// Location information of the AVL data.
type GpsElement struct {
	Longitude  `json:"longitude"`
	Latitude   `json:"latitude"`
	Altitude   `json:"altitude"`
	Angle      `json:"angle"`
	Satellites `json:"satellites"`
	Speed      `json:"speed"`
}
