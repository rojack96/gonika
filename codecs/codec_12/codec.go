package codec12

import (
	"encoding/hex"
	"fmt"

	"github.com/rojack96/teltonika-parser/models"
)

func Codec12(responseMessage *models.ResponseMessage) {

	fmt.Println("Preamble", hex.EncodeToString(responseMessage.Preamble))
	fmt.Println("DataSize", hex.EncodeToString(responseMessage.DataSize))
	fmt.Println("CodecID", responseMessage.CodecID)
	fmt.Println("ResponseQuantity1", responseMessage.ResponseQuantity1)
	fmt.Println("Type", responseMessage.Type)
	fmt.Println("ResponseSize", hex.EncodeToString(responseMessage.ResponseSize))
	fmt.Println("Response", string(responseMessage.Response))
	fmt.Println("ResponseQuantity2", responseMessage.ResponseQuantity2)
	fmt.Println("CRC16", hex.EncodeToString(responseMessage.CRC16))

}
