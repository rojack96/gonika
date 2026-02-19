package codec12

import (
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
	result.CodecID = c.parser.CodecID(data.CodecID)
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
	result.CodecID = c.parser.CodecID(data.CodecID)
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
