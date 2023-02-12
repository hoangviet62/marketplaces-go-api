package internal

import "gorm.io/gorm"

type Sku struct {
	gorm.Model
	Description string
	Quantity    uint
	Price       float64
	Status      uint
	ProductID   uint
	CartItemID  *uint `json:"CartItem,omitempty"`
	OrderItemID *uint `json:"OrderItem,omitempty"`
	Spec        Spec
	Images      []Attachment `gorm:"polymorphic:Attachment;polymorphicValue:SkuImages"`
	Medias      []Attachment `gorm:"polymorphic:Attachment;polymorphicValue:SkuMedias"`
	Attachments []Attachment `json:"Attachments,omitempty" gorm:"polymorphic:Attachment;"`
}
