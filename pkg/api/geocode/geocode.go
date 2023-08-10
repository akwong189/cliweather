package utils

import (
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/buger/jsonparser"
)

type GeoLocation struct {
	Longitude float64
	Latitude  float64
}

func RetrieveCoordinates(address string) (*GeoLocation, error) {
	url := "https://geocoding.geo.census.gov/geocoder/locations/onelineaddress?address=" + url.QueryEscape(address) + "&benchmark=2020&format=json"
	log.Println("contacting url: " + url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	addressMatches, _, _, err := jsonparser.Get(body, "result", "addressMatches")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(addressMatches))
	geolocations, err := parseAddressMatches(addressMatches)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(geolocations[0])
	return &geolocations[0], nil
}

func parseAddressMatches(addressMatches []byte) ([]GeoLocation, error) {
	geolocations := make([]GeoLocation, 0)

	_, err := jsonparser.ArrayEach(addressMatches, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if err != nil {
			log.Fatalln(err)
		}

		long, err := jsonparser.GetFloat(value, "coordinates", "x")
		if err != nil {
			log.Fatalln(err)
		}
		lat, err := jsonparser.GetFloat(value, "coordinates", "y")
		if err != nil {
			log.Fatalln(err)
		}
		geolocations = append(geolocations, GeoLocation{lat, long})
	})

	// TODO: Add code to handle empty geolocations
	// if len(geolocations) == 0 {
	// 	return geolocations, Error()
	// }
	return geolocations, err
}
