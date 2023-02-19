package internal

import "gorm.io/gorm"

type Sku struct {
	gorm.Model
	Description string `json:"description"`
	Quantity    uint `json:"quantity"`
	Price       float64 `json:"price"`
	Status      uint `json:"status"`
	ProductID   uint `json:"product_id"`
	CartItemID  *uint `json:"cart_item_id,omitempty"`
	OrderItemID *uint `json:"order_item_id,omitempty"`
	Spec        Spec `json:"spec,omitempty"`
	Images      []Attachment `json:"images,omitempty" gorm:"polymorphic:Attachment;polymorphicValue:SkuImages"`
	Medias      []Attachment `json:"medias,omitempty" gorm:"polymorphic:Attachment;polymorphicValue:SkuMedias"`
	Attachments []Attachment `json:"attachments,omitempty" gorm:"polymorphic:Attachment;"`
}
