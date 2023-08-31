package widgets

import (
	"fmt"
	"log"

	"github.com/akwong189/cliweather/pkg/model"
	"github.com/jroimartin/gocui"
)

type LocationWidget struct {
	Name string
	x, y int
	w    int
	Loc  chan *model.Geolocation
}

// creates a new location widget (initalization uses current location)
func GetLocationWidget(name string, x, y, w int, loc chan *model.Geolocation) *LocationWidget {
	return &LocationWidget{
		Name: name,
		x:    x,
		y:    y,
		w:    w,
		Loc:  loc}
}

// Format layout of the location widget
func (l *LocationWidget) Layout(g *gocui.Gui) error {
	w, _ := g.Size()
	if v, err := g.SetView(l.Name, l.w, l.y, w-l.w-1, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Location"
		v.Wrap = true

		if _, err := g.SetCurrentView(l.Name); err != nil {
			return nil
		}
	}
	return nil
}

func (l *LocationWidget) LocationUpdate(g *gocui.Gui) {
	for {
		curr_location := <-l.Loc
		log.Printf("Recieved new location: %s", curr_location.GetLocationString())
		g.Update(func(g *gocui.Gui) error {
			v, err := g.View(l.Name)
			if err != nil {
				return err
			}
			v.Clear()
			fmt.Fprint(v, curr_location.GetLocationString())
			return nil
		})
	}
}
