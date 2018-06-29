package carrera

import "fmt"

type OutOfRangeError struct {
	pos int
}

func NewOutOfRangeError(pos int) error {
	return &OutOfRangeError{pos: pos}
}

func (e *OutOfRangeError) Error() string {
	return fmt.Sprintf("position %d is out of buffer range", e.pos)
}
