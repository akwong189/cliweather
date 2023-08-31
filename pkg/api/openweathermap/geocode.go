package openweathermap

import (
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/akwong189/cliweather/pkg/model"
	"github.com/akwong189/cliweather/pkg/utils"
	"github.com/buger/jsonparser"
)

var locLimit = 5
var zipGeocodeUrl = "http://api.openweathermap.org/geo/1.0/zip?zip=%s,%s&appid=%s"
var addrGeocodeUrl = "http://api.openweathermap.org/geo/1.0/direct?q=%s,%s&limit=%d&appid=%s"

func GetGeolocation(location, countryCode, apiKey string) ([]*model.Geolocation, error) {
	if _, err := strconv.Atoi(location); err == nil {
		geolocations := make([]*model.Geolocation, 0)
		geolocation, err := zipGeoLocate(location, countryCode, apiKey)
		geolocations = append(geolocations, geolocation)
		return geolocations, err
	}
	return addressGeoLocate(location, countryCode, apiKey)
}

func zipGeoLocate(zip, countryCode, apiKey string) (*model.Geolocation, error) {
	url := fmt.Sprintf(zipGeocodeUrl, zip, countryCode, apiKey)
	log.Println(url)
	resp := utils.HttpRequest(url)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return decodeGeoLocationData(body)
}

func addressGeoLocate(location, countryCode, apiKey string) ([]*model.Geolocation, error) {
	url := fmt.Sprintf(addrGeocodeUrl, location, countryCode, locLimit, apiKey)
	log.Println(url)
	resp := utils.HttpRequest(url)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var locations []*model.Geolocation
	jsonparser.ArrayEach(body, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		log.Println(string(body))
		if err != nil {
			log.Fatalln(err)
		}
		loc, err := decodeGeoLocationData(value)
		if err != nil {
			log.Fatalln(err)
		}
		locations = append(locations, loc)
	})

	if len(location) == 0 {
		log.Fatalln("not enough positions")
	}

	return locations, nil
}

func decodeGeoLocationData(body []byte) (*model.Geolocation, error) {
	// TODO: issue with handling geolocation data as the data arrives as an array and not as a single object
	log.Println(string(body))
	name, err := jsonparser.GetString(body, "name")
	if err != nil {
		log.Fatalln(err)
	}
	lat, err := jsonparser.GetFloat(body, "lat")
	if err != nil {
		log.Fatalln(err)
	}
	long, err := jsonparser.GetFloat(body, "lon")
	if err != nil {
		log.Fatalln(err)
	}
	countryCode, err := jsonparser.GetString(body, "country")
	if err != nil {
		log.Fatalln(err)
	}

	// These may or may not exist in the json response based on OpenWeatherMap documentation
	state, err := jsonparser.GetString(body, "state")
	if err != nil {
		log.Println("State not found")
	}
	zip, err := jsonparser.GetString(body, "zip")
	if err != nil {
		log.Println("zip not found")
	}
	return &model.Geolocation{
		Longitude: strconv.FormatFloat(long, 'f', 4, 64),
		Latitude:  strconv.FormatFloat(lat, 'f', 4, 64),
		Name:      name,
		Country:   countryCode,
		Zip:       zip,
		State:     state}, nil
}
