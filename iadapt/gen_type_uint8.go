//  Generated file for Uint8 adapter
package iadapt

import "fmt"

// Adapter for uint8 types on Forward iterators. Provides a typecasting method 
// from interface{} to uint8.
// Panics if value is not of uint8 type.
func (a *ForwardItr) Uint8() uint8 {
	v, ok := a.Value().(uint8)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint8.", a.Value()))
	}
	return v
}

// Adapter for uint8 types on Forward iterators. Provides a typecasting method 
// from interface{} to uint8.
// Returns (value, false) if type is of uint8 type.
// Returns (default value, true) if type is not of uint8 type.
func (a *ForwardItr) Uint8Or(def uint8) (uint8, bool) {
	v, ok := a.Value().(uint8)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for uint8 types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to uint8.
// Panics if value is not of uint8 type.
func (a *BiDirectionalItr) Uint8() uint8 {
	v, ok := a.Value().(uint8)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint8.", a.Value()))
	}
	return v
}

// Adapter for uint8 types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to uint8.
// Returns (value, false) if type is of uint8 type.
// Returns (default value, true) if type is not of uint8 type.
func (a *BiDirectionalItr) Uint8Or(def uint8) (uint8, bool) {
	v, ok := a.Value().(uint8)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for uint8 types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to uint8.
// Panics if value is not of uint8 type.
func (a *BoundedAtStartItr) Uint8() uint8 {
	v, ok := a.Value().(uint8)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint8.", a.Value()))
	}
	return v
}

// Adapter for uint8 types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to uint8.
// Returns (value, false) if type is of uint8 type.
// Returns (default value, true) if type is not of uint8 type.
func (a *BoundedAtStartItr) Uint8Or(def uint8) (uint8, bool) {
	v, ok := a.Value().(uint8)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for uint8 types on Bounded iterators. Provides a typecasting method 
// from interface{} to uint8.
// Panics if value is not of uint8 type.
func (a *BoundedItr) Uint8() uint8 {
	v, ok := a.Value().(uint8)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint8.", a.Value()))
	}
	return v
}

// Adapter for uint8 types on Bounded iterators. Provides a typecasting method 
// from interface{} to uint8.
// Returns (value, false) if type is of uint8 type.
// Returns (default value, true) if type is not of uint8 type.
func (a *BoundedItr) Uint8Or(def uint8) (uint8, bool) {
	v, ok := a.Value().(uint8)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for uint8 types on RandomAccess iterator. Provides a typecasting method 
// from interface{} to uint8.
// Panics if value is not of uint8 type.
func (a *RandomAccessItr) Uint8() uint8 {
	v, ok := a.Value().(uint8)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint8.", a.Value()))
	}
	return v
}

// Adapter for uint8 types on RandomAccess iterators. Provides a typecasting method 
// from interface{} to uint8.
// Returns (value, false) if type is of uint8 type.
// Returns (default value, true) if type is not of uint8 type.
func (a *RandomAccessItr) Uint8Or(def uint8) (uint8, bool) {
	v, ok := a.Value().(uint8)
	if !ok {
		return def, true
	}
	return v, false
}
