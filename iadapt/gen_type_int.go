//  Generated file for Int adapter
package iadapt

import "fmt"

// Adapter for int types on Forward iterators. Provides a typecasting method 
// from interface{} to int.
// Panics if value is not of int type.
func (a *ForwardItr) Int() int {
	v, ok := a.Value().(int)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int.", a.Value()))
	}
	return v
}

// Adapter for int types on Forward iterators. Provides a typecasting method 
// from interface{} to int.
// Returns (value, false) if type is of int type.
// Returns (default value, true) if type is not of int type.
func (a *ForwardItr) IntOr(def int) (int, bool) {
	v, ok := a.Value().(int)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for int types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to int.
// Panics if value is not of int type.
func (a *BiDirectionalItr) Int() int {
	v, ok := a.Value().(int)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int.", a.Value()))
	}
	return v
}

// Adapter for int types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to int.
// Returns (value, false) if type is of int type.
// Returns (default value, true) if type is not of int type.
func (a *BiDirectionalItr) IntOr(def int) (int, bool) {
	v, ok := a.Value().(int)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for int types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to int.
// Panics if value is not of int type.
func (a *BoundedAtStartItr) Int() int {
	v, ok := a.Value().(int)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int.", a.Value()))
	}
	return v
}

// Adapter for int types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to int.
// Returns (value, false) if type is of int type.
// Returns (default value, true) if type is not of int type.
func (a *BoundedAtStartItr) IntOr(def int) (int, bool) {
	v, ok := a.Value().(int)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for int types on Bounded iterators. Provides a typecasting method 
// from interface{} to int.
// Panics if value is not of int type.
func (a *BoundedItr) Int() int {
	v, ok := a.Value().(int)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int.", a.Value()))
	}
	return v
}

// Adapter for int types on Bounded iterators. Provides a typecasting method 
// from interface{} to int.
// Returns (value, false) if type is of int type.
// Returns (default value, true) if type is not of int type.
func (a *BoundedItr) IntOr(def int) (int, bool) {
	v, ok := a.Value().(int)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for int types on RandomAccess iterator. Provides a typecasting method 
// from interface{} to int.
// Panics if value is not of int type.
func (a *RandomAccessItr) Int() int {
	v, ok := a.Value().(int)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Int.", a.Value()))
	}
	return v
}

// Adapter for int types on RandomAccess iterators. Provides a typecasting method 
// from interface{} to int.
// Returns (value, false) if type is of int type.
// Returns (default value, true) if type is not of int type.
func (a *RandomAccessItr) IntOr(def int) (int, bool) {
	v, ok := a.Value().(int)
	if !ok {
		return def, true
	}
	return v, false
}
