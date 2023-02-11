package internal

import (
	"bytes"
	"encoding/json"
	"errors"

	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	"github.com/spf13/viper"
)

type CreateConsumerResponse struct {
	Id           string `json:"Id"`
	Name         string `json:"Name"`
	ErrorMessage string `json:"Message"`
}

func CreateConsumer(username string) (bool, CreateConsumerResponse, error) {
	consumerPath := "/consumers"
	url := viper.GetString("KONG.ADMIN_URL") + consumerPath
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	payload := map[string]interface{}{
		"username":  username,
		"custom_id": username,
		"tags":      []string{"Customer"},
	}
	jsonValue, _ := json.Marshal(payload)
	consumerResponse := CreateConsumerResponse{}
	helpers.RestClient("POST", url, headers, bytes.NewBuffer(jsonValue), &consumerResponse)
	if consumerResponse.Id != "" {
		CreateConsumerJwt(consumerResponse.Id)
		return true, consumerResponse, errors.New("")
	} else {
		return false, consumerResponse, errors.New(consumerResponse.ErrorMessage)
	}
}
