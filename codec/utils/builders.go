package utils

import (
	"encoding/binary"

	"github.com/getrak/crc16"
	"github.com/rojack96/gonika/codec/models"
)

type Builders struct{}

func NewBuilders() *Builders {
	return &Builders{}
}

func (b *Builders) MergeMessage(command models.CommandMessage) []byte {
	message := make([]byte, 0)
	message = append(message, command.Preamble...)
	message = append(message, command.DataSize...)
	message = append(message, command.CodecID)
	message = append(message, command.CommandQuantity1)
	message = append(message, command.Type)
	message = append(message, command.CommandSize...)
	message = append(message, command.Command...)
	message = append(message, command.CommandQuantity2)

	return message
}

// return command size from length of command
func (b *Builders) CommandSize(command []byte) []byte {
	lenCommand := len(command)
	return b.fourBytesTransformation(lenCommand)
}

func (b *Builders) DataSize(command models.CommandMessage) []byte {
	ds := len(command.CommandSize) + len(command.Command) + 4 // 4 is equal to len of CodecId, CommandQuantity (1 & 2), Type

	return b.fourBytesTransformation(ds)
}

func (b *Builders) Crc16Builder(command []byte) []byte {
	crcTable := crc16.MakeTable(crc16.CRC16_ARC)
	crcRes := crc16.Checksum(command[8:], crcTable)

	result := b.fourBytesTransformation(int(crcRes))

	return result
}

func (b *Builders) fourBytesTransformation(data int) []byte {
	bytes := make([]byte, 4)
	num := uint32(data)

	binary.BigEndian.PutUint32(bytes, num)
	return bytes
}
