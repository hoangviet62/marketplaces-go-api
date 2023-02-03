package internal

import (
	"bytes"
	"encoding/json"

	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func CreateConsumer(username string) (res string) {
	consumerPath := "/consumers"
	url := viper.GetString("KONG.ADMIN_URL") + consumerPath
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	payload := map[string]interface{}{
		"username":  username,
		"custom_id": username,
		"tags":      []string{"User"},
	}
	jsonValue, _ := json.Marshal(payload)
	res, err := helpers.RestClient("POST", url, headers, bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Error("[KONG-CREATE-CONSUMER] FAILED TO CREATE %s", err)
	}
	return res
}
