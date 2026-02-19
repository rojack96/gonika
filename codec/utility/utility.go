package utility

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rojack96/gonika/codec/models"
	"gopkg.in/yaml.v3"
)

type Data struct {
	Codec         uint8                        `json:"codec" yaml:"codec"`
	AvlDataPacket []models.AvlDataArrayEncoder `json:"avlDataPacket" yaml:"avlDataPacket"`
}

type rawData struct {
	Codec         uint8             `json:"codec" yaml:"codec"`
	AvlDataPacket []json.RawMessage `json:"avlDataPacket" yaml:"avlDataPacket"`
}

func ReadFromFile(path string) (Data, error) {
	var result Data

	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return result, err
	}

	var raw rawData

	switch ext := filepath.Ext(path); ext {
	case ".json":
		if err := json.Unmarshal(fileBytes, &raw); err != nil {
			return result, err
		}
	case ".yaml", ".yml":
		if err := yaml.Unmarshal(fileBytes, &raw); err != nil {
			return result, err
		}
	default:
		return result, errors.New("unsupported file format")
	}

	result.Codec = raw.Codec

	for _, packet := range raw.AvlDataPacket {

		switch raw.Codec {

		case 8:
			var temp struct {
				AvlData models.Codec8Encoder     `json:"avlData" yaml:"avlData"`
				Gps     models.GpsElementEncoder `json:"gps" yaml:"gps"`
			}

			if err := json.Unmarshal(packet, &temp); err != nil {
				return result, err
			}

			result.AvlDataPacket = append(result.AvlDataPacket, models.AvlDataArrayEncoder{
				AvlDataEncoder:    temp.AvlData,
				GpsElementEncoder: temp.Gps,
			})

		case 142:
			var temp struct {
				AvlData models.Codec8ExtEncoder  `json:"avlData" yaml:"avlData"`
				Gps     models.GpsElementEncoder `json:"gps" yaml:"gps"`
			}

			if err := json.Unmarshal(packet, &temp); err != nil {
				return result, err
			}

			result.AvlDataPacket = append(result.AvlDataPacket, models.AvlDataArrayEncoder{
				AvlDataEncoder:    temp.AvlData,
				GpsElementEncoder: temp.Gps,
			})

		case 10:
			var temp struct {
				AvlData models.Codec16Encoder    `json:"avlData" yaml:"avlData"`
				Gps     models.GpsElementEncoder `json:"gps" yaml:"gps"`
			}

			if err := json.Unmarshal(packet, &temp); err != nil {
				return result, err
			}

			result.AvlDataPacket = append(result.AvlDataPacket, models.AvlDataArrayEncoder{
				AvlDataEncoder:    temp.AvlData,
				GpsElementEncoder: temp.Gps,
			})

		default:
			return result, fmt.Errorf("unsupported codec: %d", raw.Codec)
		}
	}

	return result, nil
}
