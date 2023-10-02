package models

import (
	"math/rand"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Pez struct {
	boton   *widget.Button
	ventana fyne.Window
	vida    bool
}

func NuevoPezModelo(window fyne.Window, botonFinal *widget.Button, etiquetaFinal *canvas.Text, timer *Timer) *Pez {
	t := &Pez{
		boton:   widget.NewButton("", nil),
		ventana: window,
		vida:    true,
	}

	t.boton.OnTapped = func() {
		t.vida = false
		etiquetaFinal.Text = "Tiempo: " + strconv.FormatUint(uint64(timer.tiempo), 10) + " segundos"
		etiquetaFinal.Refresh()
		botonFinal.Show()
		etiquetaFinal.Show()
		go parpadear(etiquetaFinal, botonFinal) // Tercer hilo (decorador)
		t.boton.Enable()
	}

	t.boton.Resize(fyne.NewSize(50, 70))

	return t
}

func (t *Pez) CrearContenedor(image *canvas.Image) *fyne.Container {
	contenedorPrincipal := container.NewWithoutLayout(t.boton)
	contenedorPrincipal.Add(image)

	contenedorPrincipal.Resize(fyne.NewSize(70, 60))
	contenedorPrincipal.Move(fyne.NewPos(350, 260))

	return contenedorPrincipal
}

func (t *Pez) Mover(container *fyne.Container) {
	
	for t.vida {
		newX := float32(rand.Intn(750))
		newY := float32(rand.Intn(550))

		
		container.Move(fyne.NewPos(newX, newY))

		time.Sleep(500 * time.Millisecond)
	}

	t.boton.Disable()
}

func parpadear(etiquetaFinal *canvas.Text, botonFinal *widget.Button) {
	
	for {
		time.Sleep(1 * time.Second)
		etiquetaFinal.Show()
		botonFinal.Show()

		
		time.Sleep(2 * time.Second)
		etiquetaFinal.Hide()
		botonFinal.Hide()
	}
}
