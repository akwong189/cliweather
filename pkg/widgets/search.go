package widgets

import (
	"log"

	"github.com/akwong189/cliweather/pkg/model"
	"github.com/jroimartin/gocui"
)

type SearchWidget struct {
	AppUpdator *model.UpdateChannels
}

func (sw *SearchWidget) SearchEditor(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case ch != 0 && mod == 0:
		v.EditWrite(ch)
	case key == gocui.KeySpace:
		v.EditWrite(' ')
	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
		v.EditDelete(true)
	}
}

func (sw *SearchWidget) SearchBar(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("search", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Search"
		v.Editable = true
		// v.Editor = gocui.EditorFunc(sw.SearchEditor)

		if _, err := g.SetCurrentView("search"); err != nil {
			return err
		}
	}
	return nil
}

func (sw *SearchWidget) DestroySearchBar(g *gocui.Gui, v *gocui.View) error {
	if err := sw.AddNewLocation(g, v); err != nil {
		return err
	}

	if err := g.DeleteView("search"); err != nil {
		return err
	}
	if _, err := g.SetCurrentView("location"); err != nil {
		return err
	}
	return nil
}

func (sw *SearchWidget) AddNewLocation(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}
	log.Printf("Searching for \"%s\"\n", l)
	return nil
}
