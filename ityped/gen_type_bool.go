// Generated file for Bool iterator
package ityped

import (
	"github.com/mg/i"
)	

type (
	// Declare interface for an iterator that has an 
	// Bool() method that will return the value as bool
	Bool interface {
		Bool() bool
	}

	// Declare forward iterator interface for iterators that 
	// can return a bool value
	ForwardBool interface {
		i.Forward
		Bool
	}

	// Declare bidirectional iterator interface for iterators 
	// that can return a bool value
	BiDirectionalBool interface {
		i.BiDirectional
		Bool
	}

	// Declare bounded-at-start iterator interface for iterators 
	// that can return a bool value
	BoundedAtStartBool interface {
		i.BoundedAtStart
		Bool
	}

	// Declare bounded iterator interface for iterators 
	// that can return a bool value
	BoundedBool interface {
		i.Bounded
		Bool
	}

	// Declare random access iterator interface for iterators 
	// that can return a bool value
	RandomAccessBool interface {
		i.RandomAccess
		Bool
	}
)
