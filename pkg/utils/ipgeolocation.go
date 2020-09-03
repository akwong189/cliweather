package utils

import (
    "fmt"
	"io/ioutil"
    "net/http"

	"github.com/buger/jsonparser"
)

// Gets the current locations location via IP address (may not be secure need to check)
func GetCurrentIPLocation() *Geolocation {
    url := "https://freegeoip.app/json"
    Log.Println("contacting url: " + url)

    resp, err := http.Get(url)
    if err != nil {
        Log.Fatalln(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        Log.Fatalln(err)
    }

    long_float, err := jsonparser.GetFloat(body, "longitude")
    if err != nil {
        Log.Fatalln("failed to retrieve longitude data")
    }

    lat_float, err := jsonparser.GetFloat(body, "latitude")
    if err != nil {
        Log.Fatalln("failed to retieve latitude data")
    }
    
    city, err := jsonparser.GetString(body, "city")
    if err != nil {
        Log.Fatalln("failed to retrieve city data")
    }
    
    state, err := jsonparser.GetString(body, "region_name")
    if err != nil {
        Log.Fatalln("failed to retrieve state data")
    }

    country, err := jsonparser.GetString(body, "country_name")
    if err != nil {
        Log.Fatalln("failed to retrieve country data")
    }

    Log.Println("retrieved current city, state, country: ", city, state, country)
    Log.Println("retrieved current location lattitude, longitude: ", lat_float, long_float)
    return &Geolocation{City: city, Region: state, Country: country, Long: fmt.Sprintf("%f", long_float), Lat: fmt.Sprintf("%f", lat_float)}
    
}
