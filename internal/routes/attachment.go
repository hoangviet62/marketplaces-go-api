package internal

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/hoangviet62/marketplaces-go-api/internal/controllers"
)

func AttachmentRoutes(router *gin.Engine) {
	// router.GET("/products", controllers.GetProducts)
	router.GET("/attachments", controllers.GetAttachments)
	// router.GET("/products/:id", controllers.GetProductById)
	// router.PATCH("/products/:id", controllers.UpdateProduct)
	// router.DELETE("/products/:id", controllers.UpdateProduct)
}
