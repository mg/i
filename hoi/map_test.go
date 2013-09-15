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
	"github.com/mg/i"
	"testing"
)

func mapfunc(itr i.Iterator) interface{} {
	if _, ok := itr.Value().(bool); ok {
		return "bool"
	}
	if _, ok := itr.Value().(int); ok {
		return "int"
	}
	if _, ok := itr.Value().(string); ok {
		return "string"
	}
	if _, ok := itr.Value().(float64); ok {
		return "float64"
	}
	return "unkown"
}

func TestMap(t *testing.T) {
	i.AssertForward(t, Map(mapfunc, List(123, true, "this", 45.4, -1, 1+1i)), 6, i.Strict)
	i.AssertIteration(
		t, Map(mapfunc, List(123, true, "this", 45.4, -1, 1+1i)),
		"int", "bool", "string", "float64", "int", "unkown")
}
