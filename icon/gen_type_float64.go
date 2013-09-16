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

type float64slice struct {
	slice []float64
	pos  int
	err error
}

// Construct a typed iterator to iterate through a slice of float64 values.
// Accepts a slice of float64 values.
func Float64s(slice []float64) ityped.RandomAccessFloat64 {
	return &float64slice {slice: slice, err:nil}
}

// Construct a typed iterator to iterate through a slice of float64 values.
// Accepts a list of float64 values.
func Float64List(v ...float64) ityped.RandomAccessFloat64 {
	return &float64slice {slice: v, err:nil}
}

func (s *float64slice) AtStart() bool {
	return s.pos == 0
}

func (s *float64slice) AtEnd() bool {
	return s.pos >= len(s.slice)
}

func (s *float64slice) Next() error {
	if s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos++
	}
	return s.err
}

func (s *float64slice) Prev() error {
	if s.pos < 0 {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	} else {
		s.pos--
	}
	return s.err
}

func (s *float64slice) First() error {
	s.pos= 0
	return nil
}

func (s *float64slice) Last() error {
	s.pos= len(s.slice) - 1
	return nil
}

func (s *float64slice) Goto(pos int) error {
	s.pos= pos
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err= fmt.Errorf("Index out of bounds: %d.", s.pos)
	}
	return s.err
}

func (s *float64slice) Len() int {
	return len(s.slice)
}

func (s *float64slice) Value() interface{} {
	if s.pos < 0 || s.pos >= len(s.slice) {
		s.err = fmt.Errorf("Index out of bounds: %d.", s.pos)
		return nil
	}
	return s.slice[s.pos]
}

func (s *float64slice) Float64() float64 {
	return s.slice[s.pos]
}

func (s *float64slice) Error() error {
	return s.err
}

func (s *float64slice) SetError(err error) {
	s.err= err
}

