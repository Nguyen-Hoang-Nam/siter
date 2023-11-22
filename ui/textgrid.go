package ui

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
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
}

type UnderlineStyle int

const (
	NoUnderline UnderlineStyle = iota
	StraightUnderline
	DoubleUnderline
	CurlyUnderline
	DottedUnderline
	DashedUnderline
)

type TextGridStyle struct {
	FGColor, BGColor color.Color
	Bold, Italic     bool
	Underline        UnderlineStyle
	UnderLineColor   color.Color
}

type TextGrid struct {
	widget.BaseWidget
	Rows []TextGridRow
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
	render := &textGridRenderer{text: t, lastCursor: fyne.Position{X: 0, Y: 0}}
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

	cellSize         fyne.Size
	objects          []fyne.CanvasObject
	current          fyne.Canvas
	currentTextStyle TextStyle
	lastCursor       fyne.Position
}

func (t *textGridRenderer) appendTextCell(str rune) {
	text := canvas.NewText(string(str), theme.ForegroundColor())
	text.TextStyle.Monospace = true

	bg := canvas.NewRectangle(color.Transparent)
	style := t.currentTextStyle
	ras := style.GenerateObject(int(bg.Size().Width), int(bg.Size().Height))
	t.objects = append(t.objects, bg, text, ras)
}

func (t *textGridRenderer) setCellRune(str rune, pos int, style *TextGridStyle) {
	ulWidth := 1
	var ulColor color.Color = color.Transparent

	text := t.objects[pos*3+1].(*canvas.Text)
	text.TextSize = theme.TextSize()
	fg := theme.ForegroundColor()
	if style != nil && style.FGColor != nil {
		fg = style.FGColor
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

	rect := t.objects[pos*3].(*canvas.Rectangle)
	bg := color.Color(color.Transparent)
	if style != nil && style.BGColor != nil {
		bg = style.BGColor
	}
	if rect.FillColor != bg {
		rect.FillColor = bg
		t.refresh(rect)
	}

	if style != nil && style.Underline != NoUnderline {
		if style.UnderLineColor != color.Transparent {
			ulColor = style.UnderLineColor
		} else {
			ulColor = fg
		}

		if style.Bold {
			ulWidth = 2
		}

		ras := t.objects[pos*3+2].(*canvas.Raster)
		textStyle := TextStyle{Underline: style.Underline, UnderlineWidth: ulWidth, UnderlineColor: ulColor}
		rasUpdate := textStyle.GenerateObject(int(rect.Size().Width), int(rect.Size().Height))
		rasUpdate.Move(ras.Position())
		rasUpdate.Resize(t.cellSize)
		t.objects[pos*3+2] = rasUpdate
		t.currentTextStyle = textStyle
		t.refresh(rasUpdate)
	}
}

func (t *textGridRenderer) refreshGrid() {
	line := 1
	x := int(t.lastCursor.X) + int(t.lastCursor.Y)*t.cols

	for j := int(t.lastCursor.Y); j < len(t.text.Rows); j++ {
		row := t.text.Rows[j]
		i := 0
		if j == int(t.lastCursor.Y) {
			i = int(t.lastCursor.X)
		}
		for k := i; k < len(row.Cells); k++ {
			r := row.Cells[k]
			if i >= t.cols {
				continue
			}

			t.setCellRune(r.Rune, x, r.Style)
			i++
			x++
		}

		for ; i < t.cols; i++ {
			t.setCellRune(' ', x, TextGridStyleDefault)
			x++
		}

		line++
	}

	for ; x < len(t.objects)/3; x++ {
		t.setCellRune(' ', x, TextGridStyleDefault)
	}

	x = len(t.text.Rows[len(t.text.Rows)-1].Cells)
	if x > 0 {
		x--
	}

	t.lastCursor = fyne.Position{X: float32(x), Y: float32(len(t.text.Rows) - 1)}
}

func (t *textGridRenderer) updateGridSize(size fyne.Size) {
	bufRows := len(t.text.Rows)
	bufCols := 0
	for _, row := range t.text.Rows {
		bufCols = int(math.Max(float64(bufCols), float64(len(row.Cells))))
	}
	sizeCols := math.Floor(float64(size.Width) / float64(t.cellSize.Width))
	sizeRows := math.Floor(float64(size.Height) / float64(t.cellSize.Height))

	t.cols = int(math.Max(sizeCols, float64(bufCols)))
	t.rows = int(math.Max(sizeRows, float64(bufRows)))

	cellCount := t.cols * t.rows
	if len(t.objects) == cellCount*3 {
		return
	}
	for i := len(t.objects); i < cellCount*3; i += 2 {
		t.appendTextCell(' ')
	}
}

func (t *textGridRenderer) Layout(size fyne.Size) {
	t.updateGridSize(size)

	i := 0
	cellPos := fyne.NewPos(0, 0)
	for y := 0; y < t.rows; y++ {
		for x := 0; x < t.cols; x++ {
			// rect
			t.objects[i*3].Resize(t.cellSize)
			t.objects[i*3].Move(cellPos)

			// text
			t.objects[i*3+1].Move(cellPos)

			// underline
			t.objects[i*3+2].Move(cellPos)
			// t.objects[i*3+2].Move(cellPos.Add(fyne.Position{X: 0, Y: t.cellSize.Height}))
			// t.objects[i*3+2].Resize(fyne.Size{Width: t.cellSize.Width})
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

	// TextGridStyleWhitespace = &TextGridStyle{FGColor: theme.DisabledColor()}
	t.updateGridSize(t.text.BaseWidget.Size())
	t.refreshGrid()
}

func (t *textGridRenderer) ApplyTheme() {}

func (t *textGridRenderer) Objects() []fyne.CanvasObject {
	return t.objects
}

func (t *textGridRenderer) Destroy() {}

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
