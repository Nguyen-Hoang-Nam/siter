package ui

type UnderlineStyle int

const (
	NoUnderline UnderlineStyle = iota
	StraightUnderline
	DoubleUnderline
	CurlyUnderline
	DottedUnderline
	DashedUnderline
)

type ColorIntensity int

const (
	NormalIntenisty ColorIntensity = iota
	BoldIntensity
	DimIntensity
)

type BlinkStyle int

const (
	NoBlink BlinkStyle = iota
	NormalBlink
	RapidBlink
)

type VerticalAlignStyle int

const (
	NormalBaseline VerticalAlignStyle = iota
	SuperScript
	SubScript
)
