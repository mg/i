package i

// Map iterator
type MapFunc func(Iterator) interface{}

type imap struct {
	WForward
	fmap MapFunc
	val  interface{}
}

func Map(fmap MapFunc, itr Forward) Forward {
	m := imap{fmap: fmap}
	m.WForward = *(WrapForward(itr))
	return &m
}

func (m *imap) Next() error {
	m.val = nil
	return m.WForward.Next()
}

func (m *imap) Value() interface{} {
	if m.val == nil {
		m.val = m.fmap(&m.WForward)
	}
	return m.val
}
