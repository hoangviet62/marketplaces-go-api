package internal

import (
	"bytes"
	"encoding/json"
	"errors"
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

func CreateService(serviceName string) (bool, CreateServiceResponse, error) {
	var serviceResponse CreateServiceResponse
	// services := FetchServices()
	// isExisted := false

	// for _, service := range services {
	// 	if serviceName == service.Name {
	// 		isExisted = true
	// 		serviceResponse = CreateServiceResponse{Id: service.Id, Name: service.Name}
	// 		break
	// 	}
	// }

	// if isExisted {
	// 	errMsg := "SERVICE " + serviceName + " already existed"
	// 	return false, serviceResponse, errors.New(errMsg)
	// }

	path := "/services/" + serviceName
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
	helpers.RestClient("PUT", url, headers, bytes.NewBuffer(jsonValue), &serviceResponse)
	if serviceResponse.Id != "" {
		return true, serviceResponse, errors.New("")
	} else {
		return false, serviceResponse, errors.New(serviceResponse.ErrorMessage)
	}
}
