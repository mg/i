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

package iadapt

import (
	"github.com/mg/i"
	"github.com/mg/i/itk"
)

type ForwardItr struct {
	itk.WForward
}

func Forward(itr i.Forward) *ForwardItr {
	f := ForwardItr{}
	f.WForward = *(itk.WrapForward(itr))
	return &f
}

type BoundedAtStartItr struct {
	itk.WBoundedAtStart
}

func BoundedAtStart(itr i.BoundedAtStart) *BoundedAtStartItr {
	b := BoundedAtStartItr{}
	b.WBoundedAtStart = *(itk.WrapBoundedAtStart(itr))
	return &b
}

type BiDirectionalItr struct {
	itk.WBiDirectional
}

func BiDirectional(itr i.BiDirectional) *BiDirectionalItr {
	b := BiDirectionalItr{}
	b.WBiDirectional = *(itk.WrapBiDirectional(itr))
	return &b
}

type BoundedItr struct {
	itk.WBounded
}

func Bounded(itr i.Bounded) *BoundedItr {
	b := BoundedItr{}
	b.WBounded = *(itk.WrapBounded(itr))
	return &b
}

type RandomAccessItr struct {
	itk.WRandomAccess
}

func RandomAccess(itr i.RandomAccess) *RandomAccessItr {
	r := RandomAccessItr{}
	r.WRandomAccess = *(itk.WrapRandomAccess(itr))
	return &r
}
