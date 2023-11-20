package rendering

import (
	"image/color"
	"siter/ui"
	"strconv"
	"strings"
)

func (r *Rendering) handleControlFunction(functionName string, rs []rune) {
	switch functionName {
	case "LF":
		r.handleLF()
	case "CR":
		r.handleLF()
	case "BS":
		r.handleBS()
	case "SGR":
		r.handleSGR(rs)
	}
}

func (r *Rendering) handleLF() {
	r.cells = append(r.cells, ui.TextGridCell{
		Rune:  '\n',
		Style: r.nextStyle,
	})
	r.cells = make([]ui.TextGridCell, 0)
	r.rows = append(r.rows, ui.TextGridRow{Cells: r.cells})
	r.textGrid.Rows = r.rows
	r.rowIndex++

	if !r.isNewLine {
		r.isNewLine = true
	}
}

func (r *Rendering) handleBS() {
	r.cells = r.cells[:len(r.cells)-1]
	r.rows[r.rowIndex] = ui.TextGridRow{Cells: r.cells}

	if !r.isNewOutput {
		r.isNewOutput = true
	}
}

func (r *Rendering) handleSGR(rs []rune) {
	params := strings.Split(string(rs[2:len(rs)-1]), ";")

	// CSI m equal CSI 0 m
	if len(params) == 1 && params[0] == "" {
		params[0] = "0"
	}

	isFgBoldColor := false
	isBgBoldColor := false
	isItalic := false
	fgMode := 0
	bgMode := 0
	fg := -1
	bg := -1
	fgR := -1
	fgG := -1
	fgB := -1
	bgR := -1
	bgG := -1
	bgB := -1
	fgColor := r.termColor.Foreground
	bgColor := r.termColor.Background

	for index := 0; index < len(params); index++ {
		i := parseInt(params[index])

		if i == 0 {
			isFgBoldColor = false
			isBgBoldColor = false
			isItalic = false
			fgMode = 0
			bgMode = 0
			fg = -1
			bg = -1
		} else if i == 1 {
			isFgBoldColor = true
		} else if i == 3 {
			isItalic = true
		} else if i > 29 && i < 38 {
			fgMode = 0
			fg = i - 30
		} else if i > 39 && i < 48 {
			bgMode = 0
			bg = i - 40
		} else if i > 89 && i < 98 {
			fgMode = 0
			fg = i - 90
			isFgBoldColor = true
		} else if i > 99 && i < 108 {
			bgMode = 0
			bg = i - 100
			isBgBoldColor = true
		} else if i == 38 {
			index++
			i = parseInt(params[index])

			if i == 5 {
				fgMode = 1
				if index+1 < len(params) {
					index++
					fg = parseInt(params[index])
				} else {
					panic("wrong format")
				}
			} else if i == 2 {
				fgMode = 2
				fg = -1
				if index+3 < len(params) {
					index++
					fgR = parseInt(params[index])

					index++
					fgG = parseInt(params[index])

					index++
					fgB = parseInt(params[index])
				} else {
					panic("wrong format")
				}
			}
		} else if i == 48 {
			index++
			i, err := strconv.Atoi(params[index])
			if err != nil {
				panic(err)
			}

			if i == 5 {
				bgMode = 1
				if index+1 < len(params) {
					index++
					bg = parseInt(params[index])
				} else {
					panic("wrong format")
				}
			} else if i == 2 {
				bgMode = 2
				if index+3 < len(params) {
					index++
					bgR = parseInt(params[index])

					index++
					bgG = parseInt(params[index])

					index++
					bgB = parseInt(params[index])
				} else {
					panic("wrong format")
				}
			}
		}
	}

	if fgMode == 0 {
		if fg != -1 {
			if isFgBoldColor {
				fg += 8
			}
			fgColor = r.termColor.Color16[fg]
		}
	} else if fgMode == 1 {
		fgColor = r.termColor.Color256[fg]
	} else if fgMode == 2 {
		fgColor = color.RGBA{R: uint8(fgR), G: uint8(fgG), B: uint8(fgB)}
	}

	if bgMode == 0 {
		if bg != -1 {
			if isBgBoldColor {
				bg += 8
			}
			bgColor = r.termColor.Color16[bg]
		}
	} else if bgMode == 1 {
		bgColor = r.termColor.Color16[bg]
	} else if bgMode == 2 {
		bgColor = color.RGBA{R: uint8(bgR), G: uint8(bgG), B: uint8(bgB)}
	}

	r.nextStyle = &ui.TextGridStyle{FGColor: fgColor, BGColor: bgColor, Italic: isItalic, Bold: isFgBoldColor}
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return n
}
