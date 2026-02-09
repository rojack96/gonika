package codecs

import (
	"encoding/json"
	"fmt"

	"github.com/rojack96/gonika/codec"
)

func main() {

	imei, found, err := codec.ImeiChecker("000F333536333037303432343431303133")
	if err != nil {
		fmt.Println("Error checking IMEI:", err)
		return
	}
	if found {
		fmt.Println("IMEI found:", string(imei))
	} else {
		fmt.Println("IMEI not found in the data.")
	}

	c, err := codec.DecoderFactory("000000000000003608010000016B40D8EA3001000000000000000000000001425E5F2C77C93A8A23C7CF")
	if err != nil {
		fmt.Println(err)
		return
	}

	t := c.(codec.AvlDecoder).Decode()

	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}
