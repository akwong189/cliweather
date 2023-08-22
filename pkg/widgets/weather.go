package widgets

import (
	"fmt"

	"github.com/akwong189/cliweather/pkg/utils"
	"github.com/jroimartin/gocui"
)

type CurrentWeatherWidget struct {
	weather *utils.Weather
}

func GetWeatherWidget(curr_weather *utils.Weather) *CurrentWeatherWidget {
	return &CurrentWeatherWidget{weather: curr_weather}
}

func (c *CurrentWeatherWidget) Layout(g *gocui.Gui) error {
	w, h := g.Size()
	if v, err := g.SetView("weather", 1, 3, w-2, h-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Current Weather"

		fmt.Fprint(v, "Current Weather")
	}
	return nil
}
