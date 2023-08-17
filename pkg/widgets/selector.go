package widgets

import (
	"fmt"
	"log"

	"github.com/akwong189/cliweather/pkg/utils"
	"github.com/jroimartin/gocui"
)

type SelectorWidget struct {
	handler   func(g *gocui.Gui, v *gocui.View) error
	locations []*utils.Geolocation
}

func GetSelectorWidget(handler func(g *gocui.Gui, v *gocui.View) error, locations []*utils.Geolocation) *SelectorWidget {
	return &SelectorWidget{handler, locations}
}

func (s *SelectorWidget) Layout(g *gocui.Gui) error {
	w, h := g.Size()
	v, err := g.SetView("selector", w/2-60, h/2-20, w/2+60, h/2+20)

	if err != nil {
		v.Title = "Select Location"
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		if err != gocui.ErrUnknownView {
			return err
		}

		if _, err := g.SetCurrentView("selector"); err != nil {
			return nil
		}
	}

	for _, loc := range s.locations {
		fmt.Fprint(v, loc.GetLocationString()+"\n")
	}

	// setKeybindings(g)

	return nil
}

func CursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		log.Printf("%d %d", cx, cy)
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func CursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}
