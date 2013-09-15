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
	"testing"
)

func TestInterfaces(t *testing.T) {
	i.AssertRandomAccess(t, Interfaces([]interface{}{1, "string", 3.56, 4}), i.Strict)
	i.AssertIteration(
		t, Interfaces([]interface{}{1, "string", 3.56, 4}),
		1, "string", 3.56, 4)
}

func TestList(t *testing.T) {
	i.AssertRandomAccess(t, List(1, "string", 3.56, 4), i.Strict)
	i.AssertIteration(t, List(1, "string", 3.56, 4), 1, "string", 3.56, 4)
}

func ExampleList() {
	itr := List(1, 2, "three", "four", 5.5, 6.6)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Value())
	}
}
