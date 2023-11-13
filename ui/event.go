package ui

import (
	"runtime"
	"siter/pty"

	"fyne.io/fyne/v2"
)

type Event struct {
	process pty.PTYProcess
}

func LoadEvent(canvas fyne.Canvas, process pty.PTYProcess) {
	event := Event{process: process}

	canvas.AddShortcut(&fyne.ShortcutCopy{}, event.onCtrlC)
	canvas.SetOnTypedKey(event.onTypedKey)
	canvas.SetOnTypedRune(event.onTypedRune)
}

func (event *Event) onTypedKey(e *fyne.KeyEvent) {
	if e.Name == fyne.KeyEnter || e.Name == fyne.KeyReturn {
		if runtime.GOOS == "windows" {
			_, _ = event.process.Write([]byte("\r\n"))
		} else {
			_, _ = event.process.Write([]byte{'\r'})
		}
	}

	if e.Name == fyne.KeyBackspace {
		_, _ = event.process.Write([]byte{'\b'})
	}
}

func (event *Event) onTypedRune(r rune) {
	_, _ = event.process.Write([]byte(string(r)))
}

func (event *Event) onCtrlC(_ fyne.Shortcut) {
	_, _ = event.process.Write([]byte("\x03"))
}
