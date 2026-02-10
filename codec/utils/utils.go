package utils

import (
	"encoding/hex"
	"fmt"
)

func TransformData(dataBuffer any) ([]byte, error) {
	switch v := dataBuffer.(type) {

	case []byte:
		return v, nil

	case string:
		data, err := hex.DecodeString(v)
		if err != nil {
			return nil, fmt.Errorf("invalid hex string: %w", err)
		}
		return data, nil

	default:
		return nil, fmt.Errorf(
			"invalid packet type: expected []byte or hex string, got %T",
			dataBuffer,
		)
	}
}
