package codec8

// import (
// 	// golang import

// 	// project import
// 	"encoding/hex"
// 	api "sipli/device/teltonika/tcp_server/api"
// 	c12 "sipli/device/teltonika/tcp_server/codecs/codec_12"
// 	c8 "sipli/device/teltonika/tcp_server/codecs/codec_8"
// 	c8e "sipli/device/teltonika/tcp_server/codecs/codec_8E"
// 	conf "sipli/device/teltonika/tcp_server/configuration"
// 	models "sipli/device/teltonika/tcp_server/models"

// 	// external import

// 	"go.uber.org/zap"
// )

// func responseParser(responseMessage *[]byte, IMEI uint64, sh *conf.ServiceHub) byte {

// 	var response models.ResponseMessage

// 	data := *responseMessage
// 	sh.Log.Info("responseMessageComplete", zap.String("data", hex.EncodeToString(data)))
// 	response.IMEI = IMEI
// 	response.Preamble = data[0:4]
// 	response.DataSize = data[4:8]
// 	response.CodecID = data[8]
// 	response.ResponseQuantity1 = data[9]
// 	response.Type = data[10]
// 	response.ResponseSize = data[11:15]
// 	response.Response = data[15 : len(data)-5]
// 	response.ResponseQuantity2 = data[len(data)-5]
// 	response.CRC16 = data[len(data)-4:]

// 	if response.CodecID == 12 { // 12 corresponds to Codec12

// 		// codec8parsed := c12.Codec12(&response, sh)
// 		c12.Codec12(&response, sh)
// 		// go api.SaveAvlData(codec8parsed.AVLDataCodec8, IMEI, sh)

// 	}

// 	result := response.ResponseQuantity1

// 	return result
// }

// func commandParser(c []byte, sh *conf.ServiceHub) map[string]string {
// 	sh.Log.Info("Entro")
// 	singleCommand := make(map[string]string)

// 	singleCommand[string(c[:15])] = string(c[15:])

// 	return singleCommand
// }
