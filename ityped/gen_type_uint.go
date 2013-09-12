// Generated file for Uint iterator
package ityped

import (
	"github.com/mg/i"
)	

type (
	// Declare interface for an iterator that has an 
	// Uint() method that will return the value as uint
	Uint interface {
		Uint() uint
	}

	// Declare forward iterator interface for iterators that 
	// can return a uint value
	ForwardUint interface {
		i.Forward
		Uint
	}

	// Declare bidirectional iterator interface for iterators 
	// that can return a uint value
	BiDirectionalUint interface {
		i.BiDirectional
		Uint
	}

	// Declare bounded-at-start iterator interface for iterators 
	// that can return a uint value
	BoundedAtStartUint interface {
		i.BoundedAtStart
		Uint
	}

	// Declare bounded iterator interface for iterators 
	// that can return a uint value
	BoundedUint interface {
		i.Bounded
		Uint
	}

	// Declare random access iterator interface for iterators 
	// that can return a uint value
	RandomAccessUint interface {
		i.RandomAccess
		Uint
	}
)
