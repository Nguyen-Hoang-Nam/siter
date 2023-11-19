package ui

import (
	"image/color"
	"siter/config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type Theme struct {
	config *config.Config
}

func NewTheme(c *config.Config) *Theme {
	return &Theme{c}
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
	return theme.DefaultTheme().Font(style)
}

func (t *Theme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
