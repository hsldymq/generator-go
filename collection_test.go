//go:build goexperiment.rangefunc

package goiter

import (
	"fmt"
	"maps"
	"slices"
	"testing"
)

func TestSlice(t *testing.T) {
	actual := make([]int, 0, 3)
	for v := range SliceIdx([]int{7, 8, 9}) {
		actual = append(actual, v)
	}
	expect := []int{0, 1, 2}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	actual2 := make([]int, 0, 3)
	for v := range SliceElem([]int{7, 8, 9}, true) {
		actual2 = append(actual2, v)
	}
	expect2 := []int{9, 8, 7}
	if !slices.Equal(expect2, actual2) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect2, actual2))
	}

	actual3 := make([]int, 0, 1)
	for _, v := range Slice([]int{7, 8, 9}) {
		if v == 8 {
			break
		}
		actual3 = append(actual3, v)
	}
	expect3 := []int{7}
	if !slices.Equal(expect3, actual3) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect3, actual3))
	}

	actual4 := make([]int, 0, 3)
	for v := range SliceElem([]int{7, 8, 9}, true) {
		if v == 7 {
			break
		}
		actual4 = append(actual4, v)
	}
	expect4 := []int{9, 8}
	if !slices.Equal(expect4, actual4) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect4, actual4))
	}
}

func TestMap(t *testing.T) {
	input := map[string]int{"foo": 1, "bar": 2}
	actual := make([]int, 0, 2)
	for v := range MapVal(input) {
		actual = append(actual, v)
	}
	slices.Sort(actual)
	expect := []int{1, 2}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	actual2 := make(map[string]bool)
	for k := range MapKey(input) {
		actual2[k] = true
	}
	expect2 := map[string]bool{"foo": true, "bar": true}
	if !maps.Equal(expect2, actual2) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect2, actual2))
	}
}

func TestChannel(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	actual := make([]int, 0, 3)
	for v := range Channel(ch) {
		if v == 3 {
			break
		}
		actual = append(actual, v)
	}
	expect := []int{1, 2}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestConcat(t *testing.T) {
	c1 := []int{1, 2, 3}
	c2 := []int{4, 5, 6}
	actual := make([]int, 0, 6)
	for v := range Concat(SliceElem(c1), SliceElem(c2)) {
		actual = append(actual, v)
	}
	expect := []int{1, 2, 3, 4, 5, 6}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	c1 = []int{4, 5, 6}
	c2 = []int{7, 8, 9}
	actual = make([]int, 0, 6)
	for v := range Concat(SliceElem(c1), SliceElem(c2)) {
		if v == 8 {
			break
		}
		actual = append(actual, v)
	}
	expect = []int{4, 5, 6, 7}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestConcat2(t *testing.T) {
	type person struct {
		name string
		age  int
	}
	p1 := []person{{"john", 25}, {"jane", 20}}
	i1 := T12(SliceElem(p1), func(p person) (string, int) {
		return p.name, p.age
	})
	p2 := []person{{"joe", 35}, {"ann", 30}, {"josh", 15}}
	i2 := T12(SliceElem(p2), func(p person) (string, int) {
		return p.name, p.age
	})

	actual := make(map[string]int)
	for name, age := range Concat2(i1, i2) {
		actual[name] = age
	}
	expect := map[string]int{"john": 25, "jane": 20, "joe": 35, "ann": 30, "josh": 15}
	if !maps.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	actual = make(map[string]int)
	for name, age := range Concat2(i1, i2) {
		if name == "ann" {
			break
		}
		actual[name] = age
	}
	expect = map[string]int{"john": 25, "jane": 20, "joe": 35}
	if !maps.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestFilter(t *testing.T) {
	predicate := func(v int) bool {
		return v%2 == 0
	}
	actual := make([]int, 0, 3)
	for v := range Filter(Range(0, 10), predicate) {
		actual = append(actual, v)
	}
	expect := []int{0, 2, 4, 6, 8}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestFilter2(t *testing.T) {
	predicate := func(name string, age int) bool {
		return name == "john"
	}
	input := map[string]int{"john": 20, "jane": 18}
	actual := map[string]int{}
	for k, v := range Filter2(Map(input), predicate) {
		actual[k] = v
	}
	expect := map[string]int{"john": 20}
	if !maps.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestCount(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	seq := Filter(SliceElem(input), func(v int) bool {
		return v%2 == 0
	})

	expect := 2
	actual := Count(seq)
	if actual != expect {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestCount2(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	seq := Filter2(Slice(input), func(idx int, v int) bool {
		return idx != 0
	})

	expect := 4
	actual := Count2(seq)
	if actual != expect {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}
