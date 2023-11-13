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

	window := app.New().NewWindow("Siter")
	textGrid := widget.NewTextGrid()
	scrollContainer := container.NewVScroll(textGrid)

	conf := config.Load()
	process, err := pty.Start(conf)
	if err != nil {
		fyne.LogError("Failed to open pty", err)
		os.Exit(1)
	}
	defer process.Close()

	ui.LoadEvent(window.Canvas(), process)
	ui.Render(scrollContainer, textGrid, process, conf)

	window.SetContent(container.New(layout.NewMaxLayout(), scrollContainer))
	window.ShowAndRun()
}
