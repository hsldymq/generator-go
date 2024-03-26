//go:build goexperiment.rangefunc

package goiter

import (
	"fmt"
	"maps"
	"slices"
	"testing"
)

func TestCombine(t *testing.T) {
	actual := make([]Combined[int, string], 0, 3)
	for v := range Combine(Slice([]string{"1", "2", "3"})) {
		actual = append(actual, *v)
	}
	expect := []Combined[int, string]{
		{V1: 0, V2: "1"},
		{V1: 1, V2: "2"},
		{V1: 2, V2: "3"},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestZip(t *testing.T) {
	// case 1
	seq1 := SliceElem([]string{"Alice", "Bob", "Eve"})
	seq2 := SliceElem([]int{20, 21, 22, 23}) // seq2 has one more element than seq1
	actual := make([]Combined[string, int], 0, 3)
	for v := range Zip(seq1, seq2) {
		actual = append(actual, *v)
	}
	expect := []Combined[string, int]{
		{V1: "Alice", V2: 20},
		{V1: "Bob", V2: 21},
		{V1: "Eve", V2: 22},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	// case 2
	seq1 = SliceElem([]string{"Alice", "Bob", "Eve"})
	seq2 = SliceElem([]int{20, 21, 22, 23})
	actual = make([]Combined[string, int], 0, 2)
	i := 0
	for v := range Zip(seq1, seq2) {
		actual = append(actual, *v)
		i++
		if i >= 2 {
			break
		}
	}
	expect = []Combined[string, int]{
		{V1: "Alice", V2: 20},
		{V1: "Bob", V2: 21},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestZipAs(t *testing.T) {
	type person struct {
		Name string
		Age  int
	}
	transformer := func(zipped *ZippedE[string, int]) person {
		p := person{
			Name: zipped.V1,
			Age:  zipped.V2,
		}
		if !zipped.OK1 {
			p.Name = "?"
		}
		if !zipped.OK2 {
			p.Age = -1
		}
		return p
	}

	// case 1
	nameSeq := SliceElem([]string{"Alice", "Bob", "Eve"})
	ageSeq := SliceElem([]int{20, 21})
	zipSeq := ZipAs(nameSeq, ageSeq, transformer, true)
	actual := make([]person, 0, 3)
	for each := range zipSeq {
		actual = append(actual, each)
	}
	expect := []person{
		{Name: "Alice", Age: 20},
		{Name: "Bob", Age: 21},
		{Name: "Eve", Age: -1},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	// case 2
	nameSeq = SliceElem([]string{"Alice", "Bob", "Eve"})
	ageSeq = SliceElem([]int{20, 21, 22, 23})
	zipSeq = ZipAs(nameSeq, ageSeq, transformer, true)
	actual = make([]person, 0, 4)
	for each := range zipSeq {
		actual = append(actual, each)
	}
	expect = []person{
		{Name: "Alice", Age: 20},
		{Name: "Bob", Age: 21},
		{Name: "Eve", Age: 22},
		{Name: "?", Age: 23},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestConcat(t *testing.T) {
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

	for _ = range SliceElem([]int{10, 11, 12}).Concat(SliceElem(c2)) {
		break
	}
}

func TestConcat2(t *testing.T) {
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

	for _, _ = range i1.Concat(i2) {
		break
	}
}
