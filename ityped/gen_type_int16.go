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

package ityped

import (
	"github.com/mg/i"
)	

type (
	// Declare interface for an iterator that has an 
	// Int16() method that will return the value as int16
	Int16 interface {
		Int16() int16
	}

	// Declare forward iterator interface for iterators that 
	// can return a int16 value
	ForwardInt16 interface {
		i.Forward
		Int16
	}

	// Declare bidirectional iterator interface for iterators 
	// that can return a int16 value
	BiDirectionalInt16 interface {
		i.BiDirectional
		Int16
	}

	// Declare bounded-at-start iterator interface for iterators 
	// that can return a int16 value
	BoundedAtStartInt16 interface {
		i.BoundedAtStart
		Int16
	}

	// Declare bounded iterator interface for iterators 
	// that can return a int16 value
	BoundedInt16 interface {
		i.Bounded
		Int16
	}

	// Declare random access iterator interface for iterators 
	// that can return a int16 value
	RandomAccessInt16 interface {
		i.RandomAccess
		Int16
	}
)
