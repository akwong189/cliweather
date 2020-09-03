package utils

import (
    "fmt"
	"io/ioutil"
	"log"
    "net/http"

	"github.com/buger/jsonparser"
)

// Gets the current locations location via IP address (may not be secure need to check)
func GetCurrentIPLocation() *geolocation {
    url := "https://freegeoip.app/json"
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

    long_float, err := jsonparser.GetFloat(body, "longitude")
    if err != nil {
        log.Fatalln("failed to retrieve longitude data")
    }

    lat_float, err := jsonparser.GetFloat(body, "latitude")
    if err != nil {
        log.Fatalln("failed to retieve latitude data")
    }

    log.Println("retrieved current location lattitude, longitude: ", lat_float, long_float)
    return &geolocation{long: fmt.Sprintf("%f", long_float), lat: fmt.Sprintf("%f", lat_float)}
    
}
