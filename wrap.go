package i

import ()

type wforward struct {
	value func() interface{}
	err   func() error
	next  func() error
	atEnd func() bool
}

func WrapForward(itr Forward) *wforward {
	wf := wforward{}
	if wrapped, ok := itr.(*wforward); ok {
		wf.value = wrapped.value
		wf.err = wrapped.err
		wf.next = wrapped.next
		wf.atEnd = wrapped.atEnd
	} else {
		wf.value = itr.Value
		wf.err = itr.Error
		wf.next = itr.Next
		wf.atEnd = itr.AtEnd
	}
	return &wf
}

func (wf *wforward) Value() interface{} {
	return wf.value()
}

func (wf *wforward) Error() error {
	return wf.err()
}

func (wf *wforward) Next() error {
	return wf.next()
}

func (wf *wforward) AtEnd() bool {
	return wf.atEnd()
}

// WrapBiDirectional
type wbidirectional struct {
	wforward
	prev    func() error
	atStart func() bool
}

func WrapBiDirectional(itr BiDirectional) *wbidirectional {
	wbd := wbidirectional{}
	wbd.wforward = *(WrapForward(itr))

	if wrapped, ok := itr.(*wbidirectional); ok {
		wbd.prev = wrapped.prev
		wbd.atStart = wrapped.atStart
	} else {
		wbd.prev = itr.Prev
		wbd.atStart = itr.AtStart
	}

	return &wbd
}

func (wbd *wbidirectional) Prev() error {
	return wbd.prev()
}

func (wbd *wbidirectional) AtStart() bool {
	return wbd.atStart()
}

// WrapBoundedAtStart
type wboundedatstart struct {
	wforward
	first func() error
}

func WrapBoundedAtStart(itr BoundedAtStart) *wboundedatstart {
	wbas := wboundedatstart{}
	wbas.wforward = *(WrapForward(itr))

	if wrapped, ok := itr.(*wboundedatstart); ok {
		wbas.first = wrapped.first
	} else {
		wbas.first = itr.First
	}

	return &wbas
}

func (wbas *wboundedatstart) First() error {
	return wbas.first()
}

// WrapBounded
type wbounded struct {
	wbidirectional
	first func() error
	last  func() error
}

func WrapBounded(itr Bounded) *wbounded {
	wb := wbounded{}
	wb.wbidirectional = *(WrapBiDirectional(itr))

	if wrapped, ok := itr.(*wbounded); ok {
		wb.first = wrapped.first
		wb.last = wrapped.last
	} else {
		wb.first = itr.First
		wb.last = itr.Last
	}

	return &wb
}

func (wb *wbounded) First() error {
	return wb.first()
}

func (wb *wbounded) Last() error {
	return wb.last()
}

// WrapRandomAccess
type wrandomaccess struct {
	wbounded
	goTo   func(int) error
	length func() int
}

func WrapRandomAccess(itr RandomAccess) *wrandomaccess {
	wra := wrandomaccess{}
	wra.wbounded = *(WrapBounded(itr))

	if wrapped, ok := itr.(*wrandomaccess); ok {
		wra.goTo = wrapped.goTo
		wra.length = wrapped.length
	} else {
		wra.goTo = itr.Goto
		wra.length = itr.Len
	}

	return &wra
}

func (wra *wrandomaccess) Goto(idx int) error {
	return wra.goTo(idx)
}

func (wra *wrandomaccess) Len() int {
	return wra.length()
}
