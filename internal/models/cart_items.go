package internal

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
