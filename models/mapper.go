package models

type Mapper struct {
	Data []byte
}

func (m *Mapper) CodecDataSending() AvlDataPacket {
	data := m.Data

	return AvlDataPacket{
		Preamble:        data[0:4],
		DataFieldLength: data[4:8],
		CodecId:         data[8],
		NumberOfData1:   data[9],
		AVLdata:         data[10 : len(data)-8],
		NumberOfData2:   data[len(data)-5],
		CRC16:           data[len(data)-4:],
	}
}
