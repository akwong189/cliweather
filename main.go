package main

import (
    "log"
    "github.com/akwong189/cliweather/pkg/utils"
)

func main() {
    log.Println("Program starting")
    api_keys, _ := utils.GetApi()
    // loc := utils.GetGeolocation(api_keys.Geolocation, "SAN JOSE")
    loc := utils.GetCurrentIPLocation()
    utils.GetWeather(api_keys.Weather, loc)
    log.Println("Program ended")
}
