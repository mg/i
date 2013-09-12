// []complex64 iterator
package islice

import (
	"fmt"
	"github.com/mg/i"
)	

type complex64s struct {
	slice []complex64
	pos  int
	err error
}

func Complex64s(slice []complex64) i.RandomAccess {
	return &complex64s {slice: slice, err:nil}
}

func Complex64List(v ...complex64) i.RandomAccess {
	return &complex64s {slice: v, err:nil}
}

func (s *complex64s) AtStart() bool {
	return s.pos == 0
}

func (s *complex64s) AtEnd() bool {
	return s.pos >= len(s.slice)
}

func (s *complex64s) Next() error {
	if s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos++
	}
	return s.err
}

func (s *complex64s) Prev() error {
	if s.pos < 0 {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos--
	}
	return s.err
}

func (s *complex64s) First() error {
	s.pos= 0
	return nil
}

func (s *complex64s) Last() error {
	s.pos= len(s.slice) - 1
	return nil
}

func (s *complex64s) Goto(pos int) error {
	s.pos= pos
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	}
	return s.err
}

func (s *complex64s) Len() int {
	return len(s.slice)
}

func (s *complex64s) Value() interface{} {
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err = fmt.Errorf("Index out of bounds: %d.", s.pos)
		return nil
	}
	return s.slice[s.pos]
}

func (s *complex64s) Complex64() complex64 {
	return s.slice[s.pos]
}

func (s *complex64s) Error() error {
	return s.err
}

func (s *complex64s) SetError(err error) {
	s.err= err
}

