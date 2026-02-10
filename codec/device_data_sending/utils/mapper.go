package utils

import "github.com/rojack96/gonika/codec/models"

func DataMapping(avlDataPacket []byte) *models.AvlDataPacketByte {
	var result models.AvlDataPacketByte

	if len(avlDataPacket) == 0 {
		return nil
	}

	result = models.AvlDataPacketByte{
		Preamble:        avlDataPacket[0:4],
		DataFieldLength: avlDataPacket[4:8],
		CodecID:         avlDataPacket[8],
		NumberOfData1:   avlDataPacket[9],
		Avldata:         avlDataPacket[10 : len(avlDataPacket)-8],
		NumberOfData2:   avlDataPacket[len(avlDataPacket)-5],
		Crc16:           avlDataPacket[len(avlDataPacket)-4:],
	}

	return &result
}
