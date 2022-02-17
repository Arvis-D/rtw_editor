package interfaces

import "fyne.io/fyne/v2"

type Fragment interface {
	OnDestroy()
	Display(canvas fyne.Canvas)
}
