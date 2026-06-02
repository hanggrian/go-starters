package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/johndoe/library/library"
	library_extension "github.com/johndoe/library/library-extension"
)

func main() {
	app := app.New()
	window := app.NewWindow("My Library")
	window.Resize(fyne.NewSize(400, 300))

	label := widget.NewLabel("Hello, World!")
	container := container.New(layout.NewCenterLayout(), label)

	impl := &library_extension.LabelExtImpl{LabelImpl: library.NewLabelImpl(label)}
	window.SetContent(container)
	label.SetText(fmt.Sprintf("%d pixels at %s", impl.GetSize(), impl.GetPosition()))

	window.ShowAndRun()
}
