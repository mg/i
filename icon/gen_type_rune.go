// []rune iterator
package icon

import (
	"fmt"
	"github.com/mg/i"
)	

type runes struct {
	slice []rune
	pos  int
	err error
}

func Runes(slice []rune) i.RandomAccess {
	return &runes {slice: slice, err:nil}
}

func RuneList(v ...rune) i.RandomAccess {
	return &runes {slice: v, err:nil}
}

func (s *runes) AtStart() bool {
	return s.pos == 0
}

func (s *runes) AtEnd() bool {
	return s.pos >= len(s.slice)
}

func (s *runes) Next() error {
	if s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos++
	}
	return s.err
}

func (s *runes) Prev() error {
	if s.pos < 0 {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos--
	}
	return s.err
}

func (s *runes) First() error {
	s.pos= 0
	return nil
}

func (s *runes) Last() error {
	s.pos= len(s.slice) - 1
	return nil
}

func (s *runes) Goto(pos int) error {
	s.pos= pos
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	}
	return s.err
}

func (s *runes) Len() int {
	return len(s.slice)
}

func (s *runes) Value() interface{} {
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err = fmt.Errorf("Index out of bounds: %d.", s.pos)
		return nil
	}
	return s.slice[s.pos]
}

func (s *runes) Rune() rune {
	return s.slice[s.pos]
}

func (s *runes) Error() error {
	return s.err
}

func (s *runes) SetError(err error) {
	s.err= err
}

