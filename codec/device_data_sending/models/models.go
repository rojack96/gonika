package models

type AvlDataPacketHeader struct {
	// Preamble (4 bytes)
	//
	// the packet starts with four zero bytes.
	Preamble [4]byte
	// Data Field Length (4 bytes)
	//
	// size is calculated starting from Codec ID to Number of Data 2.
	DataFieldLength [4]byte
}

type AvlDataArray struct {
	// Codec ID (1 byte)
	CodecID byte
	// Number of Data 1 (1 byte)
	//
	// a number which defines how many records is in the packet.
	NumberOfData1 byte
	// AVL Data (X bytes)
	//
	// actual data in the packet.
	AvlData []byte
	// Number of Data 2 (1 byte)
	//
	// number which defines how many records is in the packet.
	// This number must be the same as “Number of Data 1”.
	NumberOfData2 byte
}

type AvlDataPacketByte struct {
	AvlDataPacketHeader AvlDataPacketHeader
	AvlDataArray        AvlDataArray
	// CRC-16 (4 bytes)
	//
	// calculated from Codec ID to the Second Number of Data.
	// CRC (Cyclic Redundancy Check) is an error-detecting code using for detect accidental changes to RAW data.
	// For calculation we are using CRC-16/IBM.
	Crc16 [4]byte
}

type UdpChannelHeader struct {
	Length        [2]byte
	PacketId      [2]byte
	NotUsableByte byte
}

type UdpAvlPacketHeader struct {
	AvlPacketId byte
	ImeiLength  [2]byte
	Imei        [15]byte
}

type UdpAvlDataPacketByte struct {
	UdpChannelHeader   UdpChannelHeader
	UdpAvlPacketHeader UdpAvlPacketHeader
	AvlDataArray       AvlDataArray
}
