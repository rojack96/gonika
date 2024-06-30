package codecs

import (
	"encoding/hex"
	"strconv"
	"time"

	"github.com/getrak/crc16"
	modelsCodec12 "github.com/rojack96/gonika/models/codec_12"
	modelsCodec13 "github.com/rojack96/gonika/models/codec_13"
	modelsCodec14 "github.com/rojack96/gonika/models/codec_14"
)

func timestampBuilder() string {
	now := time.Now().String()

	nowBytes, e := hex.DecodeString(now)

	if e != nil {
		return ""
	}

	nowByteString := hex.EncodeToString(nowBytes)

	return nowByteString
}

func commandBuilder(cmd string) string {
	command := hex.EncodeToString([]byte(cmd))
	return command
}

func commandSize(cmd string) string {
	cmdSize := len(cmd) / 2

	commandSizeResult := padLeft(strconv.FormatInt(int64(cmdSize), 16), 8)

	return commandSizeResult
}

func padLeft(str string, length int) string {
	// ad zero's on the left
	for len(str) < length {
		str = "0" + str
	}
	return str
}

func c12dataSize(cmd *modelsCodec12.CommandMessage) string {
	c := *cmd

	ds := (len(c.CodecID) / 2) +
		(len(c.CommandQuantity1) / 2) +
		(len(c.Type) / 2) +
		(len(c.CommandSize) / 2) +
		(len(c.Command) / 2) +
		(len(c.CommandQuantity2) / 2)

	dataSizeResult := padLeft(strconv.FormatInt(int64(ds), 16), 8)

	return dataSizeResult
}

func c13dataSize(cmd *modelsCodec13.CommandMessage) string {
	c := *cmd

	ds := (len(c.CodecID) / 2) +
		(len(c.CommandQuantity1) / 2) +
		(len(c.Type) / 2) +
		(len(c.CommandSize) / 2) +
		(len(c.Timestamp) / 2) +
		(len(c.Command) / 2) +
		(len(c.CommandQuantity2) / 2)

	dataSizeResult := padLeft(strconv.FormatInt(int64(ds), 16), 8)

	return dataSizeResult
}

func c14dataSize(cmd *modelsCodec14.CommandMessage) string {
	c := *cmd

	ds := (len(c.CodecID) / 2) +
		(len(c.CommandQuantity1) / 2) +
		(len(c.Type) / 2) +
		(len(c.CommandAndImeiSize) / 2) +
		(len(c.IMEI) / 2) +
		(len(c.Command) / 2) +
		(len(c.CommandQuantity2) / 2)

	dataSizeResult := padLeft(strconv.FormatInt(int64(ds), 16), 8)

	return dataSizeResult
}

func c12crc16builder(command *modelsCodec12.CommandMessage) string {
	c := *command

	crc := c.CodecID + c.CommandQuantity1 + c.Type + c.CommandSize + c.Command + c.CommandQuantity2

	resCRC, e := hex.DecodeString(crc)
	if e != nil {
		return "00000000"
	}
	checksum := crc16.Checksum(resCRC, crc16.MakeTable(
		crc16.CRC16_ARC,
	))

	crcResult := padLeft(strconv.FormatInt(int64(checksum), 16), 8)
	return crcResult
}

func c13crc16builder(command *modelsCodec13.CommandMessage) string {
	c := *command

	crc := c.CodecID + c.CommandQuantity1 + c.Type + c.CommandSize + c.Timestamp + c.Command + c.CommandQuantity2

	resCRC, e := hex.DecodeString(crc)
	if e != nil {
		return "00000000"
	}
	checksum := crc16.Checksum(resCRC, crc16.MakeTable(
		crc16.CRC16_ARC,
	))

	crcResult := padLeft(strconv.FormatInt(int64(checksum), 16), 8)
	return crcResult
}

func c14crc16builder(command *modelsCodec14.CommandMessage) string {
	c := *command

	crc := c.CodecID + c.CommandQuantity1 + c.Type + c.CommandAndImeiSize + c.IMEI + c.Command + c.CommandQuantity2

	resCRC, e := hex.DecodeString(crc)
	if e != nil {
		return "00000000"
	}
	checksum := crc16.Checksum(resCRC, crc16.MakeTable(
		crc16.CRC16_ARC,
	))

	crcResult := padLeft(strconv.FormatInt(int64(checksum), 16), 8)
	return crcResult
}
