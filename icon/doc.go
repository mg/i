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

// Package icon provides iterators that give access to containers such as slices.
// Each slices iterator is a typed variand supporting the corresponding typed
// interface in ityped, e.g. the iterator for boolean slices returns the
// ityped.RandomAccessBool interface, which in addition to the general
// i.RandomAccess interface supports the Bool() bool method, which returns
// the current value as a boolean type.
//
// The following typed slices are supported: bool, byte, complex64, complex128,
// float32, float64, int, int8, int16, int32, int64, rune, string, uint, uint8,
// uint16, uint32, uint64.
//
// The typed iterators for each type can be constructed in two ways, eitther
// from a typed slice (e.g. Bool([]bool{ boolean values...})) or with a variable
// argument function (e.g. BoolList(boolean values...)).
//
// The iterator for the generic interface{} can be constructed from a slice of
// interfaces with the Interfaces([]interface{interface values...}) function or
// the List function (List(values)).
package icon
