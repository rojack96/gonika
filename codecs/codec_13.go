package codecs

import (
	"encoding/hex"

	"github.com/rojack96/teltonika-parser/constant"
	modelsCodec13 "github.com/rojack96/teltonika-parser/models/codec_13"
)

func c13CreateCommand(command string) []byte {
	var commandMessage modelsCodec13.CommandMessage

	commandMessage.Preamble = constant.PREAMBLE
	commandMessage.CodecID = hex.EncodeToString([]byte{constant.CODEC_12})
	commandMessage.Type = constant.TYPE_COMMAND
	commandMessage.CommandQuantity1 = constant.COMMAND_QUANTITY
	commandMessage.CommandQuantity2 = commandMessage.CommandQuantity1

	commandMessage.Command = commandBuilder(command)
	commandMessage.CommandSize = commandSize(commandMessage.Command)
	commandMessage.Timestamp = timestampBuilder()
	commandMessage.DataSize = c13dataSize(&commandMessage)

	commandMessage.CRC16 = c13crc16builder(&commandMessage)

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
