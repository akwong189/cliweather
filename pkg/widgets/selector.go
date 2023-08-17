package widgets

import (
	"fmt"

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

	return nil
}

func setKeybindings() error {
	return nil
}
