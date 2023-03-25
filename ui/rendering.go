package ui

import (
	"fmt"
	"siter/config"
	"siter/utils"
	"time"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Rendering struct {
	config          *config.Config
	textGrid        *widget.TextGrid
	scrollContainer *container.Scroll
	buffer          *[][]rune
}

func NewRendering(scrollContainer *container.Scroll, textGrid *widget.TextGrid, buffer *[][]rune, config *config.Config) *Rendering {
	return &Rendering{config: config, textGrid: textGrid, buffer: buffer, scrollContainer: scrollContainer}
}

func (r *Rendering) Render() {
	length := 0
	lineNo := 0

	go func() {
		for {
			time.Sleep(time.Duration(r.config.RepaintDeplay) * time.Millisecond)

			var lines string
			for _, line := range *r.buffer {
				lines = lines + string(line)
			}

			if length != len(lines) {
				cleanText := utils.ClearBackspace(lines)
				// r.Clear()
				r.Set(cleanText)
				// r.style(lines)

				length = len(lines)
			}

			if lineNo != len(*r.buffer) {
				r.scrollContainer.ScrollToBottom()

				lineNo = len(*r.buffer)
			}
		}
	}()
}

func (r *Rendering) style(text string) {
	row := 0
	col := 0

	for i := range text {
		if text[i] == '\n' {
			row++
			col = 0
			fmt.Printf("\n\n")
		} else {
			fmt.Printf("%d ", text[i])
			col++
		}
	}
}

func (r *Rendering) Set(text string) {
	r.textGrid.SetText(text)
}

func (r *Rendering) Clear() {
	r.Set("")
}
