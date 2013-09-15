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

func TestZip(t *testing.T) {
	itr := Zip(
		icon.List(1, 2, 3, 4, 5, 6),
		icon.List(6.4, 7.1, 8.2, 9.9),
		icon.List("A", "B", "C", "D", "E"))
	itesting.AssertForward(t, itr, 4, itesting.RelaxValueEqual)
}

func TestZipLongest(t *testing.T) {
	itesting.AssertForward(t, ZipLongest(
		icon.List(1, 2, 3, 4, 5, 6),
		icon.List(6.4, 7.1, 8.2, 9.9),
		icon.List("A", "B", "C", "D", "E")),
		6, itesting.RelaxValueEqual)
}

func ExampleZip() {
	// The data streams
	list1 := icon.List(1, 2, 3)
	list2 := icon.List("one", "two", "three")
	list3 := icon.List(10, 11, 12, 13)

	// The Zip iterator
	itrZip := Zip(list1, list2, list3)
	for ; !itrZip.AtEnd(); itrZip.Next() {
		// Prints out: [[1,"one",10], [2,"two",11], [3,"three",12]]
		fmt.Printf("%v, ", itrZip.Value())
	}
}

func ExampleZipLongest() {
	// The data streams
	list1 := icon.List(1, 2, 3)
	list2 := icon.List("one", "two", "three")
	list3 := icon.List(10, 11, 12, 13)

	// The Zip iterator
	itrZipLongest := ZipLongest(list1, list2, list3)
	for ; !itrZipLongest.AtEnd(); itrZipLongest.Next() {
		// Prints out: [[1,"one",10], [2,"two",11], [3,"three",12], [nil,nil,13]]
		fmt.Printf("%v, ", itrZipLongest.Value())
	}
}
