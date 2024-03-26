//go:build goexperiment.rangefunc

package goiter

import (
    "fmt"
    "slices"
    "testing"
)

func TestDistinct(t *testing.T) {
	actual := []int{}
	for each := range Distinct(SliceElem([]int{1, 2, 3, 4, 4, 3, 2, 1})) {
		actual = append(actual, each)
	}
	expect := []int{1, 2, 3, 4}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	//
	for _ = range Distinct(SliceElem([]int{1, 2, 3, 4, 4, 3, 2, 1})) {
		break
	}
}

func TestDistinctV1(t *testing.T) {
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
		{"john", 23}, // repeated name, so DistinctV1 will ignore this
	}
	actual := []student{}
	for name, age := range DistinctV1(Transform12(SliceElem(input), transFunc)) {
		actual = append(actual, student{name, age})
	}

	expect := []student{
		{"john", 20},
		{"jane", 18},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	//
	for _, _ = range DistinctV1(Transform12(SliceElem(input), transFunc)) {
		break
	}
}

func TestDistinctV2(t *testing.T) {
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
		{"alex", 20}, // alex has the same age as john, so DistinctV2 will ignore this
	}
	actual := []student{}
	for name, age := range DistinctV2(Transform12(SliceElem(input), transFunc)) {
		actual = append(actual, student{name, age})
	}

	expect := []student{
		{"john", 20},
		{"jane", 18},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	//
	for _, _ = range DistinctV2(Transform12(SliceElem(input), transFunc)) {
		break
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

	//
	for _ = range DistinctBy(SliceElem(input), transFunc) {
		break
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
	for _, each := range DistinctBy2(Transform12(SliceElem(input), transFunc), keySelector) {
		actual = append(actual, each)
	}

	expect := []student{
		{"john", 20},
		{"jane", 18},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	//
	for _, _ = range DistinctBy2(Transform12(SliceElem(input), transFunc), keySelector) {
		break
	}
}
