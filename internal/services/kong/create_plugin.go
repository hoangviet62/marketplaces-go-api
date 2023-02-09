package internal

import (
	"bytes"
	"encoding/json"

	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	"github.com/spf13/viper"
)

type CreatePluginResponse struct {
	Id      string `json:"Id"`
	Name    string `json:"Name"`
	Enabled bool   `json:"Enabled"`
}

func CreatePlugin(resourceName string, resourceId string, pluginName string) (status bool) {
	// plugins := FetchPlugins()
	// isExisted := false

	// for _, plugin := range plugins {
	// 	if pluginName == plugin.Name {
	// 		isExisted = true
	// 		break
	// 	}
	// }

	// if isExisted {
	// 	return false
	// }
	path := resourceName + "/" + resourceId + "/plugins"
	url := viper.GetString("KONG.ADMIN_URL") + path
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	payload := map[string]interface{}{
		"config": map[string]interface{}{
			"key_claim_name":     "iss",
			"maximum_expiration": 86400,
			"run_on_preflight":   true,
		},
		"name": pluginName,
	}
	jsonValue, _ := json.Marshal(payload)
	pluginResponse := CreatePluginResponse{}
	helpers.RestClient("POST", url, headers, bytes.NewBuffer(jsonValue), &pluginResponse)

	if pluginResponse.Id != "" {
		return true
	} else {
		return false
	}
}
