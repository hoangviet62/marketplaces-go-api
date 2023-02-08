package internal

import (
	"bytes"
	"encoding/json"

	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	"github.com/k0kubun/pp"
	"github.com/spf13/viper"
)

type CreateRouteResponse struct {
	Id           string `json:"Id"`
	Name         string `json:"Name"`
	Enabled      bool   `json:"Enabled"`
	ErrorMessage string `json:"Message"`
}

func CreateRoute(serviceId string, payload map[string]interface{}) (status bool) {
	path := "/services/" + serviceId + "/routes"
	url := viper.GetString("KONG.ADMIN_URL") + path
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	jsonValue, _ := json.Marshal(payload)
	routeResponse := CreateRouteResponse{}
	helpers.RestClient("POST", url, headers, bytes.NewBuffer(jsonValue), &routeResponse)
	pp.Print(routeResponse)
	if routeResponse.Id != "" {
		return true
	} else {
		return false
	}
}
