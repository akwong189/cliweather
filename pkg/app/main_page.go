package app

import (
	"log"

	// geocode "github.com/akwong189/cliweather/pkg/api/geocode"
	// geojs "github.com/akwong189/cliweather/pkg/api/geojs"

	"github.com/akwong189/cliweather/pkg/data"
	"github.com/akwong189/cliweather/pkg/widgets"
	"github.com/jroimartin/gocui"
)

// Starts the app
func StartApp() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln("Failed to init GUI")
	}

	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorBlue

	// api_keys, err := utils.GetApi()
	// loc := utils.GetCurrentIPLocation()
	// forcast := utils.GetWeather(api_keys.Weather, loc)

	// loc := GrabCurrentLocation()
	locs := data.GenerateGeolocations(100)
	sel_widget := widgets.GetSelectorWidget(nil, locs)
	loc_widget := widgets.GetLocationWidget("loc", 1, 0, nil, locs[0])
	// curr_widget := NewCurrentWeatherWidget("curr", 1, 3, forcast.CurrentWeather)

	g.SetManager(loc_widget, sel_widget)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("selector", gocui.KeyArrowDown, gocui.ModNone, widgets.CursorDown); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("selector", gocui.KeyArrowUp, gocui.ModNone, widgets.CursorUp); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

// Quits the program
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
