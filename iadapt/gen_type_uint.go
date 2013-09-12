//  Generated file for Uint adapter
package iadapt

import "fmt"

// Adapter for uint types on Forward iterators. Provides a typecasting method 
// from interface{} to uint.
// Panics if value is not of uint type.
func (a *ForwardItr) Uint() uint {
	v, ok := a.Value().(uint)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint.", a.Value()))
	}
	return v
}

// Adapter for uint types on Forward iterators. Provides a typecasting method 
// from interface{} to uint.
// Returns (value, false) if type is of uint type.
// Returns (default value, true) if type is not of uint type.
func (a *ForwardItr) UintOr(def uint) (uint, bool) {
	v, ok := a.Value().(uint)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for uint types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to uint.
// Panics if value is not of uint type.
func (a *BiDirectionalItr) Uint() uint {
	v, ok := a.Value().(uint)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint.", a.Value()))
	}
	return v
}

// Adapter for uint types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to uint.
// Returns (value, false) if type is of uint type.
// Returns (default value, true) if type is not of uint type.
func (a *BiDirectionalItr) UintOr(def uint) (uint, bool) {
	v, ok := a.Value().(uint)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for uint types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to uint.
// Panics if value is not of uint type.
func (a *BoundedAtStartItr) Uint() uint {
	v, ok := a.Value().(uint)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint.", a.Value()))
	}
	return v
}

// Adapter for uint types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to uint.
// Returns (value, false) if type is of uint type.
// Returns (default value, true) if type is not of uint type.
func (a *BoundedAtStartItr) UintOr(def uint) (uint, bool) {
	v, ok := a.Value().(uint)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for uint types on Bounded iterators. Provides a typecasting method 
// from interface{} to uint.
// Panics if value is not of uint type.
func (a *BoundedItr) Uint() uint {
	v, ok := a.Value().(uint)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint.", a.Value()))
	}
	return v
}

// Adapter for uint types on Bounded iterators. Provides a typecasting method 
// from interface{} to uint.
// Returns (value, false) if type is of uint type.
// Returns (default value, true) if type is not of uint type.
func (a *BoundedItr) UintOr(def uint) (uint, bool) {
	v, ok := a.Value().(uint)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for uint types on RandomAccess iterator. Provides a typecasting method 
// from interface{} to uint.
// Panics if value is not of uint type.
func (a *RandomAccessItr) Uint() uint {
	v, ok := a.Value().(uint)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint.", a.Value()))
	}
	return v
}

// Adapter for uint types on RandomAccess iterators. Provides a typecasting method 
// from interface{} to uint.
// Returns (value, false) if type is of uint type.
// Returns (default value, true) if type is not of uint type.
func (a *RandomAccessItr) UintOr(def uint) (uint, bool) {
	v, ok := a.Value().(uint)
	if !ok {
		return def, true
	}
	return v, false
}
