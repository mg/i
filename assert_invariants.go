package i

import (
	"testing"
)

const (
	Strict          = 0
	RelaxValue      = 1 << iota
	RelaxAtEnd      = 2 << iota
	RelaxValueAtEnd = 3 << iota
	RelaxValueEqual = 4 << iota
)

func isRelaxAtEnd(strictness int) bool {
	return strictness&RelaxAtEnd == RelaxAtEnd
}

func isRelaxValue(strictness int) bool {
	return strictness&RelaxValue == RelaxValue
}

func isRelaxValueAtEnd(strictness int) bool {
	return strictness&RelaxValueAtEnd == RelaxValueAtEnd
}

func isRelaxValueEqual(strictness int) bool {
	return strictness&RelaxValueEqual == RelaxValueEqual
}

// Make assertions that must hold for Forward iterators.
// If the datastream is infite, pass -1 for count.
func AssertForward(t *testing.T, itr Forward, count, strictness int) {

	// Go for 10 iterations if stream is infinite
	length := 10
	if count >= 0 {
		length = count
	}

	// Iterator must have value at start
	itr.Value()
	if e := itr.Error(); e != nil {
		t.Fatalf("Iterator reprots error  %v at start after calling Value()", e)
	}

	for i := 0; i < length; i++ {
		// Iterator should not be at end
		if itr.AtEnd() {
			t.Fatalf("Iterator reports at end after %d iteratons of %d", i, length)
		}

		// Value() should be indempotent
		for j := 0; j < 3; j++ {
			v := itr.Value()
			if e := itr.Error(); e != nil {
				t.Fatalf("Iterator reports error %v after fetching value %v for the %d time at %d iteration of %d", e, v, j, i, length)
			}
		}

		// Value() should be indempotent
		if !isRelaxValueEqual(strictness) {
			v1 := itr.Value()
			v2 := itr.Value()
			if v1 != v2 {
				t.Fatalf("Iterator returns different values %v and %v at iteration %d", v1, v2, i)
			}
		}

		itr.Next()
	}

	// Check if iterator is at end
	if count >= 0 && !itr.AtEnd() {
		t.Fatalf("Iterator should be at end, is not after %d iterations", length)
	}

	// Check if there where any errors
	if e := itr.Error(); e != nil {
		t.Fatalf("Iterator reports error %q at end after %d iterations", e, length)
	}

	// Iterator must return error if we try to access value at end
	if !isRelaxValueAtEnd(strictness) {
		itr.Value()
		if e := itr.Error(); count >= 0 && e == nil {
			t.Fatalf("Iterator does not report error at end after %d iterations after calling Value()", length)
		}
	}

	// Iterator must return error if we try to move past end
	if e := itr.Next(); count >= 0 && e == nil {
		t.Fatalf("Iterator does not report error at end after %d iterations after calling Next()", length)
	}

}

// Make assertions that must hold for BiDirectional iterators.
// If the datastream is infite, pass -1 for count.
func AssertBiDirectional(t *testing.T, itr BiDirectional, count, strictness int) {

	// Go for 10 iterations if stream is infinite
	length := 10
	if count >= 0 {
		length = count
	}

	// Iterator must be at start
	if !itr.AtStart() {
		t.Fatalf("Iterator reports not at start after 0 iterations")
	}

	// Iterator must have value at start
	itr.Value()
	if e := itr.Error(); e != nil {
		t.Fatalf("Iterator reprots error  %v at start after calling Value()", e)
	}

	for i := 0; i < length; i++ {
		// Iterator should not be at end
		if itr.AtEnd() {
			t.Fatalf("Iterator reports at end after %d iteratons of %d", i, length)
		}

		// Value() should be indempotent
		for j := 0; j < 3; j++ {
			v := itr.Value()
			if e := itr.Error(); e != nil {
				t.Fatalf("Iterator reports error %v after fetching value %v for the %d time at %d iteration of %d", e, v, j, i, length)
			}
		}

		// Value() should be indempotent
		if !isRelaxValueEqual(strictness) {
			v1 := itr.Value()
			v2 := itr.Value()
			if v1 != v2 {
				t.Fatalf("Iterator returns different values %v and %v at iteration %d", v1, v2, i)
			}
		}

		// Check if iterator returns same value after moving forward and backward
		if !isRelaxValueEqual(strictness) {
			v1 := itr.Value()
			itr.Next()
			itr.Prev()
			v2 := itr.Value()
			if v1 != v2 {
				t.Fatalf("%v is not equal to %v after %d iterations", v1, v2, i)

			}
		}
		itr.Next()
	}

	// Check if iterator is at end
	if count >= 0 && !itr.AtEnd() {
		t.Fatalf("Iterator should be at end, is not after %d iterations", length)
	}

	// Check if there where any errors
	if e := itr.Error(); e != nil {
		t.Fatalf("Iterator reports error %v at end after %d iterations", e, length)
	}

	// Iterator must return error if we try to access value at end
	if !isRelaxValueAtEnd(strictness) {
		itr.Value()
		if e := itr.Error(); count >= 0 && e == nil {
			t.Fatalf("Iterator does not report error at end after %d iterations after calling Value()", length)
		}
	}

	// Iterator must return error if we try to move past end
	if e := itr.Next(); count >= 0 && e == nil {
		t.Fatalf("Iterator does not report error at end after %d iterations after calling Next()", length)
	}

}

// Make assertions that must hold for BoundedAtStart iterators.
// If the datastream is infite, pass -1 for count.
func AssertBoundedAtStart(t *testing.T, itr BoundedAtStart, count, strictness int) {
	// Iterator should pass all tests for Forward iterators twice
	AssertForward(t, itr, count, strictness)
	itr.First()
	AssertForward(t, itr, count, strictness)
}

// Make assertions that must hold for BoundedAtStart iterators.
func AssertBounded(t *testing.T, itr Bounded, count, strictness int) {
	itr.Last()
	itr.Next()
	// Check if iterator is at end
	if !itr.AtEnd() {
		t.Fatalf("Iterator should be at end, is not after calling Last()")
	}

	// Iterator should pass all tests for AssertBiDirectional iterators
	itr.SetError(nil)
	itr.First()
	AssertBiDirectional(t, itr, count, strictness)

}

// Make assertions that must hold for RandomAccess iterators.
func AssertRandomAccess(t *testing.T, itr RandomAccess, strictness int) {
	AssertBounded(t, itr, itr.Len(), strictness)

}

// Make assertions of iterations
func AssertIteration(t *testing.T, itr Forward, values ...interface{}) {
	for i, v := range values {
		v2 := itr.Value()
		if v != v2 {
			t.Fatalf("Iteration assertion failed, expected %v, got %v at iteration %d", v, v2, i)
		}
		itr.Next()
	}
}
