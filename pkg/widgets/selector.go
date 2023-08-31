package widgets

import (
	"fmt"

	"github.com/akwong189/cliweather/pkg/model"
	"github.com/jroimartin/gocui"
)

type SelectorWidget struct {
	AppData *model.AppData
}

func (s *SelectorWidget) OpenSelector(g *gocui.Gui, v *gocui.View) error {
	w, h := g.Size()
	if v, err := g.SetView("selector", w/2-60, h/2-20, w/2+60, h/2+20); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Select Location"
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack

		if _, err := g.SetCurrentView("selector"); err != nil {
			return nil
		}
		for _, loc := range *s.AppData.Locations {
			fmt.Fprint(v, loc.GetLocationString()+"\n")
		}
	}

	return nil
}

func (s *SelectorWidget) CloseSelector(g *gocui.Gui, v *gocui.View) error {
	if err := s.GetLine(g, v); err != nil {
		return err
	}

	if err := g.DeleteView("selector"); err != nil {
		return err
	}
	if _, err := g.SetCurrentView("location"); err != nil {
		return err
	}
	return nil
}

// TODO: handle how to determine line value based on either index or the object
func (sw *SelectorWidget) GetLine(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	return sw.AppData.ChangeCurrSelectedLocation(cy)
}

func CursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
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
