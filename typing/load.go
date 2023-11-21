package typing

import (
	"runtime"
	"siter/config"
	"siter/pty"

	"fyne.io/fyne/v2"
)

type mappingStruct struct {
	process pty.PTYProcess
	canvas  fyne.Canvas
	conf    *config.Config
}

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
		fyne.KeyF1:        {'\x1b', 'O', 'P'},
		fyne.KeyF2:        {'\x1b', 'O', 'Q'},
		fyne.KeyF3:        {'\x1b', 'O', 'R'},
		fyne.KeyF4:        {'\x1b', 'O', 'S'},
		fyne.KeyF5:        {'\x1b', '[', '1', '5', '~'},
		fyne.KeyF6:        {'\x1b', '[', '1', '7', '~'},
		fyne.KeyF7:        {'\x1b', '[', '1', '8', '~'},
		fyne.KeyF8:        {'\x1b', '[', '1', '9', '~'},
		fyne.KeyF9:        {'\x1b', '[', '2', '0', '~'},
		fyne.KeyF10:       {'\x1b', '[', '2', '1', '~'},
		fyne.KeyF11:       {'\x1b', '[', '2', '3', '~'},
		fyne.KeyF12:       {'\x1b', '[', '2', '4', '~'},
	}
}

func Load(canvas fyne.Canvas, process pty.PTYProcess, conf *config.Config) {
	m := mappingStruct{process, canvas, conf}
	m.loadShortcut()

	cc := controlCharacters()
	canvas.SetOnTypedKey(func(e *fyne.KeyEvent) {
		if text, ok := cc[e.Name]; ok {
			process.Write(text)
		}
	})

	canvas.SetOnTypedRune(func(r rune) {
		process.Write([]byte(string(r)))
	})
}
