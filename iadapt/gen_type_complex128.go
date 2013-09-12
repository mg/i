//  Generated file for Complex128 adapter
package iadapt

import "fmt"

// Adapter for complex128 types on Forward iterators. Provides a typecasting method 
// from interface{} to complex128.
// Panics if value is not of complex128 type.
func (a *ForwardItr) Complex128() complex128 {
	v, ok := a.Value().(complex128)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Complex128.", a.Value()))
	}
	return v
}

// Adapter for complex128 types on Forward iterators. Provides a typecasting method 
// from interface{} to complex128.
// Returns (value, false) if type is of complex128 type.
// Returns (default value, true) if type is not of complex128 type.
func (a *ForwardItr) Complex128Or(def complex128) (complex128, bool) {
	v, ok := a.Value().(complex128)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for complex128 types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to complex128.
// Panics if value is not of complex128 type.
func (a *BiDirectionalItr) Complex128() complex128 {
	v, ok := a.Value().(complex128)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Complex128.", a.Value()))
	}
	return v
}

// Adapter for complex128 types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to complex128.
// Returns (value, false) if type is of complex128 type.
// Returns (default value, true) if type is not of complex128 type.
func (a *BiDirectionalItr) Complex128Or(def complex128) (complex128, bool) {
	v, ok := a.Value().(complex128)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for complex128 types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to complex128.
// Panics if value is not of complex128 type.
func (a *BoundedAtStartItr) Complex128() complex128 {
	v, ok := a.Value().(complex128)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Complex128.", a.Value()))
	}
	return v
}

// Adapter for complex128 types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to complex128.
// Returns (value, false) if type is of complex128 type.
// Returns (default value, true) if type is not of complex128 type.
func (a *BoundedAtStartItr) Complex128Or(def complex128) (complex128, bool) {
	v, ok := a.Value().(complex128)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for complex128 types on Bounded iterators. Provides a typecasting method 
// from interface{} to complex128.
// Panics if value is not of complex128 type.
func (a *BoundedItr) Complex128() complex128 {
	v, ok := a.Value().(complex128)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Complex128.", a.Value()))
	}
	return v
}

// Adapter for complex128 types on Bounded iterators. Provides a typecasting method 
// from interface{} to complex128.
// Returns (value, false) if type is of complex128 type.
// Returns (default value, true) if type is not of complex128 type.
func (a *BoundedItr) Complex128Or(def complex128) (complex128, bool) {
	v, ok := a.Value().(complex128)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for complex128 types on RandomAccess iterator. Provides a typecasting method 
// from interface{} to complex128.
// Panics if value is not of complex128 type.
func (a *RandomAccessItr) Complex128() complex128 {
	v, ok := a.Value().(complex128)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Complex128.", a.Value()))
	}
	return v
}

// Adapter for complex128 types on RandomAccess iterators. Provides a typecasting method 
// from interface{} to complex128.
// Returns (value, false) if type is of complex128 type.
// Returns (default value, true) if type is not of complex128 type.
func (a *RandomAccessItr) Complex128Or(def complex128) (complex128, bool) {
	v, ok := a.Value().(complex128)
	if !ok {
		return def, true
	}
	return v, false
}
