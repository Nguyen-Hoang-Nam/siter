package ui

import (
	"image"
	"image/color"
	"image/draw"
	"math"

	"fyne.io/fyne/v2"
	"github.com/go-text/render"
	"github.com/go-text/typesetting/di"
	"github.com/go-text/typesetting/font"
	"github.com/go-text/typesetting/shaping"
	"github.com/srwiley/rasterx"
	"golang.org/x/image/math/fixed"
)

const (
	DefaultTabWidth = 4

	fontTabSpaceSize = 10
)

type RuneCellStyle struct {
	FontStyle       FontStyle
	FontSize        float32
	ForegroundColor color.Color
	BackgroundColor color.Color

	Underline      UnderlineStyle
	UnderlineColor color.Color
	UnderlineWidth int
	Strike         bool
	StrikeColor    color.Color
	StrikeWidth    int
	Overline       bool
	OverlineColor  color.Color
	OverlineWidth  int
	Blink          BlinkStyle
	Invisible      bool // Not support due to security risk kitty#266
	VerticalAlign  VerticalAlignStyle
}

type RuneCell struct {
	Rune   rune
	Width  int
	Height int
	Style  RuneCellStyle
}

func NewRuneCellStyle() RuneCellStyle {
	size := float32(0)
	if fyne.CurrentApp() != nil { // nil app possible if app not started
		size = fyne.CurrentApp().Settings().Theme().Size("text") // manually name the size to avoid import loop
	}

	return RuneCellStyle{
		FontSize: size,
	}
}

type subImg interface {
	SubImage(r image.Rectangle) image.Image
}

func (t *RuneCell) Generate() func(w int, h int) image.Image {
	raw := image.NewRGBA(image.Rect(0, 0, t.Width, t.Height))
	scanner := rasterx.NewScannerGV(t.Width, t.Height, raw, raw.Bounds())

	if t.Style.BackgroundColor != color.Transparent {
		t.drawBackground(scanner)
	}

	if t.Style.Underline != NoUnderline {
		t.drawUnderline(scanner)
	}

	if t.Style.Overline {
		t.drawOverline(scanner)
	}

	if t.Style.Strike {
		t.drawStrike(scanner)
	}

	t.drawRune(raw)

	img := image.Image(raw)

	return func(w int, h int) image.Image {
		bounds := img.Bounds()
		rect := image.Rect(0, 0, w, h)

		switch {
		case w == bounds.Max.X && h == bounds.Max.Y:
			return img
		case w >= bounds.Max.X && h >= bounds.Max.Y:
			if sub, ok := img.(subImg); ok {
				return sub.SubImage(image.Rectangle{
					Min: bounds.Min,
					Max: image.Point{
						X: bounds.Min.X + w,
						Y: bounds.Min.Y + h,
					},
				})
			}
		default:
			if !rect.Overlaps(bounds) {
				return image.NewUniform(color.RGBA{})
			}
			bounds = bounds.Intersect(rect)
		}

		dst := image.NewRGBA(rect)
		draw.Draw(dst, bounds, img, bounds.Min, draw.Over)
		return dst
	}
}

func (r *RuneCell) drawBackground(scanner rasterx.Scanner) {
	filler := rasterx.NewFiller(r.Width, r.Height, scanner)
	filler.SetColor(r.Style.BackgroundColor)
	rasterx.AddRect(0, 0, float64(r.Width), float64(r.Height), 0, filler)
	filler.Draw()
}

func (t *RuneCell) drawRune(dst draw.Image) {
	face := CachedFontFace(t.Style.FontStyle)

	r := render.Renderer{
		FontSize: t.Style.FontSize,
		Color:    t.Style.ForegroundColor,
	}

	sh := &shaping.HarfbuzzShaper{}
	out := sh.Shape(shaping.Input{
		Text:     []rune{t.Rune},
		RunStart: 0,
		RunEnd:   1,
		Face:     face.Fonts[0],
		Size:     fixed.I(int(t.Style.FontSize)),
	})

	advance := float32(0)
	y := int(math.Ceil(float64(fixed266ToFloat32(out.LineBounds.Ascent))))
	walkString(face.Fonts, t.Rune, float32ToFixed266(t.Style.FontSize), &advance, func(run shaping.Output, x float32) {
		if len(run.Glyphs) == 1 {
			if run.Glyphs[0].GlyphID == 0 {
				r.DrawStringAt(string([]rune{0xfffd}), dst, int(x), y, face.Fonts[0])
				return
			}
		}

		r.DrawShapedRunAt(run, dst, int(x), y)
	})
}

func (t *RuneCell) drawUnderline(scanner rasterx.Scanner) {
	switch t.Style.Underline {
	case StraightUnderline:
		dasher := rasterx.NewDasher(t.Width, t.Height, scanner)
		dasher.SetColor(t.Style.UnderlineColor)
		dasher.SetStroke(fixed.Int26_6(float64(t.Style.UnderlineWidth)*64), 0, nil, nil, nil, 0, nil, 0)
		p1x, p1y := float32(0), float32(t.Height-3)
		p2x, p2y := float32(t.Width), float32(t.Height-3)

		dasher.Start(rasterx.ToFixedP(float64(p1x), float64(p1y)))
		dasher.Line(rasterx.ToFixedP(float64(p2x), float64(p2y)))
		dasher.Stop(true)
		dasher.Draw()
	case DoubleUnderline:
		dasher := rasterx.NewDasher(t.Width, t.Height, scanner)
		dasher.SetColor(t.Style.UnderlineColor)
		dasher.SetStroke(fixed.Int26_6(float64(t.Style.UnderlineWidth)*64), 0, nil, nil, nil, 0, nil, 0)
		p1x, p1y := float32(0), float32(t.Height-3)
		p2x, p2y := float32(t.Width), float32(t.Height-3)
		p3x, p3y := float32(0), float32(t.Height)
		p4x, p4y := float32(t.Width), float32(t.Height)

		dasher.Start(rasterx.ToFixedP(float64(p1x), float64(p1y)))
		dasher.Line(rasterx.ToFixedP(float64(p2x), float64(p2y)))
		dasher.Stop(true)
		dasher.Start(rasterx.ToFixedP(float64(p3x), float64(p3y)))
		dasher.Line(rasterx.ToFixedP(float64(p4x), float64(p4y)))
		dasher.Stop(true)
		dasher.Draw()
	case CurlyUnderline:
		dasher := rasterx.NewDasher(t.Width, t.Height, scanner)
		dasher.SetColor(t.Style.UnderlineColor)
		dasher.SetStroke(fixed.Int26_6(float64(t.Style.UnderlineWidth)*64), 0, nil, nil, nil, 0, nil, 0)
		p1x, p1y := float32(0), float32(t.Height)
		p2x, p2y := float32(t.Width/2), float32(t.Height-3)
		p3x, p3y := float32(t.Width), float32(t.Height)

		dasher.Start(rasterx.ToFixedP(float64(p1x), float64(p1y)))
		dasher.Line(rasterx.ToFixedP(float64(p2x), float64(p2y)))
		dasher.Stop(true)
		dasher.Start(rasterx.ToFixedP(float64(p2x), float64(p2y)))
		dasher.Line(rasterx.ToFixedP(float64(p3x), float64(p3y)))
		dasher.Stop(true)
		dasher.Draw()
	case DottedUnderline:
		dasher := rasterx.NewDasher(t.Width, t.Height, scanner)
		dasher.SetColor(t.Style.UnderlineColor)
		dasher.SetStroke(fixed.Int26_6(float64(t.Style.UnderlineWidth)*64), 0, nil, nil, nil, 0, nil, 0)
		for i := 0; i < t.Width; i++ {
			p1x, p1y := float32(i), float32(t.Height-3)
			p2x, p2y := float32(i+1), float32(t.Height-3)
			i += 2

			dasher.Start(rasterx.ToFixedP(float64(p1x), float64(p1y)))
			dasher.Line(rasterx.ToFixedP(float64(p2x), float64(p2y)))
			dasher.Stop(true)
		}

		dasher.Draw()
	case DashedUnderline:
		dasher := rasterx.NewDasher(t.Width, t.Height, scanner)
		dasher.SetColor(t.Style.UnderlineColor)
		dasher.SetStroke(fixed.Int26_6(float64(t.Style.UnderlineWidth)*64), 0, nil, nil, nil, 0, nil, 0)
		p1x, p1y := float32(0), float32(t.Height-3)
		p2x, p2y := float32(3), float32(t.Height-3)
		p3x, p3y := float32(t.Width-3), float32(t.Height-3)
		p4x, p4y := float32(t.Width), float32(t.Height-3)

		dasher.Start(rasterx.ToFixedP(float64(p1x), float64(p1y)))
		dasher.Line(rasterx.ToFixedP(float64(p2x), float64(p2y)))
		dasher.Stop(true)
		dasher.Start(rasterx.ToFixedP(float64(p3x), float64(p3y)))
		dasher.Line(rasterx.ToFixedP(float64(p4x), float64(p4y)))
		dasher.Stop(true)
		dasher.Draw()
	}
}

func (t *RuneCell) drawOverline(scanner rasterx.Scanner) {
	dasher := rasterx.NewDasher(t.Width, t.Height, scanner)
	dasher.SetColor(t.Style.OverlineColor)
	dasher.SetStroke(fixed.Int26_6(float64(t.Style.OverlineWidth)*64), 0, nil, nil, nil, 0, nil, 0)
	p1x, p1y := float32(0), float32(0)
	p2x, p2y := float32(t.Width), float32(0)

	dasher.Start(rasterx.ToFixedP(float64(p1x), float64(p1y)))
	dasher.Line(rasterx.ToFixedP(float64(p2x), float64(p2y)))
	dasher.Stop(true)
	dasher.Draw()
}

func (t *RuneCell) drawStrike(scanner rasterx.Scanner) {
	dasher := rasterx.NewDasher(t.Width, t.Height, scanner)
	dasher.SetColor(t.Style.StrikeColor)
	dasher.SetStroke(fixed.Int26_6(float64(t.Style.StrikeWidth)*64), 0, nil, nil, nil, 0, nil, 0)
	p1x, p1y := float32(0), float32(t.Height/2)
	p2x, p2y := float32(t.Width), float32(t.Height/2)

	dasher.Start(rasterx.ToFixedP(float64(p1x), float64(p1y)))
	dasher.Line(rasterx.ToFixedP(float64(p2x), float64(p2y)))
	dasher.Stop(true)
	dasher.Draw()
}

func fixed266ToFloat32(i fixed.Int26_6) float32 {
	return float32(float64(i) / (1 << 6))
}

func float32ToFixed266(f float32) fixed.Int26_6 {
	return fixed.Int26_6(float64(f) * (1 << 6))
}

func walkString(faces []font.Face, r rune, textSize fixed.Int26_6, advance *float32,
	cb func(run shaping.Output, x float32)) (size fyne.Size, base float32) {
	in := shaping.Input{
		Text:      []rune{' '},
		RunStart:  0,
		RunEnd:    1,
		Direction: di.DirectionLTR,
		Face:      faces[0],
		Size:      textSize,
	}
	shaper := &shaping.HarfbuzzShaper{}
	out := shaper.Shape(in)

	in.Text = []rune{r}
	in.RunStart = 0
	in.RunEnd = 1

	x := float32(0)
	spacew := float32(fontTabSpaceSize)
	ins := shaping.SplitByFontGlyphs(in, faces)
	for _, in := range ins {
		inEnd := in.RunEnd

		pending := false
		for i, r := range in.Text[in.RunStart:in.RunEnd] {
			if r == '\t' {
				if pending {
					in.RunEnd = i
					out = shaper.Shape(in)
					x = shapeCallback(shaper, in, out, x, cb)
				}
				x = tabStop(spacew, x)

				in.RunStart = i + 1
				in.RunEnd = inEnd
				pending = false
			} else {
				pending = true
			}
		}

		x = shapeCallback(shaper, in, out, x, cb)
	}

	*advance = x
	return fyne.NewSize(*advance, fixed266ToFloat32(out.LineBounds.LineHeight())),
		fixed266ToFloat32(out.LineBounds.Ascent)
}

func tabStop(spacew, x float32) float32 {
	tabw := spacew * float32(DefaultTabWidth)
	tabs, _ := math.Modf(float64((x + tabw) / tabw))
	return tabw * float32(tabs)
}

func shapeCallback(shaper shaping.Shaper, in shaping.Input, out shaping.Output, x float32, cb func(shaping.Output, float32)) float32 {
	out = shaper.Shape(in)
	glyphs := out.Glyphs
	start := 0
	pending := false
	adv := fixed.I(0)
	for i, g := range out.Glyphs {
		if g.GlyphID == 0 {
			if pending {
				out.Glyphs = glyphs[start:i]
				cb(out, x)
				x += fixed266ToFloat32(adv)
				adv = 0
			}

			out.Glyphs = glyphs[i : i+1]
			cb(out, x)
			x += fixed266ToFloat32(glyphs[i].XAdvance)
			adv = 0

			start = i + 1
			pending = false
		} else {
			pending = true
		}
		adv += g.XAdvance
	}

	if pending {
		out.Glyphs = glyphs[start:]
		cb(out, x)
		x += fixed266ToFloat32(adv)
		adv = 0
	}
	return x + fixed266ToFloat32(adv)
}
