package carrera

import (
	"strings"

	"github.com/ktnyt/carrera/syntax/golang"
)

func NewBufferService(buffer Buffer) BufferService {
	return bufferService{ptr: &bufferServiceStorage{buffer: buffer}}
}

type bufferServiceStorage struct {
	buffer Buffer
}

type bufferService struct {
	ptr   *bufferServiceStorage
	cells []Cell
}

func (s bufferService) update() {
	cells := make([]Cell, len(s.ptr.buffer))
	r := strings.NewReader(string(s.ptr.buffer))
	colors := Highlight(r, golang.Highlighter)
	for i := range s.ptr.buffer {
		cells[i].Ch = s.ptr.buffer[i]
		cells[i].Fg = colors[i]
		cells[i].Bg = Default
	}
}

func (s bufferService) Cell(pos int) Cell {
	return s.cells[pos]
}

func (s bufferService) Insert(pos int, r rune) {
	s.ptr.buffer = append(s.ptr.buffer, 0)
	copy(s.ptr.buffer[pos+1:], s.ptr.buffer[pos:])
	s.ptr.buffer[pos] = r
	s.update()
}

func (s bufferService) Delete(pos int) rune {
	r := s.ptr.buffer[pos]
	s.ptr.buffer = append(s.ptr.buffer[:pos], s.ptr.buffer[pos+1:]...)
	s.update()
	return r
}

func (s bufferService) Set(rs []rune) {
	s.ptr.buffer = rs
	s.update()
}

func (s bufferService) Clear() {
	s.ptr.buffer = NewBuffer()
	s.update()
}

func (s bufferService) Length() int {
	return len(s.ptr.buffer)
}
