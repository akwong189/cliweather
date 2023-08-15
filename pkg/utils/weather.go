package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/buger/jsonparser"
)

type Weather struct {
	Summary   string  `json:"summary"`
	Icon      string  `json:"icon"`
	Temp      float64 `json:"temperature"`
	AppTemp   float64 `json:"apparentTemperature"`
	DewPoint  float64 `json:"dewPoint"`
	Humidity  float64 `json:"humidity"`
	Pressure  float64 `json:"pressure"`
	WindSpeed float64 `json:"windSpeed"`
}

type Forcast struct {
	CurrentWeather *Weather
	ForcastHourly  []Weather
	ForcastDaily   []Weather
}

// Weather using darksky api, may change it when dark sky stops providing api support
func GetWeather(weatherApiKey string, location *Geolocation) *Forcast {
	resp, err := callOpenWeatherMapAPI(location.Latitude, location.Longitude, weatherApiKey)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("collected data for current forcast")

	curr_data, _, _, err := jsonparser.Get(body, "currently")
	if err != nil {
		log.Fatalln("Input didn't provide any current forcast information!")
	}
	curr_weather := parseWeather(curr_data)

	hourly_data, _, _, err := jsonparser.Get(body, "hourly", "data")
	if err != nil {
		log.Fatalln("Input didn't provide any hourly forcast information!")
	}
	hourly_forcast := parseMultipleForcast(hourly_data)

	daily_data, _, _, err := jsonparser.Get(body, "hourly", "data")
	if err != nil {
		log.Fatalln("Input didn't provide any hourly forcast information!")
	}
	daily_forcast := parseMultipleForcast(daily_data)

	return &Forcast{curr_weather, hourly_forcast, daily_forcast}
}

// Parses an array of weather forcast provided from the Darksky API
func parseMultipleForcast(data []byte) []Weather {
	arr := make([]Weather, 0)
	if err := json.Unmarshal(data, &arr); err != nil {
		log.Fatalln("Failed to parse into individual bytes of data", err)
	}

	log.Println("Parsed array output", arr)

	return arr
}

func callOpenWeatherMapAPI(long, lat string, apiKey string) (*http.Response, error) {
	url := "https://api.openweathermap.org/data/3.0/onecall?lat=" + fmt.Sprintf("%v", lat) + "&lon=" + fmt.Sprintf("%v", long) + "&exclude=minutely" + "&appid=" + apiKey
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	return resp, err
}

// Parses the weather provided from the Darksky API
func parseWeather(data []byte) *Weather {
	log.Println("current data recieved: ", string(data))

	summary, err := jsonparser.GetString(data, "summary")
	if err != nil {
		log.Fatalln("Summary is missing!")
	}

	icon, err := jsonparser.GetString(data, "icon")
	if err != nil {
		log.Fatalln("Icon is missing!")
	}

	temp, err := jsonparser.GetFloat(data, "temperature")
	if err != nil {
		log.Fatalln("Temperature is missing!")
	}

	appTemp, err := jsonparser.GetFloat(data, "apparentTemperature")
	if err != nil {
		log.Fatalln("Apparent Temperature is missing!")
	}

	dewPoint, err := jsonparser.GetFloat(data, "dewPoint")
	if err != nil {
		log.Fatalln("Dew Point is missing!")
	}

	humidity, err := jsonparser.GetFloat(data, "humidity")
	if err != nil {
		log.Fatalln("Humidity is missing!")
	}

	pressure, err := jsonparser.GetFloat(data, "pressure")
	if err != nil {
		log.Fatalln("Pressure is missing!")
	}

	windSpeed, err := jsonparser.GetFloat(data, "windSpeed")
	if err != nil {
		log.Fatalln("Wind Speed is missing!")
	}

	log.Println("parsing completed", summary, icon, temp, appTemp, dewPoint, humidity, pressure, windSpeed)

	return &Weather{summary, icon, temp, appTemp, dewPoint, humidity, pressure, windSpeed}
}
