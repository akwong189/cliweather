package app

import (
	"log"

	"github.com/akwong189/cliweather/pkg/data"
	"github.com/akwong189/cliweather/pkg/utils"
	"github.com/akwong189/cliweather/pkg/widgets"
	"github.com/jroimartin/gocui"
)

type App struct {
	locations []*utils.Geolocation
	updators  *utils.UpdateChannels
}

// Starts the app
func StartApp() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln("Failed to init GUI")
	}

	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorCyan

	// api_keys, err := utils.GetApi()
	// loc := utils.GetCurrentIPLocation()
	// forcast := utils.GetWeather(api_keys.Weather, loc)

	updators := utils.InitalizeUpdators()
	// loc := GrabCurrentLocation()
	locs := data.GenerateGeolocations(100)

	loc_widget := widgets.GetLocationWidget("location", 0, 0, 1, updators.Location)
	curr_widget := widgets.GetWeatherWidget(data.GenerateWeatherData())

	g.SetManager(loc_widget, curr_widget)

	log.Printf("Layout built")

	go loc_widget.LocationUpdate(g)
	go initalizeDefault(updators, locs[0])

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := keyBindings(g, locs, updators); err != nil {
		log.Panicln(err)
	}

	log.Printf("GUI closed\n")
}

func initalizeDefault(updator *utils.UpdateChannels, location *utils.Geolocation) {
	updator.UpdateLocation(location)
}

func keyBindings(g *gocui.Gui, locations []*utils.Geolocation, updator *utils.UpdateChannels) error {
	sel := &widgets.SelectorWidget{Locations: locations, Updator: updator}

	if err := g.SetKeybinding("", gocui.KeyCtrlS, gocui.ModNone, widgets.SearchBar); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("location", gocui.KeyEnter, gocui.ModNone, sel.OpenSelector); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("selector", gocui.KeyEnter, gocui.ModNone, sel.CloseSelector); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("selector", gocui.KeyArrowDown, gocui.ModNone, widgets.CursorDown); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("selector", gocui.KeyArrowUp, gocui.ModNone, widgets.CursorUp); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("search", gocui.KeyEnter, gocui.ModNone, widgets.DestroySearchBar); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
	return nil
}

// Quits the program
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
