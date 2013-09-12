//  Generated file for Int16 adapter
package iadapt

import "fmt"

// Adapter for int16 types on Forward iterators. Provides a typecasting method 
// from interface{} to int16.
// Panics if value is not of int16 type.
func (a *ForwardItr) Int16() int16 {
	v, ok := a.Value().(int16)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int16.", a.Value()))
	}
	return v
}

// Adapter for int16 types on Forward iterators. Provides a typecasting method 
// from interface{} to int16.
// Returns (value, false) if type is of int16 type.
// Returns (default value, true) if type is not of int16 type.
func (a *ForwardItr) Int16Or(def int16) (int16, bool) {
	v, ok := a.Value().(int16)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for int16 types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to int16.
// Panics if value is not of int16 type.
func (a *BiDirectionalItr) Int16() int16 {
	v, ok := a.Value().(int16)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int16.", a.Value()))
	}
	return v
}

// Adapter for int16 types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to int16.
// Returns (value, false) if type is of int16 type.
// Returns (default value, true) if type is not of int16 type.
func (a *BiDirectionalItr) Int16Or(def int16) (int16, bool) {
	v, ok := a.Value().(int16)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for int16 types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to int16.
// Panics if value is not of int16 type.
func (a *BoundedAtStartItr) Int16() int16 {
	v, ok := a.Value().(int16)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int16.", a.Value()))
	}
	return v
}

// Adapter for int16 types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to int16.
// Returns (value, false) if type is of int16 type.
// Returns (default value, true) if type is not of int16 type.
func (a *BoundedAtStartItr) Int16Or(def int16) (int16, bool) {
	v, ok := a.Value().(int16)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for int16 types on Bounded iterators. Provides a typecasting method 
// from interface{} to int16.
// Panics if value is not of int16 type.
func (a *BoundedItr) Int16() int16 {
	v, ok := a.Value().(int16)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int16.", a.Value()))
	}
	return v
}

// Adapter for int16 types on Bounded iterators. Provides a typecasting method 
// from interface{} to int16.
// Returns (value, false) if type is of int16 type.
// Returns (default value, true) if type is not of int16 type.
func (a *BoundedItr) Int16Or(def int16) (int16, bool) {
	v, ok := a.Value().(int16)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for int16 types on RandomAccess iterator. Provides a typecasting method 
// from interface{} to int16.
// Panics if value is not of int16 type.
func (a *RandomAccessItr) Int16() int16 {
	v, ok := a.Value().(int16)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int16.", a.Value()))
	}
	return v
}

// Adapter for int16 types on RandomAccess iterators. Provides a typecasting method 
// from interface{} to int16.
// Returns (value, false) if type is of int16 type.
// Returns (default value, true) if type is not of int16 type.
func (a *RandomAccessItr) Int16Or(def int16) (int16, bool) {
	v, ok := a.Value().(int16)
	if !ok {
		return def, true
	}
	return v, false
}
