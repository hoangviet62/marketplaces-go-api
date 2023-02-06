package internal

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Tag         string     `json:"tag"`
	Images      string     `json:"images"`
	Medias      string     `json:"medias"`
	CategoryID  uint       `json:"category_id"`
	CartItemID  uint       `json:"cart_item_id"`
	OrderItemID uint       `json:"order_item_id"`
	Spec        Spec       `json:"spec"`
	Sku         []Sku      `json:"skus"`
	Attachment  Attachment `gorm:"polymorphic:Attachment;"`
}

type CreateProductInput struct {
	Name        string `json:"name"`
	Tag         string `json:"tag"`
	Description string `json:"description"`
	Images      string `json:"images"`
	Medias      string `json:"medias"`
}

type UpdateProductInput struct {
	Name        string `json:"name"`
	Tag         string `json:"tag"`
	Description string `json:"description"`
	Images      string `json:"images"`
	Medias      string `json:"medias"`
}
