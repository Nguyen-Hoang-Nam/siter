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
	buffer   *[][]rune
}

func NewRendering(textGrid *widget.TextGrid, buffer *[][]rune, config *config.Config) *Rendering {
	return &Rendering{config: config, textGrid: textGrid, buffer: buffer}
}

func (r *Rendering) Render() {
	length := 0

	go func() {
		for {
			time.Sleep(time.Duration(r.config.RepaintDeplay) * time.Millisecond)

			var lines string
			for _, line := range *r.buffer {
				lines = lines + string(line)
			}

			if length != len(lines) {
				r.Clear()
				r.Set(utils.Clear_backspace(string(lines)))

				length = len(lines)
			}
		}
	}()
}

func (r *Rendering) Set(text string) {
	r.textGrid.SetText(text)
}

func (r *Rendering) Clear() {
	r.Set("")
}
