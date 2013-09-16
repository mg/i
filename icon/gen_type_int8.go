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

type int8slice struct {
	slice []int8
	pos  int
	err error
}

// Construct a typed iterator to iterate through a slice of int8 values.
// Accepts a slice of int8 values.
func Int8s(slice []int8) ityped.RandomAccessInt8 {
	return &int8slice {slice: slice, err:nil}
}

// Construct a typed iterator to iterate through a slice of int8 values.
// Accepts a list of int8 values.
func Int8List(v ...int8) ityped.RandomAccessInt8 {
	return &int8slice {slice: v, err:nil}
}

func (s *int8slice) AtStart() bool {
	return s.pos == 0
}

func (s *int8slice) AtEnd() bool {
	return s.pos >= len(s.slice)
}

func (s *int8slice) Next() error {
	if s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos++
	}
	return s.err
}

func (s *int8slice) Prev() error {
	if s.pos < 0 {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos--
	}
	return s.err
}

func (s *int8slice) First() error {
	s.pos= 0
	return nil
}

func (s *int8slice) Last() error {
	s.pos= len(s.slice) - 1
	return nil
}

func (s *int8slice) Goto(pos int) error {
	s.pos= pos
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	}
	return s.err
}

func (s *int8slice) Len() int {
	return len(s.slice)
}

func (s *int8slice) Value() interface{} {
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err = fmt.Errorf("Index out of bounds: %d.", s.pos)
		return nil
	}
	return s.slice[s.pos]
}

func (s *int8slice) Int8() int8 {
	return s.slice[s.pos]
}

func (s *int8slice) Error() error {
	return s.err
}

func (s *int8slice) SetError(err error) {
	s.err= err
}

