package internal

type Sku struct {
	Base
	Description string       `json:"description"`
	Quantity    uint         `json:"quantity"`
	Price       float64      `json:"price"`
	Status      uint         `json:"status"`
	ProductID   uint         `json:"product_id"`
	Product     *Product     `json:"product,omitempty"`
	SpecID      *uint        `json:"spec_id,omitempty"`
	Spec        *Spec        `json:"spec,omitempty"`
	Images      []Attachment `json:"images,omitempty" gorm:"polymorphic:Attachment;polymorphicValue:sku_images"`
	Medias      []Attachment `json:"medias,omitempty" gorm:"polymorphic:Attachment;polymorphicValue:sku_medias"`
	Attachments []Attachment `json:"attachments,omitempty" gorm:"polymorphic:Attachment;"`
}
