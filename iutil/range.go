package iutil

import "fmt"

type rng struct {
	m, n, cur int
	err       error
}

func Range(m, n int) *rng {
	return &rng{m: m, n: n, cur: m}
}

func (r *rng) Value() interface{} {
	return r.cur
}

func (r *rng) Int() int {
	return r.cur
}

func (r *rng) Error() error {
	return r.err
}

func (r *rng) Next() error {
	r.cur++
	return r.err
}

func (r *rng) AtEnd() bool {
	return r.cur >= r.n
}

func (r *rng) Prev() error {
	r.cur--
	return nil
}

func (r *rng) AtStart() bool {
	return r.cur <= r.m
}

func (r *rng) First() error {
	r.cur = r.m
	return nil
}

func (r *rng) Last() error {
	r.cur = r.n
	return nil
}

func (r *rng) Goto(idx int) error {
	if idx < r.m || idx >= r.n {
		r.err = fmt.Errorf("Index %d out of bounds [%d, %d)", idx, r.m, r.n)
		return r.err
	}
	r.cur = idx
	return nil
}

func (r *rng) Len() int {
	return r.n - r.m
}
