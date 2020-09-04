package app

import (
    "github.com/akwong189/cliweather/pkg/utils"
    "github.com/jroimartin/gocui"
)

// Starts the app
func StartApp() {
    g, err := gocui.NewGui(gocui.OutputNormal)
    if err != nil {
        utils.Log.Panicln("Failed to init GUI")
    }

    defer g.Close()

    g.Highlight = true
    g.SelFgColor = gocui.ColorBlue

    api_keys, err := utils.GetApi()
    loc := utils.GetCurrentIPLocation()
    forcast := utils.GetWeather(api_keys.Weather, loc)

    loc_widget := NewLocationWidget("loc", 1, 0, nil, loc)
    curr_widget := NewCurrentWeatherWidget("curr", 1, 3, forcast.CurrentWeather)

    g.SetManager(loc_widget, curr_widget)

    if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
        utils.Log.Panicln(err)
    }

    if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
        utils.Log.Panicln(err)
    }
}

// Quits the program
func quit(g *gocui.Gui, v *gocui.View) error {
    return gocui.ErrQuit
}
