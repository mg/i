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

package icon

import (
	"fmt"
	"github.com/mg/i"
)

type iinterfaces struct {
	slice []interface{}
	pos   int
	err   error
}

// Create a RandomAccess iterator from a []interface{} type.
func Interfaces(slice []interface{}) i.RandomAccess {
	return &iinterfaces{slice: slice, err: nil}
}

// Create a RandomAccess iterator from a list of elements.
func List(list ...interface{}) i.RandomAccess {
	return &iinterfaces{slice: list, err: nil}
}

// Check if the iterator is at the start of the list.
func (i *iinterfaces) AtStart() bool {
	return i.pos == 0
}

// Check if the iterator is beyond the end of the list.
func (i *iinterfaces) AtEnd() bool {
	return i.pos >= len(i.slice)
}

// Move the iterator to the next position in the list.
func (i *iinterfaces) Next() error {
	if i.pos >= len(i.slice) {
		i.err = fmt.Errorf("Next: Index out of bounds: %d.", i.pos)
	} else {
		i.pos++
	}
	return i.err
}

// Move the iterator to the previous position in the list.
func (i *iinterfaces) Prev() error {
	if i.pos < 0 {
		i.err = fmt.Errorf("Prev: Index out of bounds: %d.", i.pos)
	} else {
		i.pos--
	}
	return i.err
}

// Jump to the first position in the list.
func (i *iinterfaces) First() error {
	i.pos = 0
	return nil
}

// Jump to the last position in the list.
func (i *iinterfaces) Last() error {
	i.pos = len(i.slice) - 1
	return nil
}

// Jump to position pos in the list.
func (i *iinterfaces) Goto(pos int) error {
	i.pos = pos
	if i.pos < 0 || i.pos >= len(i.slice) {
		i.err = fmt.Errorf("Goto: Index out of bounds: %d.", i.pos)
	}
	return i.err
}

// Return the length of the list.
func (i *iinterfaces) Len() int {
	return len(i.slice)
}

// Return the value at the current position in the list.
func (i *iinterfaces) Value() interface{} {
	if i.pos < 0 || i.pos >= len(i.slice) {
		i.err = fmt.Errorf("Value: Index out of bounds: %d.", i.pos)
		return nil
	}
	return i.slice[i.pos]
}

// Return last error occurred.
func (i *iinterfaces) Error() error {
	return i.err
}

// Set the error to err.
func (i *iinterfaces) SetError(err error) {
	i.err = err
}
