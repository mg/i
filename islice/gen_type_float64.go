// []float64 iterator
package islice

import (
	"fmt"
	"github.com/mg/i"
)	

type float64s struct {
	slice []float64
	pos  int
	err error
}

func Float64s(slice []float64) i.RandomAccess {
	return &float64s {slice: slice, err:nil}
}

func Float64List(v ...float64) i.RandomAccess {
	return &float64s {slice: v, err:nil}
}

func (s *float64s) AtStart() bool {
	return s.pos == 0
}

func (s *float64s) AtEnd() bool {
	return s.pos >= len(s.slice)
}

func (s *float64s) Next() error {
	if s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos++
	}
	return s.err
}

func (s *float64s) Prev() error {
	if s.pos < 0 {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos--
	}
	return s.err
}

func (s *float64s) First() error {
	s.pos= 0
	return nil
}

func (s *float64s) Last() error {
	s.pos= len(s.slice) - 1
	return nil
}

func (s *float64s) Goto(pos int) error {
	s.pos= pos
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	}
	return s.err
}

func (s *float64s) Len() int {
	return len(s.slice)
}

func (s *float64s) Value() interface{} {
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err = fmt.Errorf("Index out of bounds: %d.", s.pos)
		return nil
	}
	return s.slice[s.pos]
}

func (s *float64s) Float64() float64 {
	return s.slice[s.pos]
}

func (s *float64s) Error() error {
	return s.err
}

func (s *float64s) SetError(err error) {
	s.err= err
}

