package i

import (
	"testing"
)

func filterfunc(itr Iterator) bool {
	if _, ok := itr.Value().(bool); ok {
		return true
	}
	if _, ok := itr.Value().(int); ok {
		return true
	}
	if _, ok := itr.Value().(string); ok {
		return false
	}
	if _, ok := itr.Value().(float64); ok {
		return true
	}
	return false
}

func TestFilter(t *testing.T) {
	AssertForward(t, Filter(filterfunc, List(123, true, "this", 45.4, -1, 1+1i)), 4, Strict)
	AssertIteration(
		t, Filter(filterfunc, List(123, true, "ssss", 45.4, -1, 1+1i)),
		123, true, 45.4, -1)
}
