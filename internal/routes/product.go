package internal

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/hoangviet62/marketplaces-go-api/internal/controllers"
)

func ProductRoutes(router *gin.Engine) {
	router.GET("/products", controllers.GetProducts)
	router.POST("/products", controllers.CreateProduct)
	router.GET("/products/:id", controllers.GetProductById)
	router.PATCH("/products/:id", controllers.UpdateProduct)
	router.DELETE("/products/:id", controllers.DeleteProduct)
}
