package widgets

import "github.com/jroimartin/gocui"

func SearchBar(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("search", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Search"
		v.Editable = true

		if _, err := g.SetCurrentView("search"); err != nil {
			return err
		}
	}
	return nil
}

func DestroySearchBar(g *gocui.Gui, v *gocui.View) error {
	if err := g.DeleteView("search"); err != nil {
		return err
	}
	if _, err := g.SetCurrentView("location"); err != nil {
		return err
	}
	return nil
}
