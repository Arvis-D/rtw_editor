package troppEditor

import (
	"arvis/rtw/common/context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TroopEditorFragment struct {
	context *context.Context
	view    fyne.CanvasObject
}

func CreateTroopEditorFragment(context *context.Context) *TroopEditorFragment {
	t := &TroopEditorFragment{}
	t.view = widget.NewLabel("Troop Editor")

	return t
}

func (t *TroopEditorFragment) View() fyne.CanvasObject {
	return t.view
}

func (t *TroopEditorFragment) OnDestroy() {

}
