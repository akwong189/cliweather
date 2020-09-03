package utils

import (
    "io/ioutil"
    "net/http"
    "encoding/json"

    "github.com/buger/jsonparser"
)

type Weather struct {
    Summary string `json:"summary"`
    Icon string `json:"icon"`
    Temp float64 `json:"temperature"`
    AppTemp float64 `json:"apparentTemperature"`
    DewPoint float64 `json:"dewPoint"`
    Humidity float64 `json:"humidity"`
    Pressure float64 `json:"pressure"`
    WindSpeed float64 `json:"windSpeed"`
}

type Forcast struct {
    current_weather Weather
    forcast_hourly []Weather
    forcast_daily []Weather
}

// Weather using darksky api, may change it when dark sky stops providing api support
func GetWeather(weather_api_key string, location *Geolocation) *Forcast {
    url := "https://api.darksky.net/forecast/" + weather_api_key + "/" + location.Lat + "," + location.Long
    resp, err := http.Get(url)
    if err != nil {
        Log.Fatalln(err)
    }
    Log.Println(url)

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        Log.Fatalln(err)
    }

    Log.Println("collected data for current forcast")

    curr_data, _, _, err := jsonparser.Get(body, "currently")
    if err != nil {
        Log.Fatalln("Input didn't provide any current forcast information!")
    }
    curr_weather := parseWeather(curr_data)

    hourly_data, _, _, err := jsonparser.Get(body, "hourly", "data")
    if err != nil {
        Log.Fatalln("Input didn't provide any hourly forcast information!")
    }
    hourly_forcast := parseMultipleForcast(hourly_data)

    daily_data, _, _, err := jsonparser.Get(body, "hourly", "data")
    if err != nil {
        Log.Fatalln("Input didn't provide any hourly forcast information!")
    }
    daily_forcast := parseMultipleForcast(daily_data)

    return &Forcast{curr_weather, hourly_forcast, daily_forcast}
}

// Parses an array of weather forcast provided from the Darksky API
func parseMultipleForcast(data []byte) []Weather {
    arr := make([]Weather, 0)
    if err := json.Unmarshal(data, &arr); err != nil {
        Log.Fatalln("Failed to parse into individual bytes of data", err)
    }

    Log.Println("Parsed array output", arr)

    return arr
}

// Parses the weather provided from the Darksky API
func parseWeather(data []byte) Weather {
    Log.Println("current data recieved: ", string(data))

    summary, err := jsonparser.GetString(data, "summary")
    if err != nil {
        Log.Fatalln("Summary is missing!")
    }

    icon, err := jsonparser.GetString(data, "icon")
    if err != nil {
        Log.Fatalln("Icon is missing!")
    }

    temp, err := jsonparser.GetFloat(data, "temperature")
    if err != nil {
        Log.Fatalln("Temperature is missing!")
    }

    appTemp, err := jsonparser.GetFloat(data, "apparentTemperature")
    if err != nil {
        Log.Fatalln("Apparent Temperature is missing!")
    }

    dewPoint, err := jsonparser.GetFloat(data, "dewPoint")
    if err != nil {
        Log.Fatalln("Dew Point is missing!")
    }

    humidity, err := jsonparser.GetFloat(data, "humidity")
    if err != nil {
        Log.Fatalln("Humidity is missing!")
    }

    pressure, err := jsonparser.GetFloat(data, "pressure")
    if err != nil {
        Log.Fatalln("Pressure is missing!")
    }

    windSpeed, err := jsonparser.GetFloat(data, "windSpeed")
    if err != nil {
        Log.Fatalln("Wind Speed is missing!")
    }

    Log.Println("parsing completed", summary, icon, temp, appTemp, dewPoint, humidity, pressure, windSpeed)

    return Weather{summary, icon, temp, appTemp, dewPoint, humidity, pressure, windSpeed}
}
