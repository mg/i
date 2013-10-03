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
