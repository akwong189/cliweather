package geojs

import (
	"io"
	"log"
	"net/http"

	"github.com/buger/jsonparser"
)

type IPAddress = string

func GetPublicIPAddress() IPAddress {
	return getPublicIPAddress()
}

func getPublicIPAddress() IPAddress {
	url := "https://get.geojs.io/v1/ip.json"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return retrieveIPAddress(body)
}

func retrieveIPAddress(body []byte) IPAddress {
	addr, err := jsonparser.GetString(body, "ip")

	if err != nil {
		log.Fatalln("No IP Address Found")
	}

	return addr
}
