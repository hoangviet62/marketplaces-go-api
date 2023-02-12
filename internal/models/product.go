package internal

import (
	"mime/multipart"

	"gorm.io/gorm"
)

type MultiString []string

type Product struct {
	gorm.Model
	Name        string
	Description string
	Tag         string
	CategoryID  uint
	Spec        Spec
	Skus        []Sku
	CartItemID  *uint        `json:"CartItem,omitempty"`
	OrderItemID *uint        `json:"OrderItem,omitempty"`
	Images      []Attachment `gorm:"polymorphic:Attachment;polymorphicValue:ProductImages"`
	Medias      []Attachment `gorm:"polymorphic:Attachment;polymorphicValue:ProductMedias"`
	Attachments []Attachment `json:"Attachments,omitempty" gorm:"polymorphic:Attachment;"`
}

type CreateProductInput struct {
	Name        string                `form:"name" binding:"required"`
	Tag         string                `form:"tag"`
	Description string                `form:"description"`
	Images      *multipart.FileHeader `form:"images"`
	Medias      *multipart.FileHeader `form:"medias"`
}

type UpdateProductInput struct {
	Name        string                `form:"name" binding:"required"`
	Tag         string                `form:"tag"`
	Description string                `form:"description"`
	Images      *multipart.FileHeader `form:"images"`
	Medias      *multipart.FileHeader `form:"medias"`
	CategoryID  int
	Attachments []Attachment `gorm:"polymorphic:Attachment;"`
}
