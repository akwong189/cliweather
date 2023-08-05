package utils

import (
	"log"
	"net/http"
)

func HttpRequest(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}
