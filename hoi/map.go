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

package hoi

import (
	"github.com/mg/i"
	"github.com/mg/i/itk"
)

// A function that will provide the mapping computation for
// the Map iterator.
type MapFunc func(i.Iterator) interface{}

type imap struct {
	mapf MapFunc
	val  interface{}
}

type fmap struct {
	imap
	itk.WForward
}

// A forward map iterator that maps a underlying data stream to a new
// data stream according to the computation provied by the
// mapping function.
func Map(mapf MapFunc, itr i.Forward) i.Forward {
	m := fmap{}
	m.mapf = mapf
	m.WForward = *(itk.WrapForward(itr))
	return &m
}

func (m *fmap) Next() error {
	m.val = nil
	return m.WForward.Next()
}

func (m *fmap) Value() interface{} {
	if m.val == nil {
		m.val = m.mapf(&m.WForward)
	}
	return m.val
}

type bimap struct {
	imap
	itk.WBiDirectional
}

// A bidirectional map iterator that maps a underlying data stream to a new
// data stream according to the computation provied by the
// mapping function.
func BiMap(mapf MapFunc, itr i.BiDirectional) i.BiDirectional {
	m := bimap{}
	m.mapf = mapf
	m.WBiDirectional = *(itk.WrapBiDirectional(itr))
	return &m
}

func (m *bimap) Next() error {
	m.val = nil
	return m.WBiDirectional.Next()
}

func (m *bimap) Prev() error {
	m.val = nil
	return m.WBiDirectional.Prev()
}

func (m *bimap) Value() interface{} {
	if m.val == nil {
		m.val = m.mapf(&m.WBiDirectional)
	}
	return m.val
}

type basmap struct {
	imap
	itk.WBoundedAtStart
}

// A bounded at start map iterator that maps a underlying data stream to a new
// data stream according to the computation provied by the
// mapping function.
func BasMap(mapf MapFunc, itr i.BoundedAtStart) i.BoundedAtStart {
	m := basmap{}
	m.mapf = mapf
	m.WBoundedAtStart = *(itk.WrapBoundedAtStart(itr))
	return &m
}

func (m *basmap) First() error {
	m.val = nil
	return m.WBoundedAtStart.First()
}

func (m *basmap) Next() error {
	m.val = nil
	return m.WBoundedAtStart.Next()
}

func (m *basmap) Value() interface{} {
	if m.val == nil {
		m.val = m.mapf(&m.WBoundedAtStart)
	}
	return m.val
}

type bmap struct {
	imap
	itk.WBounded
}

// A bounded map iterator that maps a underlying data stream to a new
// data stream according to the computation provied by the
// mapping function.
func BMap(mapf MapFunc, itr i.Bounded) i.Bounded {
	m := bmap{}
	m.mapf = mapf
	m.WBounded = *(itk.WrapBounded(itr))
	return &m
}

func (m *bmap) First() error {
	m.val = nil
	return m.WBounded.First()
}

func (m *bmap) Last() error {
	m.val = nil
	return m.WBounded.Last()
}

func (m *bmap) Next() error {
	m.val = nil
	return m.WBounded.Next()
}

func (m *bmap) Prev() error {
	m.val = nil
	return m.WBounded.Prev()
}

func (m *bmap) Value() interface{} {
	if m.val == nil {
		m.val = m.mapf(&m.WBounded)
	}
	return m.val
}

type ramap struct {
	imap
	itk.WRandomAccess
}

// A random access map iterator that maps a underlying data stream to a new
// data stream according to the computation provied by the
// mapping function.
func RaMap(mapf MapFunc, itr i.RandomAccess) i.RandomAccess {
	m := ramap{}
	m.mapf = mapf
	m.WRandomAccess = *(itk.WrapRandomAccess(itr))
	return &m
}

func (m *ramap) First() error {
	m.val = nil
	return m.WRandomAccess.First()
}

func (m *ramap) Last() error {
	m.val = nil
	return m.WRandomAccess.Last()
}

func (m *ramap) Next() error {
	m.val = nil
	return m.WRandomAccess.Next()
}

func (m *ramap) Prev() error {
	m.val = nil
	return m.WRandomAccess.Prev()
}

func (m *ramap) Goto(idx int) error {
	m.val = nil
	return m.WRandomAccess.Goto(idx)
}

func (m *ramap) Value() interface{} {
	if m.val == nil {
		m.val = m.mapf(&m.WRandomAccess)
	}
	return m.val
}
