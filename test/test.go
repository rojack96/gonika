package main

import (
	"encoding/json"
	"fmt"

	cdc "github.com/rojack96/gonika/codec"
)

func main() {
	const raw = "000000000000003608010000016B40D8EA30010000000000000000000000000000000105021503010101425E0F01F10000601A014E0000000000000000010000C7CF"
	imei, found, err := cdc.ImeiChecker("000F333536333037303432343431303133")
	if err != nil {
		fmt.Println("Error checking IMEI:", err)
		return
	}
	if found {
		fmt.Println("IMEI found:", string(imei))
	} else {
		fmt.Println("IMEI not found in the data.")
	}

	decoder, err := cdc.AvlDataDecoderFactory(raw)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("decoder", decoder)
	resp := decoder.Decode()
	fmt.Println("resp", resp)
	jsonData, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))

}
