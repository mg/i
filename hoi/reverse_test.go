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
	"github.com/mg/i/icon"
	"github.com/mg/i/itesting"
	"testing"
)

func TestReverse(t *testing.T) {
	itesting.AssertForward(t, Reverse(icon.List(1, 2, 3, 4, 5, 6, 7, 8, 9)), 9, itesting.Strict)
	itesting.AssertIteration(
		t, Reverse(icon.List(1, 2, 3, 4, 5, 6, 7, 8, 9)),
		9, 8, 7, 6, 5, 4, 3, 2, 1)
}

func ExampleReverse() {
	// The data stream
	list := icon.List(1, 2, 3, 4, 5)

	// The reversing iterator
	itrReverse := Reverse(list)
	for ; !itrReverse.AtEnd(); itrReverse.Next() {
		// Prints out: 5, 4, 3, 2, 1,
		fmt.Printf("%v, ", itrReverse.Value())
	}
}
