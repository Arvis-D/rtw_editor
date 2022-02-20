package ui

import (
	"arvis/rtw/common/context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type NavElement struct {
	context *context.Context
	action  string
	view    *widget.Button
	name    string
}

func CreateNavElement(context *context.Context, action string, name string) *NavElement {
	n := &NavElement{context, action, nil, name}
	n.view = widget.NewButton(name, n.onClick)

	return n
}

func (n *NavElement) View() fyne.CanvasObject {
	return n.view
}

func (n *NavElement) OnDestroy() {

}

func (n *NavElement) onClick() {

}
