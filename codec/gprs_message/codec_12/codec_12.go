package codec12

import (
	"github.com/rojack96/gonika/codec/constant"
	mapper "github.com/rojack96/gonika/codec/gprs_message/models"
	"github.com/rojack96/gonika/codec/gprs_message/utils"
	"github.com/rojack96/gonika/codec/models"
	"github.com/rojack96/gonika/codec/parsers"
)

type codec12 struct {
	command       string
	avlDataPacket []byte
	builders      *utils.Builders
	parser        parsers.BaseParser
}

func New(avlDataPacket []byte) *codec12 {
	return &codec12{
		avlDataPacket: avlDataPacket,
		builders:      utils.NewBuilders(),
		parser:        parsers.NewBaseParser(),
	}
}

/* ---------- Response ----------*/

func (c *codec12) DecodeResponse() *models.ResponseMessage {
	var result models.ResponseMessage

	data := utils.ResponseDataMappingCodec12(c.avlDataPacket)

	result.Preamble = c.parser.Preamble(data.Preamble)
	result.DataSize = c.parser.Parse4bytes(data.DataSize)
	result.CodecID = c.parser.CodecId(data.CodecID)
	result.ResponseQuantity1 = c.parser.Quantity(data.ResponseQuantity1)
	result.Type = c.parser.Type(data.Type)
	result.ResponseSize = c.parser.Parse4bytes(data.ResponseSize)
	result.Response = string(data.Response)
	result.ResponseQuantity2 = c.parser.Quantity(data.ResponseQuantity2)
	result.Crc16 = c.parser.Crc16(data.Crc16)

	return &result
}

/* ---------- Command ----------*/

func (c *codec12) DecodeCommand() (*models.CommandMessage, error) {
	var result models.CommandMessage

	data := utils.CommandDataMappingCodec12(c.avlDataPacket)

	result.Preamble = c.parser.Preamble(data.Preamble)
	result.DataSize = c.parser.Parse4bytes(data.DataSize)
	result.CodecID = c.parser.CodecId(data.CodecID)
	result.CommandQuantity1 = c.parser.Quantity(data.CommandQuantity1)
	result.Type = c.parser.Type(data.Type)
	result.CommandSize = c.parser.Parse4bytes(data.CommandSize)
	result.Command = string(data.Command)
	result.CommandQuantity2 = c.parser.Quantity(data.CommandQuantity2)
	result.Crc16 = c.parser.Crc16(data.Crc16)

	return &result, nil
}

func (c *codec12) SetCommand(cmd string) {
	c.command = cmd
}

// Encode build the message received in a message to write in Codec12
func (c *codec12) Encode() []byte {
	var cmd mapper.CommandMessage

	cmd.Preamble = [4]byte{0x00, 0x00, 0x00, 0x00}
	cmd.CodecID = constant.Codec12
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
