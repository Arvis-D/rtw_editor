package ui

import "fyne.io/fyne/v2"

type Fragment interface {
	OnDestroy()
	View() fyne.CanvasObject
}
