package rendering

import (
	"image/color"
	"siter/ui"
	"strconv"
	"strings"
)

func (r *Rendering) handleSGR(rs []rune) {
	text := string(rs[2 : len(rs)-1])
	separator := ";"
	if strings.Contains(text, ":") {
		separator = ":"
	}

	params := strings.Split(text, separator)

	// CSI m equal CSI 0 m
	if len(params) == 1 && params[0] == "" {
		params[0] = "0"
	}

	isFgBoldColor := false
	isBgBoldColor := false
	isItalic := false
	underline := ui.NoUnderline
	fgMode := 0
	bgMode := 0
	underlineColorMode := 0
	fg := -1
	bg := -1
	underlineColor := -1
	fgR := -1
	fgG := -1
	fgB := -1
	bgR := -1
	bgG := -1
	bgB := -1
	underlineColorR := -1
	underlineColorG := -1
	underlineColorB := -1
	fgColor := r.termColor.Foreground
	bgColor := r.termColor.Background

	for index := 0; index < len(params); index++ {
		i := parseInt(params[index])

		if i == 0 {
			isFgBoldColor = false
			isBgBoldColor = false
			isItalic = false
			underline = ui.NoUnderline
			fgMode = 0
			bgMode = 0
			fg = -1
			bg = -1
		} else if i == 1 {
			isFgBoldColor = true
		} else if i == 3 {
			isItalic = true
		} else if i == 4 {
			if index+1 < len(params) {
				index++
				i = parseInt(params[index])
				if i == 0 {
					underline = ui.NoUnderline
				} else if i == 1 {
					underline = ui.StraightUnderline
				} else if i == 2 {
					underline = ui.DoubleUnderline
				} else if i == 3 {
					underline = ui.CurlyUnderline
				} else if i == 4 {
					underline = ui.DottedUnderline
				} else if i >= 5 {
					underline = ui.DashedUnderline
				}
			} else {
				underline = ui.StraightUnderline
			}
		} else if i == 24 {
			underline = ui.NoUnderline
		} else if i > 29 && i < 38 {
			fgMode = 0
			fg = i - 30
		} else if i == 38 {
			index, fgMode, fg, fgR, fgG, fgB = parseColor(index, params)
		} else if i > 39 && i < 48 {
			bgMode = 0
			bg = i - 40
		} else if i == 48 {
			index, bgMode, bg, bgR, bgG, bgB = parseColor(index, params)
		} else if i == 58 {
			index, underlineColorMode, underlineColor, underlineColorR, underlineColorG, underlineColorB = parseColor(index, params)
		} else if i > 89 && i < 98 {
			fgMode = 0
			fg = i - 90
			isFgBoldColor = true
		} else if i > 99 && i < 108 {
			bgMode = 0
			bg = i - 100
			isBgBoldColor = true
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
		fgColor = color.RGBA{R: uint8(fgR), G: uint8(fgG), B: uint8(fgB), A: 255}
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
		bgColor = color.RGBA{R: uint8(bgR), G: uint8(bgG), B: uint8(bgB), A: 255}
	}

	var underlineColorVal color.Color = color.Transparent
	if underlineColorMode == 1 {
		underlineColorVal = r.termColor.Color16[underlineColor]
	} else if underlineColorMode == 2 {
		underlineColorVal = color.RGBA{R: uint8(underlineColorR), G: uint8(underlineColorG), B: uint8(underlineColorB)}
	}

	r.nextStyle = &ui.TextGridStyle{
		FGColor:        fgColor,
		BGColor:        bgColor,
		Italic:         isItalic,
		Bold:           isFgBoldColor,
		Underline:      underline,
		UnderLineColor: underlineColorVal,
	}
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return n
}

func parseColor(index int, params []string) (ind int, mode int, c256 int, r int, g int, b int) {
	index++
	i, err := strconv.Atoi(params[index])
	if err != nil {
		panic(err)
	}

	if i == 5 {
		if index+1 >= len(params) {
			panic("wrong format")
		}

		mode = 1
		index++
		c256 = parseInt(params[index])
	} else if i == 2 {
		mode = 2
		if index+3 >= len(params) {
			panic("wrong format")
		}

		index++
		// Colorspace ID
		if params[index] == "" {
			if index+3 >= len(params) {
				panic("wrong format")
			}

			index++
		}

		r = parseInt(params[index])

		index++
		g = parseInt(params[index])

		index++
		b = parseInt(params[index])
	}

	ind = index

	return
}
