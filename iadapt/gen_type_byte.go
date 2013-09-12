//  Generated file for Byte adapter
package iadapt

import "fmt"

// Adapter for byte types on Forward iterators. Provides a typecasting method 
// from interface{} to byte.
// Panics if value is not of byte type.
func (a *ForwardItr) Byte() byte {
	v, ok := a.Value().(byte)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Byte.", a.Value()))
	}
	return v
}

// Adapter for byte types on Forward iterators. Provides a typecasting method 
// from interface{} to byte.
// Returns (value, false) if type is of byte type.
// Returns (default value, true) if type is not of byte type.
func (a *ForwardItr) ByteOr(def byte) (byte, bool) {
	v, ok := a.Value().(byte)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for byte types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to byte.
// Panics if value is not of byte type.
func (a *BiDirectionalItr) Byte() byte {
	v, ok := a.Value().(byte)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Byte.", a.Value()))
	}
	return v
}

// Adapter for byte types on BiDirectional iterators. Provides a typecasting method 
// from interface{} to byte.
// Returns (value, false) if type is of byte type.
// Returns (default value, true) if type is not of byte type.
func (a *BiDirectionalItr) ByteOr(def byte) (byte, bool) {
	v, ok := a.Value().(byte)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for byte types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to byte.
// Panics if value is not of byte type.
func (a *BoundedAtStartItr) Byte() byte {
	v, ok := a.Value().(byte)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Byte.", a.Value()))
	}
	return v
}

// Adapter for byte types on BoundedAtStart iterators. Provides a typecasting method 
// from interface{} to byte.
// Returns (value, false) if type is of byte type.
// Returns (default value, true) if type is not of byte type.
func (a *BoundedAtStartItr) ByteOr(def byte) (byte, bool) {
	v, ok := a.Value().(byte)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for byte types on Bounded iterators. Provides a typecasting method 
// from interface{} to byte.
// Panics if value is not of byte type.
func (a *BoundedItr) Byte() byte {
	v, ok := a.Value().(byte)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Byte.", a.Value()))
	}
	return v
}

// Adapter for byte types on Bounded iterators. Provides a typecasting method 
// from interface{} to byte.
// Returns (value, false) if type is of byte type.
// Returns (default value, true) if type is not of byte type.
func (a *BoundedItr) ByteOr(def byte) (byte, bool) {
	v, ok := a.Value().(byte)
	if !ok {
		return def, true
	}
	return v, false
}

// Adapter for byte types on RandomAccess iterator. Provides a typecasting method 
// from interface{} to byte.
// Panics if value is not of byte type.
func (a *RandomAccessItr) Byte() byte {
	v, ok := a.Value().(byte)
	if !ok {
		panic(fmt.Errorf("Failed to typeassert value %v to Byte.", a.Value()))
	}
	return v
}

// Adapter for byte types on RandomAccess iterators. Provides a typecasting method 
// from interface{} to byte.
// Returns (value, false) if type is of byte type.
// Returns (default value, true) if type is not of byte type.
func (a *RandomAccessItr) ByteOr(def byte) (byte, bool) {
	v, ok := a.Value().(byte)
	if !ok {
		return def, true
	}
	return v, false
}
