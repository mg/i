// Generated file for Byte iterator
package ityped

import (
	"github.com/mg/i"
)	

type (
	// Declare interface for an iterator that has an 
	// Byte() method that will return the value as byte
	Byte interface {
		Byte() byte
	}

	// Declare forward iterator interface for iterators that 
	// can return a byte value
	ForwardByte interface {
		i.Forward
		Byte
	}

	// Declare bidirectional iterator interface for iterators 
	// that can return a byte value
	BiDirectionalByte interface {
		i.BiDirectional
		Byte
	}

	// Declare bounded-at-start iterator interface for iterators 
	// that can return a byte value
	BoundedAtStartByte interface {
		i.BoundedAtStart
		Byte
	}

	// Declare bounded iterator interface for iterators 
	// that can return a byte value
	BoundedByte interface {
		i.Bounded
		Byte
	}

	// Declare random access iterator interface for iterators 
	// that can return a byte value
	RandomAccessByte interface {
		i.RandomAccess
		Byte
	}
)
