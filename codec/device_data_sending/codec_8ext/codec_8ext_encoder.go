package codec8ext

import (
	"encoding/hex"

	"github.com/rojack96/gonika/codec/constant"
	"github.com/rojack96/gonika/codec/device_data_sending/models"
	"github.com/rojack96/gonika/codec/device_data_sending/utils"
	m "github.com/rojack96/gonika/codec/models"
	"github.com/rojack96/gonika/codec/parsers"
)

func NewEncoder() *codec8ext {
	return &codec8ext{
		parser:  parsers.NewBaseParser(),
		builder: utils.NewBuilders(),
	}
}

func (c *codec8ext) EncodeTCP(avlDataArray []m.AvlDataArrayEncoder) ([]byte, error) {
	var packet models.AvlDataPacketByteTCP

	nOfData := len(avlDataArray)

	packet.AvlDataPacketHeader.Preamble = [4]byte{0x00, 0x00, 0x00, 0x00}
	packet.AvlDataArray.CodecID = constant.Codec8ext
	packet.AvlDataArray.NumberOfData1 = byte(nOfData)
	packet.AvlDataArray.NumberOfData2 = byte(nOfData)

	for n := range nOfData {
		data, err := avlDataArrayBuilderExt(*c.builder, avlDataArray[n])
		if err != nil {
			return nil, err
		}
		packet.AvlDataArray.AvlData = append(packet.AvlDataArray.AvlData, data...)
	}

	packet.AvlDataPacketHeader.DataFieldLength = c.builder.DataFieldLength(packet)

	result := c.builder.MergeDataTCP(packet)
	packet.Crc16 = c.builder.Crc16Builder(result)

	result = append(result, packet.Crc16[:]...)

	return result, nil
}

func (c *codec8ext) EncodeUDP(imei string, avlDataArray []m.AvlDataArrayEncoder) ([]byte, error) {
	var packet models.AvlDataPacketByteUDP

	nOfData := len(avlDataArray)

	packet.UdpChannelHeader.Length = [2]byte{}
	packet.UdpChannelHeader.PacketID = [2]byte{}
	packet.UdpChannelHeader.NotUsableByte = 0x01
	packet.UdpAvlPacketHeader.AvlPacketID = 0x05
	packet.UdpAvlPacketHeader.ImeiLength = [2]byte{0x00, 0x0F}
	imeiByte, err := hex.DecodeString(imei)
	if err != nil {
		return nil, err
	}
	packet.UdpAvlPacketHeader.Imei = [15]byte(imeiByte)
	packet.AvlDataArray.CodecID = constant.Codec8ext
	packet.AvlDataArray.NumberOfData1 = byte(nOfData)
	packet.AvlDataArray.NumberOfData2 = byte(nOfData)

	for n := range nOfData {
		data, err := avlDataArrayBuilderExt(*c.builder, avlDataArray[n])
		if err != nil {
			return nil, err
		}
		packet.AvlDataArray.AvlData = append(packet.AvlDataArray.AvlData, data...)
	}

	result := c.builder.MergeDataUDP(packet)

	return result, nil
}

func avlDataArrayBuilderExt(b utils.Builders, avlData m.AvlDataArrayEncoder) ([]byte, error) {
	result := make([]byte, 0)

	gps := avlData.GpsElementEncoder
	io := avlData.CodecEncoder.(m.Codec8ExtEncoder)

	nOfOneByte := uint16(len(io.OneByte))
	nOfTwoByte := uint16(len(io.TwoByte))
	nOfFourByte := uint16(len(io.FourByte))
	nOfEightByte := uint16(len(io.EightByte))
	nOfXByte := uint16(len(io.XByte))

	timestamp := b.Timestamp()
	priority := b.Priority()
	eventIo := b.EventIo2Byte()
	nOfTotalIo := nOfOneByte + nOfTwoByte + nOfFourByte + nOfEightByte + nOfXByte

	result = append(result, timestamp[:]...)
	result = append(result, priority)

	if err := b.GpsElement(&result, gps); err != nil {
		return nil, err
	}

	// Event IO (2 bytes)
	result = append(result, eventIo[:]...)

	// Number of Total IO (2 bytes)
	nOfTotalIoBytes := b.Uint16ToBytes(nOfTotalIo)
	result = append(result, nOfTotalIoBytes[:]...)

	// One Byte IO
	oneByte := b.Uint16ToBytes(nOfOneByte)
	result = append(result, oneByte[:]...)
	for k, v := range io.OneByte {
		id := b.Uint16ToBytes(k)
		result = append(result, id[:]...)
		result = append(result, v)
	}

	// Two Byte IO
	twoByte := b.Uint16ToBytes(nOfTwoByte)
	result = append(result, twoByte[:]...)
	for k, v := range io.TwoByte {
		id := b.Uint16ToBytes(k)
		result = append(result, id[:]...)
		value := b.Uint16ToBytes(v)
		result = append(result, value[:]...)
	}

	// Four Byte IO
	fourByte := b.Uint16ToBytes(nOfFourByte)
	result = append(result, fourByte[:]...)
	for k, v := range io.FourByte {
		id := b.Uint16ToBytes(k)
		result = append(result, id[:]...)
		value := b.Uint32ToBytes(v)
		result = append(result, value[:]...)
	}

	// Eight Byte IO
	eightByte := b.Uint16ToBytes(nOfEightByte)
	result = append(result, eightByte[:]...)
	for k, v := range io.EightByte {
		id := b.Uint16ToBytes(k)
		result = append(result, id[:]...)
		value := b.Uint64ToBytes(v)
		result = append(result, value[:]...)
	}

	// X Byte IO
	xByte := b.Uint16ToBytes(nOfXByte)
	result = append(result, xByte[:]...)
	for k, v := range io.XByte {
		id := b.Uint16ToBytes(k)
		result = append(result, id[:]...)
		// Convert string to bytes
		valueBytes := []byte(v)
		// Write value length (2 bytes)
		lengthBytes := b.Uint16ToBytes(uint16(len(valueBytes)))
		result = append(result, lengthBytes[:]...)
		result = append(result, valueBytes...)
	}

	return result, nil
}
