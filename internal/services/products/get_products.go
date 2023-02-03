package internal

import (
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func GetProducts() []model.Product {
	var products []model.Product
	DB.Find(&products)
	return products
}
