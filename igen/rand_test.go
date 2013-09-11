package igen

import (
	"testing"
)

func TestRand(t *testing.T) {
	itr := Rand()
	for i := 0; i < 5; i++ {
		t.Log(itr.Value())
		itr.Next()
	}
}
