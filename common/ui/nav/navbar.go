package nav

import (
	"arvis/rtw/common/observer"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type navbar struct {
	observable *observer.Observable
	view       fyne.CanvasObject
	buttons    []fyne.CanvasObject
}

func createNavbar(observable *observer.Observable, nav []*Nav) *navbar {
	n := &navbar{}
	n.observable = observable
	n.setUpNavbar(nav)

	return n
}

func (n *navbar) setUpNavbar(nav []*Nav) {
	for _, navItem := range nav {
		n.buttons = append(n.buttons, widget.NewButton(navItem.Name, n.createButtonOnClick(navItem)))
	}

	n.view = container.New(layout.NewHBoxLayout(), n.buttons...)
}

func (n *navbar) createButtonOnClick(navItem *Nav) func() {
	return func() {
		n.observable.Notify(navItem)
	}
}

func (n *navbar) View() fyne.CanvasObject {
	return n.view
}
