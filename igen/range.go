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

package igen

import (
	"fmt"
	"github.com/mg/i/ityped"
)

type rng struct {
	m, n, cur int
	err       error
}

// The Range iterator generates a range of number from m to n, not including
// n. It is a random access iterator.
func Range(m, n int) ityped.RandomAccessInt {
	return &rng{m: m, n: n, cur: m}
}

func (r *rng) Value() interface{} {
	if r.cur < r.m || r.cur >= r.n {
		r.err = fmt.Errorf("Calling Value() at end")
		return nil
	}
	return r.cur
}

func (r *rng) Int() int {
	return r.cur
}

func (r *rng) Error() error {
	return r.err
}

func (r *rng) SetError(err error) {
	r.err = err
}

func (r *rng) Next() error {
	if r.AtEnd() {
		r.err = fmt.Errorf("Calling Next() at end")
		return r.err
	}
	r.cur++
	return r.err
}

func (r *rng) AtEnd() bool {
	return r.cur >= r.n
}

func (r *rng) Prev() error {
	if r.cur < r.m {
		r.err = fmt.Errorf("Calling Prev() before start")
		return r.err
	}
	r.cur--
	return nil
}

func (r *rng) AtStart() bool {
	return r.cur == r.m
}

func (r *rng) First() error {
	r.cur = r.m
	return nil
}

func (r *rng) Last() error {
	r.cur = r.n
	return nil
}

func (r *rng) Goto(idx int) error {
	if idx < r.m || idx >= r.n {
		r.err = fmt.Errorf("Index %d out of bounds [%d, %d)", idx, r.m, r.n)
		return r.err
	}
	r.cur = idx
	return nil
}

func (r *rng) Len() int {
	return r.n - r.m
}
