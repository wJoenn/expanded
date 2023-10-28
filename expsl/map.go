package expsl

func Map[T any, R any](s []T, f func(T) R) []R {
	r := make([]R, len(s))
	for i, e := range s {
		r[i] = f(e)
	}

	return r
}
