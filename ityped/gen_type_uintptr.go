// Generated file for UintPtr iterator
package ityped

import (
	"github.com/mg/i"
)	

type (
	// Declare interface for an iterator that has an 
	// UintPtr() method that will return the value as uintptr
	UintPtr interface {
		UintPtr() uintptr
	}

	// Declare forward iterator interface for iterators that 
	// can return a uintptr value
	FowardUintPtr interface {
		i.Forward
		UintPtr
	}

	// Declare bidirectional iterator interface for iterators 
	// that can return a uintptr value
	BiDirectionalUintPtr interface {
		i.BiDirectional
		UintPtr
	}

	// Declare bounded-at-start iterator interface for iterators 
	// that can return a uintptr value
	BoundedAtStartUintPtr interface {
		i.BoundedAtStart
		UintPtr
	}

	// Declare bounded iterator interface for iterators 
	// that can return a uintptr value
	BoundedUintPtr interface {
		i.Bounded
		UintPtr
	}

	// Declare random access iterator interface for iterators 
	// that can return a uintptr value
	RandomAccessUintPtr interface {
		i.RandomAccess
		UintPtr
	}
)
