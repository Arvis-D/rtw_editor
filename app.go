package main

import (
	"arvis/rtw/common/context"
	"arvis/rtw/common/observer"
	"arvis/rtw/common/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

const width = 800
const height = 600
const name = "RTW editor"

type App struct {
	mainWindow fyne.Window
	context    *context.Context
}

func CreateApp() *App {
	a := &App{}
	a.context = context.CreateContext(app.New(), observer.CreateObservable())

	a.mainWindow = a.context.App().NewWindow(name)
	a.mainWindow.Resize(fyne.Size{Width: float32(width), Height: float32(height)})

	return a
}

func (a *App) Context() *context.Context {
	return a.context
}

func (a *App) Run(mainFragment ui.Fragment) {
	a.mainWindow.SetMaster()
	a.mainWindow.SetContent(mainFragment.View())
	a.mainWindow.Show()
	a.context.App().Run()
	mainFragment.OnDestroy()
}
