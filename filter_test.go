//go:build goexperiment.rangefunc

package goiter

import (
	"fmt"
	"maps"
	"slices"
	"testing"
)

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

func TestDistinct(t *testing.T) {
	actual := []int{}
	for each := range Distinct(SliceElem([]int{1, 2, 3, 4, 4, 3, 2, 1})) {
		actual = append(actual, each)
	}

	expect := []int{1, 2, 3, 4}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestDistinctK(t *testing.T) {
	type student struct {
		Name string
		Age  int
	}

	transFunc := func(s student) (string, int) {
		return s.Name, s.Age // name as key, age as value
	}
	input := []student{
		{"john", 20},
		{"jane", 18},
		{"john", 23}, // repeated name, so DistinctK will ignore this
	}
	actual := []student{}
	for name, age := range DistinctK(T12(SliceElem(input), transFunc)) {
		actual = append(actual, student{name, age})
	}

	expect := []student{
		{"john", 20},
		{"jane", 18},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestDistinctV(t *testing.T) {
	type student struct {
		Name string
		Age  int
	}

	transFunc := func(s student) (string, int) {
		return s.Name, s.Age // name as key, age as value
	}
	input := []student{
		{"john", 20},
		{"jane", 18},
		{"alex", 20}, // alex has the same age as john, so DistinctV will ignore this
	}
	actual := []student{}
	for name, age := range DistinctV(T12(SliceElem(input), transFunc)) {
		actual = append(actual, student{name, age})
	}

	expect := []student{
		{"john", 20},
		{"jane", 18},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestDistinctBy(t *testing.T) {
	type student struct {
		Name string
		Age  int
	}

	input := []student{
		{"john", 20},
		{"jane", 18},
		{"john", 23},
	}
	transFunc := func(s student) string { return s.Name }
	actual := []student{}
	for each := range DistinctBy(SliceElem(input), transFunc) {
		actual = append(actual, each)
	}

	expect := []student{
		{"john", 20},
		{"jane", 18},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestDistinctBy2(t *testing.T) {
	type student struct {
		Name string
		Age  int
	}

	input := []student{
		{"john", 20},
		{"jane", 18},
		{"alex", 20},
	}
	transFunc := func(s student) (int, student) { return s.Age, s }
	keySelector := func(age int, s student) int { return s.Age }
	actual := []student{}
	for _, each := range DistinctBy2(T12(SliceElem(input), transFunc), keySelector) {
		actual = append(actual, each)
	}

	expect := []student{
		{"john", 20},
		{"jane", 18},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}
