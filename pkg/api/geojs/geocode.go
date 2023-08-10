package geojs

import (
	"io"
	"log"
	"net/http"

	"github.com/akwong189/cliweather/pkg/api/openweathermap"
	"github.com/buger/jsonparser"
)

func GetGeolocation() *openweathermap.GeoLocation {
	resp := retrieveGeolocation()
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return parseGeolocation(body)
}

func retrieveGeolocation() *http.Response {
	url := "https://get.geojs.io/v1/ip/geo.json"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	return resp
}

func parseGeolocation(body []byte) *openweathermap.GeoLocation {
	city, err := jsonparser.GetString(body, "city")
	if err != nil {
		log.Fatalln(err)
	}
	long, err := jsonparser.GetString(body, "longitude")
	if err != nil {
		log.Fatalln(err)
	}
	lat, err := jsonparser.GetString(body, "latitude")
	if err != nil {
		log.Fatalln(err)
	}
	country, err := jsonparser.GetString(body, "country_code")
	if err != nil {
		log.Fatalln(err)
	}
	state, err := jsonparser.GetString(body, "region")
	if err != nil {
		log.Fatalln(err)
	}

	return &openweathermap.GeoLocation{
		Longitude: long,
		Latitude:  lat,
		Name:      city,
		Country:   country,
		State:     state,
	}
}
