package iadapt

import (
	"github.com/mg/i"
	"github.com/mg/i/islice"
	"testing"
)

func TestAdapt(t *testing.T) {
	list := islice.IntList(1, 2, 3, 4, 5, 6, 7, 8, 9)
	i.AssertForward(t, Forward(list), 9, i.Strict)
	list.First()
	i.AssertIteration(t, Forward(list), 1, 2, 3, 4, 5, 6, 7, 8, 9)
}
