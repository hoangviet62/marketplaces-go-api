package internal

import (
	"gorm.io/gorm"
	"mime/multipart"
)

type MultiString []string

type Product struct {
	gorm.Model
	Name        string
	Description string
	Tag         string
	CategoryID  int
	Spec        Spec
	Skus        []Sku
	Images      []Attachment `gorm:"polymorphic:Attachment;polymorphicValue:ProductImages"`
	Medias      []Attachment `gorm:"polymorphic:Attachment;polymorphicValue:ProductMedias"`
	Attachments []Attachment `gorm:"polymorphic:Attachment;"`
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
