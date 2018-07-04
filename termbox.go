package carrera

import termbox "github.com/nsf/termbox-go"

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

	BrightBlack   = Color(termbox.AttrBold | termbox.ColorBlack)
	BrightRed     = Color(termbox.AttrBold | termbox.ColorRed)
	BrightGreen   = Color(termbox.AttrBold | termbox.ColorGreen)
	BrightYellow  = Color(termbox.AttrBold | termbox.ColorYellow)
	BrightBlue    = Color(termbox.AttrBold | termbox.ColorBlue)
	BrightMagenta = Color(termbox.AttrBold | termbox.ColorMagenta)
	BrightCyan    = Color(termbox.AttrBold | termbox.ColorCyan)
	BrightWhite   = Color(termbox.AttrBold | termbox.ColorWhite)
)

type Cell struct {
	Ch rune
	Fg Color
	Bg Color
}

func SetCell(x, y int, c Cell) {
	termbox.SetCell(x, y, c.Ch, termbox.Attribute(c.Fg), termbox.Attribute(c.Bg))
}
