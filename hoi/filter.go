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
	"github.com/mg/i/itk"
)

// The FilterFunc is a function that will be called on each item in the
// data stream by the Filter iterator. If the function returns true, the item will be included
// in the resulting stream, otherwise it will be omitted.
type FilterFunc func(i.Iterator) bool

type filter struct {
	itk.WForward
	ff FilterFunc
}

// Constructs a filtering iterator from a filtering function and a iterator
// providing access to the data stream that we wish to filter.
func Filter(ff FilterFunc, itr i.Forward) i.Forward {
	f := filter{ff: ff}
	f.WForward = *(itk.WrapForward(itr))
	return &f
}

func (f *filter) Next() error {
	if f.WForward.AtEnd() {
		f.WForward.SetError(fmt.Errorf("Calling Next() after end"))
		return f.WForward.Error()
	}
	for !f.WForward.AtEnd() {
		f.WForward.Next()
		if !f.WForward.AtEnd() && f.ff(&f.WForward) {
			return nil
		}
	}
	return f.WForward.Error()
}
