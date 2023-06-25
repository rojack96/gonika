package codec8e

import (
	// golang import

	// project import
	parser "sipli/device/teltonika/tcp_server/codecs"
	conf "sipli/device/teltonika/tcp_server/configuration"
	models "sipli/device/teltonika/tcp_server/models"
	"strings"

	// external import
	"go.uber.org/zap"
)

func codec8E(avlDataPacket *models.AVLDataPacket, sh *conf.ServiceHub) models.AVLDataArray {
	sh.Log.Debug("Codec8E function")

	var body []byte
	var avlDataArray models.AVLDataArray

	avlDataArray.CodecID = avlDataPacket.CodecID
	avlDataArray.NumberOfData1 = avlDataPacket.NumberOfData1
	avlDataArray.NumberOfData2 = avlDataPacket.NumberOfData2

	body = avlDataPacket.AvlData

	startIndex := 0

	if avlDataArray.NumberOfData1 == avlDataArray.NumberOfData2 {
		for i := 0; i < int(avlDataArray.NumberOfData1); i++ {
			var IORawData []string
			avlData := models.AVLDataCodec8E{}

			avlData.IMEI = avlDataPacket.IMEI
			avlData.RawData.IMEI = avlData.IMEI
			avlData.RawData.CodecID = "8e"
			sh.Log.Debug("Codec 8 Extended",
				zap.Uint64("avlData.RawData.IMEI", avlData.RawData.IMEI),
				zap.String("avlData.RawData.CodecID", avlData.RawData.CodecID),
			)

			var timestampEndIndex int
			avlData.Timestamp, avlData.RawData.Timestamp, timestampEndIndex = parser.ParseTimestamp(startIndex, body)
			sh.Log.Debug("Codec 8 Extended",
				zap.Any("avlData.Timestamp", avlData.Timestamp),
				zap.String("avlData.RawData.Timestamp", avlData.RawData.Timestamp),
				zap.Int("timestampEndIndex", timestampEndIndex),
			)

			var priorityIndex int
			avlData.Priority, priorityIndex = parser.ParsePriority(timestampEndIndex, body)
			avlData.RawData.Priority = avlData.Priority
			sh.Log.Debug("Codec 8 Extended",
				zap.Any("avlData.Priority", avlData.Priority),
				zap.Uint8("avlData.RawData.Priority", avlData.RawData.Priority),
				zap.Int("priorityIndex", priorityIndex),
			)

			var gpsEndIndex int
			avlData.GPSElement, avlData.RawData.Gps, gpsEndIndex = parser.ParseGPSElement(priorityIndex, body)
			sh.Log.Debug("Codec 8 Extended",
				zap.Any("avlData.GPSElement", avlData.GPSElement),
				zap.String("avlData.RawData.Gps", avlData.RawData.Gps),
				zap.Int("gpsEndIndex", gpsEndIndex),
			)

			var eventIOIDIndexEnd int
			avlData.EventIOID, eventIOIDIndexEnd = parseEventIO(gpsEndIndex, body)
			avlData.RawData.EventIOID = uint(avlData.EventIOID)
			sh.Log.Debug("Codec 8 Extended",
				zap.Any("avlData.EventIOID", avlData.EventIOID),
				zap.Uint("avlData.RawData.EventIOID", avlData.RawData.EventIOID),
				zap.Int("eventIOIDIndexEnd", eventIOIDIndexEnd),
			)

			var noOfTotalIOIndexEnd int
			avlData.NoOfTotalIO, noOfTotalIOIndexEnd = parseTotalNumberOfIO(eventIOIDIndexEnd, body)
			sh.Log.Debug("Codec 8 Extended",
				zap.Any("avlData.NoOfTotalIO", avlData.NoOfTotalIO),
				zap.Int("noOfTotalIOIndexEnd", noOfTotalIOIndexEnd),
			)

			var oneByteIOEndIndex int
			avlData.NoOfOneByte, avlData.OneByteIO, IORawData, oneByteIOEndIndex = parseOneByteIO(noOfTotalIOIndexEnd, body, IORawData)
			sh.Log.Debug("Codec 8 Extended",
				zap.Any("avlData.NoOfOneByte", avlData.NoOfOneByte),
				zap.Any("avlData.OneByteIO", avlData.OneByteIO),
				zap.Int("oneByteIOEndIndex", oneByteIOEndIndex),
			)

			var twoByteIOEndIndex int
			avlData.NoOfTwoByte, avlData.TwoByteIO, IORawData, twoByteIOEndIndex = parseTwoByteIO(oneByteIOEndIndex, body, IORawData)
			sh.Log.Debug("Codec 8 Extended",
				zap.Any("avlData.NoOfTwoByte", avlData.NoOfTwoByte),
				zap.Any("avlData.TwoByteIO", avlData.TwoByteIO),
				zap.Int("twoByteIOEndIndex", twoByteIOEndIndex),
			)

			var fourByteIOEndIndex int
			avlData.NoOfFourByte, avlData.FourByteIO, IORawData, fourByteIOEndIndex = parseFourByteIO(twoByteIOEndIndex, body, IORawData)
			sh.Log.Debug("Codec 8 Extended",
				zap.Any("avlData.NoOfFourByte", avlData.NoOfFourByte),
				zap.Any("avlData.FourByteIO", avlData.FourByteIO),
				zap.Int("fourByteIOEndIndex", fourByteIOEndIndex),
			)

			var eightByteIOEndIndex int
			avlData.NoOfEightByte, avlData.EightByteIO, IORawData, eightByteIOEndIndex = parseEightByteIO(fourByteIOEndIndex, body, IORawData)
			sh.Log.Debug("Codec 8 Extended",
				zap.Any("avlData.NoOfEightByte", avlData.NoOfEightByte),
				zap.Any("avlData.EightByteIO", avlData.EightByteIO),
				zap.Int("eightByteIOEndIndex", eightByteIOEndIndex),
			)

			var xByteIOEndIndex int
			avlData.NoOfXByte, avlData.XByteIO, IORawData, xByteIOEndIndex = parseXByteIO(eightByteIOEndIndex, body, IORawData, sh)
			sh.Log.Debug("Codec 8 Extended",
				zap.Any("avlData.NoOfXByte", avlData.NoOfXByte),
				zap.Any("avlData.XByteIO", avlData.XByteIO),
				zap.Int("xByteIOEndIndex", xByteIOEndIndex),
			)

			startIndex = xByteIOEndIndex
			sh.Log.Debug("Codec 8 Extended",
				zap.Int("startIndex", startIndex),
			)

			avlData.TotalByteIO = mergeMaps(avlData.OneByteIO, avlData.TwoByteIO, avlData.FourByteIO, avlData.EightByteIO, avlData.XByteIO)
			avlData.RawData.IO = strings.Join(IORawData, ",")
			avlDataArray.AVLDataCodec8E = append(avlDataArray.AVLDataCodec8E, avlData)
		}

	}

	return avlDataArray
}

func mergeMaps(n1 map[uint16]uint8, n2 map[uint16]uint16, n4 map[uint16]uint32, n8 map[uint16]uint64, nx map[uint16]uint) map[uint16]uint {
	// This function return a total of N's ID

	totalIO := make(map[uint16]uint)

	if len(nx) != 0 {
		for k, v := range nx {
			totalIO[k] = uint(v)
		}
	}

	if len(n8) != 0 {
		for k, v := range n8 {
			totalIO[k] = uint(v)
		}
	}

	if len(n4) != 0 {
		for k, v := range n4 {
			totalIO[k] = uint(v)
		}
	}

	if len(n2) != 0 {
		for k, v := range n2 {
			totalIO[k] = uint(v)
		}
	}

	if len(n1) != 0 {
		for k, v := range n1 {
			totalIO[k] = uint(v)
		}
	}

	return totalIO
}
