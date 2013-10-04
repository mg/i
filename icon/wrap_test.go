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
	"github.com/mg/i"
	"github.com/mg/i/itesting"
	"github.com/mg/i/itk"
	"testing"
)

type wrap1s struct {
	itk.WRandomAccess
}

func wrap1(itr i.RandomAccess) i.RandomAccess {
	var w wrap1s
	w.WRandomAccess = *(itk.WrapRandomAccess(itr))
	return &w
}

func (w *wrap1s) Value() interface{} {
	i, _ := w.WRandomAccess.Value().(int)
	return i + 1
}

type wrap2s struct {
	itk.WRandomAccess
}

func wrap2(itr i.RandomAccess) i.RandomAccess {
	var w wrap2s
	w.WRandomAccess = *(itk.WrapRandomAccess(itr))
	return &w
}

func (w *wrap2s) Next() error {
	return w.WRandomAccess.Next()
}

func TestWrapForward(t *testing.T) {
	itesting.AssertRandomAccess(t, wrap1(List(1, 2, 3, 4, 5)), itesting.Strict)
	itesting.AssertRandomAccess(t, wrap2(List(1, 2, 3, 4, 5)), itesting.Strict)
	itesting.AssertRandomAccess(t, wrap2(wrap1(wrap1(List(1, 2, 3, 4, 5)))), itesting.Strict)
	itesting.AssertIteration(
		t, wrap2(List(1, 2, 3, 4, 5)),
		1, 2, 3, 4, 5)
	itesting.AssertIteration(
		t, wrap1(wrap1(List(1, 2, 3, 4, 5))),
		3, 4, 5, 6, 7)
	itesting.AssertIteration(
		t, wrap1(wrap2(wrap1(wrap1(List(1, 2, 3, 4, 5))))),
		4, 5, 6, 7, 8)
}
