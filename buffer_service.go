package carrera

func NewBufferService(buffer Buffer) BufferService {
	return bufferService{ptr: &bufferServiceStorage{buffer: buffer}}
}

type bufferServiceStorage struct {
	buffer Buffer
}

type bufferService struct {
	ptr *bufferServiceStorage
}

func (s bufferService) Rune(pos int) Rune {
	return Rune{Value: s.ptr.buffer[pos], Color: White}
}

func (s bufferService) Insert(pos int, r rune) {
	s.ptr.buffer = append(s.ptr.buffer, 0)
	copy(s.ptr.buffer[pos+1:], s.ptr.buffer[pos:])
	s.ptr.buffer[pos] = r
}

func (s bufferService) Delete(pos int) rune {
	r := s.ptr.buffer[pos]
	s.ptr.buffer = append(s.ptr.buffer[:pos], s.ptr.buffer[pos+1:]...)
	return r
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
