//  Generated file for Bool adapter
package iadapt

import "fmt"

// Adapter for bool types on Forward iterators. Provides a typecasting method 
// from interface{} to bool.
// Panics if value is not of bool type.
func (a *ForwardItr) Bool() bool {
	v, ok := a.Value().(bool)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Bool.", a.Value()))
	}
	return v
}

// Adapter for bool types on Forward iterators. Provides a typecasting method 
// from interface{} to bool.
// Returns (value, false) if type is of bool type.
// Returns (default value, true) if type is not of bool type.
func (a *ForwardItr) BoolOr(def bool) (bool, bool) {
	v, ok := a.Value().(bool)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for bool types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to bool.
// Panics if value is not of bool type.
func (a *BiDirectionalItr) Bool() bool {
	v, ok := a.Value().(bool)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Bool.", a.Value()))
	}
	return v
}

// Adapter for bool types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to bool.
// Returns (value, false) if type is of bool type.
// Returns (default value, true) if type is not of bool type.
func (a *BiDirectionalItr) BoolOr(def bool) (bool, bool) {
	v, ok := a.Value().(bool)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for bool types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to bool.
// Panics if value is not of bool type.
func (a *BoundedAtStartItr) Bool() bool {
	v, ok := a.Value().(bool)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Bool.", a.Value()))
	}
	return v
}

// Adapter for bool types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to bool.
// Returns (value, false) if type is of bool type.
// Returns (default value, true) if type is not of bool type.
func (a *BoundedAtStartItr) BoolOr(def bool) (bool, bool) {
	v, ok := a.Value().(bool)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for bool types on Bounded iterators. Provides a typecasting method 
// from interface{} to bool.
// Panics if value is not of bool type.
func (a *BoundedItr) Bool() bool {
	v, ok := a.Value().(bool)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Bool.", a.Value()))
	}
	return v
}

// Adapter for bool types on Bounded iterators. Provides a typecasting method 
// from interface{} to bool.
// Returns (value, false) if type is of bool type.
// Returns (default value, true) if type is not of bool type.
func (a *BoundedItr) BoolOr(def bool) (bool, bool) {
	v, ok := a.Value().(bool)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for bool types on RandomAccess iterator. Provides a typecasting method 
// from interface{} to bool.
// Panics if value is not of bool type.
func (a *RandomAccessItr) Bool() bool {
	v, ok := a.Value().(bool)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Bool.", a.Value()))
	}
	return v
}

// Adapter for bool types on RandomAccess iterators. Provides a typecasting method 
// from interface{} to bool.
// Returns (value, false) if type is of bool type.
// Returns (default value, true) if type is not of bool type.
func (a *RandomAccessItr) BoolOr(def bool) (bool, bool) {
	v, ok := a.Value().(bool)
	if !ok {
		return def, true
	}
	return v, false
}
