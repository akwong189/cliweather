package utils

import (
    "io/ioutil"
    "log"
    "net/http"
    "github.com/buger/jsonparser"
)

type geolocation struct {
    long string
    lat string
}

// collect geolocation data from a particular location, using geocode api for faster reads (optional, leave black for none)
func GetGeolocation(geolocation_api, location string) *geolocation {
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

    log.Println("retrieved longitude, lattitude: ", long, lat)
    
    return &geolocation{long, lat}
}
