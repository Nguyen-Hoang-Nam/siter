package rendering

import (
	"image/color"
	"siter/termcolor"
	"siter/ui"
	"strconv"
	"strings"
)

type SGRColorMode int

const (
	Color16Mode SGRColorMode = iota
	Color256Mode
	ColorRGBMode
	ColorRGBAMode
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
	fgMode := Color16Mode
	bgMode := Color16Mode
	underlineColorMode := Color16Mode
	fg := -1
	bg := -1
	underlineColor := -1
	fgR := -1
	fgG := -1
	fgB := -1
	fgA := -1
	bgR := -1
	bgG := -1
	bgB := -1
	bgA := -1
	underlineColorR := -1
	underlineColorG := -1
	underlineColorB := -1
	underlineColorA := -1
	fgColor := r.termColor.Foreground
	bgColor := r.termColor.Background

	for index := 0; index < len(params); index++ {
		i := parseInt(params[index])

		if i == 0 {
			isFgBoldColor = false
			isBgBoldColor = false
			isItalic = false
			underline = ui.NoUnderline
			fgMode = Color16Mode
			bgMode = Color16Mode
			underlineColorMode = Color16Mode
			fg = -1
			bg = -1
			underlineColor = -1
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
			index, fgMode, fg, fgR, fgG, fgB, fgA = parseColor(index, params)
		} else if i > 39 && i < 48 {
			bgMode = 0
			bg = i - 40
		} else if i == 48 {
			index, bgMode, bg, bgR, bgG, bgB, bgA = parseColor(index, params)
		} else if i == 58 {
			index, underlineColorMode, underlineColor, underlineColorR, underlineColorG, underlineColorB, underlineColorA = parseColor(index, params)
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

	if fgMode == Color16Mode {
		if fg != -1 {
			if isFgBoldColor {
				fg += 8
			}
			fgColor = r.termColor.Color16[fg]
		}
	} else {
		fgColor = generateColor(fgMode, fg, fgR, fgG, fgB, fgA, r.termColor)
	}

	if bgMode == Color16Mode {
		if bg != -1 {
			if isBgBoldColor {
				bg += 8
			}
			bgColor = r.termColor.Color16[bg]
		}
	} else {
		bgColor = generateColor(bgMode, bg, bgR, bgG, bgB, bgA, r.termColor)
	}

	var underlineColorVal color.Color = color.Transparent
	underlineColorVal = generateColor(underlineColorMode, underlineColor, underlineColorR, underlineColorG, underlineColorB, underlineColorA, r.termColor)

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

func parseColor(index int, params []string) (ind int, mode SGRColorMode, c256 int, r int, g int, b int, a int) {
	index++
	i, err := strconv.Atoi(params[index])
	if err != nil {
		panic(err)
	}

	if i == 5 { // 256
		if index+1 >= len(params) {
			panic("wrong format")
		}

		mode = Color256Mode
		index++
		c256 = parseInt(params[index])
	} else if i == 2 { // RGB
		mode = ColorRGBMode
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
	} else if i == 6 { // RGBA
		mode = ColorRGBAMode
		if index+4 >= len(params) {
			panic("wrong format")
		}

		index++
		// Colorspace ID
		if params[index] == "" {
			if index+4 >= len(params) {
				panic("wrong format")
			}

			index++
		}

		r = parseInt(params[index])

		index++
		g = parseInt(params[index])

		index++
		b = parseInt(params[index])

		index++
		a = parseInt(params[index])
	}

	ind = index

	return
}

func generateColor(mode SGRColorMode, c256 int, r int, g int, b int, a int, termColor termcolor.TermColor) (c color.RGBA) {
	if mode == Color256Mode {
		c = termColor.Color16[c256]
	} else if mode == ColorRGBMode {
		c = color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
	} else if mode == ColorRGBAMode {
		c = color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
	}

	return
}
