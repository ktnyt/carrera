package carrera

import termbox "github.com/ktnyt/termbox-go"

type Event termbox.Event

type Color termbox.Attribute

var (
	Default = Color(termbox.ColorDefault)
	Black   = Color(termbox.ColorBlack)
	Red     = Color(termbox.ColorRed)
	Green   = Color(termbox.ColorGreen)
	Yellow  = Color(termbox.ColorYellow)
	Blue    = Color(termbox.ColorBlue)
	Magenta = Color(termbox.ColorMagenta)
	Cyan    = Color(termbox.ColorCyan)
	White   = Color(termbox.ColorWhite)

	BrightBlack   = Color(termbox.AttrBright | termbox.ColorBlack)
	BrightRed     = Color(termbox.AttrBright | termbox.ColorRed)
	BrightGreen   = Color(termbox.AttrBright | termbox.ColorGreen)
	BrightYellow  = Color(termbox.AttrBright | termbox.ColorYellow)
	BrightBlue    = Color(termbox.AttrBright | termbox.ColorBlue)
	BrightMagenta = Color(termbox.AttrBright | termbox.ColorMagenta)
	BrightCyan    = Color(termbox.AttrBright | termbox.ColorCyan)
	BrightWhite   = Color(termbox.AttrBright | termbox.ColorWhite)
)

type Cell struct {
	Ch rune
	Fg Color
	Bg Color
}

func SetCell(x, y int, c Cell) {
	termbox.SetCell(x, y, c.Ch, termbox.Attribute(c.Fg), termbox.Attribute(c.Bg))
}
