package main

import (
	"os"

	"siter/config"
	"siter/events"
	"siter/pty"
	"siter/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	w := app.New().NewWindow("Siter")
	textGrid := widget.NewTextGrid()
	scrollContainer := container.NewVScroll(textGrid)

	c := config.Load()
	p, err := pty.Start(c)
	if err != nil {
		fyne.LogError("Failed to open pty", err)
		os.Exit(1)
	}
	defer p.Close()

	events.Load(w.Canvas(), p)

	ui.Render(scrollContainer, textGrid, p, c)

	w.SetContent(container.New(layout.NewMaxLayout(), scrollContainer))
	w.ShowAndRun()
}
