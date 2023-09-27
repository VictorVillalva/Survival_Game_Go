package models

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

var p *Pelon
var pv *PelonVolador

//Estructura de yunque
type Yunque struct {
	posX, posY, direction float32
	running               bool
	pel                   *canvas.Image
}

func NewYunque( posx float32, posy float32, img *canvas.Image, pelon *Pelon, pelonVolador *PelonVolador) *Yunque {
	p = pelon
	pv = pelonVolador
	return &Yunque{
		posX: posx,
		posY: posy,
		running: true,
		pel: img,
	}
}

func (y *Yunque) Caer() {
	for y.running {
		var inc float32 = 50

		if y.posY > 500 {
			y.posY = -50
			y.posX = float32((rand.Intn(12) + 1) * 50)
		}
		if y.posY >= 400 {
			if y.posX >= p.posX-50 && y.posX <= p.posX+50 {
				y.SetRunning(false)
				p.SetRunning(false)
				pv.SetRunning(false)
				inc = 0
			}
		} 
		
		y.posY += inc
		y.pel.Move(fyne.NewPos(y.posX,y.posY))
		time.Sleep(100 * time.Millisecond)
		}
}

func (y *Yunque) SetRunning(pause bool) {
	y.running = pause
}
func (y *Yunque) GetRunning() bool {
	return y.running
}
