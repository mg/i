// The MIT License (MIT)
//
// Copyright (c) 2013 Magnús Örn Gylfason
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package i

type (
	Iterator interface {
		// Return value at current position in data.
		Value() interface{}

		// Return last error occurred.
		Error() error

		// Set last error
		SetError(error)
	}

	Forward interface {
		// Forward iterator is an Iterator.
		Iterator

		// Move to the next position in the data. Moving after the end
		// is reached is an error.
		Next() error

		// Check if the iterator is beyond the end of the data. This is an
		// invalid position in the data, e.g. fetching the value at the
		// end is an error.
		AtEnd() bool
	}

	BiDirectional interface {
		// BiDirectional is a Forward iterator.
		Forward

		// Move to the previous position in the data. Moving once beyond the
		// starting position is valid. Moveing after that is an error.
		Prev() error

		// Check if the iterator is at the start of the data. This is a valid
		// position in the data, e.g. fetching the value will return the first
		// value.
		AtStart() bool
	}

	BoundedAtStart interface {
		// BoundedAtStart is a Forward iterator.
		Forward

		// Jump to the first position in the stream. This is a valid position
		// in the stream. At this position AtStart() returns true.
		First() error
	}

	Bounded interface {
		// Bounded is a BiDirectional iterator.
		BiDirectional

		// Jump to the first position in the stream. This is a valid position
		// in the stream. At this position AtStart() returns true.
		First() error

		// jump to the last position in the stream. This is a valid position
		// in the stream. At this position AtEnd() returns false.
		Last() error
	}

	RandomAccess interface {
		// RandomAccess is a Bounded iterator.
		Bounded

		// Jump to any position in the stream. If the position is less than
		// First position or greater than Last position the method returns
		// an error.
		Goto(int) error

		// Get the length of the data stream.
		Len() int
	}
)
