// Generated file for Uint64 iterator
package ityped

import (
	"github.com/mg/i"
)	

type (
	// Declare interface for an iterator that has an 
	// Uint64() method that will return the value as uint64
	Uint64 interface {
		Uint64() uint64
	}

	// Declare forward iterator interface for iterators that 
	// can return a uint64 value
	ForwardUint64 interface {
		i.Forward
		Uint64
	}

	// Declare bidirectional iterator interface for iterators 
	// that can return a uint64 value
	BiDirectionalUint64 interface {
		i.BiDirectional
		Uint64
	}

	// Declare bounded-at-start iterator interface for iterators 
	// that can return a uint64 value
	BoundedAtStartUint64 interface {
		i.BoundedAtStart
		Uint64
	}

	// Declare bounded iterator interface for iterators 
	// that can return a uint64 value
	BoundedUint64 interface {
		i.Bounded
		Uint64
	}

	// Declare random access iterator interface for iterators 
	// that can return a uint64 value
	RandomAccessUint64 interface {
		i.RandomAccess
		Uint64
	}
)
