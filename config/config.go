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

var defaultConfig = Config{
	FontFamily:      "monospace",
	BoldFont:        "monospace",
	ItalicFont:      "monospace",
	BoldItalicFont:  "monospace",
	FontSize:        16.0,
	CursorColor:     parseColor("#6796E6"),
	CursorTextColor: parseColor("#111"),
	CursorShape:     CURSOR_SHAPE_BLOCK,
	ScrollbackLines: 2000,
	UrlColor:        parseColor("#0087bd"),
	UrlStyle:        URL_STYLE_CURLY,
	OpenUrlWith:     "default",
	UrlPrefixes:     []string{"http", "https", "ssh"},
	DetectUrls:      true,
	RepaintDeplay:   100,
	ForegroundColor: parseColor("#F1F1F0"),
	BackgroundColor: parseColor("#1F212A"),
	Color0:          parseColor("#282A36"),
	Color1:          parseColor("#FF5C57"),
	Color2:          parseColor("#5AF78E"),
	Color3:          parseColor("#F3F99D"),
	Color4:          parseColor("#57C7FF"),
	Color5:          parseColor("#FF6AC1"),
	Color6:          parseColor("#9AEDFE"),
	Color7:          parseColor("#F1F1F0"),
	Color8:          parseColor("#686868"),
	Color9:          parseColor("#F66151"),
	Color10:         parseColor("#33DA7A"),
	Color11:         parseColor("#E9AD0C"),
	Color15:         parseColor("#FFF"),
	Color12:         parseColor("#2A7BDE"),
	Color13:         parseColor("#C061CB"),
	Color14:         parseColor("#33C7DE"),
	Shell:           ".",
}

func Load() *Config {
	conf, err := parseFile()
	if err != nil {
		conf = &defaultConfig
	}

	return conf
}

func parseColor(str string) (c color.RGBA) {
	c, _ = utils.ParseColor(str)

	return
}
