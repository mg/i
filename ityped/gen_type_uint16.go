// Generated file for Uint16 iterator
package ityped

import (
	"github.com/mg/i"
)	

type (
	// Declare interface for an iterator that has an 
	// Uint16() method that will return the value as uint16
	Uint16 interface {
		Uint16() uint16
	}

	// Declare forward iterator interface for iterators that 
	// can return a uint16 value
	ForwardUint16 interface {
		i.Forward
		Uint16
	}

	// Declare bidirectional iterator interface for iterators 
	// that can return a uint16 value
	BiDirectionalUint16 interface {
		i.BiDirectional
		Uint16
	}

	// Declare bounded-at-start iterator interface for iterators 
	// that can return a uint16 value
	BoundedAtStartUint16 interface {
		i.BoundedAtStart
		Uint16
	}

	// Declare bounded iterator interface for iterators 
	// that can return a uint16 value
	BoundedUint16 interface {
		i.Bounded
		Uint16
	}

	// Declare random access iterator interface for iterators 
	// that can return a uint16 value
	RandomAccessUint16 interface {
		i.RandomAccess
		Uint16
	}
)
