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
	config *config.Config
}

func Render(scrollContainer *container.Scroll, textGrid *widget.TextGrid, process io.Reader, config *config.Config) {
	r := &Rendering{config}

	rows := make([]widget.TextGridRow, 1)
	cells := make([]widget.TextGridCell, 0)
	textGrid.Rows = rows

	rows[0] = widget.TextGridRow{Cells: cells}
	rowIndex := 0

	isNewOutput := false
	isNewLine := false

	go func() {
		reader := bufio.NewReader(process)

		isCr := false
		isEsc := false
		isECMA := false
		preBuffer := ""
		style := widget.CustomTextGridStyle{FGColor: r.config.ForegroundColor.RGBA}

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
				textGrid.Rows = rows
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
					scrollContainer.ScrollToBottom()
				}

				textGrid.Refresh()
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
		r.config.Color0.RGBA,
		r.config.Color1.RGBA,
		r.config.Color2.RGBA,
		r.config.Color3.RGBA,
		r.config.Color4.RGBA,
		r.config.Color5.RGBA,
		r.config.Color6.RGBA,
		r.config.Color7.RGBA,
		r.config.Color8.RGBA,
		r.config.Color9.RGBA,
		r.config.Color10.RGBA,
		r.config.Color11.RGBA,
		r.config.Color12.RGBA,
		r.config.Color13.RGBA,
		r.config.Color14.RGBA,
		r.config.Color15.RGBA,
	}

	fgColor := r.config.ForegroundColor.RGBA
	bgColor := r.config.BackgroundColor.RGBA

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
