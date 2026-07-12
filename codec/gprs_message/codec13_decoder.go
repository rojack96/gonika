package gprs

import (
	"errors"

	utils "github.com/rojack96/gonika/codec/gprs_message/helpers"
	"github.com/rojack96/gonika/codec/helpers"
	"github.com/rojack96/gonika/codec/models"
	"github.com/rojack96/gonika/codec/parsers"
)

type codec13 struct {
	command  string
	builders *utils.Builders
	parser   *parsers.BaseParser
}

func NewCodec13Decoder() *codec13 {
	return &codec13{
		builders: utils.NewBuilders(),
		parser:   parsers.NewBaseParser(),
	}
}

/* ---------- Response ----------*/

func (c *codec13) DecodeResponse(avlDataPacket any) (*models.ResponseMessage, error) {
	var result models.ResponseMessage

	transformedData, err := helpers.TransformData(avlDataPacket)
	if err != nil {
		return nil, err
	}

	data := utils.ResponseDataMappingCodec13(transformedData)

	result.Preamble = c.parser.Preamble(data.Preamble)
	result.DataSize = c.parser.Parse4bytes(data.DataSize)
	result.CodecID = c.parser.CodecID(data.CodecID)
	result.ResponseQuantity1 = c.parser.Quantity(data.ResponseQuantity1)
	result.Type = c.parser.Type(data.Type)
	result.ResponseSize = c.parser.Parse4bytes(data.ResponseSize)
	result.Response = string(data.Response)
	result.ResponseQuantity2 = c.parser.Quantity(data.ResponseQuantity2)
	result.Crc16 = c.parser.Crc16(data.Crc16)
	timestamp := data.CodeSpecificMapperParam.(models.Codec13Byte).Timestamp
	codecSpefificParam := models.Codec13{Timestamp: c.parser.Parse4bytes(timestamp)}
	result.CodecSpecificParam = codecSpefificParam

	return &result, nil
}

/* ---------- Command ----------*/

func (c *codec13) DecodeCommand(avlDataPacket any) (*models.CommandMessage, error) {
	return nil, errors.New("DecodeCommand is not supported for codec13")
}

func (c *codec13) SetCommand(cmd string) {
	c.command = cmd
}
