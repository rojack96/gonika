package constant

// Codecs
const (
	CODEC_8  = 8
	CODEC_8E = 142
	CODEC_16 = 16
	CODEC_12 = 12
	CODEC_13 = 13
	CODEC_14 = 14
)

const (
	PREAMBLE         = "00000000"
	TYPE_COMMAND     = "05"
	TYPE_RESPONSE    = "06"
	COMMAND_QUANTITY = "01"
)

const (
	IMEI_PREFIX        = "\x00\x0F"
	DATA_PACKET_PREFIX = "\x00\x00\x00\x00"
)
