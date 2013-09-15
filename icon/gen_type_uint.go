// []uint iterator
package icon

import (
	"fmt"
	"github.com/mg/i"
)	

type uints struct {
	slice []uint
	pos  int
	err error
}

func Uints(slice []uint) i.RandomAccess {
	return &uints {slice: slice, err:nil}
}

func UintList(v ...uint) i.RandomAccess {
	return &uints {slice: v, err:nil}
}

func (s *uints) AtStart() bool {
	return s.pos == 0
}

func (s *uints) AtEnd() bool {
	return s.pos >= len(s.slice)
}

func (s *uints) Next() error {
	if s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos++
	}
	return s.err
}

func (s *uints) Prev() error {
	if s.pos < 0 {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos--
	}
	return s.err
}

func (s *uints) First() error {
	s.pos= 0
	return nil
}

func (s *uints) Last() error {
	s.pos= len(s.slice) - 1
	return nil
}

func (s *uints) Goto(pos int) error {
	s.pos= pos
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	}
	return s.err
}

func (s *uints) Len() int {
	return len(s.slice)
}

func (s *uints) Value() interface{} {
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err = fmt.Errorf("Index out of bounds: %d.", s.pos)
		return nil
	}
	return s.slice[s.pos]
}

func (s *uints) Uint() uint {
	return s.slice[s.pos]
}

func (s *uints) Error() error {
	return s.err
}

func (s *uints) SetError(err error) {
	s.err= err
}

