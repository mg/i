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

// []interface{} iterator
package i

import (
	"fmt"
)

type iinterfaces struct {
	slice []interface{}
	pos   int
	err   error
}

func Interfaces(slice []interface{}) RandomAccess {
	return &iinterfaces{slice: slice, err: nil}
}

func List(list ...interface{}) RandomAccess {
	return &iinterfaces{slice: list, err: nil}
}

func (i *iinterfaces) AtStart() bool {
	return i.pos == 0
}

func (i *iinterfaces) AtEnd() bool {
	return i.pos >= len(i.slice)
}

func (i *iinterfaces) Next() error {
	if i.pos >= len(i.slice) {
		i.err = fmt.Errorf("Next: Index out of bounds: %d.", i.pos)
	} else {
		i.pos++
	}
	return i.err
}

func (i *iinterfaces) Prev() error {
	if i.pos < 0 {
		i.err = fmt.Errorf("Prev: Index out of bounds: %d.", i.pos)
	} else {
		i.pos--
	}
	return i.err
}

func (i *iinterfaces) First() error {
	i.pos = 0
	return nil
}

func (i *iinterfaces) Last() error {
	i.pos = len(i.slice) - 1
	return nil
}

func (i *iinterfaces) Goto(pos int) error {
	i.pos = pos
	if i.pos < 0 || i.pos >= len(i.slice) {
		i.err = fmt.Errorf("Goto: Index out of bounds: %d.", i.pos)
	}
	return i.err
}

func (i *iinterfaces) Len() int {
	return len(i.slice)
}

func (i *iinterfaces) Value() interface{} {
	if i.pos < 0 || i.pos >= len(i.slice) {
		i.err = fmt.Errorf("Value: Index out of bounds: %d.", i.pos)
		return nil
	}
	return i.slice[i.pos]
}

func (i *iinterfaces) Error() error {
	return i.err
}

func (i *iinterfaces) SetError(err error) {
	i.err = err
}
