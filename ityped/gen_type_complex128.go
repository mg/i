// Generated file for Complex128 iterator
package ityped

import (
	"github.com/mg/i"
)	

type (
	// Declare interface for an iterator that has an 
	// Complex128() method that will return the value as complex128
	Complex128 interface {
		Complex128() complex128
	}

	// Declare forward iterator interface for iterators that 
	// can return a complex128 value
	ForwardComplex128 interface {
		i.Forward
		Complex128
	}

	// Declare bidirectional iterator interface for iterators 
	// that can return a complex128 value
	BiDirectionalComplex128 interface {
		i.BiDirectional
		Complex128
	}

	// Declare bounded-at-start iterator interface for iterators 
	// that can return a complex128 value
	BoundedAtStartComplex128 interface {
		i.BoundedAtStart
		Complex128
	}

	// Declare bounded iterator interface for iterators 
	// that can return a complex128 value
	BoundedComplex128 interface {
		i.Bounded
		Complex128
	}

	// Declare random access iterator interface for iterators 
	// that can return a complex128 value
	RandomAccessComplex128 interface {
		i.RandomAccess
		Complex128
	}
)
