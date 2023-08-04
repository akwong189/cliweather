package utils

import (
	"io"
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
	Log.Println("contacting url: " + url)

	resp, err := http.Get(url)
	if err != nil {
		Log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		Log.Fatalln(err)
	}

	addressMatches, _, _, err := jsonparser.Get(body, "result", "addressMatches")
	if err != nil {
		Log.Fatalln(err)
	}

	Log.Println(string(addressMatches))
	geolocations, err := parseAddressMatches(addressMatches)
	if err != nil {
		Log.Fatalln(err)
	}

	Log.Println(geolocations[0])
	return &geolocations[0], nil
}

func parseAddressMatches(addressMatches []byte) ([]GeoLocation, error) {
	geolocations := make([]GeoLocation, 0)

	_, err := jsonparser.ArrayEach(addressMatches, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if err != nil {
			Log.Fatalln(err)
		}

		log, err := jsonparser.GetFloat(value, "coordinates", "x")
		if err != nil {
			Log.Fatalln(err)
		}
		lat, err := jsonparser.GetFloat(value, "coordinates", "y")
		if err != nil {
			Log.Fatalln(err)
		}
		geolocations = append(geolocations, GeoLocation{lat, log})
	})

	// TODO: Add code to handle empty geolocations
	// if len(geolocations) == 0 {
	// 	return geolocations, Error()
	// }
	return geolocations, err
}
