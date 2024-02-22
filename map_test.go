//go:build goexperiment.rangefunc

package goiter

import (
	"fmt"
	"maps"
	"testing"
)

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
