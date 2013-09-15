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

type repeat struct {
	itk.WForward
	cur, count int
}

// The Repeat iteator will iterate over a data stream and repeat each
// item in the list. E.g give the data stream [1,2,3] the Count(3, stream)
// will result in the data stream [1, 1, 1, 2, 2, 2, 3, 3, 3].
func Repeat(count int, itr i.Forward) i.Forward {
	r := repeat{count: count}
	r.WForward = *(itk.WrapForward(itr))
	return &r
}

func (r *repeat) Next() error {
	if r.WForward.AtEnd() {
		return r.WForward.Next()
	}
	r.cur++
	if r.cur >= r.count {
		r.cur = 0
		return r.WForward.Next()
	}
	return nil
}
