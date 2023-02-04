package internal

import "gorm.io/gorm"

type Sku struct {
	gorm.Model
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Status      int     `json:"status"`
	ProductID   int
	Product     Product
	SpecID      int
	Spec        Spec
}
