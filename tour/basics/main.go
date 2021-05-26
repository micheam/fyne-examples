package main

import (
	"fmt"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	go showAnother(myApp, "another 1")
	go showAnother(myApp, "another 2")
	go showAnother(myApp, "another 3")

	primaryWin := myApp.NewWindow("primary window")

	primaryWin.Resize(fyne.NewSize(400, 400))
	label := widget.NewLabel("Hello, World")
	primaryWin.SetContent(label)

	go func() {
		time.Sleep(time.Second * 10)
		label.SetText("Changed!")
		canvas.Refresh(label) // or `label.Refresh()`
	}()

	// ShowAndRun はウィンドウが閉じられるまで
	// カレントのgorutineをブロックします。
	primaryWin.ShowAndRun()
	fmt.Println("Bye ノシ")
}

func showAnother(a fyne.App, title string) {
	defer func() { log.Printf("%s has closed", title) }()
	time.Sleep(time.Second * 1)

	win := a.NewWindow("Shown later")
	win.SetContent(widget.NewLabel("will close 3 seconds later"))
	win.Resize(fyne.Size{Width: 200, Height: 100})
	win.Show()

	time.Sleep(time.Second * 5)
	win.Close()
}
