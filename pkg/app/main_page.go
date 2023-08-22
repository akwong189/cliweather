package app

import (
	"log"

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

	// api_keys, err := utils.GetApi()
	// loc := utils.GetCurrentIPLocation()
	// forcast := utils.GetWeather(api_keys.Weather, loc)

	// loc := GrabCurrentLocation()
	locs := data.GenerateGeolocations(100)
	loc_widget := widgets.GetLocationWidget(locs[0])
	curr_widget := widgets.GetWeatherWidget(data.GenerateWeatherData())

	g.SetManager(loc_widget, curr_widget)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := KeyBindings(g, locs); err != nil {
		log.Panicln(err)
	}

	log.Printf("GUI closed\n")
}

// Quits the program
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
