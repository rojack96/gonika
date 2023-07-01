package codec12

import (
	"encoding/hex"
	"strconv"

	"github.com/getrak/crc16"
	models "github.com/rojack96/teltonika-parser/models/codec_12"
)

func dataSize(command *models.CommandMessage) string {
	c := *command

	ds := (len(c.CodecID) / 2) +
		(len(c.CommandQuantity1) / 2) +
		(len(c.Type) / 2) +
		(len(c.CommandSize) / 2) +
		(len(c.Command) / 2) +
		(len(c.CommandQuantity2) / 2)

	dataSizeResult := padLeft(strconv.FormatInt(int64(ds), 16), 8)

	return dataSizeResult
}

func commandBuilder(command string) string {
	return hex.EncodeToString([]byte(command))
}

func commandSize(command string) string {
	cs := (len(command) / 2)

	commandSizeResult := padLeft(strconv.FormatInt(int64(cs), 16), 8)

	return commandSizeResult
}

func crc16builder(command *models.CommandMessage) string {
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

func padLeft(str string, length int) string {
	// ad zero's on the left
	for len(str) < length {
		str = "0" + str
	}
	return str
}
