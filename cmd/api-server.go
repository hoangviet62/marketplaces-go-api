package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/hoangviet62/marketplaces-go-api/internal/routes"
	"github.com/spf13/viper"
)

func StartApiServer() {
	router := gin.Default()

	// log.Infof("Start listening on 0.0.0.0:%s", viper.GetString("PORT"))
	server := "0.0.0.0:" + viper.GetString("PORT")

	// need to be revised
	routes.PingRoutes(router)
	routes.ProductRoutes(router)

	router.Run(server)
}
