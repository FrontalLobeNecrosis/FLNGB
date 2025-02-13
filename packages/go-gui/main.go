package GBGUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Application that manages the emulator and its settings
	FLNGB := app.New()
	gbWindow := FLNGB.NewWindow("FLNGB")

	// This is where we will set the options for the window

	// This sets the window where the emulator is running as master
	// so that if it is closed the app stops running
	gbWindow.SetMaster()
	// TODO: Make the new sizes a variable read and wrote to a settings file
	gbWindow.Resize((fyne.NewSize(800, 600)))

	// MenuItem is labeled Exit and will exit the window weird behaviour that there
	// is already a MenuItem that is labeled Quit and does the same task
	// renaming to Quit will make it appear as the same MenuItem
	// TODO: figure out if this default Quit can be removed
	exitOption := fyne.NewMenuItem("Exit", func() { FLNGB.Quit() })
	fileMenu := fyne.NewMenu("File", exitOption)
	FLNGBMainMenu := fyne.NewMainMenu(fileMenu)
	// This is a button that will quit the app if pressed
	// it's a placeholder just to have the window display something
	// TODO: Remove and replace with the actual emulation of the GB
	quitButton := widget.NewButton("Exit", func() {
		FLNGB.Quit()
	})
	buttons := container.NewVBox(
		quitButton,
	)

	icon, _ := fyne.LoadResourceFromPath("./assets/GameBoy.png")

	// This is where to add elements to the window

	gbWindow.SetContent(
		buttons,
	)
	gbWindow.SetIcon(icon)
	gbWindow.SetMainMenu(FLNGBMainMenu)
	gbWindow.ShowAndRun()
}
