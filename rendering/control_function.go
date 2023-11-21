package rendering

import (
	controlfunction "siter/control_function"
	"siter/ui"
)

func (r *Rendering) handleControlFunction(functionName controlfunction.FunctionName, rs []rune) {
	switch functionName {
	case controlfunction.LF:
		r.handleLF()
	case controlfunction.CR:
		r.handleLF()
	case controlfunction.BS:
		r.handleBS()
	case controlfunction.SGR:
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
