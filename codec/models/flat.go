package models

type AvlDataFlat interface {
	isAvlDataFlat()
}

func (AvlData8Flat) isAvlDataFlat()    {}
func (AvlData8extFlat) isAvlDataFlat() {}
func (AvlData16Flat) isAvlDataFlat()   {}

type AvlDataPacketFlat struct {
	Preamble        Preamble        `json:"preamble"`
	DataFieldLength DataFieldLength `json:"dataFieldLength"`
	CodecID         CodecID         `json:"codecId"`
	NumberOfData1   NumberOfData    `json:"numberOfData1"`
	AvlData         []AvlDataFlat   `json:"avlData"`
	NumberOfData2   NumberOfData    `json:"numberOfData2"`
	Crc16           Crc16           `json:"crc16"`
}

type AvlData8Flat struct {
	Timestamp  Timestamp          `json:"timestamp"`
	Priority   Priority           `json:"priority"`
	Longitude  Longitude          `json:"longitude"`
	Latitude   Latitude           `json:"latitude"`
	Altitude   Altitude           `json:"altitude"`
	Angle      Angle              `json:"angle"`
	Satellites Satellites         `json:"satellites"`
	Speed      Speed              `json:"speed"`
	Elements   []map[uint8]string `json:"elements"`
}

type AvlData8extFlat struct {
	Timestamp  Timestamp           `json:"timestamp"`
	Priority   Priority            `json:"priority"`
	Longitude  Longitude           `json:"longitude"`
	Latitude   Latitude            `json:"latitude"`
	Altitude   Altitude            `json:"altitude"`
	Angle      Angle               `json:"angle"`
	Satellites Satellites          `json:"satellites"`
	Speed      Speed               `json:"speed"`
	Elements   []map[uint16]string `json:"elements"`
}

type AvlData16Flat struct {
	Timestamp  Timestamp           `json:"timestamp"`
	Priority   Priority            `json:"priority"`
	Longitude  Longitude           `json:"longitude"`
	Latitude   Latitude            `json:"latitude"`
	Altitude   Altitude            `json:"altitude"`
	Angle      Angle               `json:"angle"`
	Satellites Satellites          `json:"satellites"`
	Speed      Speed               `json:"speed"`
	Elements   []map[uint16]string `json:"elements"`
}
