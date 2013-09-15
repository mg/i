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
)

type count struct {
	cur, max int
	itr      i.Forward
	err      error
}

// Given the data stream [1,2,3,4,5,6,7,8,...] the Count(5, stream) iterator
// will result n the data stream [1,2,3,4,5]
func Count(max int, itr i.Forward) i.Forward {
	return &count{max: max, itr: itr}
}

func (c *count) AtEnd() bool {
	return c.itr.AtEnd() || c.cur >= c.max
}

func (c *count) Next() error {
	if c.AtEnd() {
		c.err = fmt.Errorf("Calling Next after end")
		return c.err
	}
	c.cur++
	c.err = c.itr.Next()
	return c.err
}

func (c *count) Value() interface{} {
	if c.AtEnd() {
		c.err = fmt.Errorf("Calling Value after end")
		return nil
	}
	return c.itr.Value()
}

func (c *count) Error() error {
	return c.err
}

func (c *count) SetError(err error) {
	c.err = err
}
