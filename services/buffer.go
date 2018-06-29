package services

import (
	"unicode/utf8"

	"github.com/ktnyt/carrera"
)

func NewBufferService(buffer carrera.Buffer) carrera.BufferService {
	return bufferService{ptr: &bufferServiceStorage{buffer: buffer}}
}

type bufferServiceStorage struct {
	buffer carrera.Buffer
}

type bufferService struct {
	ptr *bufferServiceStorage
}

func (s bufferService) Rune(pos int) (rune, error) {
	if pos > len(s.ptr.buffer) {
		return utf8.RuneError, carrera.NewOutOfRangeError(pos)
	}
	return s.ptr.buffer[pos], nil
}

func (s bufferService) Insert(pos int, r rune) error {
	if pos > len(s.ptr.buffer) {
		return carrera.NewOutOfRangeError(pos)
	}
	s.ptr.buffer = append(s.ptr.buffer, 0)
	copy(s.ptr.buffer[pos+1:], s.ptr.buffer[pos:])
	s.ptr.buffer[pos] = r
	return nil
}

func (s bufferService) Delete(pos int) (rune, error) {
	if pos > len(s.ptr.buffer) {
		return utf8.RuneError, carrera.NewOutOfRangeError(pos)
	}
	r := s.ptr.buffer[pos]
	s.ptr.buffer = append(s.ptr.buffer[:pos], s.ptr.buffer[pos+1:]...)
	return r, nil
}

func (s bufferService) Length() int {
	return len(s.ptr.buffer)
}
