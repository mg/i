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
	"reflect"
)

// Gives access to wrapped methods
type wrappedForward interface {
	getValueMethod() func() interface{}
	getErrorMethod() func() error
	getSetErrorMethod() func(error)
	getNextMethod() func() error
	getAtEndMethod() func() bool
}

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

	t := reflect.TypeOf(itr)
	isWrapper := false
	if t.Kind() == reflect.Ptr {
		_, isWrapper = t.Elem().FieldByName("WForward")
	} else {
		_, isWrapper = t.FieldByName("WForward")
	}
	if isWrapper {
		wrapped, _ := itr.(wrappedForward)
		if _, ok := t.MethodByName("Value"); !ok {
			wf.value = wrapped.getValueMethod()
		} else {
			wf.value = itr.Value
		}
		if _, ok := t.MethodByName("Error"); !ok {
			wf.err = wrapped.getErrorMethod()
		} else {
			wf.err = itr.Error
		}
		if _, ok := t.MethodByName("SetError"); !ok {
			wf.seterr = wrapped.getSetErrorMethod()
		} else {
			wf.seterr = itr.SetError
		}
		if _, ok := t.MethodByName("Next"); !ok {
			wf.next = wrapped.getNextMethod()
		} else {
			wf.next = itr.Next
		}
		if _, ok := t.MethodByName("AtEnd"); !ok {
			wf.atEnd = wrapped.getAtEndMethod()
		} else {
			wf.atEnd = itr.AtEnd
		}
	} else {
		wf.value = itr.Value
		wf.err = itr.Error
		wf.seterr = itr.SetError
		wf.next = itr.Next
		wf.atEnd = itr.AtEnd
	}
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

// Methods to access wrapped functions
func (wf *WForward) getValueMethod() func() interface{} {
	return wf.value
}

func (wf *WForward) getErrorMethod() func() error {
	return wf.err
}

func (wf *WForward) getSetErrorMethod() func(error) {
	return wf.seterr
}

func (wf *WForward) getNextMethod() func() error {
	return wf.next
}

func (wf *WForward) getAtEndMethod() func() bool {
	return wf.atEnd
}

// Gives access to wrapped methods
type wrappedBiDirectional interface {
	getPrevMethod() func() error
	getAtStartMethod() func() bool
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

	t := reflect.TypeOf(itr)
	isWrapper := false
	if t.Kind() == reflect.Ptr {
		_, isWrapper = t.Elem().FieldByName("WBiDirectional")
	} else {
		_, isWrapper = t.FieldByName("WBiDirectional")
	}
	if isWrapper {
		wrapped, _ := itr.(wrappedBiDirectional)
		if _, ok := t.MethodByName("Prev"); !ok {
			wbd.prev = wrapped.getPrevMethod()
		} else {
			wbd.prev = itr.Prev
		}
		if _, ok := t.MethodByName("AtStart"); !ok {
			wbd.atStart = wrapped.getAtStartMethod()
		} else {
			wbd.atStart = itr.AtStart
		}
	} else {
		wbd.prev = itr.Prev
		wbd.atStart = itr.AtStart
	}

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

// Methods to access wrapped functions
func (wbd *WBiDirectional) getPrevMethod() func() error {
	return wbd.prev
}

func (wbd *WBiDirectional) getAtStartMethod() func() bool {
	return wbd.atStart
}

// Gives access to wrapped methods
type wrappedBondedAtStart interface {
	getFirstMethod() func() error
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

	t := reflect.TypeOf(itr)
	isWrapper := false
	if t.Kind() == reflect.Ptr {
		_, isWrapper = t.Elem().FieldByName("WBoundedAtStart")
	} else {
		_, isWrapper = t.FieldByName("WBoundedAtStart")
	}
	if isWrapper {
		wrapped, _ := itr.(wrappedBondedAtStart)
		if _, ok := t.MethodByName("First"); !ok {
			wbas.first = wrapped.getFirstMethod()
		} else {
			wbas.first = itr.First
		}
	} else {
		wbas.first = itr.First
	}

	return &wbas
}

// Forwards the call to the First() method of the wrapped iterator.
func (wbas *WBoundedAtStart) First() error {
	return wbas.first()
}

// Methods to access wrapped functions
func (wbas *WBoundedAtStart) getFirstMethod() func() error {
	return wbas.first
}

// Gives access to wrapped methods
type wrappedBonded interface {
	getFirstMethod() func() error
	getLastMethod() func() error
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

	t := reflect.TypeOf(itr)
	isWrapper := false
	if t.Kind() == reflect.Ptr {
		_, isWrapper = t.Elem().FieldByName("WBounded")
	} else {
		_, isWrapper = t.FieldByName("WBounded")
	}
	if isWrapper {
		wrapped, _ := itr.(wrappedBonded)
		if _, ok := t.MethodByName("First"); !ok {
			wb.first = wrapped.getFirstMethod()
		} else {
			wb.first = itr.First
		}
		if _, ok := t.MethodByName("Last"); !ok {
			wb.last = wrapped.getLastMethod()
		} else {
			wb.last = itr.Last
		}
	} else {
		wb.first = itr.First
		wb.last = itr.Last
	}

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

// Methods to access wrapped functions
func (wb *WBounded) getFirstMethod() func() error {
	return wb.first
}

func (wb *WBounded) getLastMethod() func() error {
	return wb.last
}

// Gives access to wrapped methods
type wrappedRandomAccess interface {
	getGoToMethod() func(int) error
	getLenMethod() func() int
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

	t := reflect.TypeOf(itr)
	isWrapper := false
	if t.Kind() == reflect.Ptr {
		_, isWrapper = t.Elem().FieldByName("WRandomAccess")
	} else {
		_, isWrapper = t.FieldByName("WRandomAccess")
	}
	if isWrapper {
		wrapped, _ := itr.(wrappedRandomAccess)
		if _, ok := t.MethodByName("GoTo"); !ok {
			wra.goTo = wrapped.getGoToMethod()
		} else {
			wra.goTo = itr.Goto
		}
		if _, ok := t.MethodByName("Len"); !ok {
			wra.length = wrapped.getLenMethod()
		} else {
			wra.length = itr.Len
		}
	} else {
		wra.goTo = itr.Goto
		wra.length = itr.Len
	}

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

// Methods to access wrapped functions
func (wra *WRandomAccess) getGoToMethod() func(int) error {
	return wra.goTo
}

func (wra *WRandomAccess) getLenMethod() func() int {
	return wra.length
}
