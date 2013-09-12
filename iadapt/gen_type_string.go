//  Generated file for String adapter
package iadapt

import "fmt"

// Adapter for string types on Forward iterators. Provides a typecasting method 
// from interface{} to string.
// Panics if value is not of string type.
func (a *ForwardItr) String() string {
	v, ok := a.Value().(string)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to String.", a.Value()))
	}
	return v
}

// Adapter for string types on Forward iterators. Provides a typecasting method 
// from interface{} to string.
// Returns (value, false) if type is of string type.
// Returns (default value, true) if type is not of string type.
func (a *ForwardItr) StringOr(def string) (string, bool) {
	v, ok := a.Value().(string)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for string types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to string.
// Panics if value is not of string type.
func (a *BiDirectionalItr) String() string {
	v, ok := a.Value().(string)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to String.", a.Value()))
	}
	return v
}

// Adapter for string types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to string.
// Returns (value, false) if type is of string type.
// Returns (default value, true) if type is not of string type.
func (a *BiDirectionalItr) StringOr(def string) (string, bool) {
	v, ok := a.Value().(string)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for string types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to string.
// Panics if value is not of string type.
func (a *BoundedAtStartItr) String() string {
	v, ok := a.Value().(string)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to String.", a.Value()))
	}
	return v
}

// Adapter for string types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to string.
// Returns (value, false) if type is of string type.
// Returns (default value, true) if type is not of string type.
func (a *BoundedAtStartItr) StringOr(def string) (string, bool) {
	v, ok := a.Value().(string)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for string types on Bounded iterators. Provides a typecasting method 
// from interface{} to string.
// Panics if value is not of string type.
func (a *BoundedItr) String() string {
	v, ok := a.Value().(string)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to String.", a.Value()))
	}
	return v
}

// Adapter for string types on Bounded iterators. Provides a typecasting method 
// from interface{} to string.
// Returns (value, false) if type is of string type.
// Returns (default value, true) if type is not of string type.
func (a *BoundedItr) StringOr(def string) (string, bool) {
	v, ok := a.Value().(string)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for string types on RandomAccess iterator. Provides a typecasting method 
// from interface{} to string.
// Panics if value is not of string type.
func (a *RandomAccessItr) String() string {
	v, ok := a.Value().(string)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to String.", a.Value()))
	}
	return v
}

// Adapter for string types on RandomAccess iterators. Provides a typecasting method 
// from interface{} to string.
// Returns (value, false) if type is of string type.
// Returns (default value, true) if type is not of string type.
func (a *RandomAccessItr) StringOr(def string) (string, bool) {
	v, ok := a.Value().(string)
	if !ok {
		return def, true
	}
	return v, false
}
