package iadapt

import "github.com/mg/i"

type ForwardItr struct {
	i.WForward
}

func Forward(itr i.Forward) *ForwardItr {
	f := ForwardItr{}
	f.WForward = *(i.WrapForward(itr))
	return &f
}

type BoundedAtStartItr struct {
	i.WBoundedAtStart
}

func BoundedAtStart(itr i.BoundedAtStart) *BoundedAtStartItr {
	b := BoundedAtStartItr{}
	b.WBoundedAtStart = *(i.WrapBoundedAtStart(itr))
	return &b
}

type BiDirectionalItr struct {
	i.WBiDirectional
}

func BiDirectional(itr i.BiDirectional) *BiDirectionalItr {
	b := BiDirectionalItr{}
	b.WBiDirectional = *(i.WrapBiDirectional(itr))
	return &b
}

type BoundedItr struct {
	i.WBounded
}

func Bounded(itr i.Bounded) *BoundedItr {
	b := BoundedItr{}
	b.WBounded = *(i.WrapBounded(itr))
	return &b
}

type RandomAccessItr struct {
	i.WRandomAccess
}

func RandomAccess(itr i.RandomAccess) *RandomAccessItr {
	r := RandomAccessItr{}
	r.WRandomAccess = *(i.WrapRandomAccess(itr))
	return &r
}
