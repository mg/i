// Generated file for Int64 iterator
package ityped

import (
	"github.com/mg/i"
)	

type (
	// Declare interface for an iterator that has an 
	// Int64() method that will return the value as int64
	Int64 interface {
		Int64() int64
	}

	// Declare forward iterator interface for iterators that 
	// can return a int64 value
	ForwardInt64 interface {
		i.Forward
		Int64
	}

	// Declare bidirectional iterator interface for iterators 
	// that can return a int64 value
	BiDirectionalInt64 interface {
		i.BiDirectional
		Int64
	}

	// Declare bounded-at-start iterator interface for iterators 
	// that can return a int64 value
	BoundedAtStartInt64 interface {
		i.BoundedAtStart
		Int64
	}

	// Declare bounded iterator interface for iterators 
	// that can return a int64 value
	BoundedInt64 interface {
		i.Bounded
		Int64
	}

	// Declare random access iterator interface for iterators 
	// that can return a int64 value
	RandomAccessInt64 interface {
		i.RandomAccess
		Int64
	}
)
