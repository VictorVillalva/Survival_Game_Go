package scenes

import (
	"Survive/models"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

// Estructura de la escena
type EscenaJuego struct {
	window fyne.Window
}

var p *models.Pelon
var y *models.Yunque
var pV *models.PelonVolador
var contador int

func NewEscenaJuego(window fyne.Window) *EscenaJuego {
	escena := &EscenaJuego{window: window}
	escena.Renderizar()
	escena.IniciarJuego()
	return escena
}

func (s *EscenaJuego) Renderizar() {
	backgroundImagen := canvas.NewImageFromURI(storage.NewFileURI("./assets/bkgJuego.png"))
	backgroundImagen.Resize(fyne.NewSize(900, 600))
	backgroundImagen.Move(fyne.NewPos(0,0))

	pelonPeel := crearPeel("./assets/PelonPlayer.png", 100, 100, 100, 450)
	p = models.NewPelon(0, 450, pelonPeel)
	pelonVoladorPeel := crearPeel("./assets/pelonVolador.png", 100, 100, 100, 50)
	pV = models.NewPelonVolador(0, 50, pelonVoladorPeel)
	yunquePeel := crearPeel("./assets/yunque.png", 100, 100, 100, 50)
	y = models.NewYunque(350, -50, yunquePeel, p, pV)
	
	//Botton reiniciar
	buttonStop := widget.NewButton("Salir", s.window.Close)
	buttonStop.Resize(fyne.NewSize(50, 50))
	buttonStop.Move(fyne.NewPos(850, 550))

	//Boton para ir a la izquierda
	buttonLeft := widget.NewButton("<", p.IrIzquierda)
	buttonLeft.Resize(fyne.NewSize(50, 50))
	buttonLeft.Move(fyne.NewPos(0, 550))

	//Boton para ir a la derecha
	buttonRigth := widget.NewButton(">", p.IrDerecha)
	buttonRigth.Resize(fyne.NewSize(50, 50))
	buttonRigth.Move(fyne.NewPos(50, 550))

	//Colocacion de botones
	s.window.SetContent(container.NewWithoutLayout(backgroundImagen, pelonPeel, yunquePeel,pelonVoladorPeel, buttonLeft, buttonRigth, buttonStop))
}

// Inicializacion del juego
// Implementacion de hilos pelon, yunque, pelonVolador verificador si eta jugando
func (s *EscenaJuego) IniciarJuego() {
	go p.Correr()
	go y.Caer()
	go pV.Correr()
	go s.estaJugando()
	go getTime()
}

// Creacion de skin (pelaje)
func crearPeel(fileUri string, sizeX float32, sizeY float32, posX float32, posY float32) *canvas.Image {
	imagen := canvas.NewImageFromURI(storage.NewFileURI(fileUri))
	imagen.Resize(fyne.NewSize(sizeX, sizeY))
	imagen.Move(fyne.NewPos(posX, posY))
	return imagen
}

func (s *EscenaJuego) estaJugando () {
	for {
		if !p.GetRunning() && !y.GetRunning() {
			time.Sleep(1000 * time.Millisecond)
			manager := NewManager( s.window )

			manager.Show("./assets/bkgPerdiste.png")
			
			break;
		}
	}
} 

func getTime(){
	contador =0
	for {
		contador++
		time.Sleep(5500* time.Millisecond)
	}
}