package i

// Filter iterator
type FilterFunc func(Iterator) bool

type filter struct {
	f   FilterFunc
	itr Forward
}

func Filter(f FilterFunc, itr Forward) Forward {
	return &filter{f, itr}
}

func (i *filter) AtEnd() bool {
	for !i.itr.AtEnd() {
		if i.f(i.itr) {
			return false
		}
		i.itr.Next()
	}
	return true
}

func (i *filter) Next() error {
	return i.itr.Next()
}

func (i *filter) Value() interface{} {
	return i.itr.Value()
}

func (i *filter) Error() error {
	return i.itr.Error()
}
