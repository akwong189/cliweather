package app

import (
	"github.com/akwong189/cliweather/pkg/api/geojs"
	"github.com/akwong189/cliweather/pkg/model"
)

func GrabCurrentLocation() *model.Geolocation {
	return geojs.GetGeolocation()
}

func GetSavedLocations() []*model.Geolocation {
	locs := make([]*model.Geolocation, 0)
	locs = append(locs, GrabCurrentLocation())

	return locs
}
