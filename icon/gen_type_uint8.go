// []uint8 iterator
package icon

import (
	"fmt"
	"github.com/mg/i"
)	

type uint8s struct {
	slice []uint8
	pos  int
	err error
}

func Uint8s(slice []uint8) i.RandomAccess {
	return &uint8s {slice: slice, err:nil}
}

func Uint8List(v ...uint8) i.RandomAccess {
	return &uint8s {slice: v, err:nil}
}

func (s *uint8s) AtStart() bool {
	return s.pos == 0
}

func (s *uint8s) AtEnd() bool {
	return s.pos >= len(s.slice)
}

func (s *uint8s) Next() error {
	if s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos++
	}
	return s.err
}

func (s *uint8s) Prev() error {
	if s.pos < 0 {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos--
	}
	return s.err
}

func (s *uint8s) First() error {
	s.pos= 0
	return nil
}

func (s *uint8s) Last() error {
	s.pos= len(s.slice) - 1
	return nil
}

func (s *uint8s) Goto(pos int) error {
	s.pos= pos
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	}
	return s.err
}

func (s *uint8s) Len() int {
	return len(s.slice)
}

func (s *uint8s) Value() interface{} {
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err = fmt.Errorf("Index out of bounds: %d.", s.pos)
		return nil
	}
	return s.slice[s.pos]
}

func (s *uint8s) Uint8() uint8 {
	return s.slice[s.pos]
}

func (s *uint8s) Error() error {
	return s.err
}

func (s *uint8s) SetError(err error) {
	s.err= err
}

