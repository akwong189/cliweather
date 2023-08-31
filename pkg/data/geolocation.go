package data

import (
	"fmt"

	"github.com/akwong189/cliweather/pkg/model"
)

func GenerateGeolocation() *model.Geolocation {
	return &model.Geolocation{
		Longitude: "-121.8939",
		Latitude:  "37.29.60",
		Country:   "US",
		State:     "California",
		Name:      "San Jose",
		Zip:       "95123",
	}
}

func GenerateGeolocations(length int) []*model.Geolocation {
	locs := make([]*model.Geolocation, 0)

	for i := 0; i < length; i++ {
		loc := GenerateGeolocation()
		loc.Name = fmt.Sprintf("%s %d", loc.Name, i)
		locs = append(locs, loc)
	}
	return locs
}
