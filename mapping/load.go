package mapping

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
