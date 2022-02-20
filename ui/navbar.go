package ui

import (
	"arvis/rtw/common/context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

const SWITCH_VIEWS = "SWITCH_VIEWS"

type Navbar struct {
	context     *context.Context
	navElements []*NavElement
}

func createNav(context *context.Context, navElements ...*NavElement) *Navbar {
	n := &Navbar{}
	n.context = context
	n.navElements = navElements

	return n
}

func (n *Navbar) View() fyne.CanvasObject {
	return container.New(layout.NewHBoxLayout(), n.getNavViews()...)
}

func (n *Navbar) OnDestroy() {

}

func (n *Navbar) getNavViews() []fyne.CanvasObject {
	navViews := []fyne.CanvasObject{}
	for _, navElement := range n.navElements {
		navViews = append(navViews, navElement.View())
	}

	return navViews
}
