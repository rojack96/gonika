package codec8

import "github.com/rojack96/gonika/models"

type AvlDataArray[AvlIdSize uint8 | uint16] struct {
	models.AvlDataArray[AvlIdSize]
	AVLData []models.AvlData[AvlIdSize] `json:"avl_data"`
}
