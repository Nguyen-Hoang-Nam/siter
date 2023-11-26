package rendering

import (
	controlfunction "siter/internal/control_function"
	"siter/ui"
)

func (r *Rendering) handleControlFunction(functionName controlfunction.FunctionName, rs []rune) {
	switch functionName {
	case controlfunction.LF, controlfunction.VT, controlfunction.FF:
		r.handleLF()
	case controlfunction.BS:
		r.handleBS()
	case controlfunction.SGR:
		r.handleSGR(rs)
	}
}

func (r *Rendering) handleLF() {
	r.cells = append(r.cells, ui.RuneCell{
		Rune:  '\n',
		Style: r.nextStyle,
	})
	r.cells = make([]ui.RuneCell, 0)
	if len(r.rows) == r.config.ScrollbackLines {
		r.rows = append(r.rows[1:], ui.TerminalRow{Cells: r.cells})
	} else {
		r.rows = append(r.rows, ui.TerminalRow{Cells: r.cells})
		r.rowIndex++
	}
	r.terminal.Rows = r.rows

	if !r.isNewLine {
		r.isNewLine = true
	}
}

func (r *Rendering) handleBS() {
	r.cells = r.cells[:len(r.cells)-1]
	r.rows[r.rowIndex] = ui.TerminalRow{Cells: r.cells}

	if !r.isNewOutput {
		r.isNewOutput = true
	}
}
