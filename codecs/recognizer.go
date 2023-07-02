package codecs

import (
	"bytes"
	"errors"

	"github.com/rojack96/gotlk/constant"
)

// Analyze buffer passed, returns the following values:
//
// 15=IMEI
//
// 8=Codec 8
//
// 142=Codec 8 Extended
//
// 16=Codec 16
//
// 12=Codec 12
//
// 13=Codec 13
//
// 14=Codec 14
func BufferAnalyzer(dataBuffer *[]byte) (*uint8, error) {
	var codec *uint8
	var imei *uint8
	data := *dataBuffer

	if bytes.HasPrefix(*dataBuffer, []byte(constant.IMEI_PREFIX)) {
		if len(*dataBuffer) == 17 {
			*imei = 15
			return imei, nil
		}
	} else if bytes.HasPrefix(data, []byte(constant.DATA_PACKET_PREFIX)) {
		codec = &data[8]

		return codec, nil
	}

	return nil, errors.New("buffer unanalyzable")
}
