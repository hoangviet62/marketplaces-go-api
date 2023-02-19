package internal

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	Quantity uint `json:"quantity"`
	Price    float64 `json:"price"`
	CartID   uint `json:"cart_id"`
	Cart     Cart `json:"cart"`
	Product  Product `json:"product"`
	Sku      Sku `json:"sku"`
}
