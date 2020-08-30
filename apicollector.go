package main

import (
    "encoding/json"
    "io/ioutil"
    "log"
)

type api_key struct {
    Weather string
    Geolocation string
    Location_query string
}


func getApi() *api_key {
    data, err := ioutil.ReadFile("./mykeys.json")
    log.Println("checking for mykeys.json")
    if err != nil {
        log.Println("file not found, asking user for info")
        return collectApiFromUser()
    }
    log.Println("found mykeys.json")
    return getApiFromFile(data)
}

func getApiFromFile(key_data []byte) *api_key {
    log.Println("data retrieved:\n" + string(key_data))
    var keys api_key
    json.Unmarshal(key_data, &keys)
    log.Println("collected keys, key info are:\n\tweather api: " + keys.Weather + "\n\tgeolocation api: " + keys.Geolocation + "\n\tlocation query api: " + keys.Location_query)
    return &keys
}

func collectApiFromUser() *api_key {
    return nil 
}
