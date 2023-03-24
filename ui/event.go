package ui

import (
	"siter/utils"

	"fyne.io/fyne/v2"
)

type Event struct {
	process utils.IPTY
	canvas  fyne.Canvas
}

func NewEvent(process utils.IPTY, canvas fyne.Canvas) *Event {
	return &Event{
		process: process,
		canvas:  canvas,
	}
}

func (event *Event) Load() {
	event.canvas.SetOnTypedKey(event.onTypedKey)
	event.canvas.SetOnTypedRune(event.onTypedRune)
}

func (event *Event) onTypedKey(e *fyne.KeyEvent) {
	if e.Name == fyne.KeyEnter || e.Name == fyne.KeyReturn {
		_, _ = event.process.Write([]byte{'\r'})
	}

	if e.Name == fyne.KeyBackspace {
		_, _ = event.process.Write([]byte{'\b'})
	}
}

func (event *Event) onTypedRune(r rune) {
	_, _ = event.process.Write([]byte(string(r)))
}
