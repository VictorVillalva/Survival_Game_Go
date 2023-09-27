package models

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Pelon struct {
	posX, posY, direction float32
	running               bool
	pel                   *canvas.Image
}

//Construccion de personaje
func NewPelon(posx float32, posy float32, img *canvas.Image) *Pelon {
	return &Pelon{
		posX: posx,
		posY: posy,
		running: true,
		pel: img,
		direction: 0,
	}
}

//Direccionamiento
//* d - Direccion
func (d *Pelon) IrDerecha(){
	d.direction = 1
}

func (d *Pelon) IrIzquierda(){
	d.direction = -1
}

//Correr del personaje
func (d *Pelon) Correr(){
	for d.running {
		var incX float32 = 50
		incX *= d.direction

		if d.posX < 50 {
			d.posX = 50
		}else if d.posX > 650 {
			d.posX = 650
		}

		d.posX += incX
		d.pel.Move(fyne.NewPos(d.posX, d.posY))
		time.Sleep(100 * time.Millisecond)
	}
}

func (d *Pelon) SetRunning(pause bool){
	d.running = pause
}

func (d *Pelon) GetRunning() bool {
	return d.running
}
