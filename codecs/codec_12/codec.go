package codec12

import (
	"encoding/hex"
	"fmt"

	models "github.com/rojack96/teltonika-parser/models/codec_12"
)

func responseParser(responseMessage []byte) models.ResponseMessage {

	var response models.ResponseMessage

	response.Preamble = responseMessage[0:4]
	response.DataSize = responseMessage[4:8]
	response.CodecID = responseMessage[8]
	response.ResponseQuantity1 = responseMessage[9]
	response.Type = responseMessage[10]
	response.ResponseSize = responseMessage[11:15]
	response.Response = responseMessage[15 : len(responseMessage)-5]
	response.ResponseQuantity2 = responseMessage[len(responseMessage)-5]
	response.CRC16 = responseMessage[len(responseMessage)-4:]

	return response
}

func CreateCommand() []byte {
	return []byte{}
}

func Codec12(responseMessage *models.ResponseMessage) {

	fmt.Println("Preamble", hex.EncodeToString(responseMessage.Preamble))
	fmt.Println("DataSize", hex.EncodeToString(responseMessage.DataSize))
	fmt.Println("CodecID", responseMessage.CodecID)
	fmt.Println("ResponseQuantity1", responseMessage.ResponseQuantity1)
	fmt.Println("Type", responseMessage.Type)
	fmt.Println("ResponseSize", hex.EncodeToString(responseMessage.ResponseSize))
	fmt.Println("Response", string(responseMessage.Response))
	fmt.Println("ResponseQuantity2", responseMessage.ResponseQuantity2)
	fmt.Println("CRC16", hex.EncodeToString(responseMessage.CRC16))

}
