package gprs

import (
	"github.com/rojack96/gonika/codec/constant"
	"github.com/rojack96/gonika/codec/gprs_message/helpers"
	"github.com/rojack96/gonika/codec/models"
	"github.com/rojack96/gonika/codec/parsers"
)

func NewCodec14Encoder() *codec14 {
	return &codec14{
		builders: helpers.NewBuilders(),
		parser:   parsers.NewBaseParser(),
	}
}

// Encode build the message received in a message to write in Codec14
func (c *codec14) Encode() []byte {
	var cmd models.CommandMessageByte

	cmd.Preamble = [4]byte{0x00, 0x00, 0x00, 0x00}
	cmd.CodecID = constant.Codec14
	cmd.Type = constant.Command
	cmd.CommandQuantity1 = 0x01
	cmd.CommandQuantity2 = cmd.CommandQuantity1
	cmd.Command = []byte(c.command)
	cmd.CommandSize = c.builders.CommandSize(cmd.Command)
	cmd.DataSize = c.builders.DataSize(cmd)

	result := c.builders.MergeMessage(cmd)

	cmd.Crc16 = c.builders.Crc16Builder(result)

	result = append(result, cmd.Crc16[:]...)

	return result
}
