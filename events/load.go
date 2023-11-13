package events

import (
	"fmt"
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

func Load(canvas fyne.Canvas, process pty.PTYProcess) {
	cc := controlCharacters()

	canvas.AddShortcut(&fyne.ShortcutCopy{}, func(_ fyne.Shortcut) {
		fmt.Println("1")
		process.Write([]byte{'\x03'})
	})

	canvas.AddShortcut(CtrlD, func(_ fyne.Shortcut) {
		process.Write([]byte{'\x04'})
	})

	canvas.AddShortcut(CtrlG, func(_ fyne.Shortcut) {
		process.Write([]byte{'\x07'})
	})

	canvas.AddShortcut(CtrlH, func(_ fyne.Shortcut) {
		process.Write([]byte{'\x08'})
	})

	canvas.AddShortcut(CtrlI, func(_ fyne.Shortcut) {
		process.Write([]byte{'\x09'})
	})

	canvas.AddShortcut(CtrlJ, func(_ fyne.Shortcut) {
		process.Write([]byte{'\x0a'})
	})

	canvas.AddShortcut(CtrlK, func(_ fyne.Shortcut) {
		process.Write([]byte{'\x0b'})
	})

	canvas.AddShortcut(CtrlL, func(_ fyne.Shortcut) {
		process.Write([]byte{'\x0c'})
	})

	canvas.AddShortcut(CtrlM, func(_ fyne.Shortcut) {
		process.Write([]byte{'\x0d'})
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
