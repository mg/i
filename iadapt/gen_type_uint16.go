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

// Adapter for uint16 types on Forward iterators. Provides a typecasting method 
// from interface{} to uint16.
// Panics if value is not of uint16 type.
func (a *ForwardItr) Uint16() uint16 {
	v, ok := a.Value().(uint16)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint16.", a.Value()))
	}
	return v
}

// Adapter for uint16 types on Forward iterators. Provides a typecasting method 
// from interface{} to uint16.
// Returns (value, false) if type is of uint16 type.
// Returns (default value, true) if type is not of uint16 type.
func (a *ForwardItr) Uint16Or(def uint16) (uint16, bool) {
	v, ok := a.Value().(uint16)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for uint16 types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to uint16.
// Panics if value is not of uint16 type.
func (a *BiDirectionalItr) Uint16() uint16 {
	v, ok := a.Value().(uint16)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint16.", a.Value()))
	}
	return v
}

// Adapter for uint16 types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to uint16.
// Returns (value, false) if type is of uint16 type.
// Returns (default value, true) if type is not of uint16 type.
func (a *BiDirectionalItr) Uint16Or(def uint16) (uint16, bool) {
	v, ok := a.Value().(uint16)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for uint16 types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to uint16.
// Panics if value is not of uint16 type.
func (a *BoundedAtStartItr) Uint16() uint16 {
	v, ok := a.Value().(uint16)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint16.", a.Value()))
	}
	return v
}

// Adapter for uint16 types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to uint16.
// Returns (value, false) if type is of uint16 type.
// Returns (default value, true) if type is not of uint16 type.
func (a *BoundedAtStartItr) Uint16Or(def uint16) (uint16, bool) {
	v, ok := a.Value().(uint16)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for uint16 types on Bounded iterators. Provides a typecasting method 
// from interface{} to uint16.
// Panics if value is not of uint16 type.
func (a *BoundedItr) Uint16() uint16 {
	v, ok := a.Value().(uint16)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint16.", a.Value()))
	}
	return v
}

// Adapter for uint16 types on Bounded iterators. Provides a typecasting method 
// from interface{} to uint16.
// Returns (value, false) if type is of uint16 type.
// Returns (default value, true) if type is not of uint16 type.
func (a *BoundedItr) Uint16Or(def uint16) (uint16, bool) {
	v, ok := a.Value().(uint16)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for uint16 types on RandomAccess iterator. Provides a typecasting method 
// from interface{} to uint16.
// Panics if value is not of uint16 type.
func (a *RandomAccessItr) Uint16() uint16 {
	v, ok := a.Value().(uint16)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint16.", a.Value()))
	}
	return v
}

// Adapter for uint16 types on RandomAccess iterators. Provides a typecasting method 
// from interface{} to uint16.
// Returns (value, false) if type is of uint16 type.
// Returns (default value, true) if type is not of uint16 type.
func (a *RandomAccessItr) Uint16Or(def uint16) (uint16, bool) {
	v, ok := a.Value().(uint16)
	if !ok {
		return def, true
	}
	return v, false
}
