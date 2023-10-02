package scences

import (
	"gamePeces/models"
	"gamePeces/views"
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type EscenaPrincipal struct {
	window fyne.Window
}

func NuevaEscenaPrincipal(window fyne.Window) *EscenaPrincipal {
	return &EscenaPrincipal{window: window}
}

var startButton *widget.Button

func (e EscenaPrincipal) Cargar() {

	text := views.NewCanvaText("Rata En Accion!!!", color.Black, 35, true, fyne.NewSize(300, 200), fyne.NewPos(200, 50), false)
	title := views.AddContainerTitle(text, fyne.NewSize(300, 200), fyne.NewPos(200, 25))

	
	img1 := canvas.NewImageFromURI(storage.NewFileURI("./assets/pez1.png"))
	img2 := canvas.NewImageFromURI(storage.NewFileURI("./assets/pez2.png"))
	img3 := canvas.NewImageFromURI(storage.NewFileURI("./assets/pez3.png"))
	img4 := canvas.NewImageFromURI(storage.NewFileURI("./assets/pez4.png"))
	img5 := canvas.NewImageFromURI(storage.NewFileURI("./assets/pez5.png"))

	img1.Resize(fyne.NewSize(70, 70))
	img2.Resize(fyne.NewSize(70, 70))
	img3.Resize(fyne.NewSize(70, 70))
	img4.Resize(fyne.NewSize(70, 70))
	img5.Resize(fyne.NewSize(70, 70))

	
	etiquetaTiempo := canvas.NewText("Reglas: El juego termina cuando atrapaz un pez", color.Black)

	
	mensajeFinal := canvas.NewText("Tu tiempo se acabo", color.Black)

	
	mensajeFinal.Resize(fyne.NewSize(120, 130))
	mensajeFinal.Move(fyne.NewPos(330, 150))

	
	botonFinal := widget.NewButton("Finalizar", func() {
		os.Exit(0)
	})

	
	botonFinal.Resize(fyne.NewSize(150, 30))
	botonFinal.Move(fyne.NewPos(320, 250))

	
	botonFinal.Hide()
	mensajeFinal.Hide()

	
	temporizadorModelo := models.NuevoTemporizadorModelo()

	
	pezModelo := models.NuevoPezModelo(e.window, botonFinal, mensajeFinal, temporizadorModelo)

	
	pez1 := pezModelo.CrearContenedor(img1)
	pez2 := pezModelo.CrearContenedor(img2)
	pez3 := pezModelo.CrearContenedor(img3)
	pez4 := pezModelo.CrearContenedor(img4)
	pez5 := pezModelo.CrearContenedor(img5)

	
	fondo := canvas.NewImageFromURI(storage.NewFileURI("./assets/fondo.jpg"))
	fondo.Resize(fyne.NewSize(800, 600))

	startButton = views.CreateButton("Start Game", func() {
		title.Hide()
		startButton.Hide()
		//Hilos logicos para mover el pez
		go pezModelo.Mover(pez1)
		go pezModelo.Mover(pez2)
		go pezModelo.Mover(pez3)
		go pezModelo.Mover(pez4)
		go pezModelo.Mover(pez5)
		// Iniciamos el temporizador en un hilo lógico
		go temporizadorModelo.IniciarTemporizador(pezModelo, etiquetaTiempo, botonFinal, mensajeFinal)
	}, fyne.NewSize(150, 30), fyne.NewPos(300, 300), false)

	
	e.window.SetContent(container.NewWithoutLayout(title, fondo, pez1, pez2, pez3, pez4, pez5, etiquetaTiempo, mensajeFinal, startButton, botonFinal))
}
