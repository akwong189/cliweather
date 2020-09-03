package app

import (
    "fmt"

    "github.com/akwong189/cliweather/pkg/utils"
    "github.com/jroimartin/gocui"
)

type LocationWidget struct {
    name string
    x, y int
    handler func(g *gocui.Gui, v *gocui.View) error
    loc *utils.Geolocation
}

// helper function to print the location data
func printLocation(loc *utils.Geolocation) string {
    return fmt.Sprintf("%s, %s, %s", loc.City, loc.Region, loc.Country)
}

// creates a new location widget (initalization uses current location)
func NewLocationWidget(name string, x, y int, handler func(g *gocui.Gui, v *gocui.View) error) *LocationWidget {
    loc := utils.GetCurrentIPLocation() 
    return &LocationWidget{name: name, x: x, y: y, handler: handler, loc: loc}
}

// Format layout of the location widget
func (l *LocationWidget) Layout(g *gocui.Gui) error {
    w, _ := g.Size()
    v, err := g.SetView(l.name, l.x, l.y, w - l.x - 1, l.y + 3)
    if err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }

        if _, err := g.SetCurrentView(l.name); err != nil {
            return nil
        }

        if err := g.SetKeybinding(l.name, gocui.KeyEnter, gocui.ModNone, l.handler); err != nil {
            return err
        }
        fmt.Fprint(v, printLocation(l.loc))
    }
    return nil
}
