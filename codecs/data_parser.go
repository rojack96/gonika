package codecs

import "github.com/rojack96/gotlk/models"

func dataParser(dataPacket *[]byte) models.DataPacket {
	var avlDataPacket models.DataPacket

	data := *dataPacket

	avlDataPacket.Preamble = data[0:4]
	avlDataPacket.DataFieldLength = data[4:8]
	avlDataPacket.CodecID = data[8]
	avlDataPacket.NumberOfData1 = data[9]
	avlDataPacket.AVLdata = data[10 : len(data)-8]
	avlDataPacket.NumberOfData2 = data[len(data)-5]
	avlDataPacket.CRC16 = data[len(data)-4:]

	return avlDataPacket
}
