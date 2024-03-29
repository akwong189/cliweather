package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	// "os"
)

type ApiKeys struct {
	Weather        string
	IP_location    string
	Geolocation    string
	Location_query string
}

// Retrieves API key from config/mykeys.json
func GetApi() (*ApiKeys, error) {
	data, err := ioutil.ReadFile("./pkg/config/mykeys.json")
	log.Println("checking for mykeys.json")
	if err != nil {
		log.Println("file not found, asking user for info")
		return collectApiFromUser(), nil
	}
	log.Println("found mykeys.json")
	return getApiFromFile(data), nil
}

// retrieve API keys from a file
func getApiFromFile(key_data []byte) *ApiKeys {
	log.Println("data retrieved:\n" + string(key_data))
	var keys ApiKeys
	json.Unmarshal(key_data, &keys)
	log.Println("collected keys, key info are:\n\tweather api: " + keys.Weather + "\n\tIP location api: " + keys.IP_location + "\n\tgeolocation api: " + keys.Geolocation + "\n\tlocation query api: " + keys.Location_query)
	return &keys
}

// retrieve API keys from the user
func collectApiFromUser() *ApiKeys {
	var apis ApiKeys

	apis.Weather = askForKey("Dark Sky API")
	apis.IP_location = askForKey("IP location API")
	apis.Geolocation = askForKey("Geolocation API")
	apis.Location_query = askForKey("Location Query API")

	data, err := json.Marshal(apis)
	if err != nil {
		log.Fatalln("Failed to parse API struct to json")
	}

	if err := ioutil.WriteFile("pkg/config/mykeys.json", data, 0644); err != nil {
		log.Fatalln("Failed to create confif/mykeys.json")
	}

	return &apis
}

// Ask user for key via CLI (will be removed when GUI is implemented)
func askForKey(api_source string) string {
	var key string
	fmt.Println("Enter your key for " + api_source + ": ")
	fmt.Scan(&key)

	if key == "_" {
		return ""
	}

	return key
}
