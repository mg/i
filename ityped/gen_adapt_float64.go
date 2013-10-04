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

// Adapter for float64 types on Forward iterators. Provides a typecasting method
// from interface{} to float64.
// Panics if value is not of float64 type.
func (a *ForwardItr) Float64() float64 {
	v, ok := a.Value().(float64)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Float64.", a.Value()))
	}
	return v
}

// Adapter for float64 types on Forward iterators. Provides a typecasting method
// from interface{} to float64.
// Returns (value, false) if type is of float64 type.
// Returns (default value, true) if type is not of float64 type.
func (a *ForwardItr) Float64Or(def float64) (float64, bool) {
	v, ok := a.Value().(float64)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for float64 types on BiDirectional iterators. Provides a typecasting method
// from interface{} to float64.
// Panics if value is not of float64 type.
func (a *BiDirectionalItr) Float64() float64 {
	v, ok := a.Value().(float64)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Float64.", a.Value()))
	}
	return v
}

// Adapter for float64 types on BiDirectional iterators. Provides a typecasting method
// from interface{} to float64.
// Returns (value, false) if type is of float64 type.
// Returns (default value, true) if type is not of float64 type.
func (a *BiDirectionalItr) Float64Or(def float64) (float64, bool) {
	v, ok := a.Value().(float64)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for float64 types on BoundedAtStart iterators. Provides a typecasting method
// from interface{} to float64.
// Panics if value is not of float64 type.
func (a *BoundedAtStartItr) Float64() float64 {
	v, ok := a.Value().(float64)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Float64.", a.Value()))
	}
	return v
}

// Adapter for float64 types on BoundedAtStart iterators. Provides a typecasting method
// from interface{} to float64.
// Returns (value, false) if type is of float64 type.
// Returns (default value, true) if type is not of float64 type.
func (a *BoundedAtStartItr) Float64Or(def float64) (float64, bool) {
	v, ok := a.Value().(float64)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for float64 types on Bounded iterators. Provides a typecasting method
// from interface{} to float64.
// Panics if value is not of float64 type.
func (a *BoundedItr) Float64() float64 {
	v, ok := a.Value().(float64)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Float64.", a.Value()))
	}
	return v
}

// Adapter for float64 types on Bounded iterators. Provides a typecasting method
// from interface{} to float64.
// Returns (value, false) if type is of float64 type.
// Returns (default value, true) if type is not of float64 type.
func (a *BoundedItr) Float64Or(def float64) (float64, bool) {
	v, ok := a.Value().(float64)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for float64 types on RandomAccess iterator. Provides a typecasting method
// from interface{} to float64.
// Panics if value is not of float64 type.
func (a *RandomAccessItr) Float64() float64 {
	v, ok := a.Value().(float64)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Float64.", a.Value()))
	}
	return v
}

// Adapter for float64 types on RandomAccess iterators. Provides a typecasting method
// from interface{} to float64.
// Returns (value, false) if type is of float64 type.
// Returns (default value, true) if type is not of float64 type.
func (a *RandomAccessItr) Float64Or(def float64) (float64, bool) {
	v, ok := a.Value().(float64)
	if !ok {
		return def, true
	}
	return v, false
}
