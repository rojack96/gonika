package utils

import (
	"encoding/hex"
	"fmt"
)

func TransformData(dataBuffer any) ([]byte, error) {
	var (
		data []byte
		err  error
	)
	// Convert string to []byte if necessary
	switch v := dataBuffer.(type) {
	case []byte:
		data = v
	case string:
		if data, err = hex.DecodeString(v); err != nil {
			return nil, fmt.Errorf("invalid hex string: %v", err)
		}
	default:
		return nil, fmt.Errorf("invalid packet type: expected []byte or string, got %T", dataBuffer)
	}
	return data, nil
}
