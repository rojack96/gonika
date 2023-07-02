package codecs

import (
	"encoding/hex"

	"github.com/rojack96/gotlk/constant"
	modelsCodec12 "github.com/rojack96/gotlk/models/codec_12"
)

// Return a response from device
func c12responseParser(responseMessage []byte) []byte {

	var response modelsCodec12.ResponseMessage

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

func c12CreateCommand(command string) []byte {
	var commandMessage modelsCodec12.CommandMessage

	commandMessage.Preamble = constant.PREAMBLE
	commandMessage.CodecID = hex.EncodeToString([]byte{constant.CODEC_12})
	commandMessage.Type = constant.TYPE_COMMAND
	commandMessage.CommandQuantity1 = constant.COMMAND_QUANTITY
	commandMessage.CommandQuantity2 = commandMessage.CommandQuantity1

	commandMessage.Command = commandBuilder(command)
	commandMessage.CommandSize = commandSize(commandMessage.Command)
	commandMessage.DataSize = c12dataSize(&commandMessage)

	commandMessage.CRC16 = c12crc16builder(&commandMessage)

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
