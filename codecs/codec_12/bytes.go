package codec12

import (
	"encoding/hex"

	models "github.com/rojack96/teltonika-parser/models/codec_12"
)

func ResponseParserBytes(responseMessage []byte) []byte {

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

	return response.Response
}

func CreateCommandBytes(command string) []byte {
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

	mes, e := hex.DecodeString(messageToSend)
	if e != nil {
		return nil
	}
	return mes
}
