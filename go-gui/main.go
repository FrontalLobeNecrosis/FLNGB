package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	FLNGB := app.New()
	gbWindow := FLNGB.NewWindow("FLNGB")

	quitButton := widget.NewButton("Quit", func() {
		FLNGB.Quit()
	})
	gbWindow.SetContent(container.NewVBox(
		quitButton,
	))

	gbWindow.ShowAndRun()
}
