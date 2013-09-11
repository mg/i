package i

import ()

type WForward struct {
	value  func() interface{}
	err    func() error
	seterr func(error)
	next   func() error
	atEnd  func() bool
}

func WrapForward(itr Forward) *WForward {
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

func WrapBiDirectional(itr BiDirectional) *WBiDirectional {
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

func WrapBoundedAtStart(itr BoundedAtStart) *WBoundedAtStart {
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

func WrapBounded(itr Bounded) *WBounded {
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

func WrapRandomAccess(itr RandomAccess) *WRandomAccess {
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
