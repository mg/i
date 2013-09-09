package i

// Filter iterator
type FilterFunc func(Iterator) bool

type filter struct {
	wforward
	ff FilterFunc
}

func Filter(ff FilterFunc, itr Forward) Forward {
	f := filter{ff: ff}
	f.wforward = *(WrapForward(itr))
	return &f
}

func (f *filter) AtEnd() bool {
	for !f.wforward.AtEnd() {
		if f.ff(&f.wforward) {
			return false
		}
		f.wforward.Next()
	}
	return true
}
