// Generated file for Complex64 iterator
package ityped

import (
	"github.com/mg/i"
)	

type (
	// Declare interface for an iterator that has an 
	// Complex64() method that will return the value as complex64
	Complex64 interface {
		Complex64() complex64
	}

	// Declare forward iterator interface for iterators that 
	// can return a complex64 value
	ForwardComplex64 interface {
		i.Forward
		Complex64
	}

	// Declare bidirectional iterator interface for iterators 
	// that can return a complex64 value
	BiDirectionalComplex64 interface {
		i.BiDirectional
		Complex64
	}

	// Declare bounded-at-start iterator interface for iterators 
	// that can return a complex64 value
	BoundedAtStartComplex64 interface {
		i.BoundedAtStart
		Complex64
	}

	// Declare bounded iterator interface for iterators 
	// that can return a complex64 value
	BoundedComplex64 interface {
		i.Bounded
		Complex64
	}

	// Declare random access iterator interface for iterators 
	// that can return a complex64 value
	RandomAccessComplex64 interface {
		i.RandomAccess
		Complex64
	}
)
