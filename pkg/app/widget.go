package app

import (
	"fmt"

	"github.com/akwong189/cliweather/pkg/utils"
	"github.com/jroimartin/gocui"
)

type LocationWidget struct {
	name    string
	x, y    int
	handler func(g *gocui.Gui, v *gocui.View) error
	loc     *utils.Geolocation
}

// creates a new location widget (initalization uses current location)
func NewLocationWidget(name string, x, y int, handler func(g *gocui.Gui, v *gocui.View) error, loc *utils.Geolocation) *LocationWidget {
	return &LocationWidget{name: name, x: x, y: y, handler: handler, loc: loc}
}

// Format layout of the location widget
func (l *LocationWidget) Layout(g *gocui.Gui) error {
	w, _ := g.Size()
	v, err := g.SetView(l.name, l.x, l.y, w-l.x-1, l.y+2)

	if err != nil {
		v.Title = "Location"
		v.Wrap = true
		if err != gocui.ErrUnknownView {
			return err
		}

		if _, err := g.SetCurrentView(l.name); err != nil {
			return nil
		}

		if err := g.SetKeybinding(l.name, gocui.KeyEnter, gocui.ModNone, l.handler); err != nil {
			return err
		}

		fmt.Fprint(v, l.loc.GetLocationString())
	}
	return nil
}

type CurrentWeatherWidget struct {
	name    string
	x, y    int
	weather *utils.Weather
}

func NewCurrentWeatherWidget(name string, x, y int, curr_weather *utils.Weather) *CurrentWeatherWidget {
	return &CurrentWeatherWidget{name: name, x: x, y: y, weather: curr_weather}
}

func (c *CurrentWeatherWidget) Layout(g *gocui.Gui) error {
	w, _ := g.Size()
	v, err := g.SetView(c.name, c.x, c.y, w-c.x-1, c.y+10)

	if err != nil {
		v.Title = "Current Weather"

		if err != gocui.ErrUnknownView {
			return err
		}

		// if _, err := g.SetCurrentView(c.name); err != nil {
		//     return nil
		// }

		fmt.Fprint(v, "Current Weather")
	}
	return nil
}
