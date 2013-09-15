// The MIT License (MIT)
//
// Copyright (c) 2013 Magnús Örn Gylfason
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package igen

import (
	"math"
	"math/rand"
	"time"
)

// Rrnd is an random access iterator that generates random numbers each time
// it is moved with Next(), Prev(), Last() or Goto(). Those number are then
// available through the Value() method (or the Foat64() or Int64() methods).
// The First() method will reset the random sequence to its originating seed
// value. It comes in two variants, Rand() which generates random floating
// numbers between 0 and 1, and RandInt(n) which generates random integers
// between 0 and n. The custom Seed() and SetSeed() methods can be used
// to save a random sequence and then rerun it at a later time. The stream of
// numbers is infinite in both directions, the AtStart() method and AtEnd()
// method always return false.
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
	return r.Next()
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

func (r *Rnd) SetError(err error) {
}

func (r *Rnd) SetSeed(seed int64) {
	r.seed = seed
}

func (r *Rnd) Seed() int64 {
	return r.seed
}
