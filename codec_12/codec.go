package codec12

import (
	"encoding/hex"
	conf "sipli/device/teltonika/tcp_server/configuration"
	"sipli/device/teltonika/tcp_server/models"

	"go.uber.org/zap"
)

func codec12(responseMessage *models.ResponseMessage, sh *conf.ServiceHub) {

	sh.Log.Info("encode", zap.String("Preamble", hex.EncodeToString(responseMessage.Preamble)))
	sh.Log.Info("encode", zap.String("DataSize", hex.EncodeToString(responseMessage.DataSize)))
	sh.Log.Info("encode", zap.Any("CodecID", responseMessage.CodecID))
	sh.Log.Info("encode", zap.Any("ResponseQuantity1", responseMessage.ResponseQuantity1))
	sh.Log.Info("encode", zap.Any("Type", responseMessage.Type))
	sh.Log.Info("encode", zap.Any("ResponseSize", hex.EncodeToString(responseMessage.ResponseSize)))
	sh.Log.Info("encode", zap.Any("Response", string(responseMessage.Response)))
	sh.Log.Info("encode", zap.Any("ResponseQuantity2", responseMessage.ResponseQuantity2))
	sh.Log.Info("encode", zap.Any("CRC16", hex.EncodeToString(responseMessage.CRC16)))

}
