package goiter

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSliceIter(t *testing.T) {
	actual := make([]int, 0, 3)
	for v := range SliceIterIdx([]int{7, 8, 9}) {
		actual = append(actual, v)
	}
	expect := []int{0, 1, 2}
	if !reflect.DeepEqual(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	actual2 := make([]int, 0, 3)
	for idx := range SliceIterElem([]int{7, 8, 9}, true) {
		actual2 = append(actual2, idx)
	}
	expect2 := []int{9, 8, 7}
	if !reflect.DeepEqual(expect2, actual2) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect2, actual2))
	}
}
