//go:build goexperiment.rangefunc

package goiter

import (
	"fmt"
	"iter"
	"maps"
	"slices"
	"testing"
)

func TestSlice(t *testing.T) {
	actual1 := make([]int, 0, 3)
	for v := range SliceElem([]int{7, 8, 9}, true) {
		actual1 = append(actual1, v)
	}
	expect1 := []int{9, 8, 7}
	if !slices.Equal(expect1, actual1) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect1, actual1))
	}

	actual2 := make([]int, 0, 1)
	for _, v := range Slice([]int{7, 8, 9}) {
		if v == 8 {
			break
		}
		actual2 = append(actual2, v)
	}
	expect2 := []int{7}
	if !slices.Equal(expect2, actual2) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect2, actual2))
	}

	actual3 := make([]int, 0, 3)
	for v := range SliceElem([]int{7, 8, 9}, true) {
		if v == 7 {
			break
		}
		actual3 = append(actual3, v)
	}
	expect3 := []int{9, 8}
	if !slices.Equal(expect3, actual3) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect3, actual3))
	}
}

func TestSliceSourceElem(t *testing.T) {
	input := []int{7, 8, 9}
	iterator := SliceSourceElem(func() []int { return input })
	actual1 := make([]int, 0, 3)
	for each := range iterator {
		actual1 = append(actual1, each)
	}
	expect1 := []int{7, 8, 9}
	if !slices.Equal(expect1, actual1) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect1, actual1))
	}

	// replace the original slice with a new one, and the iterator should traverse the new slice instead of the old one.
	input = []int{1, 2, 3}
	actual2 := make([]int, 0, 3)
	for each := range iterator {
		actual2 = append(actual2, each)
	}
	expect2 := []int{1, 2, 3}
	if !slices.Equal(expect2, actual2) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect2, actual2))
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

func TestMapSourceVal(t *testing.T) {
	input := map[string]int{"foo": 1, "bar": 2}
	iterator := MapSourceVal(func() map[string]int { return input })

	actual1 := make([]int, 0, 2)
	for v := range Order(iterator) {
		actual1 = append(actual1, v)
	}
	expect1 := []int{1, 2}
	if !slices.Equal(expect1, actual1) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect1, actual1))
	}

	input = map[string]int{"alice": 25, "bob": 20, "eve": 21}
	actual2 := make([]int, 0, 3)
	for v := range Order(iterator) {
		actual2 = append(actual2, v)
	}
	expect2 := []int{20, 21, 25}
	if !slices.Equal(expect2, actual2) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect2, actual2))
	}
}

func TestMapSourceKey(t *testing.T) {
	input := map[string]int{"foo": 1, "bar": 2}
	iterator := MapSourceKey(func() map[string]int { return input })

	actual1 := make([]string, 0, 2)
	for v := range Order(iterator) {
		actual1 = append(actual1, v)
	}
	expect1 := []string{"bar", "foo"}
	if !slices.Equal(expect1, actual1) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect1, actual1))
	}

	input = map[string]int{"eve": 21, "alice": 25, "bob": 20}
	actual2 := make([]string, 0, 3)
	for v := range Order(iterator) {
		actual2 = append(actual2, v)
	}
	expect2 := []string{"alice", "bob", "eve"}
	if !slices.Equal(expect2, actual2) {
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

func TestSeq(t *testing.T) {
	iterator := Seq(func(yield func(int) bool) {
		yield(1)
		yield(2)
		yield(3)
	})
	actual := make([]int, 0, 3)
	for v := range iterator {
		actual = append(actual, v)
	}
	expect := []int{1, 2, 3}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestSeqSource(t *testing.T) {
	itFunc := func(yield func(int) bool) {
		yield(1)
		yield(2)
		yield(3)
	}
	iterator := SeqSource(func() iter.Seq[int] {
		return itFunc
	})

	actual := make([]int, 0, 3)
	for v := range iterator {
		actual = append(actual, v)
	}
	expect := []int{1, 2, 3}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	itFunc = func(yield func(int) bool) {
		for _, each := range []int{4, 5, 6} {
			if !yield(each) {
				break
			}
		}
	}
	actual = make([]int, 0, 3)
	for v := range iterator {
		actual = append(actual, v)
		if v == 5 {
			break
		}
	}
	expect = []int{4, 5}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestSeq2(t *testing.T) {
	type person struct {
		Name string
		Age  int
	}

	itFunc := func(yield func(string, int) bool) {
		yield("alice", 20)
		yield("bob", 21)
		yield("eve", 22)
	}
	iterator := Seq2(itFunc)

	actual := make([]person, 0, 3)
	for name, age := range iterator {
		actual = append(actual, person{name, age})
	}
	expect := []person{
		{"alice", 20},
		{"bob", 21},
		{"eve", 22},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestSeq2Source(t *testing.T) {
	type person struct {
		Name string
		Age  int
	}

	itFunc := func(yield func(string, int) bool) {
		yield("alice", 20)
		yield("bob", 21)
		yield("eve", 22)
	}
	iterator := Seq2Source(func() iter.Seq2[string, int] {
		return itFunc
	})

	actual := make([]person, 0, 3)
	for name, age := range iterator {
		actual = append(actual, person{name, age})
	}
	expect := []person{
		{"alice", 20},
		{"bob", 21},
		{"eve", 22},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	itFunc = func(yield func(string, int) bool) {
		people := []person{
			{"alice", 30},
			{"bob", 31},
			{"eve", 32},
		}
		for _, each := range people {
			if !yield(each.Name, each.Age) {
				break
			}
		}
	}
	actual = make([]person, 0, 3)
	for name, age := range iterator {
		actual = append(actual, person{name, age})
		if age == 31 {
			break
		}
	}
	expect = []person{
		{"alice", 30},
		{"bob", 31},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestEmpty(t *testing.T) {
	i := 0
	for _ = range Empty[int]() {
		i++
	}

	if i != 0 {
		t.Fatal(fmt.Sprintf("expect: 0, actual: %d", i))
	}
}

func TestEmpty2(t *testing.T) {
	i := 0
	for _, _ = range Empty2[string, int]() {
		i++
	}

	if i != 0 {
		t.Fatal(fmt.Sprintf("expect: 0, actual: %d", i))
	}
}
