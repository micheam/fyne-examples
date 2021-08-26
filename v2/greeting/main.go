package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("greeting")
	w.SetContent(buildContent())
	w.ShowAndRun()
}

func buildContent() fyne.CanvasObject {
	return nil
}
