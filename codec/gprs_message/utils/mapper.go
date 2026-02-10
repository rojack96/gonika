package utils

import "github.com/rojack96/gonika/codec/models"

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
