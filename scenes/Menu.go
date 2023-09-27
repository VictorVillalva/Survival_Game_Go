package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type MenuEscena struct {
	window fyne.Window
}

func NewMenuEscena(window fyne.Window) *MenuEscena {
	escena := &MenuEscena{window: window}
	escena.Renderizar()
	return escena
}

//Renderizado
func (menuEscena *MenuEscena) Renderizar() {
	backgroundImagen := canvas.NewImageFromURI(storage.NewFileURI("./assets/bkgMenu.png"))
	backgroundImagen.Resize(fyne.NewSize(900, 600))
	backgroundImagen.Move(fyne.NewPos(0,0))

	//Boton iniciar
	botonIniciar := widget.NewButton("Comenzar", menuEscena.IniciarJuego)
	botonIniciar.Resize(fyne.NewSize(150,60))
	botonIniciar.Move(fyne.NewPos(370, 200))

	menuEscena.window.SetContent(container.NewWithoutLayout(backgroundImagen, botonIniciar)) 

}

func (menuEscena *MenuEscena) IniciarJuego(){
	NewEscenaJuego( menuEscena.window )
}
