package codec13

import (
	"errors"

	mapper "github.com/rojack96/gonika/codec/gprs_message/models"
	"github.com/rojack96/gonika/codec/gprs_message/utils"
	"github.com/rojack96/gonika/codec/models"
	"github.com/rojack96/gonika/codec/parsers"
)

type codec13 struct {
	command       string
	avlDataPacket []byte
	builders      *utils.Builders
	parser        parsers.BaseParser
}

func New(avlDataPacket []byte) *codec13 {
	return &codec13{
		avlDataPacket: avlDataPacket,
		builders:      utils.NewBuilders(),
		parser:        parsers.NewBaseParser(),
	}
}

/* ---------- Response ----------*/

func (c *codec13) DecodeResponse() *models.ResponseMessage {
	var result models.ResponseMessage

	data := utils.ResponseDataMappingCodec13(c.avlDataPacket)

	result.Preamble = c.parser.Preamble(data.Preamble)
	result.DataSize = c.parser.Parse4bytes(data.DataSize)
	result.CodecID = c.parser.CodecID(data.CodecID)
	result.ResponseQuantity1 = c.parser.Quantity(data.ResponseQuantity1)
	result.Type = c.parser.Type(data.Type)
	result.ResponseSize = c.parser.Parse4bytes(data.ResponseSize)
	result.Response = string(data.Response)
	result.ResponseQuantity2 = c.parser.Quantity(data.ResponseQuantity2)
	result.Crc16 = c.parser.Crc16(data.Crc16)
	timestamp := data.CodeSpecificMapperParam.(mapper.Codec13).Timestamp
	codecSpefificParam := models.Codec13{Timestamp: c.parser.Parse4bytes(timestamp)}
	result.CodecSpecificParam = codecSpefificParam

	return &result
}

/* ---------- Command ----------*/

func (c *codec13) DecodeCommand() (*models.CommandMessage, error) {
	return nil, errors.New("DecodeCommand is not supported for codec13")
}

func (c *codec13) SetCommand(cmd string) {
	c.command = cmd
}
