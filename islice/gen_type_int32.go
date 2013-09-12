// []int32 iterator
package islice

import (
	"fmt"
	"github.com/mg/i"
)	

type int32s struct {
	slice []int32
	pos  int
	err error
}

func Int32s(slice []int32) i.RandomAccess {
	return &int32s {slice: slice, err:nil}
}

func Int32List(v ...int32) i.RandomAccess {
	return &int32s {slice: v, err:nil}
}

func (s *int32s) AtStart() bool {
	return s.pos == 0
}

func (s *int32s) AtEnd() bool {
	return s.pos >= len(s.slice)
}

func (s *int32s) Next() error {
	if s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos++
	}
	return s.err
}

func (s *int32s) Prev() error {
	if s.pos < 0 {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos--
	}
	return s.err
}

func (s *int32s) First() error {
	s.pos= 0
	return nil
}

func (s *int32s) Last() error {
	s.pos= len(s.slice) - 1
	return nil
}

func (s *int32s) Goto(pos int) error {
	s.pos= pos
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	}
	return s.err
}

func (s *int32s) Len() int {
	return len(s.slice)
}

func (s *int32s) Value() interface{} {
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err = fmt.Errorf("Index out of bounds: %d.", s.pos)
		return nil
	}
	return s.slice[s.pos]
}

func (s *int32s) Int32() int32 {
	return s.slice[s.pos]
}

func (s *int32s) Error() error {
	return s.err
}

func (s *int32s) SetError(err error) {
	s.err= err
}

