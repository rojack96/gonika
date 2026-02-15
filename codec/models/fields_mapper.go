package models

// Preamble (4 bytes) the packet starts with four zero bytes.
type PreambleByte [4]byte

// Data Size (4 bytes)  size is calculated from Codec ID field to the second command or response quantity field.
type DataSizeByte [4]byte

// Codec ID (1 byte)
type CodecIDByte byte

// Command/Response Quantity 1 (1 byte) it is ignored when parsing the message.
type QuantityByte byte

// Type (1 byte) it can be 0x05 to denote command or 0x06 to denote response.
type TypeByte byte

// Command/Response Size (4 bytes) command or response length.
type SizeByte [4]byte

// CRC-16 (4 bytes) calculated from Codec ID to the Second Number of Data.
// CRC (Cyclic Redundancy Check) is an error-detecting code using for detect accidental changes to RAW data.
// For calculation, we are using CRC-16/IBM.
type Crc16Byte [4]byte
