package internal

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	OrderID   int
	Order     Order
	ProductID int
	Product   Product
	SkuID     int
	Sku       Sku
}
