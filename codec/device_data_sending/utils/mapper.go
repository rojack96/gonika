package utils

import (
	"github.com/rojack96/gonika/codec/device_data_sending/models"
)

func DataMapping(avlDataPacket []byte) *models.AvlDataPacketByte {
	var result models.AvlDataPacketByte

	if len(avlDataPacket) == 0 {
		return nil
	}

	avlDataPacketHeader := models.AvlDataPacketHeader{
		Preamble:        avlDataPacket[0:4],
		DataFieldLength: avlDataPacket[4:8],
	}

	avlDataArray := models.AvlDataArray{
		CodecID:       avlDataPacket[8],
		NumberOfData1: avlDataPacket[9],
		AvlData:       avlDataPacket[10 : len(avlDataPacket)-5],
		NumberOfData2: avlDataPacket[len(avlDataPacket)-5],
	}

	result = models.AvlDataPacketByte{
		AvlDataPacketHeader: avlDataPacketHeader,
		AvlDataArray:        avlDataArray,
		Crc16:               avlDataPacket[len(avlDataPacket)-4:],
	}

	return &result
}

func UdpDataMapping(avlDataPacket []byte) *models.UdpAvlDataPacketByte {
	var result models.UdpAvlDataPacketByte

	if len(avlDataPacket) == 0 {
		return nil
	}

	udpChannelHeader := models.UdpChannelHeader{
		Length:        avlDataPacket[0:4],
		PacketId:      avlDataPacket[4:8],
		NotUsableByte: avlDataPacket[8],
	}

	avlPacketHeader := models.UdpAvlPacketHeader{
		AvlPacketId: avlDataPacket[9],
		ImeiLength:  avlDataPacket[10:12],
		Imei:        avlDataPacket[12:26],
	}

	avlDataArray := models.AvlDataArray{
		CodecID:       avlDataPacket[26],
		NumberOfData1: avlDataPacket[27],
		AvlData:       avlDataPacket[28 : len(avlDataPacket)-1],
		NumberOfData2: avlDataPacket[len(avlDataPacket)-1],
	}

	result = models.UdpAvlDataPacketByte{
		UdpChannelHeader:   udpChannelHeader,
		UdpAvlPacketHeader: avlPacketHeader,
		AvlDataArray:       avlDataArray,
	}

	return &result
}
