package codec8

import "github.com/rojack96/teltonika-parser/models"

type AVLDataArray struct {
	models.AVLDataArray
	AVLData []AVLData `json:"avl_data"`
}

type AVLData struct {
	models.AVLDataTsToGps
	// Event IO ID (1 byte) this field defines which IO property has changed and generated an event.
	EventIOID uint8 `json:"event_IO_id"`
	// Number of Total IO (1 byte) a total number of properties coming with record (N = N1 + N2 + N4 + N8).
	NoOfTotalIO uint8 `json:"n_of_tolat_IO"`
	// Number of One Byte IO (1 byte) number of properties which length is 1 byte.
	NoOfOneByte uint8 `json:"-"`
	// Number of Two Byte IO (1 byte) number of properties which length is 2 bytes.
	NoOfTwoByte uint8 `json:"-"`
	// Number of Four Byte IO (1 byte) number of properties which length is 4 bytes.
	NoOfFourByte uint8 `json:"-"`
	// Number of Eight Byte IO (1 byte) number of properties which length is 8 bytes.
	NoOfEightByte uint8 `json:"-"`
	// Map id:value with properties which length is 1 byte.
	OneByteIO map[uint8]uint8 `json:"-"`
	// Map id:value with properties which length is 2 bytes.
	TwoByteIO map[uint8]uint16 `json:"-"`
	// Map id:value with properties which length is 4 bytes.
	FourByteIO map[uint8]uint32 `json:"-"`
	// Map id:value with properties which length is 8 bytes.
	EightByteIO map[uint8]uint64 `json:"-"`
}
