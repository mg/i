// []uint64 iterator
package icon

import (
	"fmt"
	"github.com/mg/i"
)	

type uint64s struct {
	slice []uint64
	pos  int
	err error
}

func Uint64s(slice []uint64) i.RandomAccess {
	return &uint64s {slice: slice, err:nil}
}

func Uint64List(v ...uint64) i.RandomAccess {
	return &uint64s {slice: v, err:nil}
}

func (s *uint64s) AtStart() bool {
	return s.pos == 0
}

func (s *uint64s) AtEnd() bool {
	return s.pos >= len(s.slice)
}

func (s *uint64s) Next() error {
	if s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos++
	}
	return s.err
}

func (s *uint64s) Prev() error {
	if s.pos < 0 {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos--
	}
	return s.err
}

func (s *uint64s) First() error {
	s.pos= 0
	return nil
}

func (s *uint64s) Last() error {
	s.pos= len(s.slice) - 1
	return nil
}

func (s *uint64s) Goto(pos int) error {
	s.pos= pos
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	}
	return s.err
}

func (s *uint64s) Len() int {
	return len(s.slice)
}

func (s *uint64s) Value() interface{} {
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err = fmt.Errorf("Index out of bounds: %d.", s.pos)
		return nil
	}
	return s.slice[s.pos]
}

func (s *uint64s) Uint64() uint64 {
	return s.slice[s.pos]
}

func (s *uint64s) Error() error {
	return s.err
}

func (s *uint64s) SetError(err error) {
	s.err= err
}

