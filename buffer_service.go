package carrera

import (
	"unicode/utf8"
)

func NewBufferService(buffer Buffer) BufferService {
	return bufferService{ptr: &bufferServiceStorage{buffer: buffer}}
}

type bufferServiceStorage struct {
	buffer Buffer
}

type bufferService struct {
	ptr *bufferServiceStorage
}

func (s bufferService) Rune(pos int) (rune, error) {
	if pos > len(s.ptr.buffer) {
		return utf8.RuneError, NewOutOfRangeError(pos)
	}
	return s.ptr.buffer[pos], nil
}

func (s bufferService) Insert(pos int, r rune) error {
	if pos > len(s.ptr.buffer) {
		return NewOutOfRangeError(pos)
	}
	s.ptr.buffer = append(s.ptr.buffer, 0)
	copy(s.ptr.buffer[pos+1:], s.ptr.buffer[pos:])
	s.ptr.buffer[pos] = r
	return nil
}

func (s bufferService) Delete(pos int) (rune, error) {
	if pos > len(s.ptr.buffer) {
		return utf8.RuneError, NewOutOfRangeError(pos)
	}
	r := s.ptr.buffer[pos]
	s.ptr.buffer = append(s.ptr.buffer[:pos], s.ptr.buffer[pos+1:]...)
	return r, nil
}

func (s bufferService) Set(rs []rune) {
	s.ptr.buffer = rs
}

func (s bufferService) Clear() {
	s.ptr.buffer = NewBuffer()
}

func (s bufferService) Length() int {
	return len(s.ptr.buffer)
}
