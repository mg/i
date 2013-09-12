// Generated file for Float32 iterator
package ityped

import (
	"github.com/mg/i"
)	

type (
	// Declare interface for an iterator that has an 
	// Float32() method that will return the value as float32
	Float32 interface {
		Float32() float32
	}

	// Declare forward iterator interface for iterators that 
	// can return a float32 value
	ForwardFloat32 interface {
		i.Forward
		Float32
	}

	// Declare bidirectional iterator interface for iterators 
	// that can return a float32 value
	BiDirectionalFloat32 interface {
		i.BiDirectional
		Float32
	}

	// Declare bounded-at-start iterator interface for iterators 
	// that can return a float32 value
	BoundedAtStartFloat32 interface {
		i.BoundedAtStart
		Float32
	}

	// Declare bounded iterator interface for iterators 
	// that can return a float32 value
	BoundedFloat32 interface {
		i.Bounded
		Float32
	}

	// Declare random access iterator interface for iterators 
	// that can return a float32 value
	RandomAccessFloat32 interface {
		i.RandomAccess
		Float32
	}
)
