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

package i

import (
	"testing"
)

func TestWrapForward(t *testing.T) {
	list := List(1, 2, 3, 4, 5)
	AssertForward(t, WrapForward(WrapForward(list)), 5, Strict)
	list.SetError(nil)
	list.First()
	AssertIteration(
		t, WrapForward(WrapForward(list)),
		1, 2, 3, 4, 5)

	list.SetError(nil)
	list.First()
	w1 := Filter(func(itr Iterator) bool {
		v, _ := itr.Value().(int)
		return v < 4
	}, list)
	w2 := Map(func(itr Iterator) interface{} {
		v, _ := itr.Value().(int)
		return v * 2
	}, w1)
	AssertForward(t, w2, 3, Strict)
	list.SetError(nil)
	list.First()
	w1 = Filter(func(itr Iterator) bool {
		v, _ := itr.Value().(int)
		return v < 4
	}, list)
	w2 = Map(func(itr Iterator) interface{} {
		v, _ := itr.Value().(int)
		return v * 2
	}, w1)
	AssertIteration(t, w2, 2, 4, 6)
}
