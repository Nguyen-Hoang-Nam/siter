package ui

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
	textAreaSpaceSymbol   = '·'
	textAreaTabSymbol     = '→'
	textAreaNewLineSymbol = '↵'
)

var (
	TextGridStyleDefault    *TextGridStyle
	TextGridStyleWhitespace *TextGridStyle
)

type TextGridCell struct {
	Rune  rune
	Style *TextGridStyle
}

type TextGridRow struct {
	Cells []TextGridCell
	Style *TextGridStyle
}

type TextGridStyle struct {
	FGColor, BGColor color.Color
	Bold, Italic     bool
}

type TextGrid struct {
	widget.BaseWidget
	Rows []TextGridRow

	ShowWhitespace bool
}

func (t *TextGrid) MinSize() fyne.Size {
	t.BaseWidget.ExtendBaseWidget(t)
	return t.BaseWidget.MinSize()
}

func (t *TextGrid) Resize(size fyne.Size) {
	t.BaseWidget.Resize(size)
	t.BaseWidget.Refresh()
}

func (t *TextGrid) Hide() {
	t.BaseWidget.Hide()
}

func (t *TextGrid) Move(pos fyne.Position) {
	t.BaseWidget.Move(pos)
}

func (t *TextGrid) Position() fyne.Position {
	return t.BaseWidget.Position()
}

func (t *TextGrid) Refresh() {
	t.BaseWidget.Refresh()
}

func (t *TextGrid) Show() {
	t.BaseWidget.Show()
}

func (t *TextGrid) Size() fyne.Size {
	return t.BaseWidget.Size()
}

func (t *TextGrid) Visible() bool {
	return t.BaseWidget.Visible()
}

func (t *TextGrid) CreateRenderer() fyne.WidgetRenderer {
	t.BaseWidget.ExtendBaseWidget(t)
	render := &textGridRenderer{text: t}
	render.updateCellSize()

	TextGridStyleDefault = &TextGridStyle{}
	TextGridStyleWhitespace = &TextGridStyle{FGColor: theme.DisabledColor()}

	return render
}

func NewTextGrid() *TextGrid {
	grid := &TextGrid{}
	grid.BaseWidget.ExtendBaseWidget(grid)
	return grid
}

type textGridRenderer struct {
	text *TextGrid

	cols, rows int

	cellSize fyne.Size
	objects  []fyne.CanvasObject
	current  fyne.Canvas
}

func (t *textGridRenderer) appendTextCell(str rune) {
	text := canvas.NewText(string(str), theme.ForegroundColor())
	text.TextStyle.Monospace = true

	bg := canvas.NewRectangle(color.Transparent)
	t.objects = append(t.objects, bg, text)
}

func (t *textGridRenderer) setCellRune(str rune, pos int, style, rowStyle *TextGridStyle) {
	if str == 0 {
		str = ' '
	}

	text := t.objects[pos*2+1].(*canvas.Text)
	text.TextSize = theme.TextSize()
	fg := theme.ForegroundColor()
	if style != nil && style.FGColor != nil {
		fg = style.FGColor
	} else if rowStyle != nil && rowStyle.FGColor != nil {
		fg = rowStyle.FGColor
	}
	newStr := string(str)
	if text.Text != newStr || text.Color != fg {
		text.Text = newStr
		text.Color = fg
		text.TextStyle = fyne.TextStyle{
			Bold:   style.Bold,
			Italic: style.Italic,
		}
		t.refresh(text)
	}

	rect := t.objects[pos*2].(*canvas.Rectangle)
	bg := color.Color(color.Transparent)
	if style != nil && style.BGColor != nil {
		bg = style.BGColor
	} else if rowStyle != nil && rowStyle.BGColor != nil {
		bg = rowStyle.BGColor
	}
	if rect.FillColor != bg {
		rect.FillColor = bg
		t.refresh(rect)
	}
}

func (t *textGridRenderer) addCellsIfRequired() {
	cellCount := t.cols * t.rows
	if len(t.objects) == cellCount*2 {
		return
	}
	for i := len(t.objects); i < cellCount*2; i += 2 {
		t.appendTextCell(' ')
	}
}

func (t *textGridRenderer) refreshGrid() {
	line := 1
	x := 0

	for rowIndex, row := range t.text.Rows {
		rowStyle := row.Style
		i := 0
		for _, r := range row.Cells {
			if i >= t.cols {
				continue
			}
			if t.text.ShowWhitespace && (r.Rune == ' ' || r.Rune == '\t') {
				sym := textAreaSpaceSymbol
				if r.Rune == '\t' {
					sym = textAreaTabSymbol
				}

				if r.Style != nil && r.Style.BGColor != nil {
					whitespaceBG := &TextGridStyle{FGColor: TextGridStyleWhitespace.FGColor,
						BGColor: r.Style.BGColor}
					t.setCellRune(sym, x, whitespaceBG, rowStyle)
				} else {
					t.setCellRune(sym, x, TextGridStyleWhitespace, rowStyle)
				}
			} else {
				t.setCellRune(r.Rune, x, r.Style, rowStyle)
			}
			i++
			x++
		}
		if t.text.ShowWhitespace && i < t.cols && rowIndex < len(t.text.Rows)-1 {
			t.setCellRune(textAreaNewLineSymbol, x, TextGridStyleWhitespace, rowStyle)
			i++
			x++
		}
		for ; i < t.cols; i++ {
			t.setCellRune(' ', x, TextGridStyleDefault, rowStyle)
			x++
		}

		line++
	}
	for ; x < len(t.objects)/2; x++ {
		t.setCellRune(' ', x, TextGridStyleDefault, nil) // trailing cells and blank lines
	}
}

func (t *textGridRenderer) updateGridSize(size fyne.Size) {
	bufRows := len(t.text.Rows)
	bufCols := 0
	for _, row := range t.text.Rows {
		bufCols = int(math.Max(float64(bufCols), float64(len(row.Cells))))
	}
	sizeCols := math.Floor(float64(size.Width) / float64(t.cellSize.Width))
	sizeRows := math.Floor(float64(size.Height) / float64(t.cellSize.Height))

	if t.text.ShowWhitespace {
		bufCols++
	}

	t.cols = int(math.Max(sizeCols, float64(bufCols)))
	t.rows = int(math.Max(sizeRows, float64(bufRows)))
	t.addCellsIfRequired()
}

func (t *textGridRenderer) Layout(size fyne.Size) {
	t.updateGridSize(size)

	i := 0
	cellPos := fyne.NewPos(0, 0)
	for y := 0; y < t.rows; y++ {
		for x := 0; x < t.cols; x++ {
			t.objects[i*2+1].Move(cellPos)

			t.objects[i*2].Resize(t.cellSize)
			t.objects[i*2].Move(cellPos)
			cellPos.X += t.cellSize.Width
			i++
		}

		cellPos.X = 0
		cellPos.Y += t.cellSize.Height
	}
}

func (t *textGridRenderer) MinSize() fyne.Size {
	longestRow := float32(0)
	for _, row := range t.text.Rows {
		longestRow = fyne.Max(longestRow, float32(len(row.Cells)))
	}
	return fyne.NewSize(t.cellSize.Width*longestRow,
		t.cellSize.Height*float32(len(t.text.Rows)))
}

func (t *textGridRenderer) Refresh() {
	if fyne.CurrentApp() != nil && fyne.CurrentApp().Driver() != nil {
		t.current = fyne.CurrentApp().Driver().CanvasForObject(t.text)
	}

	t.updateCellSize()

	TextGridStyleWhitespace = &TextGridStyle{FGColor: theme.DisabledColor()}
	t.updateGridSize(t.text.BaseWidget.Size())
	t.refreshGrid()
}

func (t *textGridRenderer) ApplyTheme() {
}

func (t *textGridRenderer) Objects() []fyne.CanvasObject {
	return t.objects
}

func (t *textGridRenderer) Destroy() {
}

func (t *textGridRenderer) refresh(obj fyne.CanvasObject) {
	if t.current == nil {
		if fyne.CurrentApp() != nil && fyne.CurrentApp().Driver() != nil {
			t.current = fyne.CurrentApp().Driver().CanvasForObject(t.text)
		}

		if t.current == nil {
			return
		}
	}

	t.current.Refresh(obj)
}

func (t *textGridRenderer) updateCellSize() {
	size := fyne.MeasureText("M", theme.TextSize(), fyne.TextStyle{Monospace: true})

	size.Width = float32(math.Round(float64((size.Width))))
	size.Height = float32(math.Round(float64((size.Height))))

	t.cellSize = size
}
