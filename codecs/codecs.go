package codecs

import (
	modelsC16 "github.com/rojack96/teltonika-parser/models/codec_16"
	modelsC8 "github.com/rojack96/teltonika-parser/models/codec_8"
	modelsC8E "github.com/rojack96/teltonika-parser/models/codec_8E"
)

type codec8 struct{}
type codec8e struct{}
type codec16 struct{}
type codec12 struct{}
type codec13 struct{}
type codec14 struct{}

var Codec8 codec8
var Codec8E codec8e
var Codec16 codec16
var Codec12 codec12
var Codec13 codec13
var Codec14 codec14

func (c *codec8) DecodeAVLData(dataPacket []byte) modelsC8.AVLDataArray {
	return c8AVLData(dataPacket)
}

func (c *codec8e) DecodeAVLData(dataPacket []byte) modelsC8E.AVLDataArray {
	return c8eAVLData(dataPacket)
}

func (c *codec16) DecodeAVLData(dataPacket []byte) modelsC16.AVLDataArray {
	return c16AVLData(dataPacket)
}

func (c *codec12) ResponseParser(responseMessage []byte) []byte {
	return c12responseParser(responseMessage)
}

func (c *codec12) CreateCommand(command string) []byte {
	return c12CreateCommand(command)
}

func (c *codec13) CreateCommand(command string) []byte {
	return c13CreateCommand(command)
}

func (c *codec14) ResponseParser(responseMessage []byte) []byte {
	return c14ResponseParser(responseMessage)
}

func (c *codec14) CreateCommand(command string) []byte {
	return c14CreateCommand(command)
}
