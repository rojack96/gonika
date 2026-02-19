package models

type CodecEncoder interface {
	isAvlDataEncoder()
}

func (Codec8Encoder) isAvlDataEncoder()    {}
func (Codec8ExtEncoder) isAvlDataEncoder() {}
func (Codec16Encoder) isAvlDataEncoder()   {}

type Codec8Encoder struct {
	OneByte   map[uint8]uint8
	TwoByte   map[uint8]uint16
	FourByte  map[uint8]uint32
	EightByte map[uint8]uint64
}

type Codec8ExtEncoder struct {
	OneByte   map[uint16]uint8
	TwoByte   map[uint16]uint16
	FourByte  map[uint16]uint32
	EightByte map[uint16]uint64
	XByte     map[uint16]string
}

type Codec16Encoder struct {
	OneByte   map[uint16]uint8
	TwoByte   map[uint16]uint16
	FourByte  map[uint16]uint32
	EightByte map[uint16]uint64
}

type GpsElementEncoder struct {
	Longitude  string `json:"longitude"`
	Latitude   string `json:"latitude"`
	Altitude   uint16 `json:"altitude"`
	Angle      uint16 `json:"angle"`
	Satellites uint8  `json:"satellites"`
	Speed      uint16 `json:"speed"`
}

type AvlDataArrayEncoder struct {
	CodecEncoder
	GpsElementEncoder
}
