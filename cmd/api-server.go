package cmd

import (
	"github.com/gin-gonic/gin"
	routes "github.com/hoangviet62/marketplaces-go-api/internal/routes"
	"github.com/spf13/viper"
	"net/http"
)

func StartApiServer() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	server := "0.0.0.0:" + viper.GetString("PORT")
	router.StaticFS("/public", http.Dir("public"))
	// need to be revised
	routes.PingRoutes(router)
	routes.ProductRoutes(router)
	routes.CategoryRoutes(router)
	routes.AttachmentRoutes(router)

	router.Run(server)
}
