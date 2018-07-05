package buffer

import (
	. "github.com/ktnyt/carrera"
)

type Middleware func([]rune) []Cell

func NewBufferService(buffer Buffer, middleware Middleware) BufferService {
	cells := middleware(buffer)
	return bufferService{
		ptr: &bufferStorage{
			buffer: buffer,
			cells:  cells,
		},
		middleware: middleware,
	}
}

type bufferStorage struct {
	buffer Buffer
	cells  []Cell
}

func (s *bufferStorage) update(middleware Middleware) {
	s.cells = middleware(s.buffer)
}

type bufferService struct {
	ptr *bufferStorage

	middleware Middleware
}

func (s bufferService) Cell(pos int) Cell {
	return s.ptr.cells[pos]
}

func (s bufferService) Insert(pos int, r rune) {
	s.ptr.buffer = append(s.ptr.buffer, 0)
	copy(s.ptr.buffer[pos+1:], s.ptr.buffer[pos:])
	s.ptr.buffer[pos] = r
	s.ptr.update(s.middleware)
}

func (s bufferService) Delete(pos int) rune {
	r := s.ptr.buffer[pos]
	s.ptr.buffer = append(s.ptr.buffer[:pos], s.ptr.buffer[pos+1:]...)
	s.ptr.update(s.middleware)
	return r
}

func (s bufferService) Set(rs []rune) {
	s.ptr.buffer = rs
	s.ptr.update(s.middleware)
}

func (s bufferService) Clear() {
	s.ptr.buffer = NewBuffer()
	s.ptr.update(s.middleware)
}

func (s bufferService) Length() int {
	return len(s.ptr.cells)
}
