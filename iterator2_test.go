package goiter

import (
	"cmp"
	"fmt"
	"maps"
	"slices"
	"testing"
)

func TestIterator2_OrderBy(t *testing.T) {
	type person struct {
		name string
		age  int
	}
	input := map[string]int{
		"bob":   20,
		"eve":   30,
		"alice": 25,
	}
	actual := []person{}
	for v1, v2 := range Map(input).OrderBy(func(a, b *Combined[string, int]) int { return cmp.Compare(a.V2, b.V2) }) {
		actual = append(actual, person{name: v1, age: v2})
	}
	expect := []person{
		{"bob", 20},
		{"alice", 25},
		{"eve", 30},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal("expect:", expect, "actual:", actual)
	}
}

func TestIterator2_StableOrderBy(t *testing.T) {
	type person struct {
		name string
		age  int
	}
	input := []person{
		{"bob", 25},
		{"eve", 30},
		{"alice", 25},
	}
	actual := []person{}
	for _, v2 := range Slice(input).StableOrderBy(func(a, b *Combined[int, person]) int { return cmp.Compare(a.V2.age, b.V2.age) }) {
		actual = append(actual, v2)
	}
	expect := []person{
		{"bob", 25},
		{"alice", 25},
		{"eve", 30},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal("expect:", expect, "actual:", actual)
	}
}

func TestIterator2_Filter(t *testing.T) {
	predicate := func(name string, age int) bool {
		return name == "john"
	}
	input := map[string]int{"john": 20, "jane": 18}
	actual := map[string]int{}
	for k, v := range Map(input).Filter(predicate) {
		actual[k] = v
	}
	expect := map[string]int{"john": 20}
	if !maps.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	for _, _ = range Map(input).Filter(predicate) {
		break
	}
}

func TestIterator2_ToSlice(t *testing.T) {
	iterator := Iterator2[string, int](func(yield func(string, int) bool) {
		yield("alice", 20)
		yield("bob", 21)
		yield("eve", 22)
	})

	actual := make([]Combined[string, int], 0, 3)
	for _, each := range iterator.ToSlice() {
		actual = append(actual, *each)
	}
	expect := []Combined[string, int]{
		{"alice", 20},
		{"bob", 21},
		{"eve", 22},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}
