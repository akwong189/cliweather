package app

import (
	"github.com/akwong189/cliweather/pkg/api/geojs"
	"github.com/akwong189/cliweather/pkg/utils"
)

func GrabCurrentLocation() *utils.Geolocation {
	return geojs.GetGeolocation()
}

func GetSavedLocations() []*utils.Geolocation {
	locs := make([]*utils.Geolocation, 0)
	locs = append(locs, GrabCurrentLocation())

	return locs
}
