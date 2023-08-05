package utils

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/buger/jsonparser"
)

type Geolocation struct {
	City    string
	Region  string
	Country string
	Long    string
	Lat     string
}

// collect geolocation data from a particular location, using geocode api for faster reads (optional, leave black for none)
//
// Deprecated: API no longer works
func GetGeolocation(geolocation_api, location string) *Geolocation {
	url := "https://geocode.xyz/" + location + "?json=1"
	if len(geolocation_api) != 0 {
		url += "auth=" + geolocation_api
	}
	log.Println("contacting url: " + url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

	long, err := jsonparser.GetString(body, "longt")
	if err != nil {
		log.Fatalln("failed to retrieve longitude data")
	}

	lat, err := jsonparser.GetString(body, "latt")
	if err != nil {
		log.Fatalln("failed to retrieve lattitude data")
	}

	city, err := jsonparser.GetString(body, "city")
	if err != nil {
		log.Fatalln("failed to retrieve city data")
	}

	state, err := jsonparser.GetString(body, "state")
	if err != nil {
		log.Fatalln("failed to retrieve state data")
	}

	country, err := jsonparser.GetString(body, "country")
	if err != nil {
		log.Fatalln("failed to retrieve country data")
	}

	log.Println("retrieved city, state, country: ", city, state, country)
	log.Println("retrieved lattitude, longitude: ", lat, long)

	return &Geolocation{city, state, country, long, lat}
}
