// []interface{} iterator
package i

import (
	"fmt"
)

type iinterfaces struct {
	slice []interface{}
	pos   int
	err   error
}

func Interfaces(slice []interface{}) RandomAccess {
	return &iinterfaces{slice: slice, err: nil}
}

func List(list ...interface{}) RandomAccess {
	return &iinterfaces{slice: list, err: nil}
}

func (i *iinterfaces) AtStart() bool {
	return i.pos == 0
}

func (i *iinterfaces) AtEnd() bool {
	return i.pos >= len(i.slice)
}

func (i *iinterfaces) Next() error {
	if i.pos >= len(i.slice) {
		i.err = fmt.Errorf("Next: Index out of bounds: %d.", i.pos)
	} else {
		i.pos++
	}
	return i.err
}

func (i *iinterfaces) Prev() error {
	if i.pos < 0 {
		i.err = fmt.Errorf("Prev: Index out of bounds: %d.", i.pos)
	} else {
		i.pos--
	}
	return i.err
}

func (i *iinterfaces) First() error {
	i.pos = 0
	return nil
}

func (i *iinterfaces) Last() error {
	i.pos = len(i.slice) - 1
	return nil
}

func (i *iinterfaces) Goto(pos int) error {
	i.pos = pos
	if i.pos < 0 || i.pos >= len(i.slice) {
		i.err = fmt.Errorf("Goto: Index out of bounds: %d.", i.pos)
	}
	return i.err
}

func (i *iinterfaces) Len() int {
	return len(i.slice)
}

func (i *iinterfaces) Value() interface{} {
	if i.pos < 0 || i.pos >= len(i.slice) {
		i.err = fmt.Errorf("Value: Index out of bounds: %d.", i.pos)
		return nil
	}
	return i.slice[i.pos]
}

func (i *iinterfaces) Error() error {
	return i.err
}

func (i *iinterfaces) SetError(err error) {
	i.err = err
}
