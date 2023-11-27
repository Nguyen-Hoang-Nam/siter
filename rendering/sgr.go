package rendering

import (
	"errors"
	"image/color"
	"siter/internal/termcolor"
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

	// Mode
	fgColorIntensity := ui.NormalIntenisty
	bgColorIntensity := ui.NormalIntenisty
	isItalic := false
	underline := ui.NoUnderline
	blink := ui.NoBlink
	isInverse := false
	isInvisible := false
	isStrike := false
	isBgDefault := true
	isOverline := false
	isUnderlineColorDefault := true
	verticalAlign := ui.NormalBaseline

	fgColorMode := Color16Mode
	fgColor, fgColorR, fgColorG, fgColorB, fgColorA := -1, -1, -1, -1, -1
	var fgColorVal color.Color = r.termColor.Foreground

	bgColorMode := Color16Mode
	bgColor, bgColorR, bgColorG, bgColorB, bgColorA := -1, -1, -1, -1, -1
	var bgColorVal color.Color = r.termColor.Background

	underlineColorMode := Color16Mode
	underlineColor, underlineColorR, underlineColorG, underlineColorB, underlineColorA := -1, -1, -1, -1, -1
	var underlineColorVal color.Color = r.termColor.Foreground

	for index := 0; index < len(params); index++ {
		i := parseInt(params[index])

		if i == 0 {
			fgColorIntensity = ui.NormalIntenisty
			bgColorIntensity = ui.NormalIntenisty
			isItalic = false
			underline = ui.NoUnderline
			blink = ui.NoBlink
			isInverse = false
			isInvisible = false
			isStrike = false
			isOverline = false
			verticalAlign = ui.NormalBaseline

			fgColorMode = Color16Mode
			fgColor = -1

			isBgDefault = true
			bgColorMode = Color16Mode
			bgColor = -1

			isUnderlineColorDefault = true
			underlineColorMode = Color16Mode
			underlineColor = -1
		} else if i == 1 {
			fgColorIntensity = ui.BoldIntensity
		} else if i == 2 {
			fgColorIntensity = ui.DimIntensity
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
		} else if i == 5 {
			blink = ui.NormalBlink
		} else if i == 6 {
			blink = ui.RapidBlink
		} else if i == 7 {
			isInverse = true
		} else if i == 8 {
			isInvisible = true
		} else if i == 9 {
			isStrike = true
		} else if i == 21 {
			underline = ui.DoubleUnderline
		} else if i == 22 {
			fgColorIntensity = ui.NormalIntenisty
		} else if i == 23 {
			isItalic = false
		} else if i == 24 {
			underline = ui.NoUnderline
		} else if i == 25 {
			blink = ui.NormalBlink
		} else if i == 27 {
			isInverse = false
		} else if i == 28 {
			isInvisible = false
		} else if i == 29 {
			isStrike = false
		} else if i > 29 && i < 38 {
			fgColorMode = 0
			fgColor = i - 30
		} else if i == 38 {
			ind, mode, c256, r, g, b, a, err := parseColor(index, params)
			if err == nil {
				index, fgColorMode, fgColor, fgColorR, fgColorG, fgColorB, fgColorA = ind, mode, c256, r, g, b, a
			}
		} else if i > 39 && i < 48 {
			isBgDefault = false
			bgColorMode = 0
			bgColor = i - 40
		} else if i == 48 {
			isBgDefault = false
			ind, mode, c256, r, g, b, a, err := parseColor(index, params)
			if err == nil {
				index, bgColorMode, bgColor, bgColorR, bgColorG, bgColorB, bgColorA = ind, mode, c256, r, g, b, a
			}
		} else if i == 49 {
			isBgDefault = true
		} else if i == 53 {
			isOverline = true
		} else if i == 55 {
			isOverline = false
		} else if i == 58 {
			isUnderlineColorDefault = false
			ind, mode, c256, r, g, b, a, err := parseColor(index, params)
			if err == nil {
				index, underlineColorMode, underlineColor, underlineColorR, underlineColorG, underlineColorB, underlineColorA = ind, mode, c256, r, g, b, a
			}
		} else if i == 59 {
			isUnderlineColorDefault = true
		} else if i == 73 {
			verticalAlign = ui.SuperScript
		} else if i == 74 {
			verticalAlign = ui.SubScript
		} else if i == 75 {
			verticalAlign = ui.NormalBaseline
		} else if i > 89 && i < 98 {
			fgColorMode = 0
			fgColor = i - 90
		} else if i > 99 && i < 108 {
			isBgDefault = false
			bgColorMode = 0
			bgColor = i - 100
			bgColorIntensity = ui.BoldIntensity
		}
	}

	if fgColorMode == Color16Mode {
		if fgColor != -1 {
			if fgColorIntensity == ui.BoldIntensity {
				fgColor += 8
			}
			fgColorVal = r.termColor.Color16[fgColor]
		}
	} else {
		fgColorVal = generateColor(fgColorMode, fgColor, fgColorR, fgColorG, fgColorB, fgColorA, r.termColor)
	}

	if bgColorMode == Color16Mode {
		if !isBgDefault {
			if bgColorIntensity == ui.BoldIntensity {
				bgColor += 8
			}
			bgColorVal = r.termColor.Color16[bgColor]
		}
	} else {
		bgColorVal = generateColor(bgColorMode, bgColor, bgColorR, bgColorG, bgColorB, bgColorA, r.termColor)
	}

	if underline != ui.NoUnderline {
		if underlineColorMode == Color16Mode {
			if !isUnderlineColorDefault {
				underlineColorVal = r.termColor.Color16[underlineColor]
			}
		} else {
			underlineColorVal = generateColor(underlineColorMode, underlineColor, underlineColorR, underlineColorG, underlineColorB, underlineColorA, r.termColor)
		}
	} else {
		underlineColorVal = color.Transparent
	}

	// Todo: optimize underlineColorVal override
	if isInverse {
		fgColorVal, bgColorVal, underlineColorVal = bgColorVal, fgColorVal, bgColorVal
	}

	if bgColorVal == r.termColor.Background {
		bgColorVal = color.Transparent
	}

	r.nextStyle.ForegroundColor = fgColorVal
	r.nextStyle.BackgroundColor = bgColorVal
	r.nextStyle.FontStyle = ui.FontStyle{
		Italic: isItalic,
		Bold:   fgColorIntensity == ui.BoldIntensity,
	}
	r.nextStyle.Underline = underline
	r.nextStyle.UnderlineColor = underlineColorVal
	r.nextStyle.Overline = isOverline
	r.nextStyle.OverlineColor = fgColorVal
	r.nextStyle.Strike = isStrike
	r.nextStyle.StrikeColor = fgColorVal
	r.nextStyle.Blink = blink
	r.nextStyle.Invisible = isInvisible
	r.nextStyle.VerticalAlign = verticalAlign

	if fgColorIntensity == ui.BoldIntensity {
		r.nextStyle.UnderlineWidth = 2
		r.nextStyle.OverlineWidth = 2
		r.nextStyle.StrikeWidth = 2
	} else {
		r.nextStyle.UnderlineWidth = 1
		r.nextStyle.OverlineWidth = 1
		r.nextStyle.StrikeWidth = 1
	}
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return n
}

func parseColor(index int, params []string) (ind int, mode SGRColorMode, c256 int, r int, g int, b int, a int, err error) {
	index++
	i, err := strconv.Atoi(params[index])
	if err != nil {
		return
	}

	if i == 5 { // 256
		if index+1 >= len(params) {
			err = errors.New("WRONG_FORMAT")
			return
		}

		mode = Color256Mode
		index++
		c256 = parseInt(params[index])
	} else if i == 2 { // RGB
		if index+3 >= len(params) {
			err = errors.New("WRONG_FORMAT")
			return
		}

		index++
		// Colorspace ID
		if params[index] == "" {
			if index+3 >= len(params) {
				err = errors.New("WRONG_FORMAT")
				return
			}

			index++
		}

		r = parseInt(params[index])
		index++
		g = parseInt(params[index])
		index++
		b = parseInt(params[index])
		mode = ColorRGBMode
	} else if i == 6 { // RGBA
		if index+4 >= len(params) {
			err = errors.New("WRONG_FORMAT")
			return
		}

		index++
		// Colorspace ID
		if params[index] == "" {
			if index+4 >= len(params) {
				err = errors.New("WRONG_FORMAT")
				return
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
		mode = ColorRGBAMode
	}

	ind = index

	return
}

func generateColor(mode SGRColorMode, c256 int, r int, g int, b int, a int, termColor termcolor.TermColor) (c color.RGBA) {
	if mode == Color256Mode {
		c = termColor.Color256[c256]
	} else if mode == ColorRGBMode {
		c = color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
	} else if mode == ColorRGBAMode {
		c = color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
	}

	return
}
