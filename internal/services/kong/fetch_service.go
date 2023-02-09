package internal

import (
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	"github.com/spf13/viper"
)

type Service struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
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
	return serviceResponse.Data
}
