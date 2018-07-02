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
)
