package codec12

import (
	"github.com/rojack96/gonika/codec/constant"
	"github.com/rojack96/gonika/codec/models"
	"github.com/rojack96/gonika/codec/parsers"
	"github.com/rojack96/gonika/codec/utils"
)

type codec12 struct {
	command       string
	avlDataPacket []byte
	builders      *utils.Builders
}

func New(avlDataPacket []byte) *codec12 {
	return &codec12{avlDataPacket: avlDataPacket, builders: utils.NewBuilders()}
}

func (c *codec12) SetCommand(cmd string) {
	c.command = cmd
}

// Encode build the message received in a message to write in Codec12
func (c *codec12) Encode() []byte {
	var cmd models.CommandMessage

	cmd.Preamble = []byte{0x00, 0x00, 0x00, 0x00}
	cmd.CodecID = constant.Codec12
	cmd.Type = constant.Command
	cmd.CommandQuantity1 = 0x01
	cmd.CommandQuantity2 = cmd.CommandQuantity1
	cmd.Command = []byte(c.command)
	cmd.CommandSize = c.builders.CommandSize(cmd.Command)
	cmd.DataSize = c.builders.DataSize(cmd)

	result := c.builders.MergeMessage(cmd)

	cmd.Crc16 = c.builders.Crc16Builder(result)

	result = append(result, cmd.Crc16...)

	return result
}

func (c *codec12) Decode() *models.ResponseMessage {
	var result models.ResponseMessage

	data := utils.ResponseDataMapping(c.avlDataPacket)

	result.Preamble = parsers.Preamble(data.Preamble)
	result.DataSize = parsers.DataSize(data.DataSize)
	result.CodecID = parsers.CodecId(data.CodecID)
	result.ResponseQuantity1 = parsers.Quantity(data.ResponseQuantity1)
	result.Type = parsers.Type(data.Type)
	result.ResponseSize = parsers.ResponseSize(data.ResponseSize)
	result.Response = string(data.Response)
	result.ResponseQuantity2 = parsers.Quantity(data.ResponseQuantity2)
	result.Crc16 = parsers.Crc16(data.Crc16)

	return &result
}
