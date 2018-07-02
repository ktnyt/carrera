package carrera

type Rune struct {
	Value rune
	Color Color
}

func NewRune(r rune, c Color) Rune {
	return Rune{Value: r, Color: c}
}
