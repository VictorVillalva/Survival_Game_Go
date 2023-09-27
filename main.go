package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"Survive/scenes"
)

func main(){
	myApp := app.New()
	myWindow := myApp.NewWindow("Juego Pelota")
	myWindow.CenterOnScreen()
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(900, 600))

	//Cargar y mostrar la escena principal
	scenes.NewMenuEscena(myWindow)
	myWindow.ShowAndRun()
}