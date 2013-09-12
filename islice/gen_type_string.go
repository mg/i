// []string iterator
package islice

import (
	"fmt"
	"github.com/mg/i"
)	

type strings struct {
	slice []string
	pos  int
	err error
}

func Strings(slice []string) i.RandomAccess {
	return &strings {slice: slice, err:nil}
}

func StringList(v ...string) i.RandomAccess {
	return &strings {slice: v, err:nil}
}

func (s *strings) AtStart() bool {
	return s.pos == 0
}

func (s *strings) AtEnd() bool {
	return s.pos >= len(s.slice)
}

func (s *strings) Next() error {
	if s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos++
	}
	return s.err
}

func (s *strings) Prev() error {
	if s.pos < 0 {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos--
	}
	return s.err
}

func (s *strings) First() error {
	s.pos= 0
	return nil
}

func (s *strings) Last() error {
	s.pos= len(s.slice) - 1
	return nil
}

func (s *strings) Goto(pos int) error {
	s.pos= pos
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	}
	return s.err
}

func (s *strings) Len() int {
	return len(s.slice)
}

func (s *strings) Value() interface{} {
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err = fmt.Errorf("Index out of bounds: %d.", s.pos)
		return nil
	}
	return s.slice[s.pos]
}

func (s *strings) String() string {
	return s.slice[s.pos]
}

func (s *strings) Error() error {
	return s.err
}

func (s *strings) SetError(err error) {
	s.err= err
}

