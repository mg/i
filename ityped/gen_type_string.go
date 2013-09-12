// Generated file for String iterator
package ityped

import (
	"github.com/mg/i"
)	

type (
	// Declare interface for an iterator that has an 
	// String() method that will return the value as string
	String interface {
		String() string
	}

	// Declare forward iterator interface for iterators that 
	// can return a string value
	ForwardString interface {
		i.Forward
		String
	}

	// Declare bidirectional iterator interface for iterators 
	// that can return a string value
	BiDirectionalString interface {
		i.BiDirectional
		String
	}

	// Declare bounded-at-start iterator interface for iterators 
	// that can return a string value
	BoundedAtStartString interface {
		i.BoundedAtStart
		String
	}

	// Declare bounded iterator interface for iterators 
	// that can return a string value
	BoundedString interface {
		i.Bounded
		String
	}

	// Declare random access iterator interface for iterators 
	// that can return a string value
	RandomAccessString interface {
		i.RandomAccess
		String
	}
)
