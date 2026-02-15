package main

import (
	"fmt"

	cdc "github.com/rojack96/gonika/codec"
)

func main() {

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

	//codec8()
	//codec8ext()
	//codec16()
	//codec12Response()
	codec13Response()
	//codec14Response()
}

func codec8() {
	// TODO use encoder instead of raw
	const raw = "000000000000003608010000016B40D8EA30010000000000000000000000000000000105021503010101425E0F01F10000601A014E0000000000000000010000C7CF"
	decoder, err := cdc.DeviceDataSendingDecoderFactory(raw)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp := decoder.DecodeTCPflat()
	jsonData, err := resp.MarshalIndent("", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}

func codec8ext() {
	// TODO use encoder instead of raw
	const raw = "00000000000000A98E020000017357633410000F0DC39B2095964A00AC00F80B00000000000B000500F00100150400C8000" +
		"04501007156000500B5000500B600040018000000430FE00044011B000100F10000601B000000000000017357633BE1000F0DC39B209" +
		"5964A00AC00F80B000001810001000000000000000000010181002D11213102030405060708090A0B0C0D0E0F104545010ABC2121020" +
		"30405060708090A0B0C0D0E0F10020B010AAD020000BF30"
	decoder, err := cdc.DeviceDataSendingDecoderFactory(raw)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp := decoder.DecodeTCPflat()
	jsonData, err := resp.MarshalIndent("", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}

func codec16() {
	// TODO implement codec16
}

func codec12Response() {
	// TODO use encoder instead of raw
	const raw = "00000000000000900C010600000088494E493A323031392F372F323220373A3232205254433A323031392F372F323220373A3533205253543A32204552523A312053523A302042523A302043463A302046473A3020464C3A302054553A302F302055543A3020534D533A30204E4F4750533A303A3330204750533A31205341543A302052533A332052463A36352053463A31204D443A30010000C78F"
	decoder, err := cdc.GprsMessageDecoderFactory(raw)
	if err != nil {
		fmt.Println(err)
		return
	}

	respCmd := decoder.DecodeResponse()
	jsonData, _ := respCmd.MarshalJSON()

	fmt.Println(string(jsonData))
}

func codec13Response() {
	// TODO use encoder instead of raw
	const raw = "000000000000001D0D01060000001564E8328168656C6C6F206C65747320746573740D0A0100003548" // Add raw data for codec13 response
	decoder, err := cdc.GprsMessageDecoderFactory(raw)
	if err != nil {
		fmt.Println(err)
		return
	}

	respCmd := decoder.DecodeResponse()
	jsonData, _ := respCmd.MarshalJSON()

	fmt.Println(string(jsonData))
}

func codec14Response() {
	// TODO use encoder instead of raw
	const raw = "00000000000000AB0E0106000000A303520930814522515665723A30332E31382E31345F3034204750533A41584E5F352E31305F333333332048773A464D42313230204D6F643A313520494D45493A33353230393330383134353232353120496E69743A323031382D31312D323220373A313320557074696D653A3137323334204D41433A363042444430303136323631205350433A312830292041584C3A30204F42443A3020424C3A312E362042543A340100007AAE" // Add raw data for codec14 response
	decoder, err := cdc.GprsMessageDecoderFactory(raw)
	if err != nil {
		fmt.Println(err)
		return
	}

	respCmd := decoder.DecodeResponse()
	jsonData, _ := respCmd.MarshalJSON()

	fmt.Println(string(jsonData))
}
