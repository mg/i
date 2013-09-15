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
	// Uint64() method that will return the value as uint64
	Uint64 interface {
		Uint64() uint64
	}

	// Declare forward iterator interface for iterators that 
	// can return a uint64 value
	ForwardUint64 interface {
		i.Forward
		Uint64
	}

	// Declare bidirectional iterator interface for iterators 
	// that can return a uint64 value
	BiDirectionalUint64 interface {
		i.BiDirectional
		Uint64
	}

	// Declare bounded-at-start iterator interface for iterators 
	// that can return a uint64 value
	BoundedAtStartUint64 interface {
		i.BoundedAtStart
		Uint64
	}

	// Declare bounded iterator interface for iterators 
	// that can return a uint64 value
	BoundedUint64 interface {
		i.Bounded
		Uint64
	}

	// Declare random access iterator interface for iterators 
	// that can return a uint64 value
	RandomAccessUint64 interface {
		i.RandomAccess
		Uint64
	}
)
