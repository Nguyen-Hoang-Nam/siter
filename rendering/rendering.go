package rendering

import (
	"bufio"
	"io"
	"os"
	"siter/config"
	controlfunction "siter/control_function"
	"siter/termcolor"
	"siter/ui"
	"time"

	"fyne.io/fyne/v2/container"
)

type Rendering struct {
	config      *config.Config
	termColor   termcolor.TermColor
	rows        []ui.TextGridRow
	cells       []ui.TextGridCell
	rowIndex    int
	textGrid    *ui.TextGrid
	isNewLine   bool
	isNewOutput bool
	nextStyle   *ui.TextGridStyle
}

func Render(scrollContainer *container.Scroll, textGrid *ui.TextGrid, process io.Reader, config *config.Config) {
	rendering := &Rendering{
		config:      config,
		termColor:   termcolor.New(config),
		rows:        make([]ui.TextGridRow, 1),
		cells:       make([]ui.TextGridCell, 0),
		rowIndex:    0,
		textGrid:    textGrid,
		isNewLine:   false,
		isNewOutput: false,
		nextStyle: &ui.TextGridStyle{
			FGColor:   config.ForegroundColor.RGBA,
			BGColor:   config.BackgroundColor.RGBA,
			Italic:    false,
			Bold:      false,
			Underline: ui.NoUnderline,
		},
	}

	rendering.textGrid.Rows = rendering.rows

	rendering.rows[0] = ui.TextGridRow{Cells: rendering.cells}

	go func() {
		reader := bufio.NewReader(process)

		for {
			r := read(reader)

			if controlfunction.IsControlCharacter(r) {
				functionName, rs := getControlFunction([]rune{r}, reader)
				rendering.handleControlFunction(functionName, rs)
			} else {
				rendering.cells = append(rendering.cells, ui.TextGridCell{
					Rune:  r,
					Style: rendering.nextStyle,
				})

				rendering.rows[rendering.rowIndex] = ui.TextGridRow{Cells: rendering.cells}

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

				rendering.textGrid.Refresh()

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
