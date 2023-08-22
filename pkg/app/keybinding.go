package app

import (
	"log"

	"github.com/akwong189/cliweather/pkg/utils"
	"github.com/akwong189/cliweather/pkg/widgets"
	"github.com/jroimartin/gocui"
)

func KeyBindings(g *gocui.Gui, locs []*utils.Geolocation) error {
	sel := &widgets.SelectorWidget{Locations: locs}

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
