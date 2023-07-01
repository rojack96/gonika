package codec8e

import "github.com/rojack96/teltonika-parser/models"

type AVLDataArray struct {
	models.AVLDataArray
	AVLData []AVLData `json:"avl_data"`
}

type AVLData struct {
	models.AVLDataTsToGps
	// Event IO ID (2 bytes) this field defines which IO property has changed and generated an event.
	EventIOID uint16 `json:"event_IO_id"`
	// Number of Total IO (2 bytes) a total number of properties coming with record (N = N1 + N2 + N4 + N8).
	NoOfTotalIO uint16 `json:"-"`
	// Number of One Byte IO (2 bytes) number of properties which length is 1 byte.
	NoOfOneByte uint16 `json:"-"`
	// Number of Two Byte IO (2 bytes) number of properties which length is 2 bytes.
	NoOfTwoByte uint16 `json:"-"`
	// Number of Four Byte IO (2 bytes) number of properties which length is 4 bytes.
	NoOfFourByte uint16 `json:"-"`
	// Number of Eight Byte IO (2 bytes) number of properties which length is 8 bytes.
	NoOfEightByte uint16 `json:"-"`
	// Number of X Byte IO (2 bytes)  a number of properties which length is defined by length element.
	NoOfXByte uint16 `json:"-"`
	// Map id:value with properties which length is 1 byte.
	OneByteIO map[uint16]uint8 `json:"-"`
	// Map id:value with properties which length is 2 bytes.
	TwoByteIO map[uint16]uint16 `json:"-"`
	// Map id:value with properties which length is 4 bytes.
	FourByteIO map[uint16]uint32 `json:"-"`
	// Map id:value with properties which length is 8 bytes.
	EightByteIO map[uint16]uint64 `json:"-"`
	// Map id:value with properties which length is defined by length element.
	XByteIO map[uint16]uint `json:"-"`
}
