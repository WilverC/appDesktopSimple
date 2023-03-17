package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hola Fyne")

	input := widget.NewEntry()
	input.SetPlaceHolder("Ingrese su nombre... ")

	hello := widget.NewLabel("")

	send := widget.NewButton("Enviar", func() {
		hello.SetText("Hola " + input.Text)
	})

	content := container.NewVBox(input, send, hello)

	//myWindow.SetContent(widget.NewLabel("Hola Mundo"))
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
