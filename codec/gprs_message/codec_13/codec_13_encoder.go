package codec13

import (
	"github.com/rojack96/gonika/codec/constant"
	"github.com/rojack96/gonika/codec/gprs_message/helpers"
	mapper "github.com/rojack96/gonika/codec/gprs_message/models"
	"github.com/rojack96/gonika/codec/parsers"
)

func NewEncoder() *codec13 {
	return &codec13{
		builders: helpers.NewBuilders(),
		parser:   parsers.NewBaseParser(),
	}
}

// Encode build the message received in a message to write in Codec13
func (c *codec13) Encode() []byte {
	var cmd mapper.CommandMessage

	cmd.Preamble = [4]byte{0x00, 0x00, 0x00, 0x00}
	cmd.CodecID = constant.Codec13
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
