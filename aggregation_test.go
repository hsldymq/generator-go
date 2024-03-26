//go:build goexperiment.rangefunc

package goiter

import (
	"fmt"
	"slices"
	"testing"
)

func TestFold(t *testing.T) {
	foldFunc := func(a int, b int) int {
		return a + b
	}
	actual := Fold(SliceElem([]int{1, 2, 3}), 0, foldFunc)
	expect := 6
	if expect != actual {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestScan(t *testing.T) {
	foldFunc := func(a int, b int) int {
		return a + b
	}
	actual := []int{}
	for each := range Scan(SliceElem([]int{1, 2, 3, 4, 5}), 0, foldFunc) {
		actual = append(actual, each)
	}
	expect := []int{1, 3, 6, 10, 15}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	type st struct {
		Sum   int
		Count int
	}
	accFunc := func(acc st, v int) st {
		acc.Sum += v
		acc.Count++
		return acc
	}
	actual = []int{}
	for each := range Scan(SliceElem([]int{100, 200, 300, 400, 500}), st{}, accFunc) {
		actual = append(actual, each.Sum/each.Count)
		if each.Sum == 600 {
			break
		}
	}
	expect = []int{100, 150, 200}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}
