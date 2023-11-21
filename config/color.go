package config

import (
	"image/color"
	"siter/utils"
)

type parsingColor struct{ color.RGBA }

func (c *parsingColor) UnmarshalText(text []byte) error {
	var err error

	c.RGBA, err = utils.ParseColor(string(text))

	return err
}
