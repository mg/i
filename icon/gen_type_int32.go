// The MIT License (MIT)
//
// Copyright (c) 2013 Magnús Örn Gylfason
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package icon

import (
	"fmt"
	"github.com/mg/i/ityped"
)	

type int32slice struct {
	slice []int32
	pos  int
	err error
}

func Int32s(slice []int32) ityped.RandomAccessInt32 {
	return &int32slice {slice: slice, err:nil}
}

func Int32List(v ...int32) ityped.RandomAccessInt32 {
	return &int32slice {slice: v, err:nil}
}

func (s *int32slice) AtStart() bool {
	return s.pos == 0
}

func (s *int32slice) AtEnd() bool {
	return s.pos >= len(s.slice)
}

func (s *int32slice) Next() error {
	if s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos++
	}
	return s.err
}

func (s *int32slice) Prev() error {
	if s.pos < 0 {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos--
	}
	return s.err
}

func (s *int32slice) First() error {
	s.pos= 0
	return nil
}

func (s *int32slice) Last() error {
	s.pos= len(s.slice) - 1
	return nil
}

func (s *int32slice) Goto(pos int) error {
	s.pos= pos
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	}
	return s.err
}

func (s *int32slice) Len() int {
	return len(s.slice)
}

func (s *int32slice) Value() interface{} {
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err = fmt.Errorf("Index out of bounds: %d.", s.pos)
		return nil
	}
	return s.slice[s.pos]
}

func (s *int32slice) Int32() int32 {
	return s.slice[s.pos]
}

func (s *int32slice) Error() error {
	return s.err
}

func (s *int32slice) SetError(err error) {
	s.err= err
}

