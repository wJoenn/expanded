package expsl

func Filter[T any](s []T, f func(T) bool) (r []T) {
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}

	return
}
