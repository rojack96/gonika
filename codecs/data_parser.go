package codecs

import "github.com/rojack96/teltonika-parser/models"

func dataBytesParser(dataPacket *[]byte) models.AVLDataPacket {
	var avlDataPacket models.AVLDataPacket

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
