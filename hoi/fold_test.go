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
	"testing"
)

func foldfunc(v interface{}, i i.Iterator) interface{} {
	return v.(int) + i.Value().(int)
}

func TestFoldl(t *testing.T) {
	itesting.AssertForward(t, Foldl(foldfunc, 1, icon.List(1, 2, 3, 4, 5, 6, 7, 8, 9)), 9, itesting.Strict)
	itesting.AssertIteration(
		t, Foldl(foldfunc, 1, icon.List(1, 2, 3, 4, 5, 6, 7, 8, 9)),
		2, 4, 7, 11, 16, 22, 29, 37, 46)
}

func TestFoldr(t *testing.T) {
	itesting.AssertForward(t, Foldr(foldfunc, 1, icon.List(1, 2, 3, 4, 5, 6, 7, 8, 9)), 9, itesting.Strict)
	itesting.AssertIteration(
		t, Foldr(foldfunc, 1, icon.List(1, 2, 3, 4, 5, 6, 7, 8, 9)),
		10, 18, 25, 31, 36, 40, 43, 45, 46)
}

func ExampleFoldl() {
	// The underying data stream
	list := icon.List(1, 2, 3, 4, 5)

	// The folding function
	foldFunc := func(v interface{}, i i.Iterator) interface{} {
		// Fold the two values into the sum of the two values
		return v.(int) + i.Value().(int)
	}

	// The Foldl iteator
	itrFoldl := Foldl(foldFunc, 1, list)
	for ; !itrFoldl.AtEnd(); itrFoldl.Next() {
		// Prints: 2, 4, 7, 11, 16
		fmt.Printf("%v, ", itrFoldl.Value())
	}
}

func ExampleFoldr() {
	// The underying data stream
	list := icon.List(1, 2, 3, 4, 5)

	// The folding function
	foldFunc := func(v interface{}, i i.Iterator) interface{} {
		// Fold the two values into the sum of the two values
		return v.(int) + i.Value().(int)
	}

	// The Foldr iteator
	itrFoldr := Foldr(foldFunc, 1, list)
	for ; !itrFoldr.AtEnd(); itrFoldr.Next() {
		// Prints: 6, 10, 13, 15, 16
		fmt.Printf("%v, ", itrFoldr.Value())
	}
}
