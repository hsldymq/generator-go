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
func TestIterator_Concat(t *testing.T) {
	c1 := []int{1, 2, 3}
	c2 := []int{4, 5, 6}
	actual := make([]int, 0, 6)
	for v := range SliceElem(c1).Concat(SliceElem(c2)) {
		actual = append(actual, v)
	}
	expect := []int{1, 2, 3, 4, 5, 6}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	c1 = []int{4, 5, 6}
	c2 = []int{7, 8, 9}
	actual = make([]int, 0, 6)
	for v := range SliceElem(c1).Concat(SliceElem(c2)) {
		if v == 8 {
			break
		}
		actual = append(actual, v)
	}
	expect = []int{4, 5, 6, 7}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	actual = make([]int, 0, 6)
	for v := range SliceElem([]int{10, 11, 12}).Concat() {
		actual = append(actual, v)
		if v == 11 {
			break
		}
	}
	expect = []int{10, 11}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
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

func TestIterator_Count(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	it := Filter(SliceElem(input), func(v int) bool {
		return v%2 == 0
	})

	expect := 2
	actual := it.Count()
	if actual != expect {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}
