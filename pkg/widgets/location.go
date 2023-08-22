package widgets

import (
	"fmt"

	"github.com/akwong189/cliweather/pkg/utils"
	"github.com/jroimartin/gocui"
)

type LocationWidget struct {
	Loc *utils.Geolocation
}

// creates a new location widget (initalization uses current location)
func GetLocationWidget(loc *utils.Geolocation) *LocationWidget {
	return &LocationWidget{loc}
}

// Format layout of the location widget
func (l *LocationWidget) Layout(g *gocui.Gui) error {
	w, _ := g.Size()
	if v, err := g.SetView("location", 1, 0, w-2, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Location"
		v.Wrap = true

		if _, err := g.SetCurrentView("location"); err != nil {
			return nil
		}

		fmt.Fprint(v, l.Loc.GetLocationString())
	}
	return nil
}
