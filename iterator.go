package i

type (
	Iterator interface {
		Value() interface{}
		Error() error
	}

	Forward interface {
		Iterator
		Next() error
		AtEnd() bool
	}

	BiDirectional interface {
		Forward
		Prev() error
		AtStart() bool
	}

	BoundedAtStart interface {
		Forward
		First() error
	}

	Bounded interface {
		BiDirectional
		First() error
		Last() error
	}

	RandomAccess interface {
		Bounded
		Goto(int) error
		Len() int
	}
)
