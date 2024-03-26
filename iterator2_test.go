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

func TestIterator2_Concat(t *testing.T) {
	type person struct {
		name string
		age  int
	}
	p1 := []person{{"john", 25}, {"jane", 20}}
	i1 := Transform12(SliceElem(p1), func(p person) (string, int) {
		return p.name, p.age
	})
	p2 := []person{{"joe", 35}, {"ann", 30}, {"josh", 15}}
	i2 := Transform12(SliceElem(p2), func(p person) (string, int) {
		return p.name, p.age
	})

	actual := make(map[string]int)
	for name, age := range i1.Concat(i2) {
		actual[name] = age
	}
	expect := map[string]int{"john": 25, "jane": 20, "joe": 35, "ann": 30, "josh": 15}
	if !maps.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	actual = make(map[string]int)
	for name, age := range i1.Concat(i2) {
		if name == "ann" {
			break
		}
		actual[name] = age
	}
	expect = map[string]int{"john": 25, "jane": 20, "joe": 35}
	if !maps.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	actual = make(map[string]int)
	for name, age := range i1.Concat() {
		actual[name] = age
		if name == "john" {
			break
		}
	}
	expect = map[string]int{"john": 25}
	if !maps.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
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

func TestIterator2_Count(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	it := Filter2(Slice(input), func(idx int, v int) bool {
		return idx != 0
	})

	expect := 4
	actual := it.Count()
	if actual != expect {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}
