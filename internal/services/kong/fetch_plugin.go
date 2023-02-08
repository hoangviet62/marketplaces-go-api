package internal

import (
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	"github.com/spf13/viper"
)

type Plugin struct {
	Name string `json:"name"`
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
	plugins = pluginResponse.Data

	return plugins
}
