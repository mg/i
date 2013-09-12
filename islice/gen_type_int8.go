// []int8 iterator
package islice

import (
	"fmt"
	"github.com/mg/i"
)	

type int8s struct {
	slice []int8
	pos  int
	err error
}

func Int8s(slice []int8) i.RandomAccess {
	return &int8s {slice: slice, err:nil}
}

func Int8List(v ...int8) i.RandomAccess {
	return &int8s {slice: v, err:nil}
}

func (s *int8s) AtStart() bool {
	return s.pos == 0
}

func (s *int8s) AtEnd() bool {
	return s.pos >= len(s.slice)
}

func (s *int8s) Next() error {
	if s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos++
	}
	return s.err
}

func (s *int8s) Prev() error {
	if s.pos < 0 {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos--
	}
	return s.err
}

func (s *int8s) First() error {
	s.pos= 0
	return nil
}

func (s *int8s) Last() error {
	s.pos= len(s.slice) - 1
	return nil
}

func (s *int8s) Goto(pos int) error {
	s.pos= pos
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	}
	return s.err
}

func (s *int8s) Len() int {
	return len(s.slice)
}

func (s *int8s) Value() interface{} {
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err = fmt.Errorf("Index out of bounds: %d.", s.pos)
		return nil
	}
	return s.slice[s.pos]
}

func (s *int8s) Int8() int8 {
	return s.slice[s.pos]
}

func (s *int8s) Error() error {
	return s.err
}

func (s *int8s) SetError(err error) {
	s.err= err
}

