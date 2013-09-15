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

package itk

import (
	"github.com/mg/i"
)

type WForward struct {
	value  func() interface{}
	err    func() error
	seterr func(error)
	next   func() error
	atEnd  func() bool
}

func WrapForward(itr i.Forward) *WForward {
	wf := WForward{}
	if wrapped, ok := itr.(*WForward); ok {
		wf.value = wrapped.value
		wf.err = wrapped.err
		wf.seterr = wrapped.seterr
		wf.next = wrapped.next
		wf.atEnd = wrapped.atEnd
	} else {
		wf.value = itr.Value
		wf.err = itr.Error
		wf.seterr = itr.SetError
		wf.next = itr.Next
		wf.atEnd = itr.AtEnd
	}
	return &wf
}

func (wf *WForward) Value() interface{} {
	return wf.value()
}

func (wf *WForward) Error() error {
	return wf.err()
}

func (wf *WForward) SetError(err error) {
	wf.seterr(err)
}

func (wf *WForward) Next() error {
	return wf.next()
}

func (wf *WForward) AtEnd() bool {
	return wf.atEnd()
}

// WrapBiDirectional
type WBiDirectional struct {
	WForward
	prev    func() error
	atStart func() bool
}

func WrapBiDirectional(itr i.BiDirectional) *WBiDirectional {
	wbd := WBiDirectional{}
	wbd.WForward = *(WrapForward(itr))

	if wrapped, ok := itr.(*WBiDirectional); ok {
		wbd.prev = wrapped.prev
		wbd.atStart = wrapped.atStart
	} else {
		wbd.prev = itr.Prev
		wbd.atStart = itr.AtStart
	}

	return &wbd
}

func (wbd *WBiDirectional) Prev() error {
	return wbd.prev()
}

func (wbd *WBiDirectional) AtStart() bool {
	return wbd.atStart()
}

// WrapBoundedAtStart
type WBoundedAtStart struct {
	WForward
	first func() error
}

func WrapBoundedAtStart(itr i.BoundedAtStart) *WBoundedAtStart {
	wbas := WBoundedAtStart{}
	wbas.WForward = *(WrapForward(itr))

	if wrapped, ok := itr.(*WBoundedAtStart); ok {
		wbas.first = wrapped.first
	} else {
		wbas.first = itr.First
	}

	return &wbas
}

func (wbas *WBoundedAtStart) First() error {
	return wbas.first()
}

// WrapBounded
type WBounded struct {
	WBiDirectional
	first func() error
	last  func() error
}

func WrapBounded(itr i.Bounded) *WBounded {
	wb := WBounded{}
	wb.WBiDirectional = *(WrapBiDirectional(itr))

	if wrapped, ok := itr.(*WBounded); ok {
		wb.first = wrapped.first
		wb.last = wrapped.last
	} else {
		wb.first = itr.First
		wb.last = itr.Last
	}

	return &wb
}

func (wb *WBounded) First() error {
	return wb.first()
}

func (wb *WBounded) Last() error {
	return wb.last()
}

// WrapRandomAccess
type WRandomAccess struct {
	WBounded
	goTo   func(int) error
	length func() int
}

func WrapRandomAccess(itr i.RandomAccess) *WRandomAccess {
	wra := WRandomAccess{}
	wra.WBounded = *(WrapBounded(itr))

	if wrapped, ok := itr.(*WRandomAccess); ok {
		wra.goTo = wrapped.goTo
		wra.length = wrapped.length
	} else {
		wra.goTo = itr.Goto
		wra.length = itr.Len
	}

	return &wra
}

func (wra *WRandomAccess) Goto(idx int) error {
	return wra.goTo(idx)
}

func (wra *WRandomAccess) Len() int {
	return wra.length()
}
