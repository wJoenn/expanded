package expsl

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

type MapTestCase[T any, R any] struct {
	slice    []T
	callback func(T) R
	result   []R
}

func TestMapIntToString(t *testing.T) {
	test := MapTestCase[int, string]{
		slice: []int{1, 2, 3},
		callback: func(e int) string {
			return strconv.Itoa(e)
		},
		result: []string{"1", "2", "3"},
	}

	got := Map(test.slice, test.callback)
	if !reflect.DeepEqual(got, test.result) {
		t.Errorf("expected: %v\n", test.result)
		t.Errorf("got: %v\n", got)
	}
}

func TestMapFloatDouble(t *testing.T) {
	test := MapTestCase[float32, float32]{
		slice: []float32{1.5, 2.5, 3.5},
		callback: func(e float32) float32 {
			return e * 2
		},
		result: []float32{3, 5, 7},
	}

	got := Map(test.slice, test.callback)
	if !reflect.DeepEqual(got, test.result) {
		t.Errorf("expected: %v\n", test.result)
		t.Errorf("got: %v\n", got)
	}
}

func TestMapStringUpper(t *testing.T) {
	test := MapTestCase[string, string]{
		slice: []string{"a", "b", "c"},
		callback: func(e string) string {
			return strings.ToUpper(e)
		},
		result: []string{"A", "B", "C"},
	}

	got := Map(test.slice, test.callback)
	if !reflect.DeepEqual(got, test.result) {
		t.Errorf("expected: %v\n", test.result)
		t.Errorf("got: %v\n", got)
	}
}
