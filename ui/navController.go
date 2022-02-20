package ui

import (
	"arvis/rtw/common/context"
	"arvis/rtw/common/observer"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type NavController struct {
	context          *context.Context
	view             *fyne.Container
	nav              *Navbar
	navEventHandlers []*observer.EventHandler
	content          fyne.CanvasObject
}

func CreateNavController(context *context.Context, navEventHandlers []*observer.EventHandler) *NavController {
	m := &NavController{}
	m.context = context

	m.nav = createNav(context)
	m.view = container.New(layout.NewVBoxLayout(), m.nav.View())
	m.addNavEventHandlers(m.navEventHandlers)

	return m
}

func (m *NavController) View() fyne.CanvasObject {
	return m.view
}

func (m *NavController) OnDestroy() {
	for _, eventHandler := range m.navEventHandlers {
		m.context.Observable().Unsubscribe(eventHandler)
	}
}

func (m *NavController) Replace(from fyne.CanvasObject, to fyne.CanvasObject) {
	for i, v := range m.view.Objects {
		if v == from {
			m.view.Objects[i] = to
		}
	}

	m.view.Refresh()
}

func (m *NavController) setFragment(fragment fyne.CanvasObject) {
	m.Replace(m.content, fragment)
	m.content = fragment
}

func (m *NavController) addNavEventHandlers(eventHanlders []*observer.EventHandler) {
	m.navEventHandlers = eventHanlders

	for _, eventHandler := range eventHanlders {
		m.context.Observable().Subscribe(eventHandler)
	}
}
