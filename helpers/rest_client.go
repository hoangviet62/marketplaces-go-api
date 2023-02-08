package helpers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

func WraperRequest(method string, url string, payload *bytes.Buffer) (*http.Request, error) {
	if payload == nil {
		return http.NewRequest(strings.ToUpper(method), url, nil)
	} else {
		return http.NewRequest(strings.ToUpper(method), url, payload)
	}
}

func RestClient(method string, url string, headers map[string]string, payload *bytes.Buffer, target interface{}) error {
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := WraperRequest(method, url, payload)

	for key, val := range headers {
		req.Header.Add(key, val)
	}

	if err != nil {
		log.Error("[REST-CLIENT] REQUEST METHOD %s - URL %s - Error %s", method, url, err)
		return err
	}

	res, err := client.Do(req)

	if err != nil {
		log.Error("[REST-CLIENT] URL %s - Error %s", url, err)
		return err
	}

	defer res.Body.Close()
	json := json.NewDecoder(res.Body).Decode(&target)

	return json
}
