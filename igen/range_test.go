package igen

import (
	"github.com/mg/i"
	"testing"
)

func TestRange(t *testing.T) {
	i.AssertRandomAccess(t, Range(5, 10), i.Strict)
	i.AssertIteration(t, Range(5, 10), 5, 6, 7, 8, 9)
}
