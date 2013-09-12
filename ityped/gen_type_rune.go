// Generated file for Rune iterator
package ityped

import (
	"github.com/mg/i"
)	

type (
	// Declare interface for an iterator that has an 
	// Rune() method that will return the value as rune
	Rune interface {
		Rune() rune
	}

	// Declare forward iterator interface for iterators that 
	// can return a rune value
	ForwardRune interface {
		i.Forward
		Rune
	}

	// Declare bidirectional iterator interface for iterators 
	// that can return a rune value
	BiDirectionalRune interface {
		i.BiDirectional
		Rune
	}

	// Declare bounded-at-start iterator interface for iterators 
	// that can return a rune value
	BoundedAtStartRune interface {
		i.BoundedAtStart
		Rune
	}

	// Declare bounded iterator interface for iterators 
	// that can return a rune value
	BoundedRune interface {
		i.Bounded
		Rune
	}

	// Declare random access iterator interface for iterators 
	// that can return a rune value
	RandomAccessRune interface {
		i.RandomAccess
		Rune
	}
)
