package goiter

import (
	"cmp"
	"fmt"
	"slices"
	"testing"
)

func TestIterator_OrderBy(t *testing.T) {
	type person struct {
		name string
		age  int
	}
	input := []person{
		{"alice", 22},
		{"bob", 20},
		{"eve", 21},
	}
	actual := []person{}
	for each := range SliceElem(input).OrderBy(func(a, b person) int { return cmp.Compare(a.age, b.age) }) {
		actual = append(actual, each)
	}
	expect := []person{
		{"bob", 20},
		{"eve", 21},
		{"alice", 22},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal("expect:", expect, "actual:", actual)
	}
}

func TestIterator_StableOrderBy(t *testing.T) {
	type person struct {
		name string
		age  int
	}
	input := []person{
		{"alice", 22},
		{"bob", 20},
		{"eve", 20},
	}
	actual := []person{}
	for each := range SliceElem(input).StableOrderBy(func(a, b person) int { return cmp.Compare(a.age, b.age) }) {
		actual = append(actual, each)
	}
	expect := []person{
		{"bob", 20},
		{"eve", 20},
		{"alice", 22},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal("expect:", expect, "actual:", actual)
	}
}

func TestIterator_Filter(t *testing.T) {
	predicate := func(v int) bool {
		return v%2 == 0
	}
	actual := []int{}
	for v := range Range(0, 10).Filter(predicate) {
		actual = append(actual, v)
	}
	expect := []int{0, 2, 4, 6, 8}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	for _ = range Range(0, 10).Filter(predicate) {
		break
	}
}

func TestIterator_ToSlice(t *testing.T) {
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
