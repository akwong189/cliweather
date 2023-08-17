package widgets

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
func GetLocationWidget(name string, x, y int, handler func(g *gocui.Gui, v *gocui.View) error, loc *utils.Geolocation) *LocationWidget {
	return &LocationWidget{name: name, x: x, y: y, handler: handler, loc: loc}
}

// Format layout of the location widget
func (l *LocationWidget) Layout(g *gocui.Gui) error {
	w, _ := g.Size()
	v, err := g.SetView("location", l.x, l.y, w-l.x-1, l.y+2)

	if err != nil {
		v.Title = "Location"
		v.Wrap = true
		if err != gocui.ErrUnknownView {
			return err
		}

		if _, err := g.SetCurrentView("location"); err != nil {
			return nil
		}

		if err := g.SetKeybinding(l.name, gocui.KeyEnter, gocui.ModNone, l.handler); err != nil {
			return err
		}

		fmt.Fprint(v, l.loc.GetLocationString())
	}
	return nil
}
