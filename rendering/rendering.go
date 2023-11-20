package rendering

import (
	"bufio"
	"image/color"
	"io"
	"os"
	"siter/config"
	termcolor "siter/term_color"
	"siter/utils"
	"time"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Rendering struct {
	termColor   termcolor.TermColor
	rows        []widget.TextGridRow
	cells       []widget.TextGridCell
	rowIndex    int
	textGrid    *widget.TextGrid
	isNewLine   bool
	nextFGColor color.RGBA
	nextBGColor color.RGBA
}

func Render(scrollContainer *container.Scroll, textGrid *widget.TextGrid, process io.Reader, config *config.Config) {
	rendering := &Rendering{
		termColor:   termcolor.New(config),
		rows:        make([]widget.TextGridRow, 1),
		cells:       make([]widget.TextGridCell, 0),
		rowIndex:    0,
		textGrid:    textGrid,
		isNewLine:   false,
		nextFGColor: config.ForegroundColor.RGBA,
		nextBGColor: config.BackgroundColor.RGBA,
	}

	rendering.textGrid.Rows = rendering.rows

	rendering.rows[0] = widget.TextGridRow{Cells: rendering.cells}

	isNewOutput := false

	go func() {
		reader := bufio.NewReader(process)

		for {
			r := read(reader)

			if utils.IsControlCharacter(r) {
				functionName, rs := getControlFunction([]rune{r}, reader)
				rendering.handleControlFunction(functionName, rs)
			} else {
				rendering.cells = append(rendering.cells, widget.TextGridCell{
					Rune: r,
					Style: &widget.CustomTextGridStyle{
						FGColor: rendering.nextFGColor,
						BGColor: rendering.nextBGColor,
					},
				})

				rendering.rows[rendering.rowIndex] = widget.TextGridRow{Cells: rendering.cells}

				if !isNewOutput {
					isNewOutput = true
				}
			}
		}
	}()

	go func() {
		deplay := time.Duration(config.RepaintDeplay) * time.Millisecond
		for {
			time.Sleep(deplay)

			if isNewOutput {
				isNewOutput = false

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

func getControlFunction(rs []rune, reader *bufio.Reader) (string, []rune) {
	functionName, isEnd := utils.ControlFunctionOneCharacter(rs[0])
	if isEnd {
		return functionName, rs
	}

	rs = append(rs, read(reader))
	functionName, isEnd = utils.ControlFunctionEsc7Bit(rs)
	if isEnd {
		return functionName, rs
	}

	for {
		rs = append(rs, read(reader))
		functionName, isEnd = utils.ControlFunctionEscSequence(rs)
		if isEnd {
			return functionName, rs
		}
	}
}
