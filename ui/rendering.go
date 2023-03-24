package ui

import (
	"siter/config"
	"siter/utils"
	"time"

	"fyne.io/fyne/v2/widget"
)

type Rendering struct {
	config   *config.Config
	textGrid *widget.TextGrid
	buffer   [][]rune
}

func NewRendering(config *config.Config, textGrid *widget.TextGrid, buffer [][]rune) *Rendering {
	return &Rendering{config: config, textGrid: textGrid, buffer: buffer}
}

func (r *Rendering) Render() {
	go func() {
		for {
			time.Sleep(time.Duration(r.config.RepaintDeplay) * time.Millisecond)
			r.Clear()

			var lines string
			for _, line := range r.buffer {
				lines = lines + string(line)
			}

			r.Set(utils.Clear_backspace(string(lines)))
		}
	}()

}

func (r *Rendering) Set(text string) {
	r.textGrid.SetText(text)
}

func (r *Rendering) Clear() {
	r.Set("")
}
