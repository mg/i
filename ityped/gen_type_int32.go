// Generated file for Int32 iterator
package ityped

import (
	"github.com/mg/i"
)	

type (
	// Declare interface for an iterator that has an 
	// Int32() method that will return the value as int32
	Int32 interface {
		Int32() int32
	}

	// Declare forward iterator interface for iterators that 
	// can return a int32 value
	ForwardInt32 interface {
		i.Forward
		Int32
	}

	// Declare bidirectional iterator interface for iterators 
	// that can return a int32 value
	BiDirectionalInt32 interface {
		i.BiDirectional
		Int32
	}

	// Declare bounded-at-start iterator interface for iterators 
	// that can return a int32 value
	BoundedAtStartInt32 interface {
		i.BoundedAtStart
		Int32
	}

	// Declare bounded iterator interface for iterators 
	// that can return a int32 value
	BoundedInt32 interface {
		i.Bounded
		Int32
	}

	// Declare random access iterator interface for iterators 
	// that can return a int32 value
	RandomAccessInt32 interface {
		i.RandomAccess
		Int32
	}
)
