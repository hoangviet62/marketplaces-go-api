package cmd

import (
	"strings"

	"github.com/gin-gonic/gin"
	kong "github.com/hoangviet62/marketplaces-go-api/internal/services/kong"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func KongMigration(routes []gin.RouteInfo) {
	plugins := viper.GetStringSlice("KONG.PLUGINS")
	for _, plugin := range plugins {
		pluginStatus := kong.CreatePlugin(plugin)
		log.Info("[KONG] PLUGIN MIGRATION ", plugin, " ", IsServiceFailed(pluginStatus))
	}

	serviceName := viper.GetString("KONG.SERVICE_NAME")
	serviceStatus := kong.CreateService(serviceName)
	log.Info("[KONG] SERVICE MIGRATION ", serviceName, " ", IsServiceFailed(serviceStatus))

	groupRoutes := make(map[string][]string)
	for _, route := range routes {
		groupRoutes[route.Path] = append(groupRoutes[route.Path], route.Method)
	}

	for routePath, methods := range groupRoutes {
		payload := map[string]interface{}{
			"name":    strings.Replace(strings.Replace(routePath, "/", "", 1), "/:id", "_id", 1),
			"paths":   []string{routePath},
			"methods": methods,
		}
		routeStatus := kong.CreateRoute(serviceName, payload)
		log.Info("[KONG] ROUTE MIGRATION ", routePath, " with ", methods, " ", IsServiceFailed(routeStatus))
	}
}

func IsServiceFailed(status bool) string {
	if status {
		return "success"
	} else {
		return "failed"
	}
}
