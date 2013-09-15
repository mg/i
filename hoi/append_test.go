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

func TestAppend(t *testing.T) {
	list1 := icon.List("one", "two", "three")
	list2 := icon.List(1, 2, 3)
	list3 := icon.List(1.1, 2.2, 3.3)

	i.AssertForward(t, Append(list1, list2, list3), 9, i.Strict)

	i.AssertIteration(t,
		Append(icon.List(1, 2), icon.List("one", "two")),
		1, 2, "one", "two")
}

func ExampleAppend() {
	// Three lists that we wish to append together
	list1 := icon.List("one", "two", "three")
	list2 := icon.List(1, 2, 3)
	list3 := icon.List(1.1, 2.2, 3.3)

	// Define Append iterator
	appendItr := Append(list1, list2, list3)
	for ; !appendItr.AtEnd(); appendItr.Next() {
		// Prints out the list: one, two, three, 1, 2, 3, 1.1, 2.2, 3.3,
		fmt.Printf("%v, ", appendItr.Value())
	}
}
