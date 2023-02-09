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
	status, service, error := kong.CreateService(serviceName)
	if status {
		log.Info("[KONG] SERVICE ", service.Name, " is created")
	} else {
		log.Info("[KONG] ", error)
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
		routeName := strings.Replace(strings.Replace(routePath, "/", "", 1), "/:id", "_id", 1)
		if slices.Contains(createdRoutes, routeName) {
			log.Info("[KONG] ROUTE ", routeName, " already existed")
			continue
		}

		payload := map[string]interface{}{
			"name":    routeName,
			"paths":   []string{routePath},
			"methods": methods,
		}
		status, service, error := kong.CreateRoute(service.Id, payload)
		if status {
			log.Info("[KONG] ROUTE ", service.Name, " is created")
		} else {
			log.Info("[KONG] ROUTE", error)
		}
	}
}

func IsServiceFailed(status bool) string {
	if status {
		return "success"
	} else {
		return "failed"
	}
}
