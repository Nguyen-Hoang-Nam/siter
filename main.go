package main

import (
	"os"

	"siter/config"
	"siter/mapping"
	"siter/pty"
	"siter/rendering"
	"siter/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	a := app.New()
	w := a.NewWindow("Siter")
	textGrid := ui.NewTextGrid()
	scrollContainer := container.NewVScroll(textGrid)

	c := config.Load()
	p, err := pty.Start(c)
	if err != nil {
		fyne.LogError("Failed to open pty", err)
		os.Exit(1)
	}
	defer p.Close()

	mapping.Load(w.Canvas(), p, c)

	rendering.Render(scrollContainer, textGrid, p, c)

	a.Settings().SetTheme(ui.NewTheme(c))

	w.SetContent(container.New(layout.NewMaxLayout(), scrollContainer))
	w.ShowAndRun()
}
