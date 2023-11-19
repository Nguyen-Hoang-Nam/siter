package rendering

import (
	"bufio"
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
	termColor termcolor.TermColor
	rows      []widget.TextGridRow
	cells     []widget.TextGridCell
	rowIndex  int
	textGrid  *widget.TextGrid
	isNewLine bool
	nextStyle widget.CustomTextGridStyle
}

func Render(scrollContainer *container.Scroll, textGrid *widget.TextGrid, process io.Reader, config *config.Config) {
	rendering := &Rendering{
		termColor: termcolor.New(config),
		rows:      make([]widget.TextGridRow, 1),
		cells:     make([]widget.TextGridCell, 0),
		rowIndex:  0,
		textGrid:  textGrid,
		isNewLine: false,
		nextStyle: widget.CustomTextGridStyle{FGColor: config.ForegroundColor.RGBA},
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
				rendering.cells = append(rendering.cells, widget.TextGridCell{Rune: r, Style: &rendering.nextStyle})

				rendering.rows[rendering.rowIndex] = widget.TextGridRow{Cells: rendering.cells}
				rendering.textGrid.SetStyle(rendering.rowIndex, len(rendering.cells)-1, &rendering.nextStyle)

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

// func (r *Rendering) style(text string) {
// 	row := 0
// 	col := 0
// 	isNewColor := false
// 	asciiColor := []byte("")
// 	currentColor := r.termianlColor["[0"]

// 	for i := range text {
// 		if text[i] == 27 {
// 			isNewColor = true
// 		} else {
// 			if isNewColor && text[i] == 'm' {
// 				currentColor = r.termianlColor[string(asciiColor)]
// 				asciiColor = []byte("")

// 				isNewColor = false
// 			} else if isNewColor {
// 				asciiColor = append(asciiColor, text[i])
// 			}
// 		}

// 		r.textGrid.SetStyle(row, col, &widget.CustomTextGridStyle{FGColor: currentColor})

// 		if !isNewColor {
// 			if text[i] == '\n' {
// 				row++
// 				col = 0
// 			} else {
// 				col++
// 			}
// 		}

// 	}
// }
