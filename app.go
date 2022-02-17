package main

import (
	"arvis/rtw/common/interfaces"

	"fyne.io/fyne/v2"
)

type App struct {
	mainWindow fyne.Window
}

func (a *App) CreateApp() {

}

func (a *App) Run(mainFragment interfaces.Fragment) {
	a.mainWindow.SetMaster()
	mainFragment.Display(a.mainWindow.Canvas())
}
