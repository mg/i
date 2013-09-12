// Generated file for Float64 iterator
package ityped

import (
	"github.com/mg/i"
)	

type (
	// Declare interface for an iterator that has an 
	// Float64() method that will return the value as float64
	Float64 interface {
		Float64() float64
	}

	// Declare forward iterator interface for iterators that 
	// can return a float64 value
	ForwardFloat64 interface {
		i.Forward
		Float64
	}

	// Declare bidirectional iterator interface for iterators 
	// that can return a float64 value
	BiDirectionalFloat64 interface {
		i.BiDirectional
		Float64
	}

	// Declare bounded-at-start iterator interface for iterators 
	// that can return a float64 value
	BoundedAtStartFloat64 interface {
		i.BoundedAtStart
		Float64
	}

	// Declare bounded iterator interface for iterators 
	// that can return a float64 value
	BoundedFloat64 interface {
		i.Bounded
		Float64
	}

	// Declare random access iterator interface for iterators 
	// that can return a float64 value
	RandomAccessFloat64 interface {
		i.RandomAccess
		Float64
	}
)
