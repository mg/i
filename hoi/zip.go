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
	"fmt"
	"github.com/mg/i"
)

type zip struct {
	itrs  []i.Forward
	err   error
	atEnd bool
}

// The Zip iterator will zip together a collection of data streams, stopping
// after the shortest data stream is finished. Given e.g. the data streams
// [1,2,3], [5,6,7] and [10,11,12,13] Zip will provide access to the data
// sream [[1,5,10], [2,6,11], [3,7,12]].
func Zip(itrs ...i.Forward) i.Forward {
	return &zip{itrs: itrs}
}

func (z *zip) Error() error {
	return z.err
}

func (z *zip) SetError(err error) {
	z.err = err
}

func (z *zip) Value() interface{} {
	if z.atEnd {
		z.err = fmt.Errorf("Calling Value() at end")
		return nil
	}
	ret := make([]interface{}, len(z.itrs))
	for idx, itr := range z.itrs {
		ret[idx] = itr.Value()
	}
	return ret
}

func (z *zip) Next() error {
	if z.atEnd {
		z.err = fmt.Errorf("Calling Next() at end")
		return z.err
	}
	for _, itr := range z.itrs {
		if !itr.AtEnd() {
			err := itr.Next()
			if err != nil {
				z.err = err
				break
			}
		}
	}
	return z.err
}

func (z *zip) AtEnd() bool {
	for _, itr := range z.itrs {
		if itr.AtEnd() {
			z.atEnd = true
			return true
		}
	}
	return false
}

type ziplongest struct {
	zip
}

// The ZipLongest iterator will zip together a collection of data streams, not
// stopping until the longest data stream is finished. Given e.g. the data
// streams [1,2,3], [5,6,7] and [10,11,12,13] Zip will provide access to the
// data sream [[1,5,10], [2,6,11], [3,7,12], [nil, nil, 13]].
func ZipLongest(itrs ...i.Forward) i.Forward {
	return &ziplongest{zip{itrs: itrs}}
}

func (z *ziplongest) Value() interface{} {
	if z.atEnd {
		z.err = fmt.Errorf("Calling Value() at end")
		return nil
	}
	ret := make([]interface{}, len(z.itrs))
	for idx, itr := range z.itrs {
		if !itr.AtEnd() {
			ret[idx] = itr.Value()
		} else {
			ret[idx] = nil
		}
	}
	return ret
}

func (z *ziplongest) AtEnd() bool {
	atEndCount := 0
	for _, itr := range z.itrs {
		if itr.AtEnd() {
			atEndCount++
		}
	}
	z.atEnd = atEndCount == len(z.itrs)
	return z.atEnd
}
