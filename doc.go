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

// Package i is an higher order library for Go that focuses on using iterators
// as an common interface for computation.
//
// The i library definecs the following interfaces for iterators:
//
//	type Iterator interface {
//		Value() interface{}
//		Error() error
//		SetError(error)
//	}
//
// The basic interface that is common to all iterators. Value() is used to
// access the value at the location that the iterator currently points to
// in the uderlying data stream.
//
// Error() returns the last error occurred while using the iterator.
//
// SetError() is used by iterators to assign an value to the last error.
//
// Both Value() and Error() should always be idempotent.
//
//	type Forward interface {
//		Iterator
//		Next() error
//		AtEnd() bool
//	}
//
// The Forward iterator is an iterator that supports moving forward over
// the underlying data with the Next() method. The iterator should always be
// valid on the range [start, end). The AtEnd() method alows the user to
// check if the end of the data has been reached, that is at the end position.
// It should always be idempotent. Next() should return error if it is called
// beyond the end of the stream.
//
//	type BiDirectional interface {
//		Forward
//		Prev() error
//		AtStart() bool
//	}
//
// The BiDirectional iterator is a Forward iterator that provides a method
// to move backwards in the stream. AtStart() should inidcate when the
// iterator is at the start of the stream. Important difference from the
// AtEnd() method is that the iterator should be valid at the AtStart position
// but not at the AtEnd position.
//
//	type BoundedAtStart interface {
//		Forward
//		First() error
//	}
//
// BoundedAtStart iterator is a Forward iterator that has an efficient method
// to reset the iterator to the starting position. The First() method should
// be idempotent.
//
//	type Bounded interface {
//		BiDirectional
//		First() error
//		Last() error
//	}
//
// The Bounded iterator is a BiDirectional iterator that has efficient ways to
// both reach the first position and the last position. Notice that while the
// the AtStart() method should return true after calling First(), the AtEnd()
// method should return false after calling Last() method, e.g. the last position
// should be a valid position. Calling Last() and then Next() should be valid and
// AtEnd() should return true after that.
//
//	type RandomAccess interface {
//		Bounded
//		Goto(int) error
//		Len() int
//	}
//
// The RandomAccess iterator is a Bounded iterator that both as a known length
// and an efficient way to jump into any valid position in the data.
package i
