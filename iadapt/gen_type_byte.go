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

package iadapt

import "fmt"

// Adapter for byte types on Forward iterators. Provides a typecasting method 
// from interface{} to byte.
// Panics if value is not of byte type.
func (a *ForwardItr) Byte() byte {
	v, ok := a.Value().(byte)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Byte.", a.Value()))
	}
	return v
}

// Adapter for byte types on Forward iterators. Provides a typecasting method 
// from interface{} to byte.
// Returns (value, false) if type is of byte type.
// Returns (default value, true) if type is not of byte type.
func (a *ForwardItr) ByteOr(def byte) (byte, bool) {
	v, ok := a.Value().(byte)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for byte types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to byte.
// Panics if value is not of byte type.
func (a *BiDirectionalItr) Byte() byte {
	v, ok := a.Value().(byte)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Byte.", a.Value()))
	}
	return v
}

// Adapter for byte types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to byte.
// Returns (value, false) if type is of byte type.
// Returns (default value, true) if type is not of byte type.
func (a *BiDirectionalItr) ByteOr(def byte) (byte, bool) {
	v, ok := a.Value().(byte)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for byte types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to byte.
// Panics if value is not of byte type.
func (a *BoundedAtStartItr) Byte() byte {
	v, ok := a.Value().(byte)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Byte.", a.Value()))
	}
	return v
}

// Adapter for byte types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to byte.
// Returns (value, false) if type is of byte type.
// Returns (default value, true) if type is not of byte type.
func (a *BoundedAtStartItr) ByteOr(def byte) (byte, bool) {
	v, ok := a.Value().(byte)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for byte types on Bounded iterators. Provides a typecasting method 
// from interface{} to byte.
// Panics if value is not of byte type.
func (a *BoundedItr) Byte() byte {
	v, ok := a.Value().(byte)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Byte.", a.Value()))
	}
	return v
}

// Adapter for byte types on Bounded iterators. Provides a typecasting method 
// from interface{} to byte.
// Returns (value, false) if type is of byte type.
// Returns (default value, true) if type is not of byte type.
func (a *BoundedItr) ByteOr(def byte) (byte, bool) {
	v, ok := a.Value().(byte)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for byte types on RandomAccess iterator. Provides a typecasting method 
// from interface{} to byte.
// Panics if value is not of byte type.
func (a *RandomAccessItr) Byte() byte {
	v, ok := a.Value().(byte)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Byte.", a.Value()))
	}
	return v
}

// Adapter for byte types on RandomAccess iterators. Provides a typecasting method 
// from interface{} to byte.
// Returns (value, false) if type is of byte type.
// Returns (default value, true) if type is not of byte type.
func (a *RandomAccessItr) ByteOr(def byte) (byte, bool) {
	v, ok := a.Value().(byte)
	if !ok {
		return def, true
	}
	return v, false
}
