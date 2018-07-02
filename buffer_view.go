package carrera

import (
	termbox "github.com/nsf/termbox-go"
)

type bufferView struct {
	presenter BufferPresenter

	width  int
	height int
}

func NewBufferView(presenter BufferPresenter, width, height int) BufferView {
	return bufferView{presenter: presenter, width: width, height: height}
}

func (v bufferView) Draw() {
	area := v.presenter.Area(0, v.width, v.height)
	for y, line := range area {
		for x, r := range line {
			termbox.SetCell(x, y, r.Value, termbox.Attribute(r.Color), termbox.ColorDefault)
		}
	}
}

func (v bufferView) Handle(event termbox.Event) {
	return
}
