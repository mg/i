// []bool iterator
package icon

import (
	"fmt"
	"github.com/mg/i"
)	

type bools struct {
	slice []bool
	pos  int
	err error
}

func Bools(slice []bool) i.RandomAccess {
	return &bools {slice: slice, err:nil}
}

func BoolList(v ...bool) i.RandomAccess {
	return &bools {slice: v, err:nil}
}

func (s *bools) AtStart() bool {
	return s.pos == 0
}

func (s *bools) AtEnd() bool {
	return s.pos >= len(s.slice)
}

func (s *bools) Next() error {
	if s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos++
	}
	return s.err
}

func (s *bools) Prev() error {
	if s.pos < 0 {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos--
	}
	return s.err
}

func (s *bools) First() error {
	s.pos= 0
	return nil
}

func (s *bools) Last() error {
	s.pos= len(s.slice) - 1
	return nil
}

func (s *bools) Goto(pos int) error {
	s.pos= pos
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	}
	return s.err
}

func (s *bools) Len() int {
	return len(s.slice)
}

func (s *bools) Value() interface{} {
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err = fmt.Errorf("Index out of bounds: %d.", s.pos)
		return nil
	}
	return s.slice[s.pos]
}

func (s *bools) Bool() bool {
	return s.slice[s.pos]
}

func (s *bools) Error() error {
	return s.err
}

func (s *bools) SetError(err error) {
	s.err= err
}

