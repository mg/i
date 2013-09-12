// []int64 iterator
package islice

import (
	"fmt"
	"github.com/mg/i"
)	

type int64s struct {
	slice []int64
	pos  int
	err error
}

func Int64s(slice []int64) i.RandomAccess {
	return &int64s {slice: slice, err:nil}
}

func Int64List(v ...int64) i.RandomAccess {
	return &int64s {slice: v, err:nil}
}

func (s *int64s) AtStart() bool {
	return s.pos == 0
}

func (s *int64s) AtEnd() bool {
	return s.pos >= len(s.slice)
}

func (s *int64s) Next() error {
	if s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos++
	}
	return s.err
}

func (s *int64s) Prev() error {
	if s.pos < 0 {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos--
	}
	return s.err
}

func (s *int64s) First() error {
	s.pos= 0
	return nil
}

func (s *int64s) Last() error {
	s.pos= len(s.slice) - 1
	return nil
}

func (s *int64s) Goto(pos int) error {
	s.pos= pos
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	}
	return s.err
}

func (s *int64s) Len() int {
	return len(s.slice)
}

func (s *int64s) Value() interface{} {
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err = fmt.Errorf("Index out of bounds: %d.", s.pos)
		return nil
	}
	return s.slice[s.pos]
}

func (s *int64s) Int64() int64 {
	return s.slice[s.pos]
}

func (s *int64s) Error() error {
	return s.err
}

func (s *int64s) SetError(err error) {
	s.err= err
}

