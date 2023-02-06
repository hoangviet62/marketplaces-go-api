package helpers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func WraperRequest(method string, url string, payload *bytes.Buffer) (*http.Request, error) {
	if payload == nil {
		return http.NewRequest(strings.ToUpper(method), url, nil)
	} else {
		return http.NewRequest(strings.ToUpper(method), url, payload)
	}
}

func RestClient(method string, url string, headers map[string]string, payload *bytes.Buffer) (result map[string]interface{}) {
	client := &http.Client{}
	req, err := WraperRequest(method, url, payload)

	for key, val := range headers {
		req.Header.Add(key, val)
	}

	if err != nil {
		log.Error("[HTTP-CLIENT-REQUEST-ERROR] - METHOD %s - URL %s - Error %s", method, url, err)
		return
	}

	res, err := client.Do(req)

	if err != nil {
		log.Error("[HTTP-CLIENT-RESPONSE-ERROR] - METHOD %s - URL %s - Error %s", method, url, err)
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	var objs map[string]interface{}
	parseErr := json.Unmarshal([]byte(body), &objs)
	if parseErr != nil {
		panic(err)
	}

	return objs
}
