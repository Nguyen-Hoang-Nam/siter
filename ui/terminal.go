package ui

import (
	"fmt"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type TerminalRow struct {
	Cells []RuneCell
}

type Cursor struct {
	Row int
	Col int
}

type Terminal struct {
	widget.BaseWidget
	Rows   []TerminalRow
	Cursor Cursor
	Index  int
}

func (t *Terminal) MinSize() fyne.Size {
	t.ExtendBaseWidget(t)
	return t.BaseWidget.MinSize()
}

func (t *Terminal) Resize(size fyne.Size) {
	t.BaseWidget.Resize(size)
	t.Refresh()
}

func (t *Terminal) CreateRenderer() fyne.WidgetRenderer {
	t.ExtendBaseWidget(t)

	cellSize := fyne.MeasureText("M", theme.TextSize(), fyne.TextStyle{Monospace: true})

	cellSize.Width = float32(math.Round(float64((cellSize.Width))))
	cellSize.Height = float32(math.Round(float64((cellSize.Height))))

	return &terminalRenderer{
		terminal:         t,
		lastCursor:       Cursor{Row: 0, Col: 0},
		lastCellPos:      fyne.Position{X: 0, Y: 0},
		lastRenderCursor: Cursor{Row: 0, Col: 0},
		cellSize:         cellSize,
	}
}

func NewTerminal() *Terminal {
	terminal := &Terminal{
		Cursor: Cursor{Row: 0, Col: 0},
	}
	terminal.ExtendBaseWidget(terminal)
	return terminal
}

type terminalRenderer struct {
	terminal *Terminal

	cols, rows int

	cellSize         fyne.Size
	current          fyne.Canvas
	objects          []fyne.CanvasObject
	lastCursor       Cursor
	lastCellPos      fyne.Position
	lastRenderCursor Cursor
}

func (t *terminalRenderer) Layout(size fyne.Size) {
	cols := t.cols
	t.cols = int(math.Floor(float64(size.Width) / float64(t.cellSize.Width)))
	if cols != 0 && cols != t.cols {
		t.resize()
	}

	t.rows = int(math.Floor(float64(size.Height) / float64(t.cellSize.Height)))
}

func (t *terminalRenderer) MinSize() fyne.Size {
	return fyne.NewSize(t.cellSize.Width*float32(t.cols),
		t.cellSize.Height*float32(t.lastRenderCursor.Row+1))
}

func (t *terminalRenderer) Refresh() {
	// if t.lastCursor.Col > t.terminal.Cursor.Col || t.lastCursor.Row > t.terminal.Cursor.Row {
	// 	if t.lastCursor.Row == t.terminal.Cursor.Row && t.lastCursor.Col >= len(t.terminal.Rows[t.lastCursor.Row].Cells) {
	// 		diff := t.lastCursor.Col - len(t.terminal.Rows[t.lastCursor.Row].Cells)
	// 		t.objects = t.objects[0 : len(t.objects)-diff]
	// 	}
	// 	t.lastCursor.Col = t.terminal.Cursor.Col
	// 	t.lastCursor.Row = t.terminal.Cursor.Row
	// }

	cellPos := t.lastCellPos
	rowIndex := t.lastCursor.Row
	colIndex := t.lastCursor.Col
	renderRowIndex := t.lastRenderCursor.Row
	renderColIndex := t.lastRenderCursor.Col
	for ; rowIndex < len(t.terminal.Rows); rowIndex++ {
		row := t.terminal.Rows[rowIndex]
		if rowIndex != int(t.lastCursor.Row) {
			cellPos.X = 0
			cellPos.Y += t.cellSize.Height
			renderRowIndex++
			renderColIndex = 0
		}

		for ; colIndex < len(row.Cells); colIndex++ {
			cell := row.Cells[colIndex]
			if renderColIndex >= t.cols {
				cellPos.X = 0
				cellPos.Y += t.cellSize.Height
				renderRowIndex++
				renderColIndex = 0
			}

			runeCell := RuneCell{
				Rune:   cell.Rune,
				Width:  int(t.cellSize.Width),
				Height: int(t.cellSize.Height),
				Style:  cell.Style,
			}
			if t.terminal.Index < len(t.objects) {
				object := t.objects[t.terminal.Index].(*canvas.Raster)
				fmt.Println(object)
				fmt.Println(runeCell)
				object.Generator = runeCell.Generate()
				t.refresh(object)
			} else {
				object := canvas.Raster{Generator: runeCell.Generate()}
				object.Move(cellPos)
				object.Resize(t.cellSize)
				t.refresh(&object)
				t.objects = append(t.objects, &object)
			}

			cellPos.X += t.cellSize.Width
			renderColIndex++
		}
		colIndex = 0
	}

	x := len(t.terminal.Rows[len(t.terminal.Rows)-1].Cells)

	t.lastCursor = Cursor{Col: x, Row: len(t.terminal.Rows) - 1}
	t.lastCellPos = cellPos
	t.lastRenderCursor = Cursor{Col: renderColIndex, Row: renderRowIndex}
}

func (t *terminalRenderer) Objects() []fyne.CanvasObject {
	return t.objects
}

func (t *terminalRenderer) Destroy() {}

func (t *terminalRenderer) refresh(obj fyne.CanvasObject) {
	if t.current == nil {
		if fyne.CurrentApp() != nil && fyne.CurrentApp().Driver() != nil {
			t.current = fyne.CurrentApp().Driver().CanvasForObject(t.terminal)
		}

		if t.current == nil {
			return
		}
	}

	t.current.Refresh(obj)
}

func (t *terminalRenderer) resize() {
	cellPos := fyne.Position{X: 0, Y: 0}
	rowIndex := 0
	colIndex := 0
	renderRowIndex := 0
	renderColIndex := 0
	t.objects = []fyne.CanvasObject{}

	for ; rowIndex < len(t.terminal.Rows); rowIndex++ {
		row := t.terminal.Rows[rowIndex]
		if rowIndex != 0 {
			cellPos.X = 0
			cellPos.Y += t.cellSize.Height
			renderRowIndex++
		}

		for ; colIndex < len(row.Cells); colIndex++ {
			cell := row.Cells[colIndex]
			if renderColIndex >= t.cols {
				cellPos.X = 0
				cellPos.Y += t.cellSize.Height
				renderRowIndex++
				renderColIndex = 0
			}

			runeCell := RuneCell{
				Rune:   cell.Rune,
				Width:  int(t.cellSize.Width),
				Height: int(t.cellSize.Height),
				Style:  cell.Style,
			}
			object := canvas.Raster{Generator: runeCell.Generate()}
			object.Move(cellPos)
			object.Resize(t.cellSize)
			t.refresh(&object)
			t.objects = append(t.objects, &object)

			cellPos.X += t.cellSize.Width
			renderColIndex++
		}
	}

	x := len(t.terminal.Rows[len(t.terminal.Rows)-1].Cells)

	t.lastCursor = Cursor{Col: x, Row: len(t.terminal.Rows) - 1}
	t.lastCellPos = cellPos
	t.lastRenderCursor = Cursor{Col: renderColIndex, Row: renderRowIndex}
}
