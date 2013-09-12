// Generated file for Uint32 iterator
package ityped

import (
	"github.com/mg/i"
)	

type (
	// Declare interface for an iterator that has an 
	// Uint32() method that will return the value as uint32
	Uint32 interface {
		Uint32() uint32
	}

	// Declare forward iterator interface for iterators that 
	// can return a uint32 value
	ForwardUint32 interface {
		i.Forward
		Uint32
	}

	// Declare bidirectional iterator interface for iterators 
	// that can return a uint32 value
	BiDirectionalUint32 interface {
		i.BiDirectional
		Uint32
	}

	// Declare bounded-at-start iterator interface for iterators 
	// that can return a uint32 value
	BoundedAtStartUint32 interface {
		i.BoundedAtStart
		Uint32
	}

	// Declare bounded iterator interface for iterators 
	// that can return a uint32 value
	BoundedUint32 interface {
		i.Bounded
		Uint32
	}

	// Declare random access iterator interface for iterators 
	// that can return a uint32 value
	RandomAccessUint32 interface {
		i.RandomAccess
		Uint32
	}
)
