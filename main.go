package main

import (
	"gamePeces/scences"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow(">>ATRAPA UN PEZ EN EL MENOR TIEMPO POSIBLE<<")
	myWindow.CenterOnScreen()
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(800, 600))
	juego := scences.NuevaEscenaPrincipal(myWindow)
	juego.Cargar()
	myWindow.ShowAndRun()
}