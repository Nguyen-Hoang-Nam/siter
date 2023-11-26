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
	rows        []ui.TerminalRow
	cells       []ui.RuneCell
	rowIndex    int
	terminal    *ui.Terminal
	isNewLine   bool
	isNewOutput bool
	nextStyle   ui.RuneCellStyle
}

func Render(scrollContainer *container.Scroll, terminal *ui.Terminal, process io.Reader, config *config.Config) {
	rendering := &Rendering{
		config:      config,
		termColor:   termcolor.New(config),
		rows:        make([]ui.TerminalRow, 1),
		cells:       make([]ui.RuneCell, 0),
		rowIndex:    0,
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

	rendering.terminal.Rows = rendering.rows

	rendering.rows[0] = ui.TerminalRow{Cells: rendering.cells}

	go func() {
		reader := bufio.NewReader(process)

		for {
			r := read(reader)
			if controlfunction.IsControlCharacter(r) {
				functionName, rs := getControlFunction([]rune{r}, reader)
				rendering.handleControlFunction(functionName, rs)
			} else {
				rendering.cells = append(rendering.cells, ui.RuneCell{
					Rune:  r,
					Style: rendering.nextStyle,
				})

				rendering.rows[rendering.rowIndex] = ui.TerminalRow{Cells: rendering.cells}

				if !rendering.isNewOutput {
					rendering.isNewOutput = true
				}
			}
		}
	}()

	go func() {
		deplay := time.Duration(config.RepaintDeplay) * time.Millisecond
		for {
			time.Sleep(deplay)

			if rendering.isNewOutput {
				rendering.isNewOutput = false

				rendering.terminal.Refresh()

				if rendering.isNewLine {
					rendering.isNewLine = false
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
