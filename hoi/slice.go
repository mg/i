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

package hoi

import (
	"fmt"
	"github.com/mg/i"
	"github.com/mg/i/itk"
)

type slice struct {
	itk.WRandomAccess
	start, end, cur int
}

// The Slice iterator will give access to a subset of an underlying
// data stream. Given the stream [1,2,3,4,5,6,7,8], the Slice(2,6,stream)
// iterator will give access to the stream [3,4,5,6].
func Slice(start, end int, itr i.RandomAccess) i.RandomAccess {
	s := slice{start: start, cur: start, end: end}
	s.WRandomAccess = *(itk.WrapRandomAccess(itr))
	s.Goto(0)
	return &s
}

func (s *slice) Next() error {
	if s.AtEnd() {
		s.WRandomAccess.SetError(fmt.Errorf("Calling Next() at end"))
		return s.WRandomAccess.Error()
	}
	s.cur++
	return s.WRandomAccess.Next()
}

func (s *slice) AtEnd() bool {
	return s.cur >= s.end || s.WRandomAccess.AtEnd()
}

func (s *slice) Prev() error {
	if s.cur < s.start-1 {
		s.WRandomAccess.SetError(fmt.Errorf("Calling Prev() before start"))
		return s.WRandomAccess.Error()
	}
	s.cur--
	return s.WRandomAccess.Prev()
}

func (s *slice) AtStart() bool {
	return s.cur == s.start || s.WRandomAccess.AtStart()
}

func (s *slice) First() error {
	s.cur = s.start
	return s.WRandomAccess.First()
}

func (s *slice) Last() error {
	s.cur = s.end - 1
	return s.WRandomAccess.Last()
}

func (s *slice) Goto(idx int) error {
	if idx+s.start < s.start || idx+s.start >= s.end {
		s.WRandomAccess.SetError(fmt.Errorf("Indx out of bounds: %d <= %d < %d", s.start, s.start+idx, s.end))
		return s.WRandomAccess.Error()
	}
	return s.WRandomAccess.Goto(idx + s.start)
}

func (s *slice) Len() int {
	return s.end - s.start
}
