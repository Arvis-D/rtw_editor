package main

import (
	"arvis/rtw/common/context"
	"arvis/rtw/common/ui"
	"arvis/rtw/common/ui/nav"
	"arvis/rtw/settlementEditor"
	troppEditor "arvis/rtw/troopEditor"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

const troop_editor = "Troop Editor"
const settlement_editor = "Settlement Editor"

type MainFragment struct {
	context       *context.Context
	content       *fyne.Container
	navbar        *fyne.Container
	view          fyne.CanvasObject
	navController nav.NavController
}

func CreateMainFragment(context *context.Context) *MainFragment {
	m := &MainFragment{}
	m.context = context
	m.content = container.New(layout.NewHBoxLayout())
	m.navbar = container.New(layout.NewHBoxLayout())
	m.view = container.New(layout.NewVBoxLayout(), m.navbar, m.content)

	m.navController = *nav.CreateNavController(troop_editor, m.content, m.navbar,
		nav.NewNav(troop_editor, func() ui.Fragment { return troppEditor.CreateTroopEditorFragment(context) }),
		nav.NewNav(settlement_editor, func() ui.Fragment { return settlementEditor.CreateSettlementEditorFragment(context) }),
	)

	return m
}

func (m *MainFragment) View() fyne.CanvasObject {
	return m.view
}

func (m *MainFragment) OnDestroy() {

}
