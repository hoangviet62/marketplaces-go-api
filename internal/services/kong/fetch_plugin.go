package internal

import (
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	"github.com/spf13/viper"
)

type Plugin struct {
	name string
}

func FetchPlugin(pluginName string) (status bool) {
	path := "/plugins"
	url := viper.GetString("KONG.ADMIN_URL") + path
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	result := helpers.RestClient("GET", url, headers, nil)
	plugins := []Plugin result["data"]
	// var data Plugin
	// json.Unmarshal(result["data"], &data)

	// var plugins []Plugin
	// json.Unmarshal(res, &plugins)
	// pp.Print(plugins)

	return true
}
