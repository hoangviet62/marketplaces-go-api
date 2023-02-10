package cmd

import (
	"strings"

	"github.com/gin-gonic/gin"
	kong "github.com/hoangviet62/marketplaces-go-api/internal/services/kong"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
)

func KongMigration(routes []gin.RouteInfo) {
	serviceName := viper.GetString("KONG.SERVICE_NAME")
	plugins := viper.GetStringSlice("KONG.PLUGINS")
	var unsecureRoutes []string
	for _, unsecureRoute := range viper.GetStringSlice("KONG.UNSECURE_ROUTES") {
		unsecureRoutes = append(unsecureRoutes, unsecureRoute)
	}

	status, service, error := kong.CreateService(serviceName)
	if status {
		log.Info("[KONG] SERVICE ", service.Name, " is created")
	} else {
		log.Error("[KONG] ", error)
	}

	var createdRoutes []string
	for _, route := range kong.FetchRoutes(service.Id) {
		createdRoutes = append(createdRoutes, route.Name)
	}

	groupRoutes := make(map[string][]string)
	for _, route := range routes {
		groupRoutes[route.Path] = append(groupRoutes[route.Path], route.Method)
	}

	for routePath, methods := range groupRoutes {
		routeName := strings.Replace(strings.Replace(strings.Replace(routePath, "/", "", 1), "/:id", "_id", 1), "/*", "_", 1)
		if slices.Contains(createdRoutes, routeName) {
			log.Error("[KONG] ROUTE ", routeName, " already existed")
			continue
		}

		payload := map[string]interface{}{
			"name":    strings.Replace(routeName, "/", "-", 1),
			"paths":   []string{routePath},
			"methods": methods,
		}

		status, route, error := kong.CreateRoute(service.Id, payload)
		if status {
			log.Info("[KONG] ROUTE ", route.Name, " is created")
			for _, pluginName := range plugins {
				if slices.Contains(unsecureRoutes, routePath) {
					continue
				}

				status := kong.CreatePlugin("/routes", route.Id, pluginName)
				if status {
					log.Info("[KONG] ASSIGNED PLUGIN ", service.Name, " to ROUTE ", route.Name)
				} else {
					log.Error("[KONG] FAILED TO ASSIGN PLUGIN ", service.Name, " to ROUTE ", route.Name)
				}
			}
		} else {
			log.Error("[KONG] ROUTE ", error)
		}
	}
}
