package goiter

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSliceIter(t *testing.T) {
	expect := []int{7, 8, 9}

	actual := make([]int, 3)
	for idx, v := range SliceIter(expect) {
		actual[idx] = v
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}
