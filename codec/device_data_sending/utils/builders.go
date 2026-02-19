package utils

import (
	"encoding/binary"
	"math/rand"
	"strconv"
	"time"

	"github.com/getrak/crc16"
	"github.com/rojack96/gonika/codec/device_data_sending/models"
	m "github.com/rojack96/gonika/codec/models"
)

type Builders struct{}

func NewBuilders() *Builders {
	return &Builders{}
}

func (b *Builders) MergeDataTCP(data models.AvlDataPacketByteTCP) []byte {
	dataLen := len(data.AvlDataPacketHeader.Preamble) +
		len(data.AvlDataPacketHeader.DataFieldLength) +
		len(data.AvlDataArray.AvlData) +
		3 //3 equals to CodecID, NumberOfData1 & 2

	message := make([]byte, 0, dataLen)
	message = append(message, data.AvlDataPacketHeader.Preamble[:]...)
	message = append(message, data.AvlDataPacketHeader.DataFieldLength[:]...)
	message = append(message, data.AvlDataArray.CodecID)
	message = append(message, data.AvlDataArray.NumberOfData1)
	message = append(message, data.AvlDataArray.AvlData...)
	message = append(message, data.AvlDataArray.NumberOfData2)

	return message
}

func (b *Builders) MergeDataUDP(data models.AvlDataPacketByteUDP) []byte {
	dataLen := len(data.UdpChannelHeader.Length) +
		len(data.UdpChannelHeader.PacketID) +
		len(data.UdpAvlPacketHeader.ImeiLength) +
		len(data.UdpAvlPacketHeader.Imei) +
		len(data.AvlDataArray.AvlData) +
		5 // 5 equals to NotUsableByte, AvlPacketID, CodecID, NumberOfData1 & 2
	message := make([]byte, 0, dataLen)
	message = append(message, data.UdpChannelHeader.Length[:]...)
	message = append(message, data.UdpChannelHeader.PacketID[:]...)
	message = append(message, data.UdpChannelHeader.NotUsableByte)
	message = append(message, data.UdpAvlPacketHeader.AvlPacketID)
	message = append(message, data.UdpAvlPacketHeader.ImeiLength[:]...)
	message = append(message, data.UdpAvlPacketHeader.Imei[:]...)
	message = append(message, data.AvlDataArray.CodecID)
	message = append(message, data.AvlDataArray.NumberOfData1)
	message = append(message, data.AvlDataArray.AvlData...)
	message = append(message, data.AvlDataArray.NumberOfData2)

	return message
}

/* ---------- Data to byte transformation ----------*/

func (b *Builders) DataFieldLength(command models.AvlDataPacketByteTCP) [4]byte {
	ds := len(command.AvlDataArray.AvlData) + 4 // 4 is equal to len of CodecID, CommandQuantity (1 & 2), Type
	return b.fourBytesTransformation(ds)
}

func (b *Builders) Timestamp() [8]byte {
	ts := time.Now().UnixMilli()
	return b.Uint64ToBytes(uint64(ts))
}

func (b *Builders) Priority() uint8 {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return uint8(r.Intn(3))
}

func (b *Builders) EventIo1Byte() uint8 {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	// more zero then other number
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

	return b.Uint16ToBytes(value)
}

func (b *Builders) GenerationType() byte {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	return byte(r.Intn(8)) // 0–7
}

func (b *Builders) GpsElement(result *[]byte, gps m.GpsElementEncoder) error {
	lat, err := b.coordinateStringToBytes(gps.Latitude)
	if err != nil {
		return err
	}
	long, err := b.coordinateStringToBytes(gps.Longitude)
	if err != nil {
		return err
	}
	altitude := b.Uint16ToBytes(gps.Altitude)
	angle := b.Uint16ToBytes(gps.Angle)
	speed := b.Uint16ToBytes(gps.Speed)

	*result = append(*result, long[:]...)
	*result = append(*result, lat[:]...)
	*result = append(*result, altitude[:]...)
	*result = append(*result, angle[:]...)
	*result = append(*result, gps.Satellites)
	*result = append(*result, speed[:]...)

	return nil
}

// coordinateStringToBytes converts a latitude or longitude string to a 4-byte representation
// The string is parsed as a float64, multiplied by 1e7, and converted to big-endian bytes.
// This is the inverse operation of decodeCoordinate from base_parsers.go
func (b *Builders) coordinateStringToBytes(coordinateStr string) ([4]byte, error) {
	var result [4]byte
	// Parse string to float64
	coord, err := strconv.ParseFloat(coordinateStr, 64)
	if err != nil {
		return result, err
	}

	// Multiply by 1e7 to get integer representation
	intCoord := int32(coord * 1e7)

	// Convert to 4 bytes in big-endian format
	result = b.Uint32ToBytes(uint32(intCoord))
	return result, nil
}

func (b *Builders) Crc16Builder(command []byte) [4]byte {
	crcTable := crc16.MakeTable(crc16.CRC16_ARC)
	crcRes := crc16.Checksum(command[8:], crcTable)

	result := b.fourBytesTransformation(int(crcRes))

	return result
}

/* ---------- Byte transformation ----------*/

func (b *Builders) fourBytesTransformation(data int) [4]byte {
	num := uint32(data)
	return b.Uint32ToBytes(num)
}

func (b *Builders) Uint16ToBytes(value uint16) [2]byte {
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, value)
	return [2]byte(bytes)
}

func (b *Builders) Uint32ToBytes(value uint32) [4]byte {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, value)
	return [4]byte(bytes)
}

func (b *Builders) Uint64ToBytes(value uint64) [8]byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, value)
	return [8]byte(bytes)
}
