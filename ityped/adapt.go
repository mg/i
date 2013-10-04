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

package ityped

import (
	"github.com/mg/i"
	"github.com/mg/i/itk"
)

// Typed adapter for a Forward iterator. Contains methods to acccess the
// value typed as any of the basic Go types.
type ForwardItr struct {
	itk.WForward
}

// Wrap a Forward iterator in a structure that has typed methods to access
// the value of the iterator.
func Forward(itr i.Forward) *ForwardItr {
	f := ForwardItr{}
	f.WForward = *(itk.WrapForward(itr))
	return &f
}

// Typed adapter for a Bounded At start iterator. Contains methods to acccess the
// value typed as any of the basic Go types.
type BoundedAtStartItr struct {
	itk.WBoundedAtStart
}

// Wrap a Bounded At Start iterator in a structure that has typed methods to access
// the value of the iterator.
func BoundedAtStart(itr i.BoundedAtStart) *BoundedAtStartItr {
	b := BoundedAtStartItr{}
	b.WBoundedAtStart = *(itk.WrapBoundedAtStart(itr))
	return &b
}

// Typed adapter for a BiDirecitonal iterator. Contains methods to acccess the
// value typed as any of the basic Go types.
type BiDirectionalItr struct {
	itk.WBiDirectional
}

// Wrap a BiDirectional iterator in a structure that has typed methods to access
// the value of the iterator.
func BiDirectional(itr i.BiDirectional) *BiDirectionalItr {
	b := BiDirectionalItr{}
	b.WBiDirectional = *(itk.WrapBiDirectional(itr))
	return &b
}

// Typed adapter for a Bounded iterator. Contains methods to acccess the
// value typed as any of the basic Go types.
type BoundedItr struct {
	itk.WBounded
}

// Wrap a Bounded iterator in a structure that has typed methods to access
// the value of the iterator.
func Bounded(itr i.Bounded) *BoundedItr {
	b := BoundedItr{}
	b.WBounded = *(itk.WrapBounded(itr))
	return &b
}

// Typed adapter for a Random Access iterator. Contains methods to acccess the
// value typed as any of the basic Go types.
type RandomAccessItr struct {
	itk.WRandomAccess
}

// Wrap a Random Access iterator in a structure that has typed methods to access
// the value of the iterator.
func RandomAccess(itr i.RandomAccess) *RandomAccessItr {
	r := RandomAccessItr{}
	r.WRandomAccess = *(itk.WrapRandomAccess(itr))
	return &r
}
