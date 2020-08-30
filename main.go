package main

import (
    "log"
<<<<<<< HEAD
=======
    "github.com/akwong189/cliweather/pkg/utils"
>>>>>>> 45673f5... Created a more Go mod like package
)

func main() {
    log.Println("Program starting")
<<<<<<< HEAD
    api_keys := getApi()
    getWeather(api_keys.Weather)
=======
    api_keys := utils.GetApi()
    utils.GetWeather(api_keys.Weather, "current location")
>>>>>>> 45673f5... Created a more Go mod like package
    log.Println("Program ended")
}
