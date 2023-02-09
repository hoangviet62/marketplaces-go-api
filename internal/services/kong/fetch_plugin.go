package internal

import (
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	"github.com/spf13/viper"
)

type Plugin struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
}

type FetchPluginResponse struct {
	Data []Plugin
}

func FetchPlugins() (plugins []Plugin) {
	path := "/plugins"
	url := viper.GetString("KONG.ADMIN_URL") + path
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	pluginResponse := FetchPluginResponse{}
	helpers.RestClient("GET", url, headers, nil, &pluginResponse)
	return pluginResponse.Data
}
