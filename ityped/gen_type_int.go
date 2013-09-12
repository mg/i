// Generated file for Int iterator
package ityped

import (
	"github.com/mg/i"
)	

type (
	// Declare interface for an iterator that has an 
	// Int() method that will return the value as int
	Int interface {
		Int() int
	}

	// Declare forward iterator interface for iterators that 
	// can return a int value
	ForwardInt interface {
		i.Forward
		Int
	}

	// Declare bidirectional iterator interface for iterators 
	// that can return a int value
	BiDirectionalInt interface {
		i.BiDirectional
		Int
	}

	// Declare bounded-at-start iterator interface for iterators 
	// that can return a int value
	BoundedAtStartInt interface {
		i.BoundedAtStart
		Int
	}

	// Declare bounded iterator interface for iterators 
	// that can return a int value
	BoundedInt interface {
		i.Bounded
		Int
	}

	// Declare random access iterator interface for iterators 
	// that can return a int value
	RandomAccessInt interface {
		i.RandomAccess
		Int
	}
)
