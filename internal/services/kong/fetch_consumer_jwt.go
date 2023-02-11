package internal

import (
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	"github.com/spf13/viper"
)

type Jwt struct {
	Id        string
	Algorithm string
	Secret    string
	Key       string
	CreatedAt uint64
}

type JwtResponse struct {
	Data []Jwt
}

func FetchConsumerJwt(consumerId string) JwtResponse {
	path := "/consumers/" + consumerId + "/jwt"
	url := viper.GetString("KONG.ADMIN_URL") + path
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	consumerJwtResponse := JwtResponse{}
	helpers.RestClient("GET", url, headers, nil, &consumerJwtResponse)
	return consumerJwtResponse
}
