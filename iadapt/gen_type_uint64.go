//  Generated file for Uint64 adapter
package iadapt

import "fmt"

// Adapter for uint64 types on Forward iterators. Provides a typecasting method 
// from interface{} to uint64.
// Panics if value is not of uint64 type.
func (a *ForwardItr) Uint64() uint64 {
	v, ok := a.Value().(uint64)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint64.", a.Value()))
	}
	return v
}

// Adapter for uint64 types on Forward iterators. Provides a typecasting method 
// from interface{} to uint64.
// Returns (value, false) if type is of uint64 type.
// Returns (default value, true) if type is not of uint64 type.
func (a *ForwardItr) Uint64Or(def uint64) (uint64, bool) {
	v, ok := a.Value().(uint64)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for uint64 types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to uint64.
// Panics if value is not of uint64 type.
func (a *BiDirectionalItr) Uint64() uint64 {
	v, ok := a.Value().(uint64)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint64.", a.Value()))
	}
	return v
}

// Adapter for uint64 types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to uint64.
// Returns (value, false) if type is of uint64 type.
// Returns (default value, true) if type is not of uint64 type.
func (a *BiDirectionalItr) Uint64Or(def uint64) (uint64, bool) {
	v, ok := a.Value().(uint64)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for uint64 types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to uint64.
// Panics if value is not of uint64 type.
func (a *BoundedAtStartItr) Uint64() uint64 {
	v, ok := a.Value().(uint64)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint64.", a.Value()))
	}
	return v
}

// Adapter for uint64 types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to uint64.
// Returns (value, false) if type is of uint64 type.
// Returns (default value, true) if type is not of uint64 type.
func (a *BoundedAtStartItr) Uint64Or(def uint64) (uint64, bool) {
	v, ok := a.Value().(uint64)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for uint64 types on Bounded iterators. Provides a typecasting method 
// from interface{} to uint64.
// Panics if value is not of uint64 type.
func (a *BoundedItr) Uint64() uint64 {
	v, ok := a.Value().(uint64)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint64.", a.Value()))
	}
	return v
}

// Adapter for uint64 types on Bounded iterators. Provides a typecasting method 
// from interface{} to uint64.
// Returns (value, false) if type is of uint64 type.
// Returns (default value, true) if type is not of uint64 type.
func (a *BoundedItr) Uint64Or(def uint64) (uint64, bool) {
	v, ok := a.Value().(uint64)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for uint64 types on RandomAccess iterator. Provides a typecasting method 
// from interface{} to uint64.
// Panics if value is not of uint64 type.
func (a *RandomAccessItr) Uint64() uint64 {
	v, ok := a.Value().(uint64)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint64.", a.Value()))
	}
	return v
}

// Adapter for uint64 types on RandomAccess iterators. Provides a typecasting method 
// from interface{} to uint64.
// Returns (value, false) if type is of uint64 type.
// Returns (default value, true) if type is not of uint64 type.
func (a *RandomAccessItr) Uint64Or(def uint64) (uint64, bool) {
	v, ok := a.Value().(uint64)
	if !ok {
		return def, true
	}
	return v, false
}
