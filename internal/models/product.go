package internal

import (
	"gorm.io/gorm"
)

type MultiString []string

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
	CategoryID  uint  `json:"category_id"`
	Spec        Spec `json:"spec"`
	Skus        []Sku `json:"skus,omitempty"`
	CartItemID  *uint        `json:"cart_item,omitempty"`
	OrderItemID *uint        `json:"order_item,omitempty"`
	Images      []Attachment `json:"images,omitempty" gorm:"polymorphic:Attachment;polymorphicValue:product_images"`
	Medias      []Attachment `json:"medias,omitempty" gorm:"polymorphic:Attachment;polymorphicValue:product_medias"`
	Attachments []Attachment `json:"attachments,omitempty" gorm:"polymorphic:Attachment;"`
}
