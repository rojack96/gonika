package codec16

import (
	"encoding/hex"

	"github.com/rojack96/gonika/codec/constant"
	"github.com/rojack96/gonika/codec/device_data_sending/models"
	"github.com/rojack96/gonika/codec/device_data_sending/utils"
	m "github.com/rojack96/gonika/codec/models"
)

func (c *codec16) EncodeTCP(avlDataArray []m.AvlDataArrayEncoder) ([]byte, error) {
	var packet models.AvlDataPacketByteTCP

	nOfData := len(avlDataArray)

	packet.AvlDataPacketHeader.Preamble = [4]byte{0x00, 0x00, 0x00, 0x00}
	packet.AvlDataArray.CodecID = constant.Codec16
	packet.AvlDataArray.NumberOfData1 = byte(nOfData)
	packet.AvlDataArray.NumberOfData2 = byte(nOfData)

	for n := range nOfData {
		data, err := avlDataArrayBuilder(*c.builder, avlDataArray[n])
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

func (c *codec16) EncodeUDP(imei string, avlDataArray []m.AvlDataArrayEncoder) ([]byte, error) {
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
	packet.AvlDataArray.CodecID = constant.Codec16
	packet.AvlDataArray.NumberOfData1 = byte(nOfData)
	packet.AvlDataArray.NumberOfData2 = byte(nOfData)

	for n := range nOfData {
		data, err := avlDataArrayBuilder(*c.builder, avlDataArray[n])
		if err != nil {
			return nil, err
		}
		packet.AvlDataArray.AvlData = append(packet.AvlDataArray.AvlData, data...)
	}

	result := c.builder.MergeDataUDP(packet)

	return result, nil
}

func avlDataArrayBuilder(b utils.Builders, avlData m.AvlDataArrayEncoder) ([]byte, error) {
	result := make([]byte, 0)

	gps := avlData.GpsElementEncoder
	io := avlData.CodecEncoder.(m.Codec16Encoder)

	nOfOneByte := uint8(len(io.OneByte))
	nOfTwoByte := uint8(len(io.TwoByte))
	nOfFourByte := uint8(len(io.FourByte))
	nOfEightByte := uint8(len(io.EightByte))

	timestamp := b.Timestamp()
	priority := b.Priority()
	eventIo := b.EventIo2Byte()
	generationType := b.GenerationType()
	nOfTotalIo := nOfOneByte + nOfTwoByte + nOfFourByte + nOfEightByte

	result = append(result, timestamp[:]...)
	result = append(result, priority)

	if err := b.GpsElement(&result, gps); err != nil {
		return nil, err
	}

	result = append(result, eventIo[:]...)
	result = append(result, generationType)
	result = append(result, nOfTotalIo)

	result = append(result, nOfOneByte)
	for k, v := range io.OneByte {
		id := b.Uint16ToBytes(k)
		result = append(result, id[:]...)
		result = append(result, v)
	}

	result = append(result, nOfTwoByte)
	for k, v := range io.TwoByte {
		id := b.Uint16ToBytes(k)
		result = append(result, id[:]...)
		value := b.Uint16ToBytes(v)
		result = append(result, value[:]...)
	}

	result = append(result, nOfFourByte)
	for k, v := range io.FourByte {
		id := b.Uint16ToBytes(k)
		result = append(result, id[:]...)
		value := b.Uint32ToBytes(v)
		result = append(result, value[:]...)
	}

	result = append(result, nOfEightByte)
	for k, v := range io.EightByte {
		id := b.Uint16ToBytes(k)
		result = append(result, id[:]...)
		value := b.Uint64ToBytes(v)
		result = append(result, value[:]...)
	}

	return result, nil
}
