package i

type EachFunc func(i Iterator) bool

func Each(i Forward, e EachFunc) {
	for ; !i.AtEnd(); i.Next() {
		if !e(i) {
			break
		}
	}
}
