package iutil

import (
	"testing"
)

func TestRange(t *testing.T) {
	itr := Range(5, 10)
	for ; !itr.AtEnd(); itr.Next() {
		t.Log(itr.Value())
	}
}
