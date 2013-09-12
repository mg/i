// []byte iterator
package islice

import (
	"fmt"
	"github.com/mg/i"
)	

type bytes struct {
	slice []byte
	pos  int
	err error
}

func Bytes(slice []byte) i.RandomAccess {
	return &bytes {slice: slice, err:nil}
}

func ByteList(v ...byte) i.RandomAccess {
	return &bytes {slice: v, err:nil}
}

func (s *bytes) AtStart() bool {
	return s.pos == 0
}

func (s *bytes) AtEnd() bool {
	return s.pos >= len(s.slice)
}

func (s *bytes) Next() error {
	if s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos++
	}
	return s.err
}

func (s *bytes) Prev() error {
	if s.pos < 0 {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos--
	}
	return s.err
}

func (s *bytes) First() error {
	s.pos= 0
	return nil
}

func (s *bytes) Last() error {
	s.pos= len(s.slice) - 1
	return nil
}

func (s *bytes) Goto(pos int) error {
	s.pos= pos
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	}
	return s.err
}

func (s *bytes) Len() int {
	return len(s.slice)
}

func (s *bytes) Value() interface{} {
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err = fmt.Errorf("Index out of bounds: %d.", s.pos)
		return nil
	}
	return s.slice[s.pos]
}

func (s *bytes) Byte() byte {
	return s.slice[s.pos]
}

func (s *bytes) Error() error {
	return s.err
}

func (s *bytes) SetError(err error) {
	s.err= err
}

