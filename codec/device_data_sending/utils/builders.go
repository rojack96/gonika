package utils

import (
	"encoding/binary"
	"math/rand"
	"time"

	"github.com/getrak/crc16"
	"github.com/rojack96/gonika/codec/device_data_sending/models"
)

type Builders struct{}

func NewBuilders() *Builders {
	return &Builders{}
}

func (b *Builders) MergeData(data models.AvlDataPacketByte) []byte {
	message := make([]byte, 0)
	message = append(message, data.AvlDataPacketHeader.Preamble[:]...)
	message = append(message, data.AvlDataPacketHeader.DataFieldLength[:]...)
	message = append(message, data.AvlDataArray.CodecID)
	message = append(message, data.AvlDataArray.NumberOfData1)
	message = append(message, data.AvlDataArray.AvlData...)
	message = append(message, data.AvlDataArray.NumberOfData2)

	return message
}

func (b *Builders) DataFieldLength(command models.AvlDataPacketByte) [4]byte {
	ds := len(command.AvlDataArray.AvlData) + 4 // 4 is equal to len of CodecId, CommandQuantity (1 & 2), Type

	return b.fourBytesTransformation(ds)
}

func (b *Builders) Timestamp() [8]byte {
	ts := time.Now().UnixMilli()
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, uint64(ts))
	return [8]byte(bytes)
}

func (b *Builders) Priority() uint8 {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return uint8(r.Intn(3))
}

func (b *Builders) EventIo1Byte() uint8 {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	if r.Float64() < 0.8 {
		return 0
	}
	return uint8(r.Intn(256))
}

func (b *Builders) EventIo2Byte() [2]byte {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	var value uint16
	if r.Float64() < 0.8 {
		value = 0
	} else {
		value = uint16(r.Intn(65536))
	}

	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, value)
	return [2]byte(bytes)
}

func (b *Builders) Crc16Builder(command []byte) [4]byte {
	crcTable := crc16.MakeTable(crc16.CRC16_ARC)
	crcRes := crc16.Checksum(command[8:], crcTable)

	result := b.fourBytesTransformation(int(crcRes))

	return result
}

func (b *Builders) fourBytesTransformation(data int) [4]byte {
	bytes := make([]byte, 4)
	num := uint32(data)

	binary.BigEndian.PutUint32(bytes, num)
	return [4]byte(bytes)
}
