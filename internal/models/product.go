package internal

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
	Images      string `json:"images"`
	Medias      string `json:"medias"`
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
