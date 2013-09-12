//  Generated file for Uint32 adapter
package iadapt

import "fmt"

// Adapter for uint32 types on Forward iterators. Provides a typecasting method 
// from interface{} to uint32.
// Panics if value is not of uint32 type.
func (a *ForwardItr) Uint32() uint32 {
	v, ok := a.Value().(uint32)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint32.", a.Value()))
	}
	return v
}

// Adapter for uint32 types on Forward iterators. Provides a typecasting method 
// from interface{} to uint32.
// Returns (value, false) if type is of uint32 type.
// Returns (default value, true) if type is not of uint32 type.
func (a *ForwardItr) Uint32Or(def uint32) (uint32, bool) {
	v, ok := a.Value().(uint32)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for uint32 types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to uint32.
// Panics if value is not of uint32 type.
func (a *BiDirectionalItr) Uint32() uint32 {
	v, ok := a.Value().(uint32)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint32.", a.Value()))
	}
	return v
}

// Adapter for uint32 types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to uint32.
// Returns (value, false) if type is of uint32 type.
// Returns (default value, true) if type is not of uint32 type.
func (a *BiDirectionalItr) Uint32Or(def uint32) (uint32, bool) {
	v, ok := a.Value().(uint32)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for uint32 types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to uint32.
// Panics if value is not of uint32 type.
func (a *BoundedAtStartItr) Uint32() uint32 {
	v, ok := a.Value().(uint32)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint32.", a.Value()))
	}
	return v
}

// Adapter for uint32 types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to uint32.
// Returns (value, false) if type is of uint32 type.
// Returns (default value, true) if type is not of uint32 type.
func (a *BoundedAtStartItr) Uint32Or(def uint32) (uint32, bool) {
	v, ok := a.Value().(uint32)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for uint32 types on Bounded iterators. Provides a typecasting method 
// from interface{} to uint32.
// Panics if value is not of uint32 type.
func (a *BoundedItr) Uint32() uint32 {
	v, ok := a.Value().(uint32)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint32.", a.Value()))
	}
	return v
}

// Adapter for uint32 types on Bounded iterators. Provides a typecasting method 
// from interface{} to uint32.
// Returns (value, false) if type is of uint32 type.
// Returns (default value, true) if type is not of uint32 type.
func (a *BoundedItr) Uint32Or(def uint32) (uint32, bool) {
	v, ok := a.Value().(uint32)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for uint32 types on RandomAccess iterator. Provides a typecasting method 
// from interface{} to uint32.
// Panics if value is not of uint32 type.
func (a *RandomAccessItr) Uint32() uint32 {
	v, ok := a.Value().(uint32)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint32.", a.Value()))
	}
	return v
}

// Adapter for uint32 types on RandomAccess iterators. Provides a typecasting method 
// from interface{} to uint32.
// Returns (value, false) if type is of uint32 type.
// Returns (default value, true) if type is not of uint32 type.
func (a *RandomAccessItr) Uint32Or(def uint32) (uint32, bool) {
	v, ok := a.Value().(uint32)
	if !ok {
		return def, true
	}
	return v, false
}
