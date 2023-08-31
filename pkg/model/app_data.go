package model

import (
	"errors"
	"log"
)

type AppData struct {
	Locations         *Geolocations
	CurrLocationIndex *int
	ApiKey            string
	Updators          *UpdateChannels
}

func InitAppData(locations *Geolocations, index int, api string, updators *UpdateChannels) (*AppData, error) {
	currIndex := index
	return &AppData{locations, &currIndex, api, updators}, nil
}

func (app AppData) UpdateLocations(new_location *Geolocation) error {
	*app.Locations = append(*app.Locations, new_location)
	return nil
}

func (app AppData) ChangeCurrSelectedLocation(index int) error {
	if index > len(*app.Locations) || index < 0 {
		return errors.New("location index out of bound")
	}

	if *app.CurrLocationIndex == index {
		return nil
	}

	log.Printf("Changed current location index to %d", index)
	*app.CurrLocationIndex = index
	app.Updators.UpdateLocation((*app.Locations)[index])
	return nil
}
