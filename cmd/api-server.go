package cmd

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	routes "github.com/hoangviet62/marketplaces-go-api/internal/routes"
	"github.com/spf13/viper"
)

func StartApiServer() {

	server := "0.0.0.0:" + viper.GetString("PORT")

	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.StaticFS("/public", http.Dir("public"))

	// need to be revised
	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.PingRoutes(router)
	routes.ProductRoutes(router)
	routes.CategoryRoutes(router)
	routes.AttachmentRoutes(router)
	routes.SkuRoutes(router)
	routes.BannerRoutes(router)
	routes.CartItemRoutes(router)
	routes.CartRoutes(router)
	routes.MenuRoutes(router)

	// Kong migration for all routes
	shouldMigrate, _ := strconv.ParseBool(viper.GetString("KONG.SHOULD_MIGRATE"))
	if shouldMigrate {
		KongMigration(router.Routes())
	}

	router.Run(server)
}
