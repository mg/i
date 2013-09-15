// []float32 iterator
package icon

import (
	"fmt"
	"github.com/mg/i"
)	

type float32s struct {
	slice []float32
	pos  int
	err error
}

func Float32s(slice []float32) i.RandomAccess {
	return &float32s {slice: slice, err:nil}
}

func Float32List(v ...float32) i.RandomAccess {
	return &float32s {slice: v, err:nil}
}

func (s *float32s) AtStart() bool {
	return s.pos == 0
}

func (s *float32s) AtEnd() bool {
	return s.pos >= len(s.slice)
}

func (s *float32s) Next() error {
	if s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos++
	}
	return s.err
}

func (s *float32s) Prev() error {
	if s.pos < 0 {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos--
	}
	return s.err
}

func (s *float32s) First() error {
	s.pos= 0
	return nil
}

func (s *float32s) Last() error {
	s.pos= len(s.slice) - 1
	return nil
}

func (s *float32s) Goto(pos int) error {
	s.pos= pos
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	}
	return s.err
}

func (s *float32s) Len() int {
	return len(s.slice)
}

func (s *float32s) Value() interface{} {
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err = fmt.Errorf("Index out of bounds: %d.", s.pos)
		return nil
	}
	return s.slice[s.pos]
}

func (s *float32s) Float32() float32 {
	return s.slice[s.pos]
}

func (s *float32s) Error() error {
	return s.err
}

func (s *float32s) SetError(err error) {
	s.err= err
}

