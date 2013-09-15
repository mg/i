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

func TestSample(t *testing.T) {
	list := icon.List(1, 2, 3, 4, 5, 6, 7, 8, 9)
	itr := Sample(list)
	itesting.AssertForward(t, itr, -1, itesting.Strict)
}

func TestShuffle(t *testing.T) {
	list := icon.List(1, 2, 3, 4, 5, 6, 7, 8, 9)
	itr := Shuffle(list)
	itesting.AssertForward(t, itr, 9, itesting.RelaxValueAtEnd)
}

func ExampleSample() {
	// The underlying data stream
	list := icon.List(1, 2, 3, 4, 5, 6, 7, 8, 9)

	// The Sample iterator
	itrSample := Sample(list)

	fmt.Println(itrSample.Value()) // Prints a random element
	fmt.Println(itrSample.Value()) // Prints the same element
	itrSample.Next()
	fmt.Println(itrSample.Value()) // Prints a random element, might be the same as before
}

func ExampleShuffle() {
	// The underlying data stream
	list := icon.List(1, 2, 3, 4, 5, 6, 7, 8, 9)

	// The Shuffle iterator
	itrShuffle := Shuffle(list)

	for ; !itrShuffle.AtEnd(); itrShuffle.Next() {
		// Prints out the underlying data stream in a random order
		fmt.Printf("%v, ", itrShuffle.Value())
	}
}
