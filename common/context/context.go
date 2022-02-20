package context

import (
	"arvis/rtw/common/observer"

	"fyne.io/fyne/v2"
)

type Context struct {
	app        fyne.App
	observable *observer.Observable
}

func CreateContext(app fyne.App, observable *observer.Observable) *Context {
	return &Context{app, observable}
}

func (c *Context) App() fyne.App {
	return c.app
}

func (c *Context) Observable() *observer.Observable {
	return c.observable
}
