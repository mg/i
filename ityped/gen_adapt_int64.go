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

import "fmt"

// Adapter for int64 types on Forward iterators. Provides a typecasting method
// from interface{} to int64.
// Panics if value is not of int64 type.
func (a *ForwardItr) Int64() int64 {
	v, ok := a.Value().(int64)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int64.", a.Value()))
	}
	return v
}

// Adapter for int64 types on Forward iterators. Provides a typecasting method
// from interface{} to int64.
// Returns (value, false) if type is of int64 type.
// Returns (default value, true) if type is not of int64 type.
func (a *ForwardItr) Int64Or(def int64) (int64, bool) {
	v, ok := a.Value().(int64)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for int64 types on BiDirectional iterators. Provides a typecasting method
// from interface{} to int64.
// Panics if value is not of int64 type.
func (a *BiDirectionalItr) Int64() int64 {
	v, ok := a.Value().(int64)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int64.", a.Value()))
	}
	return v
}

// Adapter for int64 types on BiDirectional iterators. Provides a typecasting method
// from interface{} to int64.
// Returns (value, false) if type is of int64 type.
// Returns (default value, true) if type is not of int64 type.
func (a *BiDirectionalItr) Int64Or(def int64) (int64, bool) {
	v, ok := a.Value().(int64)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for int64 types on BoundedAtStart iterators. Provides a typecasting method
// from interface{} to int64.
// Panics if value is not of int64 type.
func (a *BoundedAtStartItr) Int64() int64 {
	v, ok := a.Value().(int64)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int64.", a.Value()))
	}
	return v
}

// Adapter for int64 types on BoundedAtStart iterators. Provides a typecasting method
// from interface{} to int64.
// Returns (value, false) if type is of int64 type.
// Returns (default value, true) if type is not of int64 type.
func (a *BoundedAtStartItr) Int64Or(def int64) (int64, bool) {
	v, ok := a.Value().(int64)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for int64 types on Bounded iterators. Provides a typecasting method
// from interface{} to int64.
// Panics if value is not of int64 type.
func (a *BoundedItr) Int64() int64 {
	v, ok := a.Value().(int64)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int64.", a.Value()))
	}
	return v
}

// Adapter for int64 types on Bounded iterators. Provides a typecasting method
// from interface{} to int64.
// Returns (value, false) if type is of int64 type.
// Returns (default value, true) if type is not of int64 type.
func (a *BoundedItr) Int64Or(def int64) (int64, bool) {
	v, ok := a.Value().(int64)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for int64 types on RandomAccess iterator. Provides a typecasting method
// from interface{} to int64.
// Panics if value is not of int64 type.
func (a *RandomAccessItr) Int64() int64 {
	v, ok := a.Value().(int64)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int64.", a.Value()))
	}
	return v
}

// Adapter for int64 types on RandomAccess iterators. Provides a typecasting method
// from interface{} to int64.
// Returns (value, false) if type is of int64 type.
// Returns (default value, true) if type is not of int64 type.
func (a *RandomAccessItr) Int64Or(def int64) (int64, bool) {
	v, ok := a.Value().(int64)
	if !ok {
		return def, true
	}
	return v, false
}
