//  Generated file for Uint16 adapter
package iadapt

import "fmt"

// Adapter for uint16 types on Forward iterators. Provides a typecasting method 
// from interface{} to uint16.
// Panics if value is not of uint16 type.
func (a *ForwardItr) Uint16() uint16 {
	v, ok := a.Value().(uint16)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint16.", a.Value()))
	}
	return v
}

// Adapter for uint16 types on Forward iterators. Provides a typecasting method 
// from interface{} to uint16.
// Returns (value, false) if type is of uint16 type.
// Returns (default value, true) if type is not of uint16 type.
func (a *ForwardItr) Uint16Or(def uint16) (uint16, bool) {
	v, ok := a.Value().(uint16)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for uint16 types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to uint16.
// Panics if value is not of uint16 type.
func (a *BiDirectionalItr) Uint16() uint16 {
	v, ok := a.Value().(uint16)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint16.", a.Value()))
	}
	return v
}

// Adapter for uint16 types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to uint16.
// Returns (value, false) if type is of uint16 type.
// Returns (default value, true) if type is not of uint16 type.
func (a *BiDirectionalItr) Uint16Or(def uint16) (uint16, bool) {
	v, ok := a.Value().(uint16)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for uint16 types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to uint16.
// Panics if value is not of uint16 type.
func (a *BoundedAtStartItr) Uint16() uint16 {
	v, ok := a.Value().(uint16)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint16.", a.Value()))
	}
	return v
}

// Adapter for uint16 types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to uint16.
// Returns (value, false) if type is of uint16 type.
// Returns (default value, true) if type is not of uint16 type.
func (a *BoundedAtStartItr) Uint16Or(def uint16) (uint16, bool) {
	v, ok := a.Value().(uint16)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for uint16 types on Bounded iterators. Provides a typecasting method 
// from interface{} to uint16.
// Panics if value is not of uint16 type.
func (a *BoundedItr) Uint16() uint16 {
	v, ok := a.Value().(uint16)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint16.", a.Value()))
	}
	return v
}

// Adapter for uint16 types on Bounded iterators. Provides a typecasting method 
// from interface{} to uint16.
// Returns (value, false) if type is of uint16 type.
// Returns (default value, true) if type is not of uint16 type.
func (a *BoundedItr) Uint16Or(def uint16) (uint16, bool) {
	v, ok := a.Value().(uint16)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for uint16 types on RandomAccess iterator. Provides a typecasting method 
// from interface{} to uint16.
// Panics if value is not of uint16 type.
func (a *RandomAccessItr) Uint16() uint16 {
	v, ok := a.Value().(uint16)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Uint16.", a.Value()))
	}
	return v
}

// Adapter for uint16 types on RandomAccess iterators. Provides a typecasting method 
// from interface{} to uint16.
// Returns (value, false) if type is of uint16 type.
// Returns (default value, true) if type is not of uint16 type.
func (a *RandomAccessItr) Uint16Or(def uint16) (uint16, bool) {
	v, ok := a.Value().(uint16)
	if !ok {
		return def, true
	}
	return v, false
}
