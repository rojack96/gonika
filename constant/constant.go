package constant

// Codecs
const (
	Codec8  = 8
	Codec8e = 142
	Codec16 = 16
	Codec12 = 12
	Codec13 = 13
	Codec14 = 14
)

const (
	Preamble        = "00000000"
	TypeCommand     = "05"
	TypeResponse    = "06"
	CommandQuantity = "01"
)

const (
	ImeiPrefix       = "\x00\x0F"
	DataPacketPrefix = "\x00\x00\x00\x00"
)
