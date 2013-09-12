//  Generated file for Float32 adapter
package iadapt

import "fmt"

// Adapter for float32 types on Forward iterators. Provides a typecasting method 
// from interface{} to float32.
// Panics if value is not of float32 type.
func (a *ForwardItr) Float32() float32 {
	v, ok := a.Value().(float32)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Float32.", a.Value()))
	}
	return v
}

// Adapter for float32 types on Forward iterators. Provides a typecasting method 
// from interface{} to float32.
// Returns (value, false) if type is of float32 type.
// Returns (default value, true) if type is not of float32 type.
func (a *ForwardItr) Float32Or(def float32) (float32, bool) {
	v, ok := a.Value().(float32)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for float32 types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to float32.
// Panics if value is not of float32 type.
func (a *BiDirectionalItr) Float32() float32 {
	v, ok := a.Value().(float32)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Float32.", a.Value()))
	}
	return v
}

// Adapter for float32 types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to float32.
// Returns (value, false) if type is of float32 type.
// Returns (default value, true) if type is not of float32 type.
func (a *BiDirectionalItr) Float32Or(def float32) (float32, bool) {
	v, ok := a.Value().(float32)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for float32 types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to float32.
// Panics if value is not of float32 type.
func (a *BoundedAtStartItr) Float32() float32 {
	v, ok := a.Value().(float32)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Float32.", a.Value()))
	}
	return v
}

// Adapter for float32 types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to float32.
// Returns (value, false) if type is of float32 type.
// Returns (default value, true) if type is not of float32 type.
func (a *BoundedAtStartItr) Float32Or(def float32) (float32, bool) {
	v, ok := a.Value().(float32)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for float32 types on Bounded iterators. Provides a typecasting method 
// from interface{} to float32.
// Panics if value is not of float32 type.
func (a *BoundedItr) Float32() float32 {
	v, ok := a.Value().(float32)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Float32.", a.Value()))
	}
	return v
}

// Adapter for float32 types on Bounded iterators. Provides a typecasting method 
// from interface{} to float32.
// Returns (value, false) if type is of float32 type.
// Returns (default value, true) if type is not of float32 type.
func (a *BoundedItr) Float32Or(def float32) (float32, bool) {
	v, ok := a.Value().(float32)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for float32 types on RandomAccess iterator. Provides a typecasting method 
// from interface{} to float32.
// Panics if value is not of float32 type.
func (a *RandomAccessItr) Float32() float32 {
	v, ok := a.Value().(float32)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Float32.", a.Value()))
	}
	return v
}

// Adapter for float32 types on RandomAccess iterators. Provides a typecasting method 
// from interface{} to float32.
// Returns (value, false) if type is of float32 type.
// Returns (default value, true) if type is not of float32 type.
func (a *RandomAccessItr) Float32Or(def float32) (float32, bool) {
	v, ok := a.Value().(float32)
	if !ok {
		return def, true
	}
	return v, false
}
