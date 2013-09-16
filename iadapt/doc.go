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

// The Iterator Adapters (iadapt) package provides adaptors that give access to
// functions that return the current value from the adapted iterator in a typed
// way.
//
// The purpose of the adapters is to be added to the end of a chain of iterators
// to typecast the final value of the chain. The only function they provide are
// various typed functions that access the value of the iterator and perform
// a type assertion and return the resulting value. Every function comes in two
// forms. E.g the boolean variant is Bool(), which returns the value as a bolean
// value or panics if it fails the type assertion, or BoolOr(bool), witch returns
// the boolean value of the iterator and true if the type assertions succedes and
// returns the supplied boolean value and false if the type assertion fails. This
// pattern is repeated for all the basic types of Go.
package iadapt
