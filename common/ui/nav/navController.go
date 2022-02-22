package nav

import (
	"arvis/rtw/common/observer"
	"arvis/rtw/common/ui"

	"fyne.io/fyne/v2"
)

type NavController struct {
	observable    *observer.Observable
	nav           *navbar
	contentHolder *fyne.Container
	fragment      ui.Fragment
}

func CreateNavController(
	startingNav string,
	contentHolder *fyne.Container,
	navbarHolder *fyne.Container,
	nav ...*Nav,
) *NavController {
	n := &NavController{}
	n.contentHolder = contentHolder
	n.observable = observer.CreateObservable()
	n.observable.Subscribe(&observer.EventHandler{Action: nav_event, Callback: func(event observer.Event) {
		if nav, ok := event.(*Nav); ok {
			n.onNavChange(nav.Creator())
		}
	}})

	n.nav = createNavbar(n.observable, nav)
	navbarHolder.Objects = []fyne.CanvasObject{n.nav.view}
	navbarHolder.Refresh()

	n.Navigate(startingNav)

	return n
}

func (n *NavController) Navigate(action string) {
	n.observable.Notify(observer.Action(action))
}

func (n *NavController) onNavChange(fragment ui.Fragment) {
	if n.fragment != nil {
		n.fragment.OnDestroy()
	}

	n.fragment = fragment
	n.contentHolder.Objects = []fyne.CanvasObject{fragment.View()}
	n.contentHolder.Refresh()
}
