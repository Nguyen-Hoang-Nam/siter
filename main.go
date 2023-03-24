package main

import (
	"os"

	"siter/config"
	"siter/ui"
	"siter/utils"

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

	process, err := utils.GetShell(conf.Shell)
	if err != nil {
		fyne.LogError("Failed to open pty", err)
		os.Exit(1)
	}

	defer process.Close()

	ui.NewEvent(process, window.Canvas()).Load()

	buffer := [][]rune{}
	process.Read(buffer)

	ui.NewRendering(conf, textGrid, buffer).Render()

	window.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewGridWrapLayout(fyne.NewSize(900, 325)),
			textGrid,
		),
	)
	window.ShowAndRun()
}
