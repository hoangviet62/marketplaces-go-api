package internal

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	Quantity int
	Price    float64
	OrderID  int32
	Order    Order
	Product  Product
	Sku      Sku
}
