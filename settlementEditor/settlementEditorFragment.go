package settlementEditor

import (
	"arvis/rtw/common/context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type SettlementEditorFragment struct {
	context *context.Context
	view    fyne.CanvasObject
}

func CreateSettlementEditorFragment(context *context.Context) *SettlementEditorFragment {
	s := &SettlementEditorFragment{}
	s.view = widget.NewLabel("Settlement Editor")

	return s
}

func (s *SettlementEditorFragment) View() fyne.CanvasObject {
	return s.view
}

func (s *SettlementEditorFragment) OnDestroy() {

}
