package i

// Map iterator
type MapFunc func(Iterator) interface{}

type imap struct {
	wforward
	fmap MapFunc
	val  interface{}
}

func Map(fmap MapFunc, itr Forward) Forward {
	m := imap{fmap: fmap}
	m.wforward = *(WrapForward(itr))
	return &m
}

func (m *imap) Next() error {
	m.val = nil
	return m.wforward.Next()
}

func (m *imap) Value() interface{} {
	if m.val == nil {
		m.val = m.fmap(&m.wforward)
	}
	return m.val
}
