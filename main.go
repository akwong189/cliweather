package main

import (
    "log"
    "github.com/akwong189/cliweather/pkg/utils"
)

func main() {
    log.Println("Program starting")
    api_keys := utils.GetApi()
    utils.GetWeather(api_keys.Weather, "current location")
    log.Println("Program ended")
}
