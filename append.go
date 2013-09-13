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

package i

import "fmt"

// Append iterator
type iappend struct {
	itrs []Forward
	pos  int
	err  error
}

func Append(itrs ...Forward) Forward {
	return &iappend{itrs: itrs}
}

func (a *iappend) AtEnd() bool {
	if a.pos < len(a.itrs)-1 {
		return false
	} else if a.pos >= len(a.itrs) {
		return true
	}
	return a.itrs[a.pos].AtEnd()
}

func (a *iappend) Next() error {
	if a.pos >= len(a.itrs) {
		a.err = fmt.Errorf("Appending beyond last iterator")
		return a.err
	}
	a.itrs[a.pos].Next()
	for a.itrs[a.pos].AtEnd() {
		a.pos++
		if a.pos >= len(a.itrs) {
			return nil
		}
	}
	return nil
}

func (a *iappend) Value() interface{} {
	if a.AtEnd() {
		a.err = fmt.Errorf("Fetching value beyond last iterator")
		return nil
	}
	return a.itrs[a.pos].Value()
}

func (a *iappend) Error() error {
	if a.err != nil {
		return a.err
	}
	for _, v := range a.itrs {
		err := v.Error()
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *iappend) SetError(err error) {
	a.err = err
}
