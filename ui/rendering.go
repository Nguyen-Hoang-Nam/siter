package ui

import (
	"image/color"
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
	termianlColor   map[string]color.RGBA
}

func NewRendering(scrollContainer *container.Scroll, textGrid *widget.TextGrid, buffer *[][]rune, config *config.Config) *Rendering {
	terminalColor := map[string]color.RGBA{
		"[0":     config.ForegroundColor,
		"[0;30":  config.Color0,
		"[0;31":  config.Color1,
		"[0;32":  config.Color2,
		"[0;33":  config.Color3,
		"[0;34":  config.Color4,
		"[0;35":  config.Color5,
		"[0;36":  config.Color6,
		"[0;37":  config.Color7,
		"[01;30": config.Color8,
		"[01;31": config.Color9,
		"[01;32": config.Color10,
		"[01;33": config.Color11,
		"[01;34": config.Color12,
		"[01;35": config.Color13,
		"[01;36": config.Color14,
		"[01;37": config.Color15,
	}

	return &Rendering{config: config, textGrid: textGrid, buffer: buffer, scrollContainer: scrollContainer, termianlColor: terminalColor}
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
				cleanBackspaceText := utils.ClearBackspace(lines)
				// r.Clear()
				r.Set(utils.ClearColor(cleanBackspaceText))
				r.style(cleanBackspaceText)

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
	isNewColor := false
	asciiColor := []byte("")
	currentColor := r.termianlColor["[0"]

	for i := range text {
		if text[i] == 27 {
			isNewColor = true
		} else {
			if isNewColor && text[i] == 'm' {
				currentColor = r.termianlColor[string(asciiColor)]
				asciiColor = []byte("")

				isNewColor = false
			} else if isNewColor {
				asciiColor = append(asciiColor, text[i])
			}
		}

		r.textGrid.SetStyle(row, col, &widget.CustomTextGridStyle{FGColor: currentColor})

		if !isNewColor {
			if text[i] == '\n' {
				row++
				col = 0
			} else {
				col++
			}
		}

	}
}

func (r *Rendering) Set(text string) {
	r.textGrid.SetText(text)
}

func (r *Rendering) Clear() {
	r.Set("")
}
