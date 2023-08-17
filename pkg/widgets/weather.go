package widgets

import (
	"fmt"

	"github.com/akwong189/cliweather/pkg/utils"
	"github.com/jroimartin/gocui"
)

type CurrentWeatherWidget struct {
	name    string
	x, y    int
	weather *utils.Weather
}

func NewCurrentWeatherWidget(name string, x, y int, curr_weather *utils.Weather) *CurrentWeatherWidget {
	return &CurrentWeatherWidget{name: name, x: x, y: y, weather: curr_weather}
}

func (c *CurrentWeatherWidget) Layout(g *gocui.Gui) error {
	w, _ := g.Size()
	v, err := g.SetView(c.name, c.x, c.y, w-c.x-1, c.y+10)

	if err != nil {
		v.Title = "Current Weather"

		if err != gocui.ErrUnknownView {
			return err
		}

		// if _, err := g.SetCurrentView(c.name); err != nil {
		//     return nil
		// }

		fmt.Fprint(v, "Current Weather")
	}
	return nil
}
