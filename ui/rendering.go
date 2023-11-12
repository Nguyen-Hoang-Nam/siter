package ui

import (
	"bufio"
	"image/color"
	"io"
	"os"
	"siter/config"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Rendering struct {
	config          *config.Config
	textGrid        *widget.TextGrid
	scrollContainer *container.Scroll
	termianlColor   map[string]color.RGBA
	reader          io.Reader
}

func NewRendering(scrollContainer *container.Scroll, textGrid *widget.TextGrid, reader io.Reader, config *config.Config) *Rendering {
	terminalColor := map[string]color.RGBA{
		"[0m":     config.ForegroundColor,
		"[0;30m":  config.Color0,
		"[0;31m":  config.Color1,
		"[0;32m":  config.Color2,
		"[0;33m":  config.Color3,
		"[0;34m":  config.Color4,
		"[0;35m":  config.Color5,
		"[0;36m":  config.Color6,
		"[0;37m":  config.Color7,
		"[01;30m": config.Color8,
		"[01;31m": config.Color9,
		"[01;32m": config.Color10,
		"[01;33m": config.Color11,
		"[01;34m": config.Color12,
		"[01;35m": config.Color13,
		"[01;36m": config.Color14,
		"[01;37m": config.Color15,
	}

	return &Rendering{config: config, textGrid: textGrid, reader: reader, scrollContainer: scrollContainer, termianlColor: terminalColor}
}

func (r *Rendering) Render() {
	rows := make([]widget.TextGridRow, 1)
	cells := make([]widget.TextGridCell, 0)
	r.textGrid.Rows = rows

	rows[0] = widget.TextGridRow{Cells: cells}
	rowIndex := 0

	isNewOutput := false
	isNewLine := false

	go func() {
		reader := bufio.NewReader(r.reader)

		isCr := false
		isEsc := false
		isECMA := false
		preBuffer := ""
		style := widget.CustomTextGridStyle{FGColor: r.termianlColor["[0m"]}

		for {
			c, _, err := reader.ReadRune()
			if err != nil {
				if err == io.EOF {
					return
				}
				os.Exit(0)
			}

			if isCr {
				if c != '\n' {
					cells = make([]widget.TextGridCell, 0)
				}

				isCr = false
			}

			if isEsc {
				if !isECMA {
					if preBuffer == "" && c == ']' {
						isECMA = true
					} else if c != 'm' {
						preBuffer += string(c)
					}
				}
			}

			if c == '\b' {
				cells = cells[:len(cells)-1]
			} else if c == '\r' {
				isCr = true
			} else if c == 27 {
				isEsc = true
			} else if c != '\n' && !isEsc {
				cells = append(cells, widget.TextGridCell{Rune: c, Style: &style})
			}

			if isEsc && isECMA && c == ';' {
				isEsc = false
				isECMA = false
			}

			if isEsc && c == 'm' {
				style = r.getSGRStyle(preBuffer)
				isEsc = false
				preBuffer = ""
			}

			if isEsc && c == 'l' {
				isEsc = false
				preBuffer = ""
			}

			rows[rowIndex] = widget.TextGridRow{Cells: cells}
			if c == '\n' {
				cells = make([]widget.TextGridCell, 0)
				rows = append(rows, widget.TextGridRow{Cells: cells})
				r.textGrid.Rows = rows
				rowIndex++

				if !isNewLine {
					isNewLine = true
				}
			}

			if !isNewOutput {
				isNewOutput = true
			}
		}
	}()

	go func() {
		for {
			time.Sleep(time.Duration(r.config.RepaintDeplay) * time.Millisecond)

			if isNewOutput {
				isNewOutput = false

				if isNewLine {
					isNewLine = false
					r.scrollContainer.ScrollToBottom()
				}

				r.textGrid.Refresh()
			}
		}
	}()
}

func (r *Rendering) getSGRStyle(text string) widget.CustomTextGridStyle {
	params := strings.Split(text, ";")

	isFgBoldColor := false
	isBgBoldColor := false
	isFG8 := true
	isBG8 := true
	fg := -1
	bg := -1
	fgR := -1
	fgG := -1
	fgB := -1
	bgR := -1
	bgG := -1
	bgB := -1
	for index := 1; index < len(params); index++ {
		i, err := strconv.Atoi(params[index])
		if err != nil {
			panic(err)
		}

		if i == 0 {
			isFgBoldColor = false
			isBgBoldColor = false
			isFG8 = true
			isBG8 = true
			fg = -1
			bg = -1
			fgR = -1
			fgG = -1
			fgB = -1
			bgR = -1
			bgG = -1
			bgB = -1
		} else if i == 1 {
			isFgBoldColor = true
		} else if i > 29 && i < 38 {
			isFG8 = true
			fg = i - 30
		} else if i > 39 && i < 48 {
			isBG8 = true
			bg = i - 40
		} else if i > 89 && i < 98 {
			isFG8 = true
			fg = i - 90
			isFgBoldColor = true
		} else if i > 99 && i < 108 {
			isBG8 = true
			bg = i - 100
			isBgBoldColor = true
		} else if i == 38 {
			isFG8 = false

			index++
			i, err := strconv.Atoi(params[index])
			if err != nil {
				panic(err)
			}

			if i == 5 {
				if index+1 < len(params) {
					index++
					fg, err = strconv.Atoi(params[index+1])
					if err != nil {
						panic(err)
					}
				} else {
					panic("wrong format")
				}
			} else if i == 2 {
				fg = -1
				if index+3 < len(params) {
					index++
					fgR, err = strconv.Atoi(params[index])
					if err != nil {
						panic(err)
					}

					index++
					fgG, err = strconv.Atoi(params[index])
					if err != nil {
						panic(err)
					}

					index++
					fgB, err = strconv.Atoi(params[index])
					if err != nil {
						panic(err)
					}
				} else {
					panic("wrong format")
				}
			}
		} else if i == 48 {
			isBG8 = false

			index++
			i, err := strconv.Atoi(params[index])
			if err != nil {
				panic(err)
			}

			if i == 5 {
				if index+1 < len(params) {
					index++
					bg, err = strconv.Atoi(params[index+1])
					if err != nil {
						panic(err)
					}
				} else {
					panic("wrong format")
				}
			} else if i == 2 {
				bg = -1
				if index+3 < len(params) {
					index++
					bgR, err = strconv.Atoi(params[index])
					if err != nil {
						panic(err)
					}

					index++
					bgG, err = strconv.Atoi(params[index])
					if err != nil {
						panic(err)
					}

					index++
					bgB, err = strconv.Atoi(params[index])
					if err != nil {
						panic(err)
					}
				} else {
					panic("wrong format")
				}
			}
		}
	}

	colors := []color.RGBA{
		r.config.Color0,
		r.config.Color1,
		r.config.Color2,
		r.config.Color3,
		r.config.Color4,
		r.config.Color5,
		r.config.Color6,
		r.config.Color7,
		r.config.Color8,
		r.config.Color9,
		r.config.Color10,
		r.config.Color11,
		r.config.Color12,
		r.config.Color13,
		r.config.Color14,
		r.config.Color15,
	}

	fgColor := r.config.ForegroundColor
	bgColor := r.config.BackgroundColor

	if isFG8 {
		if fg != -1 {
			if isFgBoldColor {
				fg += 8
			}
			fgColor = colors[fg]
		}
	} else {
		if fg == -1 {
			fgColor = color.RGBA{R: uint8(fgR), G: uint8(fgG), B: uint8(fgB)}
		} else {
			panic("Not support 256 color")
		}
	}

	if isBG8 {
		if bg != -1 {
			if isBgBoldColor {
				bg += 8
			}
			bgColor = colors[bg]
		}
	} else {
		if bg == -1 {
			bgColor = color.RGBA{R: uint8(bgR), G: uint8(bgG), B: uint8(bgB)}
		} else {
			panic("Not support 256 color")
		}
	}

	return widget.CustomTextGridStyle{FGColor: fgColor, BGColor: bgColor}
}

// func (r *Rendering) style(text string) {
// 	row := 0
// 	col := 0
// 	isNewColor := false
// 	asciiColor := []byte("")
// 	currentColor := r.termianlColor["[0"]

// 	for i := range text {
// 		if text[i] == 27 {
// 			isNewColor = true
// 		} else {
// 			if isNewColor && text[i] == 'm' {
// 				currentColor = r.termianlColor[string(asciiColor)]
// 				asciiColor = []byte("")

// 				isNewColor = false
// 			} else if isNewColor {
// 				asciiColor = append(asciiColor, text[i])
// 			}
// 		}

// 		r.textGrid.SetStyle(row, col, &widget.CustomTextGridStyle{FGColor: currentColor})

// 		if !isNewColor {
// 			if text[i] == '\n' {
// 				row++
// 				col = 0
// 			} else {
// 				col++
// 			}
// 		}

// 	}
// }

func (r *Rendering) Set(text string) {
	r.textGrid.SetText(text)
}

func (r *Rendering) Clear() {
	r.Set("")
}
