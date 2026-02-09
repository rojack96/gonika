package constant

// Codecs
const (
	Codec8    = 0x08
	Codec8ext = 0x8E
	Codec16   = 0x10
	Codec12   = 0x0C
	Codec13   = 0x0D
	Codec14   = 0x0E
)

const (
	Command  = 0x05
	Response = 0x06
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
