package internal

import "gorm.io/gorm"

type Sku struct {
	gorm.Model
	Description string     `json:"description"`
	Quantity    int        `json:"quantity"`
	Price       float64    `json:"price"`
	Status      int        `json:"status"`
	ProductID   uint       `json:"product_id"`
	CartItemID  uint       `json:"cart_item_id"`
	OrderItemID uint       `json:"order_item_id"`
	Spec        Spec       `json:"spec"`
	Attachment  Attachment `gorm:"polymorphic:Attachment;"`
}
