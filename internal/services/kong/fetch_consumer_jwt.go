package internal

import (
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	"github.com/spf13/viper"
)

func GetConsumerJwt(consumerId string) {
	path := "/consumers/" + consumerId + "/jwt"
	url := viper.GetString("KONG.ADMIN_URL") + path
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	pluginResponse := FetchPluginResponse{}
	helpers.RestClient("GET", url, headers, nil, &pluginResponse)
	// return pluginResponse.Data
}
