package main

import (
    "log"
    "io/ioutil"
    "net/http"
)

type weather struct {
    location string
    current_weather int
    forcast_hourly []int
    forcast_daily []int
}

// Weather using darksky api, may change it to allow other api's to work
func getWeather(weather_api_key string) *weather {
    url := "https://api.darksky.net/forecast/" + weather_api_key + "/37.8267,-122.4233"
    log.Println(url)
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
    return nil
}
