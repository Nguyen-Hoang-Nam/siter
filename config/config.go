package config

import (
	"image/color"
	"siter/utils"
)

type CursorShapeType string

const (
	CURSOR_SHAPE_BLOCK     CursorShapeType = "BLOCK"
	CURSOR_SHAPE_BEAM      CursorShapeType = "BEAM"
	CURSOR_SHAPE_UNDERLINE CursorShapeType = "UNDERLINE"
)

type UrlStyleType string

const (
	URL_STYLE_NONE     UrlStyleType = "NONE"
	URL_STYLE_STRAIGHT UrlStyleType = "STRAIGHT"
	URL_STYLE_DOUBLE   UrlStyleType = "DOUBLE"
	URL_STYLE_CURLY    UrlStyleType = "CURLY"
	URL_STYLE_DOTTED   UrlStyleType = "DOTTED"
	URL_STYLE_DASHED   UrlStyleType = "DASHED"
)

type parsingColor struct{ color.RGBA }

func (c *parsingColor) UnmarshalText(text []byte) error {
	var err error

	c.RGBA, err = utils.ParseColor(string(text))

	return err
}

type Config struct {
	FontFamily        string          `toml:"font_family"`
	BoldFont          string          `toml:"bold_font"`
	ItalicFont        string          `toml:"italic_font"`
	BoldItalicFont    string          `toml:"bold_italic_font"`
	FontSize          float64         `toml:"font_size"`
	CursorColor       parsingColor    `toml:"cursor_color"`
	CursorTextColor   parsingColor    `toml:"cursor_text_color"`
	CursorShape       CursorShapeType `toml:"cursor_shape"`
	ScrollbackLines   int             `toml:"scrollback_lines"`
	UrlColor          parsingColor    `toml:"url_color"`
	UrlStyle          UrlStyleType    `toml:"url_style"`
	OpenUrlWith       string          `toml:"open_url_with"`
	UrlPrefixes       []string        `toml:"url_prefixes"`
	DetectUrls        bool            `toml:"detect_urls"`
	RepaintDeplay     int             `toml:"repaint_deplay"`
	ForegroundColor   parsingColor    `toml:"foreground_color"`
	BackgroundColor   parsingColor    `toml:"background_color"`
	BackgroundOpacity float64         `toml:"background_opacity"`
	BackgroundImage   string          `toml:"background_image"`
	Color0            parsingColor    `toml:"color0"`
	Color1            parsingColor    `toml:"color1"`
	Color2            parsingColor    `toml:"color2"`
	Color3            parsingColor    `toml:"color3"`
	Color4            parsingColor    `toml:"color4"`
	Color5            parsingColor    `toml:"color5"`
	Color6            parsingColor    `toml:"color6"`
	Color7            parsingColor    `toml:"color7"`
	Color8            parsingColor    `toml:"color8"`
	Color9            parsingColor    `toml:"color9"`
	Color10           parsingColor    `toml:"color10"`
	Color11           parsingColor    `toml:"color11"`
	Color12           parsingColor    `toml:"color12"`
	Color13           parsingColor    `toml:"color13"`
	Color14           parsingColor    `toml:"color14"`
	Color15           parsingColor    `toml:"color15"`
	Shell             string          `toml:"shell"`
}

func Load() *Config {
	userConfig(&defaultConfig)

	return &defaultConfig
}
