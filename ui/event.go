package ui

import (
	"runtime"
	"siter/pty"

	"fyne.io/fyne/v2"
)

func controlCharacters() map[fyne.KeyName][]byte {
	var newline []byte
	if runtime.GOOS == "windows" {
		newline = []byte("\r\n")
	} else {
		newline = []byte{'\r'}
	}

	return map[fyne.KeyName][]byte{
		fyne.KeyBackspace: {'\x08'},
		fyne.KeyTab:       {'\x09'},
		fyne.KeyEnter:     newline,
		fyne.KeyReturn:    newline,
		fyne.KeyEscape:    {'\x1b'},
	}
}

func LoadEvent(canvas fyne.Canvas, process pty.PTYProcess) {
	cc := controlCharacters()

	canvas.AddShortcut(&fyne.ShortcutCopy{}, func(_ fyne.Shortcut) {
		process.Write([]byte{'\x03'})
	})

	canvas.SetOnTypedKey(func(e *fyne.KeyEvent) {
		if text, ok := cc[e.Name]; ok {
			process.Write(text)
		}
	})

	canvas.SetOnTypedRune(func(r rune) {
		process.Write([]byte(string(r)))
	})
}
