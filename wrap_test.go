package i

import (
	"testing"
)

func TestWrapForward(t *testing.T) {
	list := List(1, 2, 3, 4, 5)
	AssertForward(t, WrapForward(WrapForward(list)), 5, Strict)
	list.SetError(nil)
	list.First()
	AssertIteration(
		t, WrapForward(WrapForward(list)),
		1, 2, 3, 4, 5)

	list.SetError(nil)
	list.First()
	w1 := Filter(func(itr Iterator) bool {
		v, _ := itr.Value().(int)
		return v < 4
	}, list)
	w2 := Map(func(itr Iterator) interface{} {
		v, _ := itr.Value().(int)
		return v * 2
	}, w1)
	AssertForward(t, w2, 3, Strict)
	list.SetError(nil)
	list.First()
	w1 = Filter(func(itr Iterator) bool {
		v, _ := itr.Value().(int)
		return v < 4
	}, list)
	w2 = Map(func(itr Iterator) interface{} {
		v, _ := itr.Value().(int)
		return v * 2
	}, w1)
	AssertIteration(t, w2, 2, 4, 6)
}
