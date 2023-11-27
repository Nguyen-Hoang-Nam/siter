package rendering

import (
	"bufio"
	"io"
	"os"
	"siter/config"
	controlfunction "siter/internal/control_function"
	"siter/internal/termcolor"
	"siter/ui"
	"time"

	"fyne.io/fyne/v2/container"
)

type Rendering struct {
	config      *config.Config
	termColor   termcolor.TermColor
	terminal    *ui.Terminal
	isNewLine   bool
	isNewOutput bool
	nextStyle   ui.RuneCellStyle
}

func Render(scrollContainer *container.Scroll, terminal *ui.Terminal, process io.Reader, config *config.Config) {
	rd := &Rendering{
		config:      config,
		termColor:   termcolor.New(config),
		terminal:    terminal,
		isNewLine:   false,
		isNewOutput: false,
		nextStyle: ui.RuneCellStyle{
			ForegroundColor: config.ForegroundColor.RGBA,
			BackgroundColor: config.BackgroundColor.RGBA,
			FontStyle: ui.FontStyle{
				Italic: false,
				Bold:   false,
			},
			FontSize:  float32(config.FontSize),
			Underline: ui.NoUnderline,
			Overline:  false,
			Strike:    false,
		},
	}

	rd.terminal.Rows = []ui.TerminalRow{
		{Cells: make([]ui.RuneCell, 0)},
	}

	go func() {
		reader := bufio.NewReader(process)

		for {
			r := read(reader)
			if controlfunction.IsControlCharacter(r) {
				functionName, rs := getControlFunction([]rune{r}, reader)
				rd.handleControlFunction(functionName, rs)
			} else {
				rd.terminal.Rows[rd.terminal.Cursor.Row].Cells = append(rd.terminal.Rows[rd.terminal.Cursor.Row].Cells, ui.RuneCell{
					Rune:  r,
					Style: rd.nextStyle,
				})

				rd.terminal.Cursor.Col++
				rd.terminal.Index++

				if !rd.isNewOutput {
					rd.isNewOutput = true
				}
			}
		}
	}()

	go func() {
		deplay := time.Duration(config.RepaintDeplay) * time.Millisecond
		for {
			time.Sleep(deplay)

			if rd.isNewOutput {
				rd.isNewOutput = false

				rd.terminal.Refresh()

				if rd.isNewLine {
					rd.isNewLine = false
					scrollContainer.ScrollToBottom()
				}
			}
		}
	}()
}

func read(reader *bufio.Reader) rune {
	r, _, err := reader.ReadRune()
	if err != nil {
		os.Exit(0)
	}

	return r
}

func getControlFunction(rs []rune, reader *bufio.Reader) (controlfunction.FunctionName, []rune) {
	functionName, isEnd := controlfunction.ControlCharacter(rs[0])
	if isEnd {
		return functionName, rs
	}

	for {
		rs = append(rs, read(reader))
		functionName, isEnd = controlfunction.ControlSequence(rs)
		if isEnd {
			return functionName, rs
		}
	}
}
