//  Generated file for Int64 adapter
package iadapt

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
