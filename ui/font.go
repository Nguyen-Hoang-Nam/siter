package ui

import (
	"bytes"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/go-text/typesetting/font"
)

type FontStyle struct {
	Bold      bool
	Italic    bool
	Monospace bool
	Symbol    bool
}

func CachedFontFace(style FontStyle) *FontCacheItem {
	val, ok := fontCache.Load(style)
	if !ok {
		var f1, f2 font.Face
		switch {
		case style.Monospace:
			f1 = loadMeasureFont(theme.TextMonospaceFont())
			f2 = loadMeasureFont(theme.DefaultTextMonospaceFont())
		case style.Bold:
			if style.Italic {
				f1 = loadMeasureFont(theme.TextBoldItalicFont())
				f2 = loadMeasureFont(theme.DefaultTextBoldItalicFont())
			} else {
				f1 = loadMeasureFont(theme.TextBoldFont())
				f2 = loadMeasureFont(theme.DefaultTextBoldFont())
			}
		case style.Italic:
			f1 = loadMeasureFont(theme.TextItalicFont())
			f2 = loadMeasureFont(theme.DefaultTextItalicFont())
		case style.Symbol:
			f1 = loadMeasureFont(theme.SymbolFont())
			f2 = loadMeasureFont(theme.DefaultSymbolFont())
		default:
			f1 = loadMeasureFont(theme.TextFont())
			f2 = loadMeasureFont(theme.DefaultTextFont())
		}

		if f1 == nil {
			f1 = f2
		}

		faces := []font.Face{f1, f2}
		if emoji := theme.DefaultEmojiFont(); emoji != nil {
			faces = append(faces, loadMeasureFont(emoji))
		}
		val = &FontCacheItem{Fonts: faces}
		fontCache.Store(style, val)
	}

	return val.(*FontCacheItem)
}

func loadMeasureFont(data fyne.Resource) font.Face {
	loaded, err := font.ParseTTF(bytes.NewReader(data.Content()))
	if err != nil {
		fyne.LogError("font load error", err)
		return nil
	}

	return loaded
}

type FontCacheItem struct {
	Fonts []font.Face
}

var fontCache = &sync.Map{} // map[FontStyle]*FontCacheItem
