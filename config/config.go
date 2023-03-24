package config

import (
	"image/color"
	"siter/utils"
)

type CursorShapeType string

type UrlStyleType string

type Config struct {
	FontFamily        string          `toml:"font_family"`
	BoldFont          string          `toml:"bold_font"`
	ItalicFont        string          `toml:"italic_font"`
	BoldItalicFont    string          `toml:"bold_italic_font"`
	FontSize          float64         `toml:"font_size"`
	CursorColor       color.RGBA      `toml:"cursor_color"`
	CursorTextColor   color.RGBA      `toml:"cursor_text_color"`
	CursorShape       CursorShapeType `toml:"cursor_shape"`
	ScrollbackLines   int             `toml:"scrollback_lines"`
	UrlColor          color.RGBA      `toml:"url_color"`
	UrlStyle          UrlStyleType    `toml:"url_style"`
	OpenUrlWith       string          `toml:"open_url_with"`
	UrlPrefixes       []string        `toml:"url_prefixes"`
	DetectUrls        bool            `toml:"detect_urls"`
	RepaintDeplay     int             `toml:"repaint_deplay"`
	ForegroundColor   color.RGBA      `toml:"foreground_color"`
	BackgroundColor   color.RGBA      `toml:"background_color"`
	BackgroundOpacity float64         `toml:"background_opacity"`
	BackgroundImage   string          `toml:"background_image"`
	Color0            color.RGBA      `toml:"color0"`
	Color1            color.RGBA      `toml:"color1"`
	Color2            color.RGBA      `toml:"color2"`
	Color3            color.RGBA      `toml:"color3"`
	Color4            color.RGBA      `toml:"color4"`
	Color5            color.RGBA      `toml:"color5"`
	Color6            color.RGBA      `toml:"color6"`
	Color7            color.RGBA      `toml:"color7"`
	Color8            color.RGBA      `toml:"color8"`
	Color9            color.RGBA      `toml:"color9"`
	Color10           color.RGBA      `toml:"color10"`
	Color11           color.RGBA      `toml:"color11"`
	Color12           color.RGBA      `toml:"color12"`
	Color13           color.RGBA      `toml:"color13"`
	Color14           color.RGBA      `toml:"color14"`
	Color15           color.RGBA      `toml:"color15"`
	Shell             string          `toml:"shell"`
}

const (
	CURSOR_SHAPE_BLOCK     CursorShapeType = "BLOCK"
	CURSOR_SHAPE_BEAM      CursorShapeType = "BEAM"
	CURSOR_SHAPE_UNDERLINE CursorShapeType = "UNDERLINE"
)

const (
	URL_STYLE_NONE     UrlStyleType = "NONE"
	URL_STYLE_STRAIGHT UrlStyleType = "STRAIGHT"
	URL_STYLE_DOUBLE   UrlStyleType = "DOUBLE"
	URL_STYLE_CURLY    UrlStyleType = "CURLY"
	URL_STYLE_DOTTED   UrlStyleType = "DOTTED"
	URL_STYLE_DASHED   UrlStyleType = "DASHED"
)

func Load() *Config {
	conf, err := parseFile()
	if err != nil {
		conf.setDefault()
	}

	return conf
}

func (c *Config) setDefault() {
	c.FontFamily = "monospace"
	c.BoldFont = "monospace"
	c.ItalicFont = "monospace"
	c.BoldItalicFont = "monospace"
	c.FontSize = 16.0
	c.CursorColor, _ = utils.ParseColor("#6796E6")
	c.CursorTextColor, _ = utils.ParseColor("#111")
	c.CursorShape = CURSOR_SHAPE_BLOCK
	c.ScrollbackLines = 2000
	c.UrlColor, _ = utils.ParseColor("#0087bd")
	c.UrlStyle = URL_STYLE_CURLY
	c.OpenUrlWith = "default"
	c.UrlPrefixes = []string{"http", "https", "ssh"}
	c.DetectUrls = true
	c.RepaintDeplay = 10
	c.ForegroundColor, _ = utils.ParseColor("#F1F1F0")
	c.BackgroundColor, _ = utils.ParseColor("#1F212A")
	c.Color0, _ = utils.ParseColor("#282A36")
	c.Color1, _ = utils.ParseColor("#FF5C57")
	c.Color2, _ = utils.ParseColor("#5AF78E")
	c.Color3, _ = utils.ParseColor("#F3F99D")
	c.Color4, _ = utils.ParseColor("#57C7FF")
	c.Color5, _ = utils.ParseColor("#FF6AC1")
	c.Color6, _ = utils.ParseColor("#9AEDFE")
	c.Color7, _ = utils.ParseColor("#F1F1F0")
	c.Color8, _ = utils.ParseColor("#686868")
	c.Color9, _ = utils.ParseColor("#F66151")
	c.Color10, _ = utils.ParseColor("#33DA7A")
	c.Color11, _ = utils.ParseColor("#E9AD0C")
	c.Color15, _ = utils.ParseColor("#FFF")
	c.Color12, _ = utils.ParseColor("#2A7BDE")
	c.Color13, _ = utils.ParseColor("#C061CB")
	c.Color14, _ = utils.ParseColor("#33C7DE")
	c.Shell = "."
}
