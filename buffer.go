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
	Rune(pos int) (rune, error)
	Insert(pos int, r rune) error
	Delete(pos int) (rune, error)
	Length() int
}

type BufferController interface {
	Insert(pos int, r rune)
	Delete(pos int) rune
}

type BufferView interface {
	Draw() error
	Handle(termbox.Event) error
}
