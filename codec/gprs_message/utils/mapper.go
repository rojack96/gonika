package utils

import "github.com/rojack96/gonika/codec/gprs_message/models"

/* ---------- Codec12 ---------- */

func CommandDataMappingCodec12(responseData []byte) *models.CommandMessage {
	var result models.CommandMessage

	if len(responseData) == 0 {
		return nil
	}

	result = models.CommandMessage{
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

func ResponseDataMappingCodec12(responseData []byte) *models.ResponseMessage {
	var result models.ResponseMessage

	if len(responseData) == 0 {
		return nil
	}

	result = models.ResponseMessage{
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

/* ---------- Codec13 ---------- */

func ResponseDataMappingCodec13(responseData []byte) *models.ResponseMessage {
	var result models.ResponseMessage

	if len(responseData) == 0 {
		return nil
	}

	result = models.ResponseMessage{
		Preamble:          responseData[0:4],
		DataSize:          responseData[4:8],
		CodecID:           responseData[8],
		ResponseQuantity1: responseData[9],
		Type:              responseData[10],
		ResponseSize:      responseData[11:15],
		Response:          responseData[19 : len(responseData)-5],
		ResponseQuantity2: responseData[len(responseData)-5],
		Crc16:             responseData[len(responseData)-4:],
	}

	timestamp := models.Codec13{Timestamp: [4]byte(responseData[15:19])}

	result.CodeSpecificMapperParam = timestamp

	return &result
}

/* ----------Codec14 ---------- */

func CommandDataMappingCodec13(responseData []byte) *models.CommandMessage {
	var result models.CommandMessage

	if len(responseData) == 0 {
		return nil
	}

	result = models.CommandMessage{
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

func ResponseDataMappingCodec14(responseData []byte) *models.ResponseMessage {
	var result models.ResponseMessage

	if len(responseData) == 0 {
		return nil
	}

	result = models.ResponseMessage{
		Preamble:          responseData[0:4],
		DataSize:          responseData[4:8],
		CodecID:           responseData[8],
		ResponseQuantity1: responseData[9],
		Type:              responseData[10],
		ResponseSize:      responseData[11:15],
		Response:          responseData[23 : len(responseData)-5],
		ResponseQuantity2: responseData[len(responseData)-5],
		Crc16:             responseData[len(responseData)-4:],
	}

	imei := models.Codec14{Imei: [8]byte(responseData[15:23])}
	result.CodeSpecificMapperParam = imei

	return &result
}
