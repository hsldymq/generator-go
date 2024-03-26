//go:build goexperiment.rangefunc

package goiter

import (
	"fmt"
	"maps"
	"slices"
	"testing"
)

func TestToSlice(t *testing.T) {
	iterator := Iterator[int](func(yield func(int) bool) {
		yield(1)
		yield(2)
		yield(3)
	})

	actual := iterator.ToSlice()
	expect := []int{1, 2, 3}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestToMap(t *testing.T) {
	seq := func(yield func(string, int) bool) {
		yield("alice", 20)
		yield("bob", 21)
		yield("eve", 22)
	}

	actual := ToMap(seq)
	expect := map[string]int{
		"alice": 20,
		"bob":   21,
		"eve":   22,
	}
	if !maps.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestToMapBy(t *testing.T) {
	seq := func(yield func(string) bool) {
		yield("alice")
		yield("bob")
		yield("eve")
	}

	actual := ToMapAs(seq, func(name string) (string, int) {
		return name, len(name)
	})
	expect := map[string]int{
		"alice": 5,
		"bob":   3,
		"eve":   3,
	}
	if !maps.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestToMapBy2(t *testing.T) {
	seq := func(yield func(string, string) bool) {
		yield("Alice", "Paris")
		yield("Bob", "Shanghai")
		yield("Eve", "Bangkok")
	}

	actual := ToMapAs2(seq, func(name string, city string) (string, string) {
		return name + "_" + city, string(name[0]) + "_" + string(city[0])
	})
	expect := map[string]string{
		"Alice_Paris":  "A_P",
		"Bob_Shanghai": "B_S",
		"Eve_Bangkok":  "E_B",
	}
	if !maps.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}
