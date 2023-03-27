package ui

import (
	"runtime"
	"siter/pty"

	"fyne.io/fyne/v2"
)

type Event struct {
	process pty.IPTY
	canvas  fyne.Canvas
}

func NewEvent(process pty.IPTY, canvas fyne.Canvas) *Event {
	return &Event{
		process: process,
		canvas:  canvas,
	}
}

func (event *Event) Load() {
	event.canvas.AddShortcut(&fyne.ShortcutCopy{}, event.onCtrlC)
	event.canvas.SetOnTypedKey(event.onTypedKey)
	event.canvas.SetOnTypedRune(event.onTypedRune)
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
