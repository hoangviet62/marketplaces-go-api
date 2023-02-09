package internal

import "gorm.io/gorm"

type Sku struct {
	gorm.Model
	Description string
	Quantity    int
	Price       float64
	Status      int
	ProductID   uint
	CartItemID  uint
	OrderItemID uint
	Spec        Spec
	Attachment  Attachment `gorm:"polymorphic:Attachment;"`
}
