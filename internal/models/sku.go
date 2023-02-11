package internal

import "gorm.io/gorm"

type Sku struct {
	gorm.Model
	ID int32 `gorm:"primaryKey"`

	Description string
	Quantity    int
	Price       float64
	Status      int
	ProductID   int32
	CartItemID  *int32 `json:"CartItem,omitempty"`
	OrderItemID *int32 `json:"OrderItem,omitempty"`
	Spec        Spec
	Images      []Attachment `gorm:"polymorphic:Attachment;polymorphicValue:SkuImages"`
	Medias      []Attachment `gorm:"polymorphic:Attachment;polymorphicValue:SkuMedias"`
	Attachments []Attachment `json:"Attachments,omitempty" gorm:"polymorphic:Attachment;"`
}
