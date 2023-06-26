package codec13

import models "github.com/rojack96/teltonika-parser/models/codec_13"

func CreateCommand(command string) string {
	var commandMessage models.CommandMessage

	commandMessage.Preamble = PREAMBLE
	commandMessage.CodecID = CODEC_ID_13
	commandMessage.Type = TYPE_COMMAND
	commandMessage.CommandQuantity1 = COMMAND_QUANTITY
	commandMessage.CommandQuantity2 = commandMessage.CommandQuantity1

	commandMessage.Command = commandBuilder(command)
	commandMessage.Timestamp = timestampBuilder()
	commandMessage.CommandSize = commandSize(commandMessage.Command)
	commandMessage.DataSize = dataSize(&commandMessage)

	commandMessage.CRC16 = crc16builder(&commandMessage)

	messageToSend := commandMessage.Preamble +
		commandMessage.DataSize +
		commandMessage.CodecID +
		commandMessage.CommandQuantity1 +
		commandMessage.Type +
		commandMessage.CommandSize +
		commandMessage.Timestamp +
		commandMessage.Command +
		commandMessage.CommandQuantity2 +
		commandMessage.CRC16

	return messageToSend
}
