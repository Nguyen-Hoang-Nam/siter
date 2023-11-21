package ui

import (
	"image/color"
	"os"
	"regexp"
	"runtime"
	"siter/config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type Theme struct {
	config                                               *config.Config
	regular, bold, italic, boldItalic, monospace, symbol fyne.Resource
}

func NewTheme(c *config.Config) *Theme {
	regular := theme.DefaultTheme().Font(fyne.TextStyle{Monospace: true})
	monospace := theme.DefaultTheme().Font(fyne.TextStyle{Monospace: true})
	bold := theme.DefaultTheme().Font(fyne.TextStyle{Bold: true})
	boldItalic := theme.DefaultTheme().Font(fyne.TextStyle{Bold: true, Italic: true})
	italic := theme.DefaultTheme().Font(fyne.TextStyle{Italic: true})
	symbol := theme.DefaultTheme().Font(fyne.TextStyle{Symbol: true})

	regular = loadCustomFont(c.FontFamily, regular)
	monospace = loadCustomFont(c.FontFamily, monospace)
	bold = loadCustomFont(c.BoldFont, bold)
	boldItalic = loadCustomFont(c.BoldItalicFont, boldItalic)
	italic = loadCustomFont(c.ItalicFont, italic)
	symbol = loadCustomFont(c.FontFamily, symbol)

	return &Theme{
		config:     c,
		regular:    regular,
		bold:       bold,
		italic:     italic,
		boldItalic: boldItalic,
		monospace:  monospace,
		symbol:     symbol,
	}
}

func loadCustomFont(fontPath string, fallback fyne.Resource) fyne.Resource {
	if fontPath == "" {
		return fallback
	}

	res, err := fyne.LoadResourceFromPath(expandEnv(fontPath))
	if err != nil {
		fyne.LogError("Error loading specified font", err)
		return fallback
	}

	return res
}

var re = regexp.MustCompile(`%(.+)%`)

func expandEnv(text string) string {
	goos := runtime.GOOS
	if goos == "windows" {
		text = os.ExpandEnv(re.ReplaceAllString(text, `${${1}}`))
	} else if goos == "linux" || goos == "darwin" {
		text = os.ExpandEnv(text)
	}

	return text
}

func (t *Theme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		return t.config.BackgroundColor.RGBA
	}

	if name == theme.ColorNameForeground {
		return t.config.ForegroundColor.RGBA
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (t *Theme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *Theme) Font(style fyne.TextStyle) fyne.Resource {
	if style.Monospace {
		return t.monospace
	}
	if style.Bold {
		if style.Italic {
			return t.boldItalic
		}
		return t.bold
	}
	if style.Italic {
		return t.italic
	}
	if style.Symbol {
		return t.symbol
	}
	return t.regular
}

func (t *Theme) Size(name fyne.ThemeSizeName) float32 {
	return float32(t.config.FontSize)
}
