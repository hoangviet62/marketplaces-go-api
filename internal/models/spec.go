package internal

import "gorm.io/gorm"

type Spec struct {
	gorm.Model
	Description string `json:"description"`
	SkuID       uint `json:"sku_id"`
	ProductID   uint `json:"product_id"`
}
