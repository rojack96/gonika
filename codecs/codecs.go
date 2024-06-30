package codecs

import (
	"github.com/rojack96/gonika/codecs/codec_16"
	codec8 "github.com/rojack96/gonika/codecs/codec_8"
	codec8e "github.com/rojack96/gonika/codecs/codec_8e"
)

type Codecs struct {
	c8  codec8.Codec8
	c8e codec8e.Codec8e
	c16 codec16.Codec16
}

func (c *Codecs) BufferAnalyzer(dataBuffer *[]byte) (*uint8, error) {
	return BufferAnalyzer(dataBuffer)
}

func (c *Codecs) Decode8(dataPacket []byte) codec8.AVLDataArray { return c.c8.Decode(dataPacket) }

func (c *Codecs) Decode8ext(dataPacket []byte) codec8e.AVLDataArray { return c.c8e.Decode(dataPacket) }

func (c *Codecs) Decode16(dataPacket []byte) codec16.AVLDataArray { return c.c16.Decode(dataPacket) }
