package main

import (
    "log"
    "github.com/akwong189/cliweather/pkg/utils"
)

func main() {
    log.Println("Program starting")
    api_keys := utils.GetApi()
    loc := utils.GetGeolocation(api_keys.Geolocation, "SAN JOSE")
    utils.GetWeather(api_keys.Weather, loc)
    log.Println("Program ended")
}
