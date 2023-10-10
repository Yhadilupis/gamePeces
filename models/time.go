package models

import (
	"strconv"
	"time"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type Timer struct {
	tiempo uint
}

func NuevoTemporizadorModelo() *Timer {
	return &Timer{
		tiempo: 60,
	}
}


func (t *Timer) IniciarTemporizador(pez *Pez, etiqueta *canvas.Text, botonFinal *widget.Button, etiquetaFinal *canvas.Text) {

	for t.tiempo > 0 && pez.vida {
		
		time.Sleep(time.Second * 1)
		
		t.tiempo--
		
		cadena := "Tiempo Restante:" + strconv.FormatUint(uint64(t.tiempo), 10) + " segundos"
		// Actualizamos la etiqueta
		etiqueta.Text = cadena
		etiqueta.Refresh()
	}

	pez.vida = false
	
	cadena := "Felicidades atrapazte uno :)"
	etiqueta.Text = cadena
	etiqueta.Refresh()
	etiquetaFinal.Text = "Tiempo restante: " + strconv.FormatUint(uint64(t.tiempo), 10) + " segundos"
	etiquetaFinal.Refresh()
	
	botonFinal.Show()
	etiquetaFinal.Show()
	go parpadear(etiquetaFinal, botonFinal)
}
