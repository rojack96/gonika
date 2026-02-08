package codecs

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/rojack96/gonika/constant"
)

// ImeiChecker This function checks if the data buffer contains a valid IMEI number and returns it if found.
func ImeiChecker(dataBuffer any) ([]byte, bool, error) {
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
			return nil, false, fmt.Errorf("invalid hex string: %v", err)
		}
	default:
		return nil, false, fmt.Errorf("invalid packet type: expected []byte or string, got %T", dataBuffer)
	}

	if len(data) == 0 {
		return nil, false, nil
	}

	if bytes.HasPrefix(data, []byte(constant.ImeiPrefix)) && len(data) == 17 {
		return data[2:], true, nil
	}

	return nil, false, nil
}
