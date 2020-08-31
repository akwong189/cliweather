package utils

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "os"
)

type api_key struct {
    Weather string
    Geolocation string
    Location_query string
}

// Retrieves API key from config/mykeys.json
func GetApi() *api_key {
    data, err := ioutil.ReadFile("./pkg/config/mykeys.json")
    log.Println("checking for mykeys.json")
    if err != nil {
        log.Println("file not found, asking user for info")
        return collectApiFromUser()
    }
    log.Println("found mykeys.json")
    return getApiFromFile(data)
}

// retrieve API keys from a file
func getApiFromFile(key_data []byte) *api_key {
    log.Println("data retrieved:\n" + string(key_data))
    var keys api_key
    json.Unmarshal(key_data, &keys)
    log.Println("collected keys, key info are:\n\tweather api: " + keys.Weather + "\n\tgeolocation api: " + keys.Geolocation + "\n\tlocation query api: " + keys.Location_query)
    return &keys
}

// retrieve API keys from the user
func collectApiFromUser() *api_key {
    os.Exit(1)
    return nil 
}
