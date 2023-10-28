package expsl

import (
	"reflect"
	"testing"
)

type FilterTestCase[T any] struct {
	slice    []T
	callback func(T) bool
	result   []T
}

func TestFilterString(t *testing.T) {
	testCases := []FilterTestCase[string]{
		{
			slice: []string{"a", "b", "c"},
			callback: func(e string) bool {
				return e == "a"
			},
			result: []string{"a"},
		},
		{
			slice: []string{"a", "b", "c"},
			callback: func(e string) bool {
				return e == "b"
			},
			result: []string{"b"},
		},
		{
			slice: []string{"a", "b", "c"},
			callback: func(e string) bool {
				return e == "c"
			},
			result: []string{"c"},
		},
	}

	for _, test := range testCases {
		t.Run("test filter result", func(t *testing.T) {
			got := Filter(test.slice, test.callback)
			if !reflect.DeepEqual(got, test.result) {
				t.Errorf("expected: %v\n", test.result)
				t.Errorf("got: %v\n", got)
			}
		})
	}
}

func TestFilterInt(t *testing.T) {
	testCases := []FilterTestCase[int]{
		{
			slice: []int{1, 2, 3},
			callback: func(e int) bool {
				return e == 1
			},
			result: []int{1},
		},
		{
			slice: []int{1, 2, 3},
			callback: func(e int) bool {
				return e == 2
			},
			result: []int{2},
		},
		{
			slice: []int{1, 2, 3},
			callback: func(e int) bool {
				return e == 3
			},
			result: []int{3},
		},
	}

	for _, test := range testCases {
		t.Run("test filter result", func(t *testing.T) {
			got := Filter(test.slice, test.callback)
			if !reflect.DeepEqual(got, test.result) {
				t.Errorf("expected: %v\n", test.result)
				t.Errorf("got: %v\n", got)
			}
		})
	}
}

func TestFilterFloat(t *testing.T) {
	testCases := []FilterTestCase[float32]{
		{
			slice: []float32{0.1, 0.2, 0.3},
			callback: func(e float32) bool {
				return e == 0.1
			},
			result: []float32{0.1},
		},
		{
			slice: []float32{0.1, 0.2, 0.3},
			callback: func(e float32) bool {
				return e == 0.2
			},
			result: []float32{0.2},
		},
		{
			slice: []float32{0.1, 0.2, 0.3},
			callback: func(e float32) bool {
				return e == 0.3
			},
			result: []float32{0.3},
		},
	}

	for _, test := range testCases {
		t.Run("test filter result", func(t *testing.T) {
			got := Filter(test.slice, test.callback)
			if !reflect.DeepEqual(got, test.result) {
				t.Errorf("expected: %v\n", test.result)
				t.Errorf("got: %v\n", got)
			}
		})
	}
}

func TestFilterBool(t *testing.T) {
	testCases := []FilterTestCase[bool]{
		{
			slice: []bool{true, false, true, false},
			callback: func(e bool) bool {
				return e
			},
			result: []bool{true, true},
		},
		{
			slice: []bool{true, false, true, false},
			callback: func(e bool) bool {
				return !e
			},
			result: []bool{false, false},
		},
	}

	for _, test := range testCases {
		t.Run("test filter result", func(t *testing.T) {
			got := Filter(test.slice, test.callback)
			if !reflect.DeepEqual(got, test.result) {
				t.Errorf("expected: %v\n", test.result)
				t.Errorf("got: %v\n", got)
			}
		})
	}
}
