package codec8e

import (
	// golang import
	"encoding/binary"
	"encoding/hex"
	"fmt"

	// project import
	conf "sipli/device/teltonika/tcp_server/configuration"

	"go.uber.org/zap"
	// external import
)

func logger(sh *conf.ServiceHub) *zap.Logger {
	return sh.Log
}

// This function parse the Event IO data from AVL data.
//
// Event IO changed
// if data is acquired on event
// this field defines which IO property has changed and generated an event.
func parseEventIO(startIndex int, body []byte) (uint16, int) {
	eventIOIDIndexStart := startIndex
	eventIOIDIndexEnd := eventIOIDIndexStart + 2
	eventIOID := binary.BigEndian.Uint16(body[eventIOIDIndexStart:eventIOIDIndexEnd])

	return eventIOID, eventIOIDIndexEnd
}

// Total Number of IO.
//
// A total number of properties coming with record (N = N1 + N2 + N4 + N8).
func parseTotalNumberOfIO(startIndex int, body []byte) (uint16, int) {
	noOfTotalIOIndexStart := startIndex
	noOfTotalIOIndexEnd := noOfTotalIOIndexStart + 2
	noOfTotalIO := binary.BigEndian.Uint16(body[noOfTotalIOIndexStart:noOfTotalIOIndexEnd])

	return noOfTotalIO, noOfTotalIOIndexEnd
}

// This function parse one byte IO.
//
// N1
// number of properties, which length is 1 byte.
func parseOneByteIO(startIndex int, body []byte, IORawData []string) (uint16, map[uint16]uint8, []string, int) {
	oneByteIO := map[uint16]uint8{}

	// One Byte IO Number
	noOfOneByteIOIndexStart := startIndex
	noOfOneByteIOIndexEnd := noOfOneByteIOIndexStart + 2
	noOfOneByteIO := binary.BigEndian.Uint16(body[noOfOneByteIOIndexStart:noOfOneByteIOIndexEnd])
	// One Byte IO Data
	oneByteIOStartIndex := noOfOneByteIOIndexEnd
	oneByteIOEndIndex := oneByteIOStartIndex + int(noOfOneByteIO)*3
	data := body[oneByteIOStartIndex:oneByteIOEndIndex]

	for i := 0; i < len(data); i += 3 {
		id := binary.BigEndian.Uint16(data[i : i+2])
		value := data[i+2]

		oneByteIO[id] = value

		IORawData = append(IORawData, fmt.Sprintf("%s=%s", hex.EncodeToString(data[i:i+2]), hex.EncodeToString([]byte{value})))
	}

	return noOfOneByteIO, oneByteIO, IORawData, oneByteIOEndIndex
}

// This function parse two byte IO.
//
// N2
// number of properties, which length is 2 byte.
func parseTwoByteIO(startIndex int, body []byte, IORawData []string) (uint16, map[uint16]uint16, []string, int) {
	twoByteIO := map[uint16]uint16{}

	// Two Byte IO Number
	noOfTwoByteIOIndexStart := startIndex
	noOfTwoByteIOIndexEnd := noOfTwoByteIOIndexStart + 2
	noOfTwoByteIO := binary.BigEndian.Uint16(body[noOfTwoByteIOIndexStart:noOfTwoByteIOIndexEnd])
	// Two Byte IO Data
	twoByteIOStartIndex := noOfTwoByteIOIndexEnd
	twoByteIOEndIndex := twoByteIOStartIndex + int(noOfTwoByteIO)*4
	data := body[twoByteIOStartIndex:twoByteIOEndIndex]

	for i := 0; i < len(data); i += 4 {
		id := binary.BigEndian.Uint16(data[i : i+2])
		value := binary.BigEndian.Uint16(data[i+2 : i+4])

		twoByteIO[id] = value

		IORawData = append(IORawData, fmt.Sprintf("%s=%s", hex.EncodeToString(data[i:i+2]), hex.EncodeToString(data[i+2:i+4])))
	}

	return noOfTwoByteIO, twoByteIO, IORawData, twoByteIOEndIndex
}

// This function parse four byte IO.
//
// N4
// number of properties, which length is 4 byte.
func parseFourByteIO(startIndex int, body []byte, IORawData []string) (uint16, map[uint16]uint32, []string, int) {
	fourByteIO := map[uint16]uint32{}

	// Four Byte IO Number
	noOfFourByteIOIndexStart := startIndex
	noOfFourByteIOIndexEnd := noOfFourByteIOIndexStart + 2
	noOfFourByteIO := binary.BigEndian.Uint16(body[noOfFourByteIOIndexStart:noOfFourByteIOIndexEnd])
	// Four Byte IO Number
	fourByteIOStartIndex := noOfFourByteIOIndexEnd
	fourByteIOEndIndex := fourByteIOStartIndex + int(noOfFourByteIO)*6
	data := body[fourByteIOStartIndex:fourByteIOEndIndex]

	for i := 0; i < len(data); i += 6 {
		id := binary.BigEndian.Uint16(data[i : i+2])
		value := binary.BigEndian.Uint32(data[i+2 : i+6])

		fourByteIO[id] = value

		IORawData = append(IORawData, fmt.Sprintf("%s=%s", hex.EncodeToString(data[i:i+2]), hex.EncodeToString(data[i+2:i+6])))
	}

	return noOfFourByteIO, fourByteIO, IORawData, fourByteIOEndIndex
}

// This function parse eight byte IO.
//
// N8
// number of properties, which length is 8 byte.
// Eight Byte IO Number
func parseEightByteIO(startIndex int, body []byte, IORawData []string) (uint16, map[uint16]uint64, []string, int) {
	eightByteIO := map[uint16]uint64{}

	// Eight Byte IO Number
	noOfEightByteIOIndexStart := startIndex
	noOfEightByteIOIndexEnd := noOfEightByteIOIndexStart + 2
	noOfEightByteIO := binary.BigEndian.Uint16(body[noOfEightByteIOIndexStart:noOfEightByteIOIndexEnd])
	// Eight Byte IO Data
	eightByteIOStartIndex := noOfEightByteIOIndexEnd
	eightByteIOEndIndex := eightByteIOStartIndex + int(noOfEightByteIO)*10
	data := body[eightByteIOStartIndex:eightByteIOEndIndex]

	for i := 0; i < len(data); i += 10 {
		id := binary.BigEndian.Uint16(data[i : i+2])
		value := binary.BigEndian.Uint64(data[i+2 : i+10])

		eightByteIO[id] = value

		IORawData = append(IORawData, fmt.Sprintf("%s=%s", hex.EncodeToString(data[i:i+2]), hex.EncodeToString(data[i+2:i+10])))
	}

	return noOfEightByteIO, eightByteIO, IORawData, eightByteIOEndIndex
}

// This function parse X byte IO.
//
// NX
// a number of properties, which length is defined by length element.
// X Byte IO Number
func parseXByteIO(startIndex int, body []byte, IORawData []string, sh *conf.ServiceHub) (uint16, map[uint16]uint, []string, int) {
	xByteIO := map[uint16]uint{}

	// Eight Byte IO Number
	noOfXByteIOIndexStart := startIndex
	sh.Log.Debug("noOfXByteIOIndexStart", zap.Any("x", noOfXByteIOIndexStart))
	noOfXByteIOIndexEnd := noOfXByteIOIndexStart + 2
	sh.Log.Debug("noOfXByteIOIndexEnd", zap.Any("x", noOfXByteIOIndexEnd))
	noOfXByteIO := binary.BigEndian.Uint16(body[noOfXByteIOIndexStart:noOfXByteIOIndexEnd])
	sh.Log.Debug("noOfXByteIO", zap.Any("x", noOfXByteIO))
	// Eight Byte IO Data
	if noOfXByteIO != 0 {
		xByteIOStartIndex := noOfXByteIOIndexEnd

		var i uint16 = 0
		j := 0

		for i < noOfXByteIO {
			xIOIDStartIndex := j + xByteIOStartIndex
			sh.Log.Debug("xIOIDStartIndex", zap.Any("x", xIOIDStartIndex))
			xIOIDEndIndex := xIOIDStartIndex + 2
			sh.Log.Debug("xIOIDEndIndex", zap.Any("x", xIOIDEndIndex))

			id := binary.BigEndian.Uint16(body[xIOIDStartIndex:xIOIDEndIndex])
			sh.Log.Debug("id", zap.Any("x", id))

			xValueLengthStartIndex := xIOIDEndIndex
			sh.Log.Debug("xValueLengthStartIndex", zap.Any("x", xValueLengthStartIndex))
			xValueLengthEndIndex := xValueLengthStartIndex + 2
			sh.Log.Debug("xValueLengthEndIndex", zap.Any("x", xValueLengthEndIndex))

			valueLength := binary.BigEndian.Uint16(body[xValueLengthStartIndex:xValueLengthEndIndex])
			sh.Log.Debug("valueLength", zap.Any("v", valueLength))

			if valueLength == 0 {
				return noOfXByteIO, xByteIO, IORawData, xValueLengthEndIndex
			}

			xValueStartIndex := xValueLengthEndIndex
			sh.Log.Debug("xValueStartIndex", zap.Any("x", xValueStartIndex))
			xValueEndIndex := xValueStartIndex + int(valueLength)
			sh.Log.Debug("xValueEndIndex", zap.Any("x", xValueEndIndex))

			value := binary.BigEndian.Uint64(body[xValueStartIndex:xValueEndIndex])

			xByteIO[id] = uint(value)

			IORawData = append(IORawData, fmt.Sprintf("%s=%s", hex.EncodeToString(body[xIOIDStartIndex:xIOIDEndIndex]), hex.EncodeToString(body[xValueStartIndex:xValueEndIndex])))
			sh.Log.Debug("xByteIO", zap.Any("x", xByteIO))
			i = i + 1
			j = j + xValueEndIndex
		}

		xByteIOEndIndex := j
		return noOfXByteIO, xByteIO, IORawData, xByteIOEndIndex
	}

	return noOfXByteIO, xByteIO, IORawData, startIndex + 2
}
