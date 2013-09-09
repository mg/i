package i

import (
	"testing"
)

func mapfunc(itr Iterator) interface{} {
	if _, ok := itr.Value().(bool); ok {
		return "bool"
	}
	if _, ok := itr.Value().(int); ok {
		return "int"
	}
	if _, ok := itr.Value().(string); ok {
		return "string"
	}
	if _, ok := itr.Value().(float64); ok {
		return "float64"
	}
	return "unkown"
}

func TestMap(t *testing.T) {
	itr := Map(mapfunc, List(123, true, "this", 45.4, -1, 1+1i))
	for ; !itr.AtEnd(); itr.Next() {
		t.Log(itr.Value())
	}
}
