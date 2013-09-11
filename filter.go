package i

import "fmt"

// Filter iterator
type FilterFunc func(Iterator) bool

type filter struct {
	WForward
	ff FilterFunc
}

func Filter(ff FilterFunc, itr Forward) Forward {
	f := filter{ff: ff}
	f.WForward = *(WrapForward(itr))
	return &f
}

func (f *filter) Next() error {
	if f.WForward.AtEnd() {
		f.WForward.SetError(fmt.Errorf("Calling Next() after end"))
		return f.WForward.Error()
	}
	for !f.WForward.AtEnd() {
		f.WForward.Next()
		if !f.WForward.AtEnd() && f.ff(&f.WForward) {
			return nil
		}
	}
	return f.WForward.Error()
}
