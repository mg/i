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

package iadapt

import (
	"fmt"
	"github.com/mg/i/icon"
	"github.com/mg/i/itesting"
	"testing"
)

func TestAdapt(t *testing.T) {
	list := icon.IntList(1, 2, 3, 4, 5, 6, 7, 8, 9)
	itesting.AssertForward(t, Forward(list), 9, itesting.Strict)
	list.First()
	itesting.AssertIteration(t, Forward(list), 1, 2, 3, 4, 5, 6, 7, 8, 9)
}

func ExampleForwardItr() {
	// Underlying data stream
	list := icon.IntList(1, 2, 3, 4, 5, 6, 7, 8, 9)

	// Create an adapter
	itr := Forward(list)

	// Loop and print
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Printf("%d, ", itr.Int())
	}
}

func ExampleBiDirectionalItr() {
	// Underlying data stream
	list := icon.BoolList(true, true, false, true)

	// Create an adapter
	itr := BiDirectional(list)

	// Loop and print
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Printf("%t, ", itr.Bool())
	}
}

func ExampleBoundedAtStartItr() {
	// Underlying data stream
	list := icon.Float64List(0.1, 1.005, 6.28, 4.233)

	// Create an adapter
	itr := BoundedAtStart(list)

	// Loop and print
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Printf("%f, ", itr.Float64())
	}
}

func ExampleBoundedItr() {
	// Underlying data stream
	list := icon.Complex64List(1+1i, 0.1-4i, 70+35i)

	// Create an adapter
	itr := Bounded(list)

	// Loop and print
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Printf("%f, ", itr.Complex64())
	}
}

func ExampleRandomAccessItr() {
	// Underlying data stream
	list := icon.StringList("aa", "bb", "cc", "dd")

	// Create an adapter
	itr := RandomAccess(list)

	// Loop and print
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Printf("%q, ", itr.String())
	}
}
