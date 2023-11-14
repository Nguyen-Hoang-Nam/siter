package config

import "image/color"

func defaultConfig() Config {
	var defaultShell = parsingShell{}

	defaultShell.UnmarshalText([]byte{'.'})

	return Config{
		FontFamily:        "monospace",
		BoldFont:          "monospace",
		ItalicFont:        "monospace",
		BoldItalicFont:    "monospace",
		FontSize:          16.0,
		CursorColor:       parsingColor{color.RGBA{R: 103, G: 150, B: 230, A: 255}},
		CursorTextColor:   parsingColor{color.RGBA{R: 17, G: 17, B: 17, A: 255}},
		CursorShape:       CURSOR_SHAPE_BLOCK,
		ScrollbackLines:   2000,
		UrlColor:          parsingColor{color.RGBA{R: 0, G: 135, B: 189, A: 255}},
		UrlStyle:          URL_STYLE_CURLY,
		OpenUrlWith:       "default",
		UrlPrefixes:       []string{"http", "https", "ssh"},
		DetectUrls:        true,
		RepaintDeplay:     100,
		ForegroundColor:   parsingColor{color.RGBA{R: 241, G: 241, B: 240, A: 255}},
		BackgroundColor:   parsingColor{color.RGBA{R: 31, G: 33, B: 42, A: 255}},
		BackgroundOpacity: 1,
		BackgroundImage:   "",
		Color0:            parsingColor{color.RGBA{R: 40, G: 42, B: 54, A: 255}},
		Color1:            parsingColor{color.RGBA{R: 255, G: 92, B: 87, A: 255}},
		Color2:            parsingColor{color.RGBA{R: 90, G: 247, B: 142, A: 255}},
		Color3:            parsingColor{color.RGBA{R: 243, G: 249, B: 157, A: 255}},
		Color4:            parsingColor{color.RGBA{R: 87, G: 199, B: 255, A: 255}},
		Color5:            parsingColor{color.RGBA{R: 255, G: 105, B: 193, A: 255}},
		Color6:            parsingColor{color.RGBA{R: 154, G: 237, B: 254, A: 255}},
		Color7:            parsingColor{color.RGBA{R: 241, G: 241, B: 240, A: 255}},
		Color8:            parsingColor{color.RGBA{R: 104, G: 104, B: 104, A: 255}},
		Color9:            parsingColor{color.RGBA{R: 246, G: 97, B: 81, A: 255}},
		Color10:           parsingColor{color.RGBA{R: 51, G: 218, B: 122, A: 255}},
		Color11:           parsingColor{color.RGBA{R: 233, G: 173, B: 12, A: 255}},
		Color15:           parsingColor{color.RGBA{R: 255, G: 255, B: 255, A: 255}},
		Color12:           parsingColor{color.RGBA{R: 42, G: 123, B: 222, A: 255}},
		Color13:           parsingColor{color.RGBA{R: 192, G: 97, B: 203, A: 255}},
		Color14:           parsingColor{color.RGBA{R: 51, G: 199, B: 222, A: 255}},
		Shell:             defaultShell,
		Map: map[string]string{
			"ctrl+c": "write \x03",
			"ctrl+d": "write \x04",
			"ctrl+g": "write \x07",
			"ctrl+h": "write \x08",
			"ctrl+i": "write \x09",
			"ctrl+j": "write \x0a",
			"ctrl+k": "write \x0b",
			"ctrl+l": "write \x0c",
			"ctrl+m": "write \x0d",
		},
	}
}
