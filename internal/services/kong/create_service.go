package internal

import (
	"bytes"
	"encoding/json"
	"strconv"

	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	"github.com/spf13/viper"
)

type CreateServiceResponse struct {
	Id           string `json:"Id"`
	Name         string `json:"Name"`
	Enabled      bool   `json:"Enabled"`
	ErrorMessage string `json:"Message"`
}

func CreateService(serviceName string) (status bool) {
	services := FetchServices()
	isExisted := false

	for _, service := range services {
		if serviceName == service.Name {
			isExisted = true
			break
		}
	}

	if isExisted {
		return false
	}

	path := "/services"
	url := viper.GetString("KONG.ADMIN_URL") + path
	port, _ := strconv.ParseInt(viper.GetString("PORT"), 10, 0)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	payload := map[string]interface{}{
		"host":               string(helpers.LocalIP()),
		"connect_timeout":    60000,
		"protocol":           "http",
		"name":               serviceName,
		"read_timeout":       60000,
		"port":               port,
		"path":               "/",
		"retries":            0,
		"write_timeout":      60000,
		"tags":               nil,
		"client_certificate": nil,
	}
	jsonValue, _ := json.Marshal(payload)
	serviceResponse := CreateServiceResponse{}
	helpers.RestClient("POST", url, headers, bytes.NewBuffer(jsonValue), &serviceResponse)

	if serviceResponse.ErrorMessage == "" {
		return true
	} else {
		return false
	}
}
