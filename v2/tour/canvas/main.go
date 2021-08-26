package main

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.LightTheme())
	win := myApp.NewWindow("canvas examples")

	go showRectangle(myApp)
	go showText(myApp)
	go showLine(myApp)
	go showLineResize(myApp)
	go showCircle(myApp)
	go showImageWithFillStretch(myApp, fyne.NewSize(400, 200))
	go showImageWithFillStretch(myApp, fyne.NewSize(200, 400))

	size := fyne.NewSize(1000, 600)
	go showImage(myApp, size, canvas.ImageFillStretch)
	go showImage(myApp, size, canvas.ImageFillOriginal)
	go showImage(myApp, size, canvas.ImageFillContain)

	win.Resize(fyne.Size{Width: 400, Height: 200})
	win.ShowAndRun()
}

func showRectangle(parent fyne.App) {
	child := parent.NewWindow("Rectangle")
	rect := canvas.NewRectangle(color.Gray16{0x9999})
	child.SetContent(rect)
	child.Resize(fyne.Size{Width: 400, Height: 200})
	child.Show()
	child.Close()
}

func showText(parent fyne.App) {
	child := parent.NewWindow("Text")
	text := canvas.NewText("Hello, World", color.Black)
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle = fyne.TextStyle{
		Italic:    true,
		Monospace: true,
		Bold:      false,
	}
	child.SetContent(text)
	child.Resize(fyne.Size{Width: 400, Height: 200})
	child.Show()
	child.Close()
}

func showLine(parent fyne.App) {
	child := parent.NewWindow("Line")
	line := canvas.NewLine(color.Black)
	line.StrokeWidth = 5
	child.SetContent(line)
	child.Resize(fyne.Size{Width: 400, Height: 200})
	child.Show()
	child.Close()
}

func showLineResize(parent fyne.App) {
	child := parent.NewWindow("Line")

	line := canvas.NewLine(color.Black)
	line.StrokeWidth = 5

	child.SetContent(line)
	child.Resize(fyne.Size{Width: 400, Height: 200})
	child.Show()

	lsize := fyne.Size{Width: 400, Height: 200}
	for {
		if lsize.Width < 0 {
			lsize.Width = 400
		}
		time.Sleep(time.Millisecond * 100)
		lsize.Width = lsize.Width - 10
		line.Resize(lsize) // Resize は、右下頂点の座標を指定
		canvas.Refresh(line)
	}
}

func showCircle(parent fyne.App) {
	child := parent.NewWindow("Circle")

	circle := canvas.NewCircle(color.Gray{0x99})
	circle.StrokeColor = color.Gray{0x55}
	circle.StrokeWidth = 5
	child.SetContent(circle)

	child.Resize(fyne.Size{Width: 400, Height: 200})
	child.Show()
}

func showImageWithFillStretch(parent fyne.App, size fyne.Size) {
	w, h := int(size.Width), int(size.Height)
	title := fmt.Sprintf("Image - FillStretch %dx%d", w, h)
	child := parent.NewWindow(title)

	img := canvas.NewImageFromFile("sample.jpg")
	child.SetContent(img)

	child.Resize(size)
	child.Show()
}

func showImage(parent fyne.App, size fyne.Size, fill canvas.ImageFill) {
	title := fmt.Sprintf("Image - %s", Fill(fill))
	child := parent.NewWindow(title)
	img := canvas.NewImageFromFile("sample.jpg")
	img.FillMode = fill
	child.SetContent(img)
	child.Resize(size)
	child.Show()
}

type Fill canvas.ImageFill

func (i Fill) String() string {
	switch canvas.ImageFill(i) {
	default:
		return "ImageFillStretch"
	case canvas.ImageFillContain:
		return "ImageFillContain"
	case canvas.ImageFillOriginal:
		return "ImageFillOriginal"
	}
}
