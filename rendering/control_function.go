package rendering

import (
	"siter/ui"
)

func (r *Rendering) handleControlFunction(functionName string, rs []rune) {
	switch functionName {
	case "LF":
		r.handleLF()
	case "CR":
		r.handleLF()
	case "BS":
		r.handleBS()
	case "SGR":
		r.handleSGR(rs)
	}
}

func (r *Rendering) handleLF() {
	r.cells = append(r.cells, ui.TextGridCell{
		Rune:  '\n',
		Style: r.nextStyle,
	})
	r.cells = make([]ui.TextGridCell, 0)
	r.rows = append(r.rows, ui.TextGridRow{Cells: r.cells})
	r.textGrid.Rows = r.rows
	r.rowIndex++

	if !r.isNewLine {
		r.isNewLine = true
	}
}

func (r *Rendering) handleBS() {
	r.cells = r.cells[:len(r.cells)-1]
	r.rows[r.rowIndex] = ui.TextGridRow{Cells: r.cells}

	if !r.isNewOutput {
		r.isNewOutput = true
	}
}
