package app

import (
	"log"

	"github.com/akwong189/cliweather/pkg/utils"
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

	// loc_widget := NewLocationWidget("loc", 1, 0, nil, loc)
	// curr_widget := NewCurrentWeatherWidget("curr", 1, 3, forcast.CurrentWeather)

	// g.SetManager(loc_widget, curr_widget)

	utils.RetrieveCoordinates("363 Alric Drive, San Jose, California 95123")

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
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
