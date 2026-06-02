package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("My Application")
	window.Resize(fyne.NewSize(400, 300))

	label := widget.NewLabel("Hello, World!")
	container := container.New(layout.NewCenterLayout(), label)

	window.SetContent(container)
	label.SetText(fmt.Sprintf("%d pixels", (&LabelImpl{label}).GetSize()))

	window.ShowAndRun()
}
