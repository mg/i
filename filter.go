package i

// Filter iterator
type FilterFunc func(Iterator) bool

type filter struct {
	ff  FilterFunc
	itr Forward
}

func Filter(ff FilterFunc, itr Forward) Forward {
	return &filter{ff, itr}
}

func (f *filter) AtEnd() bool {
	for !f.itr.AtEnd() {
		if f.ff(f.itr) {
			return false
		}
		f.itr.Next()
	}
	return true
}

func (f *filter) Next() error {
	return f.itr.Next()
}

func (f *filter) Value() interface{} {
	return f.itr.Value()
}

func (f *filter) Error() error {
	return f.itr.Error()
}
