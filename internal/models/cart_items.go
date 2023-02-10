package internal

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	Quantity int
	Price    float64
	CartID   uint
	Cart     Cart
	Product  Product
	Sku      Sku
}
