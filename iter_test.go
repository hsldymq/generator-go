//go:build goexperiment.rangefunc

package goiter

import (
	"fmt"
	"maps"
	"slices"
	"testing"
)

func TestSliceIter(t *testing.T) {
	actual := make([]int, 0, 3)
	for v := range SliceIdx([]int{7, 8, 9}) {
		actual = append(actual, v)
	}
	expect := []int{0, 1, 2}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	actual2 := make([]int, 0, 3)
	for idx := range SliceElem([]int{7, 8, 9}, true) {
		actual2 = append(actual2, idx)
	}
	expect2 := []int{9, 8, 7}
	if !slices.Equal(expect2, actual2) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect2, actual2))
	}
}

func TestMap(t *testing.T) {
	expect := map[string]int{
		"foo": 1,
		"bar": 2,
	}

	actual := map[string]int{}
	for k, v := range Map(expect) {
		actual[k] = v
	}

	if !maps.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}
