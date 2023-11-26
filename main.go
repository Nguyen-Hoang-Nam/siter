package main

import (
	"os"

	"siter/config"
	"siter/pty"
	"siter/rendering"
	"siter/typing"
	"siter/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	a := app.New()
	w := a.NewWindow("Siter")
	terminal := ui.NewTerminal()
	scrollContainer := container.NewVScroll(terminal)

	c := config.Load()
	p, err := pty.Start(c)
	if err != nil {
		fyne.LogError("Failed to open pty", err)
		os.Exit(1)
	}
	defer p.Close()

	typing.Load(w.Canvas(), p, c)

	rendering.Render(scrollContainer, terminal, p, c)

	a.Settings().SetTheme(ui.NewTheme(c))

	w.SetContent(container.New(layout.NewStackLayout(), scrollContainer))
	w.ShowAndRun()
}
