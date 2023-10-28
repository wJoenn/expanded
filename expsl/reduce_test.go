package expsl

import "testing"

type ReduceTestCase[T Number, R Number] struct {
	slice    []T
	base     R
	callback func(T, R) R
	result   R
}

func TestReduceSum(t *testing.T) {
	test := ReduceTestCase[int, int]{
		slice: []int{1, 2, 3},
		base:  0,
		callback: func(e int, sum int) int {
			return sum + e
		},
		result: 6,
	}

	got := Reduce(test.slice, test.base, test.callback)
	if got != test.result {
		t.Errorf("expected: %v\n", test.result)
		t.Errorf("got: %v\n", got)
	}
}

func TestReduceSub(t *testing.T) {
	test := ReduceTestCase[int, int]{
		slice: []int{1, 2, 3},
		base:  0,
		callback: func(e int, sum int) int {
			return sum - e
		},
		result: -6,
	}

	got := Reduce(test.slice, test.base, test.callback)
	if got != test.result {
		t.Errorf("expected: %v\n", test.result)
		t.Errorf("got: %v\n", got)
	}
}

func TestReduceMult(t *testing.T) {
	test := ReduceTestCase[int, int]{
		slice: []int{1, 2, 3},
		base:  1,
		callback: func(e int, sum int) int {
			return sum * e
		},
		result: 6,
	}

	got := Reduce(test.slice, test.base, test.callback)
	if got != test.result {
		t.Errorf("expected: %v\n", test.result)
		t.Errorf("got: %v\n", got)
	}
}

func TestReduceDiv(t *testing.T) {
	test := ReduceTestCase[int, float64]{
		slice: []int{2, 2, 2},
		base:  4,
		callback: func(e int, sum float64) float64 {
			return sum / float64(e)
		},
		result: 0.5,
	}

	got := Reduce(test.slice, test.base, test.callback)
	if got != test.result {
		t.Errorf("expected: %v\n", test.result)
		t.Errorf("got: %v\n", got)
	}
}
