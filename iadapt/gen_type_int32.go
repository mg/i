//  Generated file for Int32 adapter
package iadapt

import "fmt"

// Adapter for int32 types on Forward iterators. Provides a typecasting method 
// from interface{} to int32.
// Panics if value is not of int32 type.
func (a *ForwardItr) Int32() int32 {
	v, ok := a.Value().(int32)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int32.", a.Value()))
	}
	return v
}

// Adapter for int32 types on Forward iterators. Provides a typecasting method 
// from interface{} to int32.
// Returns (value, false) if type is of int32 type.
// Returns (default value, true) if type is not of int32 type.
func (a *ForwardItr) Int32Or(def int32) (int32, bool) {
	v, ok := a.Value().(int32)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for int32 types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to int32.
// Panics if value is not of int32 type.
func (a *BiDirectionalItr) Int32() int32 {
	v, ok := a.Value().(int32)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int32.", a.Value()))
	}
	return v
}

// Adapter for int32 types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to int32.
// Returns (value, false) if type is of int32 type.
// Returns (default value, true) if type is not of int32 type.
func (a *BiDirectionalItr) Int32Or(def int32) (int32, bool) {
	v, ok := a.Value().(int32)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for int32 types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to int32.
// Panics if value is not of int32 type.
func (a *BoundedAtStartItr) Int32() int32 {
	v, ok := a.Value().(int32)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int32.", a.Value()))
	}
	return v
}

// Adapter for int32 types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to int32.
// Returns (value, false) if type is of int32 type.
// Returns (default value, true) if type is not of int32 type.
func (a *BoundedAtStartItr) Int32Or(def int32) (int32, bool) {
	v, ok := a.Value().(int32)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for int32 types on Bounded iterators. Provides a typecasting method 
// from interface{} to int32.
// Panics if value is not of int32 type.
func (a *BoundedItr) Int32() int32 {
	v, ok := a.Value().(int32)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int32.", a.Value()))
	}
	return v
}

// Adapter for int32 types on Bounded iterators. Provides a typecasting method 
// from interface{} to int32.
// Returns (value, false) if type is of int32 type.
// Returns (default value, true) if type is not of int32 type.
func (a *BoundedItr) Int32Or(def int32) (int32, bool) {
	v, ok := a.Value().(int32)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for int32 types on RandomAccess iterator. Provides a typecasting method 
// from interface{} to int32.
// Panics if value is not of int32 type.
func (a *RandomAccessItr) Int32() int32 {
	v, ok := a.Value().(int32)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int32.", a.Value()))
	}
	return v
}

// Adapter for int32 types on RandomAccess iterators. Provides a typecasting method 
// from interface{} to int32.
// Returns (value, false) if type is of int32 type.
// Returns (default value, true) if type is not of int32 type.
func (a *RandomAccessItr) Int32Or(def int32) (int32, bool) {
	v, ok := a.Value().(int32)
	if !ok {
		return def, true
	}
	return v, false
}
