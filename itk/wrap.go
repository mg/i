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

// Compose this structure in your iterator to wrap a Forward iterator
// within your own iterator. It contains function pointers that point
// to all the methods of the wrapped Forward iterator.
type WForward struct {
	value  func() interface{}
	err    func() error
	seterr func(error)
	next   func() error
	atEnd  func() bool
}

// Wrap a Forward iterator. If ITR is your iterator type and F is
// the iterator your wish to wrap, the usage is
//
// ITR.WForward= *(WrapForward(F))
//
// Now if you don't e.g. provide an implementation for the Value()
// method, a call to ITR.Value() will be forwarded to F.Value() by
// WForward. To guard against a chain of useless forwards, the
// WrapForward checks if F is a WForward iterator. If so, it simply
// copies the values of the mehtods from F so any calls to e.g.
// ITR.Value() will result in a direct call of the Value() method
// the F iterator wrapped.
func WrapForward(itr i.Forward) *WForward {
	wf := WForward{}

	wf.value = itr.Value
	wf.err = itr.Error
	wf.seterr = itr.SetError
	wf.next = itr.Next
	wf.atEnd = itr.AtEnd

	return &wf
}

// Forwards the call to the Value() method of the wrapped iterator.
func (wf *WForward) Value() interface{} {
	return wf.value()
}

// Forwards the call to the Error() method of the wrapped iterator.
func (wf *WForward) Error() error {
	return wf.err()
}

// Forwards the call to the SetError() method of the wrapped iterator.
func (wf *WForward) SetError(err error) {
	wf.seterr(err)
}

// Forwards the call to the Next() method of the wrapped iterator.
func (wf *WForward) Next() error {
	return wf.next()
}

// Forwards the call to the AtEnd() method of the wrapped iterator.
func (wf *WForward) AtEnd() bool {
	return wf.atEnd()
}

// Compose this structure in your iterator to wrap a BiDirectional iterator
// within your own iterator. It contains function pointers that point
// to all the methods of the wrapped BiDirectional iterator.
type WBiDirectional struct {
	WForward
	prev    func() error
	atStart func() bool
}

// Wrap a BiDirectional iterator. If ITR is your iterator type and B is
// the iterator your wish to wrap, the usage is
//
// ITR.WBiDirectional= *(WrapBiDirectional(B))
//
// Now if you don't e.g. provide an implementation for the Value()
// method, a call to ITR.Value() will be forwarded to B.Value() by
// WBiDirectional. To guard against a chain of useless forwards, the
// WrapBiDirectional checks if B is a WBiDirectional iterator. If so,
// it simply copies the values of the mehtods from B so any calls to e.g.
// ITR.Value() will result in a direct call of the Value() method
// the B iterator wrapped.
func WrapBiDirectional(itr i.BiDirectional) *WBiDirectional {
	wbd := WBiDirectional{}
	wbd.WForward = *(WrapForward(itr))

	wbd.prev = itr.Prev
	wbd.atStart = itr.AtStart

	return &wbd
}

// Forwards the call to the Prev() method of the wrapped iterator.
func (wbd *WBiDirectional) Prev() error {
	return wbd.prev()
}

// Forwards the call to the AtStart() method of the wrapped iterator.
func (wbd *WBiDirectional) AtStart() bool {
	return wbd.atStart()
}

// Compose this structure in your iterator to wrap a BoundedAtStart iterator
// within your own iterator. It contains function pointers that point
// to all the methods of the wrapped BoundedAtStart iterator.
type WBoundedAtStart struct {
	WForward
	first func() error
}

// Wrap a BoundedAtStart iterator. If ITR is your iterator type and B is
// the iterator your wish to wrap, the usage is
//
// ITR.WBoundedAtStart= *(WrapBoundedAtStart(B))
//
// Now if you don't e.g. provide an implementation for the Value()
// method, a call to ITR.Value() will be forwarded to B.Value() by
// WFBoundedAtStart. To guard against a chain of useless forwards, the
// WrapBoundedAtStart checks if B is a WBoundedAtStart iterator. If so,
// it simply copies the values of the mehtods from B so any calls to e.g.
// ITR.Value() will result in a direct call of the Value() method
// the B iterator wrapped.
func WrapBoundedAtStart(itr i.BoundedAtStart) *WBoundedAtStart {
	wbas := WBoundedAtStart{}
	wbas.WForward = *(WrapForward(itr))

	wbas.first = itr.First

	return &wbas
}

// Forwards the call to the First() method of the wrapped iterator.
func (wbas *WBoundedAtStart) First() error {
	return wbas.first()
}

// Compose this structure in your iterator to wrap a Bounded iterator
// within your own iterator. It contains function pointers that point
// to all the methods of the wrapped Bounded iterator.
type WBounded struct {
	WBiDirectional
	first func() error
	last  func() error
}

// Wrap a Bounded iterator. If ITR is your iterator type and B is
// the iterator your wish to wrap, the usage is
//
// ITR.WBounded= *(WrapBounded(B))
//
// Now if you don't e.g. provide an implementation for the Value()
// method, a call to ITR.Value() will be forwarded to B.Value() by
// WBounded. To guard against a chain of useless forwards, the
// WrapBounded checks if B is a WBounded iterator. If so, it simply
// copies the values of the mehtods from B so any calls to e.g.
// ITR.Value() will result in a direct call of the Value() method
// the B iterator wrapped.
func WrapBounded(itr i.Bounded) *WBounded {
	wb := WBounded{}
	wb.WBiDirectional = *(WrapBiDirectional(itr))

	wb.first = itr.First
	wb.last = itr.Last

	return &wb
}

// Forwards the call to the First() method of the wrapped iterator.
func (wb *WBounded) First() error {
	return wb.first()
}

// Forwards the call to the Last() method of the wrapped iterator.
func (wb *WBounded) Last() error {
	return wb.last()
}

// Compose this structure in your iterator to wrap a RandomAccess iterator
// within your own iterator. It contains function pointers that point
// to all the methods of the wrapped RandomAccess iterator.
type WRandomAccess struct {
	WBounded
	goTo   func(int) error
	length func() int
}

// Wrap a RandomAccess iterator. If ITR is your iterator type and R is
// the iterator your wish to wrap, the usage is
//
// ITR.WRandomAccess= *(WrapRandomAccess(R))
//
// Now if you don't e.g. provide an implementation for the Value()
// method, a call to ITR.Value() will be forwarded to R.Value() by
// WRandomAccess. To guard against a chain of useless forwards, the
// WrapRandomAccess checks if R is a WRandomAccess iterator. If so, it simply
// copies the values of the mehtods from R so any calls to e.g.
// ITR.Value() will result in a direct call of the Value() method
// the R iterator wrapped.
func WrapRandomAccess(itr i.RandomAccess) *WRandomAccess {
	wra := WRandomAccess{}
	wra.WBounded = *(WrapBounded(itr))

	wra.goTo = itr.Goto
	wra.length = itr.Len

	return &wra
}

// Forwards the call to the Goto() method of the wrapped iterator.
func (wra *WRandomAccess) Goto(idx int) error {
	return wra.goTo(idx)
}

// Forwards the call to the Len() method of the wrapped iterator.
func (wra *WRandomAccess) Len() int {
	return wra.length()
}
