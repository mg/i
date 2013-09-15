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

// The FoldFunc will provide the folding action that the Foldl and Foldr
// iterators will apply to a data stream.
type FoldFunc func(interface{}, i.Iterator) interface{}

type foldl struct {
	itk.WForward
	ff  FoldFunc
	val interface{}
	err error
}

// Fold a data stream from the left. If the providing iterator gives access
// to the datastream [1,2,3,4,5] the Foldl iterator will apply the folding
// function to the data stream in the following way (provided the iteration
// is completed):
// f(f(f(f(f(inital value, 1), 2), 3), 4), 5)
func Foldl(ff FoldFunc, init interface{}, itr i.Forward) i.Forward {
	f := foldl{ff: ff}
	f.WForward = *(itk.WrapForward(itr))
	f.val = f.ff(init, &f.WForward)
	return &f
}

func (f *foldl) Next() error {
	f.err = f.WForward.Next()
	if f.err == nil && !f.WForward.AtEnd() {
		f.val = f.ff(f.val, &f.WForward)
	}
	return f.err
}

func (f *foldl) Value() interface{} {
	if f.WForward.AtEnd() {
		f.err = fmt.Errorf("Calling value at end")
		return nil
	}
	return f.val
}

func (f *foldl) Error() error {
	if f.err != nil {
		return f.err
	}
	return f.WForward.Error()
}

type foldr struct {
	ff    FoldFunc
	itr   i.Bounded
	val   interface{}
	err   error
	atEnd bool
}

// Fold a data stream from the right. If the providing iterator gives access
// to the datastream [1,2,3,4,5] the Foldr iterator will apply the folding
// function to the data stream in the following way (provided the iteration
// is completed):
// f(f(f(f(f(inital value, 5), 4), 3), 2), 1)
func Foldr(ff FoldFunc, init interface{}, itr i.Bounded) i.Forward {
	f := foldr{ff: ff, itr: itr}
	itr.Last()
	f.val = f.ff(init, f.itr)
	return &f
}

func (f *foldr) AtEnd() bool {
	return f.atEnd
}

func (f *foldr) Next() error {
	if f.itr.AtStart() {
		f.atEnd = true
	}
	f.err = f.itr.Prev()
	if f.err == nil && !f.atEnd {
		f.val = f.ff(f.val, f.itr)

	}
	return f.err
}

func (f *foldr) Value() interface{} {
	if f.AtEnd() {
		f.err = fmt.Errorf("Calling value at end")
		return nil
	}
	return f.val
}

func (f *foldr) Error() error {
	if f.err != nil {
		return f.err
	}
	return f.itr.Error()
}

func (f *foldr) SetError(err error) {
	f.err = err
}
