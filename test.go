package main

import (
	"encoding/json"
	"fmt"

	"github.com/rojack96/gonika/codecs"
)

func main() {

	imei, found, err := codecs.ImeiChecker("000F333536333037303432343431303133")
	if err != nil {
		fmt.Println("Error checking IMEI:", err)
		return
	}
	if found {
		fmt.Println("IMEI found:", string(imei))
	} else {
		fmt.Println("IMEI not found in the data.")
	}

	c, err := codecs.DecoderFactory("000000000000003608010000016B40D8EA30010000000000000000000000000000000105021503010101425E0F01F10000601A014E0000000000000000010000C7CF")
	if err != nil {
		fmt.Println(err)
		return
	}

	t := c.Decode()

	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}
