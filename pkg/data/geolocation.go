package data

import "github.com/akwong189/cliweather/pkg/utils"

func GenerateGeolocation() *utils.Geolocation {
	return &utils.Geolocation{
		Longitude: "-121.8939",
		Latitude:  "37.29.60",
		Country:   "US",
		State:     "California",
		Name:      "San Jose",
		Zip:       "95123",
	}
}
