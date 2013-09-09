package i

import (
	"testing"
)

func TestWrapForward(t *testing.T) {
	src := List(1, 2, 3, 4, 5)
	w1 := WrapForward(src)
	w2 := WrapForward(w1)
	for ; !w2.AtEnd(); w2.Next() {
		t.Log(w2.Value())
	}

	src.First()
	w3 := Filter(func(itr Iterator) bool {
		v, _ := itr.Value().(int)
		return v < 4
	}, src)
	w4 := Map(func(itr Iterator) interface{} {
		v, _ := itr.Value().(int)
		return v * 2
	}, w3)
	for ; !w4.AtEnd(); w4.Next() {
		t.Log(w4.Value())
	}

}
