// Generated file for Uint8 iterator
package ityped

import (
	"github.com/mg/i"
)	

type (
	// Declare interface for an iterator that has an 
	// Uint8() method that will return the value as uint8
	Uint8 interface {
		Uint8() uint8
	}

	// Declare forward iterator interface for iterators that 
	// can return a uint8 value
	ForwardUint8 interface {
		i.Forward
		Uint8
	}

	// Declare bidirectional iterator interface for iterators 
	// that can return a uint8 value
	BiDirectionalUint8 interface {
		i.BiDirectional
		Uint8
	}

	// Declare bounded-at-start iterator interface for iterators 
	// that can return a uint8 value
	BoundedAtStartUint8 interface {
		i.BoundedAtStart
		Uint8
	}

	// Declare bounded iterator interface for iterators 
	// that can return a uint8 value
	BoundedUint8 interface {
		i.Bounded
		Uint8
	}

	// Declare random access iterator interface for iterators 
	// that can return a uint8 value
	RandomAccessUint8 interface {
		i.RandomAccess
		Uint8
	}
)
