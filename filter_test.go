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
	itr := Filter(filterfunc, List(123, true, "this", 45.4, -1, 1+1i))
	for ; !itr.AtEnd(); itr.Next() {
		t.Log(itr.Value())
	}
}
