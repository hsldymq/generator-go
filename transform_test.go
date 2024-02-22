//go:build goexperiment.rangefunc

package goiter

import (
	"fmt"
	"slices"
	"testing"
)

func TestPickK(t *testing.T) {
	actual := make([]int, 0, 3)
	for idx := range PickK(Slice([]int{7, 8, 9})) {
		actual = append(actual, idx)
	}

	expect := []int{0, 1, 2}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestPickV(t *testing.T) {
	actual := make([]int, 0, 3)
	for v := range PickV(Slice([]int{7, 8, 9})) {
		actual = append(actual, v)
	}

	expect := []int{7, 8, 9}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestT11(t *testing.T) {
	transformFunc := func(v int) string {
		return fmt.Sprintf("%d", v)
	}
	actual := make([]string, 0, 3)
	for v := range T11(SliceElem([]int{1, 2, 3}), transformFunc) {
		actual = append(actual, v)
	}

	expect := []string{"1", "2", "3"}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestT12(t *testing.T) {
	transformFunc := func(v int) (int, string) {
		return v + 10, fmt.Sprintf("%d", v)
	}
	actualK := make([]int, 0, 3)
	actualV := make([]string, 0, 3)
	for k, v := range T12(SliceIdx([]int{1, 2, 3}), transformFunc) {
		actualK = append(actualK, k)
		actualV = append(actualV, v)
	}

	expectK := []int{10, 11, 12}
	if !slices.Equal(expectK, actualK) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expectK, actualK))
	}
	expectV := []string{"0", "1", "2"}
	if !slices.Equal(expectV, actualV) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expectV, actualV))
	}
}

func TestT21(t *testing.T) {
	transformFunc := func(k int, v int) string {
		return fmt.Sprintf("%d_%d", k, v)
	}
	actual := make([]string, 0, 3)
	for v := range T21(Slice([]int{1, 2, 3}), transformFunc) {
		actual = append(actual, v)
	}

	expect := []string{"0_1", "1_2", "2_3"}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestT22(t *testing.T) {
	transformFunc := func(k int, v int) (string, string) {
		return fmt.Sprintf("%d", k+10), fmt.Sprintf("%d", v+100)
	}
	actualK := make([]string, 0, 3)
	actualV := make([]string, 0, 3)
	for k, v := range T22(Slice([]int{1, 2, 3}), transformFunc) {
		actualK = append(actualK, k)
		actualV = append(actualV, v)
	}

	expectK := []string{"10", "11", "12"}
	if !slices.Equal(expectK, actualK) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expectK, actualK))
	}
	expectV := []string{"101", "102", "103"}
	if !slices.Equal(expectV, actualV) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expectV, actualV))
	}
}
