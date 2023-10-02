package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func CreateButton(label string, clickFunc func(), size fyne.Size, pos fyne.Position, val bool) *widget.Button {
	btn := widget.NewButton(label, func ()  {
		clickFunc()
	})

	btn.Resize(size)
	btn.Move(pos)
	btn.Hidden = val
	return btn
}