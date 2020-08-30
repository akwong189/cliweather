package main

import (
    "log"
)

func main() {
    log.Println("Program starting")
    api_keys := getApi()
    getWeather(api_keys.Weather)
    log.Println("Program ended")
}
