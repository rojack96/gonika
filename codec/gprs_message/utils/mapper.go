package utils

import "github.com/rojack96/gonika/codec/models"

func CommandDataMapping(responseData []byte) *models.CommandMessageByte {
	var result models.CommandMessageByte

	if len(responseData) == 0 {
		return nil
	}

	result = models.CommandMessageByte{
		Preamble:         responseData[0:4],
		DataSize:         responseData[4:8],
		CodecID:          responseData[8],
		CommandQuantity1: responseData[9],
		Type:             responseData[10],
		CommandSize:      responseData[11:15],
		Command:          responseData[15 : len(responseData)-5],
		CommandQuantity2: responseData[len(responseData)-5],
		Crc16:            responseData[len(responseData)-4:],
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

func ResponseDataMappingCodec13(responseData []byte) *models.ResponseMessageCodec13Byte {
	var result models.ResponseMessageCodec13Byte

	if len(responseData) == 0 {
		return nil
	}

	result = models.ResponseMessageCodec13Byte{
		Preamble:          responseData[0:4],
		DataSize:          responseData[4:8],
		CodecID:           responseData[8],
		ResponseQuantity1: responseData[9],
		Type:              responseData[10],
		ResponseSize:      responseData[11:15],
		Timestamp:         responseData[15:19],
		Response:          responseData[19 : len(responseData)-5],
		ResponseQuantity2: responseData[len(responseData)-5],
		Crc16:             responseData[len(responseData)-4:],
	}

	return &result
}

func ResponseDataMappingCodec14(responseData []byte) *models.ResponseMessageCodec14Byte {
	var result models.ResponseMessageCodec14Byte

	if len(responseData) == 0 {
		return nil
	}

	result = models.ResponseMessageCodec14Byte{
		Preamble:          responseData[0:4],
		DataSize:          responseData[4:8],
		CodecID:           responseData[8],
		ResponseQuantity1: responseData[9],
		Type:              responseData[10],
		ResponseSize:      responseData[11:15],
		Imei:              responseData[15:23],
		Response:          responseData[23 : len(responseData)-5],
		ResponseQuantity2: responseData[len(responseData)-5],
		Crc16:             responseData[len(responseData)-4:],
	}

	return &result
}
