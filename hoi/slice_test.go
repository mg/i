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

func TestSlice(t *testing.T) {
	itesting.AssertRandomAccess(t, Slice(4, 8, icon.List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)), itesting.RelaxValueAtEnd)
	itesting.AssertIteration(
		t, Slice(4, 8, icon.List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)),
		4, 5, 6, 7)
}

func ExampleSlice() {
	// The data stream
	list := icon.List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	// The Slice iterator
	itrSlice := Slice(4, 8, list)

	for ; !itrSlice.AtEnd(); itrSlice.Next() {
		// Produces the list 4, 5, 6, 7,
		fmt.Printf("%v, ", itrSlice.Value())
	}
}
