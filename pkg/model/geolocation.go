package model

import (
	"fmt"
)

type Geolocation struct {
	Longitude string
	Latitude  string
	Name      string
	Country   string
	Zip       string
	State     string
}

type Geolocations = []*Geolocation

func (g *Geolocation) GetLocationString() string {
	str := g.Name

	if len(g.State) != 0 {
		str += fmt.Sprintf(", %s", g.State)
	}
	if len(g.Country) != 0 {
		str += fmt.Sprintf(", %s", g.Country)
	}
	if len(g.Zip) != 0 {
		str += fmt.Sprintf(" %s", g.Zip)
	}

	return str
}
