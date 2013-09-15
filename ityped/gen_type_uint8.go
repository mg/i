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
	// Uint8() method that will return the value as uint8
	Uint8 interface {
		Uint8() uint8
	}

	// Declare forward iterator interface for iterators that 
	// can return a uint8 value
	ForwardUint8 interface {
		i.Forward
		Uint8
	}

	// Declare bidirectional iterator interface for iterators 
	// that can return a uint8 value
	BiDirectionalUint8 interface {
		i.BiDirectional
		Uint8
	}

	// Declare bounded-at-start iterator interface for iterators 
	// that can return a uint8 value
	BoundedAtStartUint8 interface {
		i.BoundedAtStart
		Uint8
	}

	// Declare bounded iterator interface for iterators 
	// that can return a uint8 value
	BoundedUint8 interface {
		i.Bounded
		Uint8
	}

	// Declare random access iterator interface for iterators 
	// that can return a uint8 value
	RandomAccessUint8 interface {
		i.RandomAccess
		Uint8
	}
)
