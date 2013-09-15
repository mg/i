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
	"testing"
)

func TestCycle(t *testing.T) {
	list := icon.List(0, 1, 2)

	i.AssertForward(t, Cycle(list), -1, i.Strict)
	list.First()
	i.AssertIteration(t, Cycle(list), 0, 1, 2, 0, 1, 2, 0, 1, 2)
}

func ExampleCycle() {
	// The underlying data stream
	list := icon.List(1, 2, 3)

	// The Cycle iterator
	itrCycle := Cycle(list)
	for i := 0; i < 10; i++ {
		// Prints out the list: 1, 2, 3, 1, 2, 3, 1, 2, 3, 1,
		fmt.Printf("%v, ", itrCycle.Value())
		itrCycle.Next()
	}

}
