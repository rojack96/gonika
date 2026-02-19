package utils

import (
	"github.com/rojack96/gonika/codec/device_data_sending/models"
)

func DataMapping(avlDataPacket []byte) *models.AvlDataPacketByteTCP {
	var result models.AvlDataPacketByteTCP

	if len(avlDataPacket) == 0 {
		return nil
	}

	avlDataPacketHeader := models.AvlDataPacketHeader{
		Preamble:        [4]byte(avlDataPacket[0:4]),
		DataFieldLength: [4]byte(avlDataPacket[4:8]),
	}

	avlDataArray := models.AvlDataArray{
		CodecID:       avlDataPacket[8],
		NumberOfData1: avlDataPacket[9],
		AvlData:       avlDataPacket[10 : len(avlDataPacket)-5],
		NumberOfData2: avlDataPacket[len(avlDataPacket)-5],
	}

	result = models.AvlDataPacketByteTCP{
		AvlDataPacketHeader: avlDataPacketHeader,
		AvlDataArray:        avlDataArray,
		Crc16:               [4]byte(avlDataPacket[len(avlDataPacket)-4:]),
	}

	return &result
}

func UdpDataMapping(avlDataPacket []byte) *models.AvlDataPacketByteUDP {
	var result models.AvlDataPacketByteUDP

	if len(avlDataPacket) == 0 {
		return nil
	}

	udpChannelHeader := models.UdpChannelHeader{
		Length:        [2]byte(avlDataPacket[0:4]),
		PacketId:      [2]byte(avlDataPacket[4:8]),
		NotUsableByte: avlDataPacket[8],
	}

	avlPacketHeader := models.UdpAvlPacketHeader{
		AvlPacketId: avlDataPacket[9],
		ImeiLength:  [2]byte(avlDataPacket[10:12]),
		Imei:        [15]byte(avlDataPacket[12:26]),
	}

	avlDataArray := models.AvlDataArray{
		CodecID:       avlDataPacket[26],
		NumberOfData1: avlDataPacket[27],
		AvlData:       avlDataPacket[28 : len(avlDataPacket)-1],
		NumberOfData2: avlDataPacket[len(avlDataPacket)-1],
	}

	result = models.AvlDataPacketByteUDP{
		UdpChannelHeader:   udpChannelHeader,
		UdpAvlPacketHeader: avlPacketHeader,
		AvlDataArray:       avlDataArray,
	}

	return &result
}
