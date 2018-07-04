package golang

import (
	"github.com/ktnyt/carrera/syntax"
	"github.com/ktnyt/pars"
)

var boundary = pars.Any(
	pars.Head,
	pars.Break,
	pars.NotLatin,
	pars.EOF,
)

func word(q pars.ParserLike) pars.Parser {
	return pars.Seq(boundary, q, boundary).Map(pars.Child(1))
}

func words(q ...pars.ParserLike) pars.Parser {
	ps := make([]pars.ParserLike, len(q))
	for i := range q {
		ps[i] = word(q[i])
	}
	return pars.Any(ps...)
}

var (
	statements = words(
		"defer", "go", "goto", "return", "break", "continue", "fallthrough",
	)

	conditionals = words("if", "else", "switch", "select")

	labels = words("case", "default")

	repeats = words("for", "range")

	builtins = words(
		"append", "cap", "close", "complex", "copy", "delete", "imag", "len",
		"make", "new", "panic", "print", "println", "real", "recover",
	)

	types = words(
		"chan", "map", "bool", "string", "error",
		"int", "int8", "int16", "int32", "int64", "rune",
		"byte", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
		"float32", "float64",
		"complex64", "complex128",
	)
)

var blue = pars.Any(statements, conditionals, labels, repeats, types)
var brightCyan = pars.Any(types)

var Highlighter = syntax.Highlighter{
	Black:   pars.Fail,
	Red:     pars.Fail,
	Green:   pars.Fail,
	Yellow:  pars.Fail,
	Blue:    blue,
	Magenta: pars.Fail,
	Cyan:    pars.Fail,
	White:   pars.Fail,

	BrightBlack:   pars.Fail,
	BrightRed:     pars.Fail,
	BrightGreen:   pars.Fail,
	BrightYellow:  pars.Fail,
	BrightBlue:    pars.Fail,
	BrightMagenta: pars.Fail,
	BrightCyan:    brightCyan,
	BrightWhite:   pars.Fail,
}
