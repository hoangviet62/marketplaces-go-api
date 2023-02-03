package helpers

import (
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func PostRequest(url string, headers string, payload io.Reader) (result any, err error) {
	resp, err := http.Post(url, headers, payload)
	if err != nil {
		log.Error("[HTTP-CLIENT] - POST METHOD - %s", err)
		return
	}

	return resp, nil
}

func GetRequest(url string, headers string) (result any, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Error("[HTTP-CLIENT] - GET METHOD - %s", err)
		return
	}

	return resp, nil
}
