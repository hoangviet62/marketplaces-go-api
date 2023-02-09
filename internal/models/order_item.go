package internal

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	Quantity  int
	Price     float64
	OrderID   uint
	Order     Order
	ProductID uint
	Product   Product
	Sku       Sku
}
