package i

import (
	"testing"
)

func TestInterfaces(t *testing.T) {
	AssertRandomAccess(t, Interfaces([]interface{}{1, "string", 3.56, 4}), Strict)
	AssertIteration(
		t, Interfaces([]interface{}{1, "string", 3.56, 4}),
		1, "string", 3.56, 4)
}

func TestList(t *testing.T) {
	AssertRandomAccess(t, List(1, "string", 3.56, 4), Strict)
	AssertIteration(t, List(1, "string", 3.56, 4), 1, "string", 3.56, 4)
}
