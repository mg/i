package iutil

import "fmt"

type rng struct {
	m, n, cur int
	err       error
}

func Range(m, n int) *rng {
	return &rng{m: m, n: n, cur: m}
}

func (i *rng) Value() interface{} {
	return i.cur
}

func (i *rng) Int() int {
	return i.cur
}

func (i *rng) Error() error {
	return i.err
}

func (i *rng) Next() error {
	i.cur++
	return i.err
}

func (i *rng) AtEnd() bool {
	return i.cur >= i.n
}

func (i *rng) Prev() error {
	i.cur--
	return nil
}

func (i *rng) AtStart() bool {
	return i.cur <= i.m
}

func (i *rng) First() error {
	i.cur = i.m
	return nil
}

func (i *rng) Last() error {
	i.cur = i.n
	return nil
}

func (i *rng) Goto(idx int) error {
	if idx < i.m || idx >= i.n {
		i.err = fmt.Errorf("Index %d out of bounds [%d, %d)", idx, i.m, i.n)
		return i.err
	}
	i.cur = idx
	return nil
}

func (i *rng) Len() int {
	return i.n - i.m
}
