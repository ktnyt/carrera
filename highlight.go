package carrera

import (
	"io"

	"github.com/ktnyt/carrera/syntax"
	"github.com/ktnyt/pars"
)

type result struct {
	Length int
	Color  Color
}

func color(c Color) pars.Map {
	return func(r *pars.Result) {
		colors := make([]Color, len(r.Value.(string)))
		for i := range colors {
			colors[i] = c
		}
		r.Value = colors
	}
}

func constructParser(h syntax.Highlighter) pars.Parser {
	return pars.Many(pars.Any(
		h.Black.Map(color(Black)),
		h.Red.Map(color(Red)),
		h.Green.Map(color(Green)),
		h.Yellow.Map(color(Yellow)),
		h.Blue.Map(color(Blue)),
		h.Magenta.Map(color(Magenta)),
		h.Cyan.Map(color(Cyan)),
		h.White.Map(color(White)),

		h.BrightBlack.Map(color(BrightBlack)),
		h.BrightRed.Map(color(BrightRed)),
		h.BrightGreen.Map(color(BrightGreen)),
		h.BrightYellow.Map(color(BrightYellow)),
		h.BrightBlue.Map(color(BrightBlue)),
		h.BrightMagenta.Map(color(BrightMagenta)),
		h.BrightCyan.Map(color(BrightCyan)),
		h.BrightWhite.Map(color(BrightWhite)),

		pars.WildRune.Map(func(r *pars.Result) {
			r.Value = []Color{White}
		}),
	)).Map(func(r *pars.Result) {
		colors := make([]Color, 0)
		for _, child := range r.Children {
			colors = append(colors, child.Value.([]Color)...)
		}
		r.Value = colors
		r.Children = nil
	})
}

func Highlight(rd io.Reader, h syntax.Highlighter) []Color {
	s := pars.NewState(rd)
	r := pars.Result{}
	constructParser(h)(s, &r)
	return r.Value.([]Color)
}
