package cmd

import (
	"github.com/gin-gonic/gin"
	c "github.com/hoangviet62/marketplaces-go-api/internal/controllers"
	routes "github.com/hoangviet62/marketplaces-go-api/internal/routes"
	"github.com/spf13/viper"
)

func StartApiServer() {
	router := gin.Default()
	controllers := c.BaseController{DB: DB}

	// log.Infof("Start listening on 0.0.0.0:%s", viper.GetString("PORT"))
	server := "0.0.0.0:" + viper.GetString("PORT")

	// need to be revised
	routes.PingRoutes(router)

	// PRODUCTS
	router.GET("/products", controllers.GetProducts)
	router.POST("/products", controllers.CreateProduct)
	router.GET("/products/:id", controllers.GetProductById)
	router.PATCH("/products/:id", controllers.UpdateProduct)
	router.DELETE("/products/:id", controllers.UpdateProduct)

	router.Run(server)
}
