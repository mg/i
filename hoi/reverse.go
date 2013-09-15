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
	"github.com/mg/i"
	"github.com/mg/i/itk"
)

type reverse struct {
	itk.WBounded
	atEnd bool
}

// The Reverse iterator will construct an iterator that will loop over the
// data stream in reverse order. Given the data stream [1,2,3,4], the
// iterator will provide the data stream [4,3,2,1]
func Reverse(itr i.Bounded) i.Forward {
	r := reverse{}
	r.WBounded = *(itk.WrapBounded(itr))
	r.WBounded.Last()
	return &r
}

func (r *reverse) AtEnd() bool {
	return r.atEnd
}

func (r *reverse) Next() error {
	if r.WBounded.AtStart() {
		r.atEnd = true
	}
	return r.WBounded.Prev()
}
