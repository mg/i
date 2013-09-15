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
	"github.com/mg/i/icon"
	"github.com/mg/i/itesting"
	"math"
	"testing"
)

func filterfunc(itr i.Iterator) bool {
	if _, ok := itr.Value().(bool); ok {
		return true
	}
	if _, ok := itr.Value().(int); ok {
		return true
	}
	if _, ok := itr.Value().(string); ok {
		return false
	}
	if _, ok := itr.Value().(float64); ok {
		return true
	}
	return false
}

func TestFilter(t *testing.T) {
	itesting.AssertForward(t, Filter(filterfunc, icon.List(123, true, "this", 45.4, -1, 1+1i)), 4, itesting.Strict)
	itesting.AssertIteration(
		t, Filter(filterfunc, icon.List(123, true, "ssss", 45.4, -1, 1+1i)),
		123, true, 45.4, -1)
}

func ExampleFilter() {
	// The underlying data stream
	list := icon.List(1, 2, 3, 4, 5, 6, 7, 8, 9)

	// The filtering function will filter out odd numbers
	filterFunc := func(itr i.Iterator) bool {
		v, _ := itr.Value().(int)
		return math.Mod(float64(v), 2) == 0
	}

	// The Filter iterator
	itrFilter := Filter(filterFunc, list)
	for ; !itrFilter.AtEnd(); itrFilter.Next() {
		// Prints out the list: 2, 4, 6, 8
		fmt.Printf("%v, ", itrFilter.Value())
	}

}
