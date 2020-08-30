package utils

import (
    "log"
    "io/ioutil"
    "net/http"
)

type weather struct {
    Summary string
    Icon string
    Temp int `json:"temperature"`
    AppTemp int `json:"apparentTemperature"`
    DewPoint int `json:"dewPoint"`
    Humidity int `json:"humidity"`
    Pressure int `json:"pressure"`
    WindSpeed int `json:"windSpeed"`
}

type forcast struct {
    location string
    current_weather weather
    forcast_hourly []weather
    forcast_daily []weather
}

// Weather using darksky api, may change it to allow other api's to work
func GetWeather(weather_api_key, location string) *forcast {
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

    log.Println("collected data for current forcast")
    log.Println(string(body))

    return nil
}
