// []complex128 iterator
package icon

import (
	"fmt"
	"github.com/mg/i"
)	

type complex128s struct {
	slice []complex128
	pos  int
	err error
}

func Complex128s(slice []complex128) i.RandomAccess {
	return &complex128s {slice: slice, err:nil}
}

func Complex128List(v ...complex128) i.RandomAccess {
	return &complex128s {slice: v, err:nil}
}

func (s *complex128s) AtStart() bool {
	return s.pos == 0
}

func (s *complex128s) AtEnd() bool {
	return s.pos >= len(s.slice)
}

func (s *complex128s) Next() error {
	if s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos++
	}
	return s.err
}

func (s *complex128s) Prev() error {
	if s.pos < 0 {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos--
	}
	return s.err
}

func (s *complex128s) First() error {
	s.pos= 0
	return nil
}

func (s *complex128s) Last() error {
	s.pos= len(s.slice) - 1
	return nil
}

func (s *complex128s) Goto(pos int) error {
	s.pos= pos
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	}
	return s.err
}

func (s *complex128s) Len() int {
	return len(s.slice)
}

func (s *complex128s) Value() interface{} {
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err = fmt.Errorf("Index out of bounds: %d.", s.pos)
		return nil
	}
	return s.slice[s.pos]
}

func (s *complex128s) Complex128() complex128 {
	return s.slice[s.pos]
}

func (s *complex128s) Error() error {
	return s.err
}

func (s *complex128s) SetError(err error) {
	s.err= err
}

