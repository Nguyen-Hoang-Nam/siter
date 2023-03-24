package utils

import (
	"errors"
	"image/color"
)

var (
	ERROR_INVALID_COLOR_FORMAT = errors.New("INVALID_FORMAT")
)

func ParseColor(colorStr string) (c color.RGBA, err error) {
	if colorStr[0] == '#' {
		return parseHexColor(colorStr)
	}

	return c, ERROR_INVALID_COLOR_FORMAT
}

// Credit https://stackoverflow.com/a/54200713
func parseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = ERROR_INVALID_COLOR_FORMAT
	}

	return
}

func hexToByte(b byte) byte {
	switch {
	case b >= '0' && b <= '9':
		return b - '0'
	case b >= 'a' && b <= 'f':
		return b - 'a' + 10
	case b >= 'A' && b <= 'F':
		return b - 'A' + 10
	}

	return 0
}
