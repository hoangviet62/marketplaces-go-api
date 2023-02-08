package internal

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	OrderID   uint    `json:"order_id"`
	Order     Order   `json:"order"`
	ProductID uint    `json:"product_id"`
	Product   Product `json:"product"`
	Sku       Sku     `json:"sku"`
}
