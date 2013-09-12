//  Generated file for Int8 adapter
package iadapt

import "fmt"

// Adapter for int8 types on Forward iterators. Provides a typecasting method 
// from interface{} to int8.
// Panics if value is not of int8 type.
func (a *ForwardItr) Int8() int8 {
	v, ok := a.Value().(int8)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int8.", a.Value()))
	}
	return v
}

// Adapter for int8 types on Forward iterators. Provides a typecasting method 
// from interface{} to int8.
// Returns (value, false) if type is of int8 type.
// Returns (default value, true) if type is not of int8 type.
func (a *ForwardItr) Int8Or(def int8) (int8, bool) {
	v, ok := a.Value().(int8)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for int8 types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to int8.
// Panics if value is not of int8 type.
func (a *BiDirectionalItr) Int8() int8 {
	v, ok := a.Value().(int8)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int8.", a.Value()))
	}
	return v
}

// Adapter for int8 types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to int8.
// Returns (value, false) if type is of int8 type.
// Returns (default value, true) if type is not of int8 type.
func (a *BiDirectionalItr) Int8Or(def int8) (int8, bool) {
	v, ok := a.Value().(int8)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for int8 types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to int8.
// Panics if value is not of int8 type.
func (a *BoundedAtStartItr) Int8() int8 {
	v, ok := a.Value().(int8)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int8.", a.Value()))
	}
	return v
}

// Adapter for int8 types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to int8.
// Returns (value, false) if type is of int8 type.
// Returns (default value, true) if type is not of int8 type.
func (a *BoundedAtStartItr) Int8Or(def int8) (int8, bool) {
	v, ok := a.Value().(int8)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for int8 types on Bounded iterators. Provides a typecasting method 
// from interface{} to int8.
// Panics if value is not of int8 type.
func (a *BoundedItr) Int8() int8 {
	v, ok := a.Value().(int8)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int8.", a.Value()))
	}
	return v
}

// Adapter for int8 types on Bounded iterators. Provides a typecasting method 
// from interface{} to int8.
// Returns (value, false) if type is of int8 type.
// Returns (default value, true) if type is not of int8 type.
func (a *BoundedItr) Int8Or(def int8) (int8, bool) {
	v, ok := a.Value().(int8)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for int8 types on RandomAccess iterator. Provides a typecasting method 
// from interface{} to int8.
// Panics if value is not of int8 type.
func (a *RandomAccessItr) Int8() int8 {
	v, ok := a.Value().(int8)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int8.", a.Value()))
	}
	return v
}

// Adapter for int8 types on RandomAccess iterators. Provides a typecasting method 
// from interface{} to int8.
// Returns (value, false) if type is of int8 type.
// Returns (default value, true) if type is not of int8 type.
func (a *RandomAccessItr) Int8Or(def int8) (int8, bool) {
	v, ok := a.Value().(int8)
	if !ok {
		return def, true
	}
	return v, false
}
