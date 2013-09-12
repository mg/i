//  Generated file for Rune adapter
package iadapt

import "fmt"

// Adapter for rune types on Forward iterators. Provides a typecasting method 
// from interface{} to rune.
// Panics if value is not of rune type.
func (a *ForwardItr) Rune() rune {
	v, ok := a.Value().(rune)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Rune.", a.Value()))
	}
	return v
}

// Adapter for rune types on Forward iterators. Provides a typecasting method 
// from interface{} to rune.
// Returns (value, false) if type is of rune type.
// Returns (default value, true) if type is not of rune type.
func (a *ForwardItr) RuneOr(def rune) (rune, bool) {
	v, ok := a.Value().(rune)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for rune types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to rune.
// Panics if value is not of rune type.
func (a *BiDirectionalItr) Rune() rune {
	v, ok := a.Value().(rune)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Rune.", a.Value()))
	}
	return v
}

// Adapter for rune types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to rune.
// Returns (value, false) if type is of rune type.
// Returns (default value, true) if type is not of rune type.
func (a *BiDirectionalItr) RuneOr(def rune) (rune, bool) {
	v, ok := a.Value().(rune)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for rune types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to rune.
// Panics if value is not of rune type.
func (a *BoundedAtStartItr) Rune() rune {
	v, ok := a.Value().(rune)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Rune.", a.Value()))
	}
	return v
}

// Adapter for rune types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to rune.
// Returns (value, false) if type is of rune type.
// Returns (default value, true) if type is not of rune type.
func (a *BoundedAtStartItr) RuneOr(def rune) (rune, bool) {
	v, ok := a.Value().(rune)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for rune types on Bounded iterators. Provides a typecasting method 
// from interface{} to rune.
// Panics if value is not of rune type.
func (a *BoundedItr) Rune() rune {
	v, ok := a.Value().(rune)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Rune.", a.Value()))
	}
	return v
}

// Adapter for rune types on Bounded iterators. Provides a typecasting method 
// from interface{} to rune.
// Returns (value, false) if type is of rune type.
// Returns (default value, true) if type is not of rune type.
func (a *BoundedItr) RuneOr(def rune) (rune, bool) {
	v, ok := a.Value().(rune)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for rune types on RandomAccess iterator. Provides a typecasting method 
// from interface{} to rune.
// Panics if value is not of rune type.
func (a *RandomAccessItr) Rune() rune {
	v, ok := a.Value().(rune)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Rune.", a.Value()))
	}
	return v
}

// Adapter for rune types on RandomAccess iterators. Provides a typecasting method 
// from interface{} to rune.
// Returns (value, false) if type is of rune type.
// Returns (default value, true) if type is not of rune type.
func (a *RandomAccessItr) RuneOr(def rune) (rune, bool) {
	v, ok := a.Value().(rune)
	if !ok {
		return def, true
	}
	return v, false
}
