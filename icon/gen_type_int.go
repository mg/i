// []int iterator
package icon

import (
	"fmt"
	"github.com/mg/i"
)	

type ints struct {
	slice []int
	pos  int
	err error
}

func Ints(slice []int) i.RandomAccess {
	return &ints {slice: slice, err:nil}
}

func IntList(v ...int) i.RandomAccess {
	return &ints {slice: v, err:nil}
}

func (s *ints) AtStart() bool {
	return s.pos == 0
}

func (s *ints) AtEnd() bool {
	return s.pos >= len(s.slice)
}

func (s *ints) Next() error {
	if s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos++
	}
	return s.err
}

func (s *ints) Prev() error {
	if s.pos < 0 {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos--
	}
	return s.err
}

func (s *ints) First() error {
	s.pos= 0
	return nil
}

func (s *ints) Last() error {
	s.pos= len(s.slice) - 1
	return nil
}

func (s *ints) Goto(pos int) error {
	s.pos= pos
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	}
	return s.err
}

func (s *ints) Len() int {
	return len(s.slice)
}

func (s *ints) Value() interface{} {
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err = fmt.Errorf("Index out of bounds: %d.", s.pos)
		return nil
	}
	return s.slice[s.pos]
}

func (s *ints) Int() int {
	return s.slice[s.pos]
}

func (s *ints) Error() error {
	return s.err
}

func (s *ints) SetError(err error) {
	s.err= err
}

