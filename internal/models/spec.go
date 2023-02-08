package internal

import "gorm.io/gorm"

type Spec struct {
	gorm.Model
	Description string `json:"description"`
	SkuID       int    `json:"sku_id"`
	ProductID   uint   `json:"product_id"`
}
