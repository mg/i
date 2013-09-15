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
	"math/rand"
	"time"
)

type sample struct {
	itk.WRandomAccess
	rnd *rand.Rand
}

// The Sample iterator will construct an infinite data stream from an underlying
// finite data stream, selecting elements from the underlying in a random
// fashion. The selection is a weak random selection, there is no guarantee that
// item A is not selected twice before item B is selected for the first time.
func Sample(itr i.RandomAccess) i.Forward {
	s := sample{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	s.WRandomAccess = *(itk.WrapRandomAccess(itr))
	s.Next()
	return &s
}

func (s *sample) AtEnd() bool {
	return false
}

func (s *sample) Next() error {
	return s.WRandomAccess.Goto(s.rnd.Intn(s.WRandomAccess.Len()))
}

type shuffle struct {
	itk.WRandomAccess
	rnd      *rand.Rand
	returned []bool
	count    int
}

// The Shuffle iterator will construct an finite data stream from an underlying
// finite data stream, selecting elements from the underlying in a random
// fashion. The selection is a strong random selection, each element will be
// selected only once.
func Shuffle(itr i.RandomAccess) i.Forward {
	s := shuffle{
		rnd:      rand.New(rand.NewSource(time.Now().UnixNano())),
		returned: make([]bool, itr.Len()),
	}
	s.WRandomAccess = *(itk.WrapRandomAccess(itr))
	s.Next()
	return &s
}

func (s *shuffle) AtEnd() bool {
	return s.count > len(s.returned)
}

func (s *shuffle) Next() error {
	if s.AtEnd() {
		s.WRandomAccess.SetError(fmt.Errorf("Shuffle beyond end"))
		return s.WRandomAccess.Error()
	}
	s.count++
	if s.AtEnd() {
		return nil
	}
	idx := s.rnd.Intn(len(s.returned))
	for s.returned[idx] {
		idx = s.rnd.Intn(len(s.returned))
	}
	s.returned[idx] = true
	return s.WRandomAccess.Goto(idx)
}
