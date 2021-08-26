package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	ap := app.New()
	win := ap.NewWindow("Box Layout")
	win.Resize(fyne.Size{Width: 800, Height: 400})

	uri := binding.NewString()
	_ = uri.Set("please choose file ...")

	showFilePicker := func() {
		fmt.Println("button tapped")
		onChosen := func(f fyne.URIReadCloser, err error) {
			if err != nil {
				fmt.Println(err)
				return
			}
			if f == nil {
				return
			}
			fmt.Printf("chosen: %v", f.URI())
			_ = uri.Set(f.URI().String())
		}
		dialog.ShowFileOpen(onChosen, win)
	}

	entry := widget.NewEntryWithData(uri)
	entry.Disable()

	button := widget.NewButtonWithIcon("OPEN", theme.FileIcon(), showFilePicker)
	card := widget.NewCard("File Picker", "",
		container.New(layout.NewVBoxLayout(), entry, button))

	win.SetContent(container.NewMax(card))
	win.ShowAndRun()
}
