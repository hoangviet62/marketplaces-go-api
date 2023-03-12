package internal

import (

	// "fmt"

	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"

	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	skuService "github.com/hoangviet62/marketplaces-go-api/internal/services/skus"
	specService "github.com/hoangviet62/marketplaces-go-api/internal/services/specs"
)

func CreateProduct(context *gin.Context) (model.Product, error) {
	// Validate input
	// var input model.CreateProductInput
	name := context.PostForm("name")
	description := context.PostForm("description")
	tag := context.PostForm("tag")
	isFeatured, _ := strconv.ParseBool(context.PostForm("is_featured"))
	status := model.ProductStatus(context.PostForm("status"))

	categoryId, _ := strconv.ParseUint(context.PostForm("category_id"), 10, 8)
	product := model.Product{Name: name, Tag: tag, Description: description, CategoryID: uint(categoryId), IsFeatured: isFeatured, Status: status}

	skuParams, isSkuParamsVisible := context.GetPostFormMap("sku")

	if err := helpers.DB.Create(&product).Error; err != nil {
		return product, errors.New(err.Error())
	}

	if isSkuParamsVisible {
		price, _ := strconv.ParseFloat(skuParams["price"], 64)
		quantity, _ := strconv.ParseUint(skuParams["quantity"], 10, 8)
		sku := model.Sku{Description: skuParams["description"], Quantity: uint(quantity), Price: price, ProductID: product.ID}
		_, error := skuService.CreateSku(context, sku)

		if error != nil {
			return product, errors.New(error.Error())
		}
	}

	specParams, isSpecParamsVisible := context.GetPostFormMap("spec")
	if isSpecParamsVisible {
		specDescription := datatypes.JSON([]byte(specParams["description"]))

		spec := model.Spec{Description: specDescription, ProductID: product.ID}

		_, specError := specService.CreateSpec(context, spec)

		if specError != nil {
			return product, errors.New(specError.Error())
		}
	}

	return product, nil
}
