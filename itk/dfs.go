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
	"fmt"
	"github.com/mg/i"
)

type ChildrenFunc func(interface{}) []interface{}

type depthfirst struct {
	children ChildrenFunc
	agenda   []interface{}
	node     interface{}
	err      error
}

func Dfs(root interface{}, children ChildrenFunc) i.Forward {
	dfs := depthfirst{children: children}
	dfs.agenda = append(dfs.agenda, root)
	dfs.Next()
	return &dfs
}

func (dfs *depthfirst) AtEnd() bool {
	return len(dfs.agenda) == 0
}

func (dfs *depthfirst) Value() interface{} {
	if dfs.AtEnd() {
		dfs.err = fmt.Errorf("Next: Beyond end")
		return nil
	}
	return dfs.node
}

func (dfs *depthfirst) Error() error {
	return dfs.err
}

func (dfs *depthfirst) SetError(err error) {
	dfs.err = err
}

func (dfs *depthfirst) Next() error {
	if dfs.AtEnd() {
		dfs.err = fmt.Errorf("Next: Beyond end")
		return dfs.err
	}
	dfs.node, dfs.agenda = dfs.agenda[0], dfs.agenda[1:len(dfs.agenda)]
	dfs.agenda = append(dfs.agenda, dfs.children(dfs.node)...)
	return nil
}
