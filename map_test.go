package goiter

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMapIter(t *testing.T) {
	expect := map[string]int{
		"foo": 1,
		"bar": 2,
	}

	actual := map[string]int{}
	for k, v := range MapIter(expect) {
		actual[k] = v
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}
