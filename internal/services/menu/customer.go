package Menu

import (
	"github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

type CustomerMenu struct {
	Id   int8   `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

type CustomerParentMenu struct {
	Id       int8           `json:"id"`
	Name     string         `json:"name"`
	Path     string         `json:"path"`
	SubItems []CustomerMenu `json:"sub_items"`
}

func Customer() []CustomerParentMenu {
	results := []CustomerParentMenu{}

	categories := []model.Category{}
	childCategories := []CustomerMenu{}
	helpers.DB.Model(&model.Category{}).Limit(10).Offset(0).Find(&categories)
	for _, category := range categories {
		child := CustomerMenu{Id: int8(category.ID), Name: category.Name, Path: "/catalog?type=Category&Name=" + category.Name}
		childCategories = append(childCategories, child)
	}
	result := CustomerParentMenu{Id: 1, Name: "Categories", SubItems: childCategories}
	results = append(results, result)

	products := []model.Product{}
	childProducts := []CustomerMenu{}
	helpers.DB.Model(&model.Product{}).Limit(10).Offset(0).Find(&products)
	for _, product := range products {
		child := CustomerMenu{Id: int8(product.ID), Name: product.Name, Path: "/catalog?type=Product&Name=" + product.Name}
		childProducts = append(childProducts, child)
	}
	result = CustomerParentMenu{Id: 2, Name: "Products", SubItems: childProducts}
	results = append(results, result)

	suppliers := []model.Supplier{}
	childSuppliers := []CustomerMenu{}
	helpers.DB.Model(&model.Supplier{}).Limit(10).Offset(0).Find(&suppliers)
	for _, supplier := range suppliers {
		child := CustomerMenu{Id: int8(supplier.ID), Name: supplier.Name, Path: "/catalog?type=Supplier&Name=" + supplier.Name}
		childSuppliers = append(childSuppliers, child)
	}
	result = CustomerParentMenu{Id: 3, Name: "Suppliers", SubItems: childSuppliers}
	results = append(results, result)

	result = CustomerParentMenu{Id: 4, Name: "News", Path: "#", SubItems: []CustomerMenu{}}
	results = append(results, result)

	return results
}
