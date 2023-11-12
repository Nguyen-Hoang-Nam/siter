package main

import (
	"os"

	"siter/config"
	"siter/pty"
	"siter/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	conf := config.Load()

	app := app.New()
	window := app.NewWindow("Siter")
	textGrid := widget.NewTextGrid()
	scrollContainer := container.NewVScroll(textGrid)

	process, err := pty.GetShell(*conf)
	if err != nil {
		fyne.LogError("Failed to open pty", err)
		os.Exit(1)
	}

	defer process.Close()

	ui.NewEvent(process, window.Canvas()).Load()

	reader := process.Read()
	ui.NewRendering(scrollContainer, textGrid, reader, conf).Render()

	window.SetContent(
		container.New(layout.NewMaxLayout(), scrollContainer),
	)

	window.ShowAndRun()
}
