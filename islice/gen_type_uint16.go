// []uint16 iterator
package islice

import (
	"fmt"
	"github.com/mg/i"
)	

type uint16s struct {
	slice []uint16
	pos  int
	err error
}

func Uint16s(slice []uint16) i.RandomAccess {
	return &uint16s {slice: slice, err:nil}
}

func Uint16List(v ...uint16) i.RandomAccess {
	return &uint16s {slice: v, err:nil}
}

func (s *uint16s) AtStart() bool {
	return s.pos == 0
}

func (s *uint16s) AtEnd() bool {
	return s.pos >= len(s.slice)
}

func (s *uint16s) Next() error {
	if s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos++
	}
	return s.err
}

func (s *uint16s) Prev() error {
	if s.pos < 0 {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos--
	}
	return s.err
}

func (s *uint16s) First() error {
	s.pos= 0
	return nil
}

func (s *uint16s) Last() error {
	s.pos= len(s.slice) - 1
	return nil
}

func (s *uint16s) Goto(pos int) error {
	s.pos= pos
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	}
	return s.err
}

func (s *uint16s) Len() int {
	return len(s.slice)
}

func (s *uint16s) Value() interface{} {
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err = fmt.Errorf("Index out of bounds: %d.", s.pos)
		return nil
	}
	return s.slice[s.pos]
}

func (s *uint16s) Uint16() uint16 {
	return s.slice[s.pos]
}

func (s *uint16s) Error() error {
	return s.err
}

func (s *uint16s) SetError(err error) {
	s.err= err
}

