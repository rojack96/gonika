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

func ResponseDataMapping(responseData []byte) *models.ResponseMessageByte {
	var result models.ResponseMessageByte

	if len(responseData) == 0 {
		return nil
	}

	result = models.ResponseMessageByte{
		Preamble:          responseData[0:4],
		DataSize:          responseData[4:8],
		CodecID:           responseData[8],
		ResponseQuantity1: responseData[9],
		Type:              responseData[10],
		ResponseSize:      responseData[11:15],
		Response:          responseData[15 : len(responseData)-5],
		ResponseQuantity2: responseData[len(responseData)-5],
		Crc16:             responseData[len(responseData)-4:],
	}

	return &result
}
