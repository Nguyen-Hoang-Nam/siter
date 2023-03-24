package main

import (
	"os"

	"siter/config"
	"siter/pty"
	"siter/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	conf := config.Load()

	app := app.New()
	window := app.NewWindow("Siter")
	textGrid := widget.NewTextGrid()

	process, err := pty.GetShell(*conf)
	if err != nil {
		fyne.LogError("Failed to open pty", err)
		os.Exit(1)
	}

	defer process.Close()

	ui.NewEvent(process, window.Canvas()).Load()

	buffer := [][]rune{}
	process.Read(&buffer)

	ui.NewRendering(textGrid, &buffer, conf).Render()

	window.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewMaxLayout(),
			textGrid,
		),
	)
	window.ShowAndRun()
}
