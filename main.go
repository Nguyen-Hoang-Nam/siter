package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	siter "siter/src"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/creack/pty"
)

const MaxBufferSize = 16

func main() {
	app := app.New()
	window := app.NewWindow("Siter")

	ui := widget.NewTextGrid()

	os.Setenv("TERM", "dumb")
	start_command := exec.Command("/bin/bash")
	process, err := pty.Start(start_command)
	if err != nil {
		fyne.LogError("Failed to open pty", err)
		os.Exit(1)
	}

	defer start_command.Process.Kill()

	onTypedKey := func(e *fyne.KeyEvent) {
		if e.Name == fyne.KeyEnter || e.Name == fyne.KeyReturn {
			_, _ = process.Write([]byte{'\r'})
		}

		if e.Name == fyne.KeyBackspace {
			_, _ = process.Write([]byte{'\b'})
		}

		fmt.Println(e.Name)
	}

	onTypedRune := func(r rune) {
		_, _ = process.WriteString(string(r))
	}

	window.Canvas().SetOnTypedKey(onTypedKey)
	window.Canvas().SetOnTypedRune(onTypedRune)

	buffer := [][]rune{}
	reader := bufio.NewReader(process)

	go func() {
		line := []rune{}
		buffer = append(buffer, line)
		for {
			r, _, err := reader.ReadRune()

			if err != nil {
				if err == io.EOF {
					return
				}
				os.Exit(0)
			}

			line = append(line, r)
			buffer[len(buffer)-1] = line
			if r == '\n' {
				if len(buffer) > MaxBufferSize {
					buffer = buffer[1:]
				}

				line = []rune{}
				buffer = append(buffer, line)
			}
		}
	}()

	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			ui.SetText("")

			var lines string
			for _, line := range buffer {
				lines = lines + string(line)
			}

			ui.SetText(siter.Clear_backspace(string(lines)))
		}
	}()

	window.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewGridWrapLayout(fyne.NewSize(900, 325)),
			ui,
		),
	)
	window.ShowAndRun()
}
