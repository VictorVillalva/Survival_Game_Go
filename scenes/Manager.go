package scenes

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type Manager struct {
	window fyne.Window
}

func NewManager(window fyne.Window) *Manager {
	return &Manager{window: window}
}

func (manager *Manager) Show(bgUrl string) {
	bg := canvas.NewImageFromURI(storage.NewFileURI(bgUrl))
	bg.Resize(fyne.NewSize(900, 600))
	bg.Move(fyne.NewPos(0, 0))

	botonReiniciar := widget.NewButton("Reiniciar", func() {
		NewEscenaJuego(manager.window)
		contador = 0
		
	})
	botonReiniciar.Resize(fyne.NewSize(150, 50))
	botonReiniciar.Move(fyne.NewPos(600, 390))

	puntaje := contador
	contador=0
	fmt.Println("tu puntaje es: ", puntaje)

	labelPuntaje := widget.NewLabel("")
	labelPuntaje.SetText(fmt.Sprintf("Tu puntaje : %d", puntaje))
	
	botonPuntaje := widget.NewButton(labelPuntaje.Text,func() {})
	botonPuntaje.Resize(fyne.NewSize(150, 50))
	botonPuntaje.Move(fyne.NewPos(600, 520))

	manager.window.SetContent(container.NewWithoutLayout(bg, botonReiniciar,botonPuntaje))
}