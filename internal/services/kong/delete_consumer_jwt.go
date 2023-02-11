package internal

import (
	"bytes"
	"encoding/json"

	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	"github.com/spf13/viper"
)

func DeleteConsumerJwt(consumerId string) ConsumerJwt {
	path := "/consumers/" + consumerId + "/jwt"
	url := viper.GetString("KONG.ADMIN_URL") + path
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	payload := map[string]interface{}{}
	jsonValue, _ := json.Marshal(payload)
	consumerJwtResponse := ConsumerJwt{}
	helpers.RestClient("DELETE", url, headers, bytes.NewBuffer(jsonValue), &consumerJwtResponse)
	return consumerJwtResponse
}
