package igen

import (
	"math"
	"math/rand"
	"time"
)

// Random iterator
type Rnd struct {
	seed int64
	cur  interface{}
	n    int64
	gen  *rand.Rand
}

func Rand() *Rnd {
	var r Rnd
	r.seed = time.Now().UnixNano()
	r.First()
	return &r
}

func RandInt(n int64) *Rnd {
	var r Rnd
	r.seed = time.Now().UnixNano()
	r.n = n
	r.First()
	return &r
}

func (r *Rnd) First() error {
	r.gen = rand.New(rand.NewSource(r.seed))
	r.Next()
	return nil
}

func (r *Rnd) Last() error {
	return nil
}

func (r *Rnd) AtEnd() bool {
	return false
}

func (r *Rnd) AtStart() bool {
	return false
}

func (r *Rnd) Next() error {
	if r.n > 0 {
		r.cur = r.gen.Int63n(r.n)
	} else {
		r.cur = r.gen.Float64()
	}
	return nil
}

func (r *Rnd) Prev() error {
	return r.Next()
}

func (r *Rnd) Goto(idx int) error {
	return r.Next()
}

func (r *Rnd) Len() int {
	return math.MaxInt64
}

func (r *Rnd) Value() interface{} {
	return r.cur
}

func (r *Rnd) Float64() float64 {
	f, _ := r.cur.(float64)
	return f
}

func (r *Rnd) Int64() int64 {
	i, _ := r.cur.(int64)
	return i
}

func (r *Rnd) Error() error {
	return nil
}

func (r *Rnd) SetSeed(seed int64) {
	r.seed = seed
}

func (r *Rnd) Seed() int64 {
	return r.seed
}
