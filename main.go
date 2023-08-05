package main

import (
	"log"

	"github.com/akwong189/cliweather/pkg/app"
	"github.com/akwong189/cliweather/pkg/utils"
)

func main() {
	// defer closing log file until after main function is done
	defer utils.F.Close()

	log.Println("Program starting")
	app.StartApp()

	// api_keys, _ := utils.GetApi()
	// loc := utils.GetGeolocation(api_keys.Geolocation, "SAN JOSE")
	// loc := utils.GetCurrentIPLocation()
	// utils.GetWeather(api_keys.Weather, loc)
	log.Println("Program ended")
}
