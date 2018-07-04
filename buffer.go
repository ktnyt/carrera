package carrera

import termbox "github.com/nsf/termbox-go"

type Buffer []rune

func NewBuffer() Buffer {
	return NewBufferWithCapacity(0)
}

func NewBufferWithCapacity(c int) Buffer {
	return make([]rune, 0, c)
}

type BufferService interface {
	Cell(pos int) Cell
	Insert(pos int, r rune)
	Delete(pos int) rune
	Set(rs []rune)
	Clear()
	Length() int
}

type BufferPresenter interface {
	Area(line, width, height int) [][]Cell
	Insert(pos int, r rune)
	Delete(pos int) rune
}

type BufferView interface {
	Draw()
	Handle(event termbox.Event)
}
