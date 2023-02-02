package cmd

import (
	"github.com/gin-gonic/gin"
	routes "github.com/hoangviet62/marketplaces-go-api/internal/routes"
	"github.com/spf13/viper"
)

func StartApiServer() {
	router := gin.Default()
	server := "0.0.0.0:" + viper.GetString("PORT")

	// need to be revised
	routes.PingRoutes(router)
	routes.ProductRoutes(router)

	router.Run(server)
}
