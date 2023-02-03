package helpers

import (
	"bytes"
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

func RestClient(method string, url string, headers map[string]string, payload *bytes.Buffer) (result string, err error) {
	client := &http.Client{}
	req, err := WraperRequest(method, url, payload)

	for key, val := range headers {
		req.Header.Add(key, val)
	}

	if err != nil {
		log.Error("[HTTP-CLIENT-REQUEST-ERROR] - METHOD %s - URL %s - Error %s", method, url, err)
		return
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Error("[HTTP-CLIENT-RESPONSE-ERROR] - METHOD %s - URL %s - Error %s", method, url, err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	log.Info("[HTTP-RESPONSE] %s", string(body))

	return string(body), err
}
