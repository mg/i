package i

import (
	"testing"
)

func TestEach(t *testing.T) {
	itr := List("one", "two", "three")
	Each(itr, func(i Iterator) bool {
		t.Log(itr.Value())
		return true
	})
}
