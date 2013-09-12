// Generated file for Int8 iterator
package ityped

import (
	"github.com/mg/i"
)	

type (
	// Declare interface for an iterator that has an 
	// Int8() method that will return the value as int8
	Int8 interface {
		Int8() int8
	}

	// Declare forward iterator interface for iterators that 
	// can return a int8 value
	ForwardInt8 interface {
		i.Forward
		Int8
	}

	// Declare bidirectional iterator interface for iterators 
	// that can return a int8 value
	BiDirectionalInt8 interface {
		i.BiDirectional
		Int8
	}

	// Declare bounded-at-start iterator interface for iterators 
	// that can return a int8 value
	BoundedAtStartInt8 interface {
		i.BoundedAtStart
		Int8
	}

	// Declare bounded iterator interface for iterators 
	// that can return a int8 value
	BoundedInt8 interface {
		i.Bounded
		Int8
	}

	// Declare random access iterator interface for iterators 
	// that can return a int8 value
	RandomAccessInt8 interface {
		i.RandomAccess
		Int8
	}
)
