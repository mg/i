//  Generated file for Complex64 adapter
package iadapt

import "fmt"

// Adapter for complex64 types on Forward iterators. Provides a typecasting method 
// from interface{} to complex64.
// Panics if value is not of complex64 type.
func (a *ForwardItr) Complex64() complex64 {
	v, ok := a.Value().(complex64)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Complex64.", a.Value()))
	}
	return v
}

// Adapter for complex64 types on Forward iterators. Provides a typecasting method 
// from interface{} to complex64.
// Returns (value, false) if type is of complex64 type.
// Returns (default value, true) if type is not of complex64 type.
func (a *ForwardItr) Complex64Or(def complex64) (complex64, bool) {
	v, ok := a.Value().(complex64)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for complex64 types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to complex64.
// Panics if value is not of complex64 type.
func (a *BiDirectionalItr) Complex64() complex64 {
	v, ok := a.Value().(complex64)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Complex64.", a.Value()))
	}
	return v
}

// Adapter for complex64 types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to complex64.
// Returns (value, false) if type is of complex64 type.
// Returns (default value, true) if type is not of complex64 type.
func (a *BiDirectionalItr) Complex64Or(def complex64) (complex64, bool) {
	v, ok := a.Value().(complex64)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for complex64 types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to complex64.
// Panics if value is not of complex64 type.
func (a *BoundedAtStartItr) Complex64() complex64 {
	v, ok := a.Value().(complex64)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Complex64.", a.Value()))
	}
	return v
}

// Adapter for complex64 types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to complex64.
// Returns (value, false) if type is of complex64 type.
// Returns (default value, true) if type is not of complex64 type.
func (a *BoundedAtStartItr) Complex64Or(def complex64) (complex64, bool) {
	v, ok := a.Value().(complex64)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for complex64 types on Bounded iterators. Provides a typecasting method 
// from interface{} to complex64.
// Panics if value is not of complex64 type.
func (a *BoundedItr) Complex64() complex64 {
	v, ok := a.Value().(complex64)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Complex64.", a.Value()))
	}
	return v
}

// Adapter for complex64 types on Bounded iterators. Provides a typecasting method 
// from interface{} to complex64.
// Returns (value, false) if type is of complex64 type.
// Returns (default value, true) if type is not of complex64 type.
func (a *BoundedItr) Complex64Or(def complex64) (complex64, bool) {
	v, ok := a.Value().(complex64)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for complex64 types on RandomAccess iterator. Provides a typecasting method 
// from interface{} to complex64.
// Panics if value is not of complex64 type.
func (a *RandomAccessItr) Complex64() complex64 {
	v, ok := a.Value().(complex64)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Complex64.", a.Value()))
	}
	return v
}

// Adapter for complex64 types on RandomAccess iterators. Provides a typecasting method 
// from interface{} to complex64.
// Returns (value, false) if type is of complex64 type.
// Returns (default value, true) if type is not of complex64 type.
func (a *RandomAccessItr) Complex64Or(def complex64) (complex64, bool) {
	v, ok := a.Value().(complex64)
	if !ok {
		return def, true
	}
	return v, false
}
