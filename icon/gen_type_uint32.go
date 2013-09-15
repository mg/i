// []uint32 iterator
package icon

import (
	"fmt"
	"github.com/mg/i"
)	

type uint32s struct {
	slice []uint32
	pos  int
	err error
}

func Uint32s(slice []uint32) i.RandomAccess {
	return &uint32s {slice: slice, err:nil}
}

func Uint32List(v ...uint32) i.RandomAccess {
	return &uint32s {slice: v, err:nil}
}

func (s *uint32s) AtStart() bool {
	return s.pos == 0
}

func (s *uint32s) AtEnd() bool {
	return s.pos >= len(s.slice)
}

func (s *uint32s) Next() error {
	if s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos++
	}
	return s.err
}

func (s *uint32s) Prev() error {
	if s.pos < 0 {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos--
	}
	return s.err
}

func (s *uint32s) First() error {
	s.pos= 0
	return nil
}

func (s *uint32s) Last() error {
	s.pos= len(s.slice) - 1
	return nil
}

func (s *uint32s) Goto(pos int) error {
	s.pos= pos
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	}
	return s.err
}

func (s *uint32s) Len() int {
	return len(s.slice)
}

func (s *uint32s) Value() interface{} {
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err = fmt.Errorf("Index out of bounds: %d.", s.pos)
		return nil
	}
	return s.slice[s.pos]
}

func (s *uint32s) Uint32() uint32 {
	return s.slice[s.pos]
}

func (s *uint32s) Error() error {
	return s.err
}

func (s *uint32s) SetError(err error) {
	s.err= err
}

