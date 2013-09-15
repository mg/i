// []int16 iterator
package icon

import (
	"fmt"
	"github.com/mg/i"
)	

type int16s struct {
	slice []int16
	pos  int
	err error
}

func Int16s(slice []int16) i.RandomAccess {
	return &int16s {slice: slice, err:nil}
}

func Int16List(v ...int16) i.RandomAccess {
	return &int16s {slice: v, err:nil}
}

func (s *int16s) AtStart() bool {
	return s.pos == 0
}

func (s *int16s) AtEnd() bool {
	return s.pos >= len(s.slice)
}

func (s *int16s) Next() error {
	if s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos++
	}
	return s.err
}

func (s *int16s) Prev() error {
	if s.pos < 0 {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos--
	}
	return s.err
}

func (s *int16s) First() error {
	s.pos= 0
	return nil
}

func (s *int16s) Last() error {
	s.pos= len(s.slice) - 1
	return nil
}

func (s *int16s) Goto(pos int) error {
	s.pos= pos
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	}
	return s.err
}

func (s *int16s) Len() int {
	return len(s.slice)
}

func (s *int16s) Value() interface{} {
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err = fmt.Errorf("Index out of bounds: %d.", s.pos)
		return nil
	}
	return s.slice[s.pos]
}

func (s *int16s) Int16() int16 {
	return s.slice[s.pos]
}

func (s *int16s) Error() error {
	return s.err
}

func (s *int16s) SetError(err error) {
	s.err= err
}

