package app

import (
    "github.com/akwong189/cliweather/pkg/utils"
    "github.com/jroimartin/gocui"
)


func StartApp() {
    g, err := gocui.NewGui(gocui.OutputNormal)
    if err != nil {
        utils.Log.Panicln("Failed to init GUI")
    }

    defer g.Close()

    loc_widget := NewLocationWidget("loc", 1, 1, nil)

    g.SetManager(loc_widget)

    if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
        utils.Log.Panicln(err)
    }

    if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
        utils.Log.Panicln(err)
    }
}

func quit(g *gocui.Gui, v *gocui.View) error {
    return gocui.ErrQuit
}
