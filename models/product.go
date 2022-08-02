package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	FarmerId        string `json:"farmerId"`
	ProductCategory string `json:"productCategory"`
	ProductName     string `json:"productName"`
	Price           string `json:"price"`
	Status          string `json:"status"`
	Qty             string `json:"qty"`
	Image           string `json:"image"`
}
