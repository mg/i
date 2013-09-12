// Generated file for Int16 iterator
package ityped

import (
	"github.com/mg/i"
)	

type (
	// Declare interface for an iterator that has an 
	// Int16() method that will return the value as int16
	Int16 interface {
		Int16() int16
	}

	// Declare forward iterator interface for iterators that 
	// can return a int16 value
	ForwardInt16 interface {
		i.Forward
		Int16
	}

	// Declare bidirectional iterator interface for iterators 
	// that can return a int16 value
	BiDirectionalInt16 interface {
		i.BiDirectional
		Int16
	}

	// Declare bounded-at-start iterator interface for iterators 
	// that can return a int16 value
	BoundedAtStartInt16 interface {
		i.BoundedAtStart
		Int16
	}

	// Declare bounded iterator interface for iterators 
	// that can return a int16 value
	BoundedInt16 interface {
		i.Bounded
		Int16
	}

	// Declare random access iterator interface for iterators 
	// that can return a int16 value
	RandomAccessInt16 interface {
		i.RandomAccess
		Int16
	}
)
