package ui

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2/canvas"
	"github.com/srwiley/rasterx"
	"golang.org/x/image/math/fixed"
)

type TextStyle struct {
	Underline      UnderlineStyle
	UnderlineColor color.Color
	UnderlineWidth int
	Strike         bool
	StrikeWidth    int
	Overline       bool
	OverlineColor  color.Color
	OverlineWidth  int
}

func (t *TextStyle) GenerateObject(width, height int) *canvas.Raster {
	raw := image.NewRGBA(image.Rect(0, 0, width, height))
	scanner := rasterx.NewScannerGV(width, height, raw, raw.Bounds())

	switch t.Underline {
	case StraightUnderline:
		dasher := rasterx.NewDasher(width, height, scanner)
		dasher.SetColor(t.UnderlineColor)
		dasher.SetStroke(fixed.Int26_6(float64(t.UnderlineWidth)*64), 0, nil, nil, nil, 0, nil, 0)
		p1x, p1y := float32(0), float32(height-3)
		p2x, p2y := float32(width), float32(height-3)

		dasher.Start(rasterx.ToFixedP(float64(p1x), float64(p1y)))
		dasher.Line(rasterx.ToFixedP(float64(p2x), float64(p2y)))
		dasher.Stop(true)
		dasher.Draw()
	case DoubleUnderline:
		dasher := rasterx.NewDasher(width, height, scanner)
		dasher.SetColor(t.UnderlineColor)
		dasher.SetStroke(fixed.Int26_6(float64(t.UnderlineWidth)*64), 0, nil, nil, nil, 0, nil, 0)
		p1x, p1y := float32(0), float32(height-3)
		p2x, p2y := float32(width), float32(height-3)
		p3x, p3y := float32(0), float32(height)
		p4x, p4y := float32(width), float32(height)

		dasher.Start(rasterx.ToFixedP(float64(p1x), float64(p1y)))
		dasher.Line(rasterx.ToFixedP(float64(p2x), float64(p2y)))
		dasher.Stop(true)
		dasher.Start(rasterx.ToFixedP(float64(p3x), float64(p3y)))
		dasher.Line(rasterx.ToFixedP(float64(p4x), float64(p4y)))
		dasher.Stop(true)
		dasher.Draw()
	case CurlyUnderline:
		dasher := rasterx.NewDasher(width, height, scanner)
		dasher.SetColor(t.UnderlineColor)
		dasher.SetStroke(fixed.Int26_6(float64(t.UnderlineWidth)*64), 0, nil, nil, nil, 0, nil, 0)
		p1x, p1y := float32(0), float32(height)
		p2x, p2y := float32(width/2), float32(height-3)
		p3x, p3y := float32(width), float32(height)

		dasher.Start(rasterx.ToFixedP(float64(p1x), float64(p1y)))
		dasher.Line(rasterx.ToFixedP(float64(p2x), float64(p2y)))
		dasher.Stop(true)
		dasher.Start(rasterx.ToFixedP(float64(p2x), float64(p2y)))
		dasher.Line(rasterx.ToFixedP(float64(p3x), float64(p3y)))
		dasher.Stop(true)
		dasher.Draw()
	case DottedUnderline:
		dasher := rasterx.NewDasher(width, height, scanner)
		dasher.SetColor(t.UnderlineColor)
		dasher.SetStroke(fixed.Int26_6(float64(t.UnderlineWidth)*64), 0, nil, nil, nil, 0, nil, 0)
		for i := 0; i < width; i++ {
			p1x, p1y := float32(i), float32(height-3)
			p2x, p2y := float32(i+1), float32(height-3)
			i += 2

			dasher.Start(rasterx.ToFixedP(float64(p1x), float64(p1y)))
			dasher.Line(rasterx.ToFixedP(float64(p2x), float64(p2y)))
			dasher.Stop(true)
		}

		dasher.Draw()
	case DashedUnderline:
		dasher := rasterx.NewDasher(width, height, scanner)
		dasher.SetColor(t.UnderlineColor)
		dasher.SetStroke(fixed.Int26_6(float64(t.UnderlineWidth)*64), 0, nil, nil, nil, 0, nil, 0)
		p1x, p1y := float32(0), float32(height-3)
		p2x, p2y := float32(3), float32(height-3)
		p3x, p3y := float32(width-3), float32(height-3)
		p4x, p4y := float32(width), float32(height-3)

		dasher.Start(rasterx.ToFixedP(float64(p1x), float64(p1y)))
		dasher.Line(rasterx.ToFixedP(float64(p2x), float64(p2y)))
		dasher.Stop(true)
		dasher.Start(rasterx.ToFixedP(float64(p3x), float64(p3y)))
		dasher.Line(rasterx.ToFixedP(float64(p4x), float64(p4y)))
		dasher.Stop(true)
		dasher.Draw()
	}

	return canvas.NewRasterFromImage(raw)
}
