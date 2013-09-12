//  Generated file for UintPtr adapter
package iadapt

import "fmt"

// Adapter for uintptr types. Provides a typecasting method 
// from interface{} to uintptr.
// Panics if value is not of uintptr type.
func (a *ForwardItr) UintPtr() uintptr {
	v, ok := a.Value().(uintptr)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to UintPtr.", a.Value()))
	}
	return v
}

// Adapter for uintptr types. Provides a typecasting method 
// from interface{} to uintptr.
// Returns (value, false) if type is of uintptr type.
// Returns (default value, true) if type is not of uintptr type.
func (a *ForwardItr) UintPtrOr(def uintptr) (uintptr, bool) {
	v, ok := a.Value().(uintptr)
	if !ok {
		return def, true
	}
	return v, false
}
