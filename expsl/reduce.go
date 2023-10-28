package expsl

type Number = interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Reduce[T Number, R Number](s []T, b R, f func(T, R) R) (r R) {
	r = b
	for _, e := range s {
		r = f(e, r)
	}

	return
}
