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

// The Iterator ToolKit package (itk) provides reusable components
// to help with the construction of iterators.
//
// Often when we write iterators we are simply writing adapters that wrap
// around existing iterators and change some but not all of the methods.
// The wrappers in itk exists to reduce the amount of boilerplate code we
// need to write in those cases. Simply use composition to wrap the iterator,
// provide the logic for the methods you whish to change, and let the wrapper
// handle the rest.
//
// E.g. LyingIterator is a RandomAccess iterator that lies about its length.
// It supports all other methods defined in i.RandomAccess through the wrapper.
//
//   type lyingiterator struct {
//     itk.WRandomAccess
//   }
//
//   func LyingIterator( itr i.RandomAccess) i.RandomAccess {
//     var li lyingiterator
//     li.WRandomAccess= *(itk.WrapRandomAccess(itr))
//     return &li
//   }
//
//   func (li *lyingiterator) {
//     return li.WRandomAccess.Len()/2
//   }
//
package itk
