package internal

import (
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	"github.com/spf13/viper"
)

type Route struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type FetchRouteResponse struct {
	Data []Route
}

func FetchRoutes(serviceId string) (routes []Route) {
	path := "/services/" + serviceId + "/routes"
	url := viper.GetString("KONG.ADMIN_URL") + path
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	routesResponse := FetchRouteResponse{}
	helpers.RestClient("GET", url, headers, nil, &routesResponse)
	return routesResponse.Data
}
