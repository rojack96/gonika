package codec12

import (
	"encoding/hex"

	models "github.com/rojack96/teltonika-parser/models/codec_12"
)

const (
	PREAMBLE     = "00000000"
	CODEC_ID_12  = "0C"
	TYPE_COMMAND = "05"
	COMMAND_QUANTITY
)

func ResponseParser(responseMessage []byte) string {

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

	return hex.EncodeToString(response.Response)
}

func CreateCommand(command string) string {
	var commandMessage models.CommandMessage

	commandMessage.Preamble = PREAMBLE
	commandMessage.CodecID = CODEC_ID_12
	commandMessage.Type = TYPE_COMMAND
	commandMessage.CommandQuantity1 = COMMAND_QUANTITY
	commandMessage.CommandQuantity2 = commandMessage.CommandQuantity1

	commandMessage.Command = commandBuilder(command)
	commandMessage.CommandSize = commandSize(commandMessage.Command)
	commandMessage.DataSize = dataSize(&commandMessage)

	commandMessage.CRC16 = crc16builder(&commandMessage)

	messageToSend := commandMessage.Preamble +
		commandMessage.DataSize +
		commandMessage.CodecID +
		commandMessage.CommandQuantity1 +
		commandMessage.Type +
		commandMessage.CommandSize +
		commandMessage.Command +
		commandMessage.CommandQuantity2 +
		commandMessage.CRC16

	return messageToSend
}

// func Codec12(responseMessage *models.ResponseMessage) {

// 	fmt.Println("Preamble", hex.EncodeToString(responseMessage.Preamble))
// 	fmt.Println("DataSize", hex.EncodeToString(responseMessage.DataSize))
// 	fmt.Println("CodecID", responseMessage.CodecID)
// 	fmt.Println("ResponseQuantity1", responseMessage.ResponseQuantity1)
// 	fmt.Println("Type", responseMessage.Type)
// 	fmt.Println("ResponseSize", hex.EncodeToString(responseMessage.ResponseSize))
// 	fmt.Println("Response", string(responseMessage.Response))
// 	fmt.Println("ResponseQuantity2", responseMessage.ResponseQuantity2)
// 	fmt.Println("CRC16", hex.EncodeToString(responseMessage.CRC16))

// }
