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
	"github.com/mg/i/itk"
)

type cycle struct {
	itk.WBoundedAtStart
}

// The Cycle iterator will transform any BoundedAtStart iterator into a
// infinitie data stream that will cycle to the start of the original
// stream when it reaches the end.
func Cycle(itr i.BoundedAtStart) i.Forward {
	c := cycle{}
	c.WBoundedAtStart = *(itk.WrapBoundedAtStart(itr))
	return &c
}

func (c *cycle) Next() error {
	err := c.WBoundedAtStart.Next()
	if c.WBoundedAtStart.AtEnd() {
		return c.WBoundedAtStart.First()
	}
	return err
}

func (c *cycle) AtEnd() bool {
	return false
}
