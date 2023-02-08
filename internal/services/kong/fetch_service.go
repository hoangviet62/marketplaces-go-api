package internal

import (
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	"github.com/spf13/viper"
)

type Service struct {
	Name string `json:"name"`
}

type FetchServiceResponse struct {
	Data []Service
}

func FetchServices() (services []Service) {
	path := "/services"
	url := viper.GetString("KONG.ADMIN_URL") + path
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	serviceResponse := FetchServiceResponse{}
	helpers.RestClient("GET", url, headers, nil, &serviceResponse)
	services = serviceResponse.Data

	return services
}
