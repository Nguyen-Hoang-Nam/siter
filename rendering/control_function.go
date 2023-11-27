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

func (rd *Rendering) handleLF() {
	newline := make([]ui.RuneCell, 0)
	if len(rd.terminal.Rows) == rd.config.ScrollbackLines {
		rd.terminal.Rows = append(rd.terminal.Rows[1:], ui.TerminalRow{Cells: newline})
	} else {
		rd.terminal.Rows = append(rd.terminal.Rows, ui.TerminalRow{Cells: newline})
		rd.terminal.Cursor.Row++
	}
	rd.terminal.Cursor.Col = 0

	if !rd.isNewLine {
		rd.isNewLine = true
	}
}

func (rd *Rendering) handleBS() {
	if rd.terminal.Cursor.Col > 0 {
		rd.terminal.Rows[rd.terminal.Cursor.Row].Cells = rd.terminal.Rows[rd.terminal.Cursor.Row].Cells[:len(rd.terminal.Rows[rd.terminal.Cursor.Row].Cells)-1]
		rd.terminal.Cursor.Col--
		rd.terminal.Index--

		if !rd.isNewOutput {
			rd.isNewOutput = true
		}
	}
}
