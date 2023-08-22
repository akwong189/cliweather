package data

import "github.com/akwong189/cliweather/pkg/utils"

func GenerateWeatherData() *utils.Weather {
	return &utils.Weather{
		Summary:   "",
		Icon:      "",
		Temp:      1,
		AppTemp:   1,
		DewPoint:  1,
		Humidity:  1,
		Pressure:  1,
		WindSpeed: 1,
	}
}
