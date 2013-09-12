//  Generated file for Float64 adapter
package iadapt

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
