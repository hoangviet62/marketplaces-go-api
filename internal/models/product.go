package internal

import (
	"gorm.io/gorm"
	"mime/multipart"
)

type MultiString []string

type Product struct {
	gorm.Model
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Tag         string       `json:"tag"`
	CategoryID  int          `json:"category_id"`
	Spec        Spec         `json:"spec"`
	Skus        []Sku        `json:"skus"`
	Attachments []Attachment `json:"attachments" gorm:"polymorphic:Attachment;"`
}

type CreateProductInput struct {
	Name        string                `form:"name" binding:"required"`
	Tag         string                `form:"tag"`
	Description string                `form:"description"`
	Images      *multipart.FileHeader `form:"images"`
	Medias      *multipart.FileHeader `form:"medias"`
}

type UpdateProductInput struct {
	Name        string `json:"name"`
	Tag         string `json:"tag"`
	Description string `json:"description"`
	Images      string `json:"images"`
	Medias      string `json:"medias"`
}
