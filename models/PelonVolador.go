package models

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type PelonVolador struct {
	posX, posY float32
	running               bool
	pel                   *canvas.Image
}

//Construccion de personaje
func NewPelonVolador(posx float32, posy float32, img *canvas.Image) *PelonVolador {
	return &PelonVolador{
		posX: posx,
		posY: posy,
		running: true,
		pel: img,
	}
}

//Correr del personaje
func (d *PelonVolador) Correr(){
	var incX float32 = 50
	for d.running {
		
		if d.posX < 00 || d.posX >840 {
			incX *= -1	
		}
		
		d.posX += incX
		d.pel.Move(fyne.NewPos(d.posX,d.posY))
		time.Sleep(80 * time.Millisecond)
	}
}

func (d *PelonVolador) SetRunning(pause bool){
	d.running = pause
}

func (d *PelonVolador) GetRunning() bool {
	return d.running
}
