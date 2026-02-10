package main

import (
	"encoding/json"
	"fmt"

	cdc "github.com/rojack96/gonika/codec"
)

func main() {
	const raw = "000000000000003608010000016B40D8EA30010000000000000000000000000000000105021503010101425E0F01F10000601A014E0000000000000000010000C7CF"
	const raw8ext = "00000000000000A98E020000017357633410000F0DC39B2095964A00AC00F80B00000000000B000500F00100150400C8000" +
		"04501007156000500B5000500B600040018000000430FE00044011B000100F10000601B000000000000017357633BE1000F0DC39B209" +
		"5964A00AC00F80B000001810001000000000000000000010181002D11213102030405060708090A0B0C0D0E0F104545010ABC2121020" +
		"30405060708090A0B0C0D0E0F10020B010AAD020000BF30"
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

	decoder, err := cdc.AvlDataDecoderFactory(raw8ext)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println("decoder", decoder)
	resp := decoder.Decode()
	fmt.Println("resp", resp)
	jsonData, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))

}
