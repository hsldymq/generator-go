//go:build goexperiment.rangefunc

package goiter

import (
	"fmt"
	"maps"
	"slices"
	"testing"
)

func TestPick1(t *testing.T) {
	actual := make([]int, 0, 3)
	for idx := range PickV1(Slice([]int{7, 8, 9})) {
		actual = append(actual, idx)
	}

	expect := []int{0, 1, 2}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestPick2(t *testing.T) {
	actual := make([]int, 0, 3)
	for v := range PickV2(Slice([]int{7, 8, 9})) {
		actual = append(actual, v)
	}

	expect := []int{7, 8, 9}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestSwap(t *testing.T) {
	input := map[string]int{"1": 1, "2": 2}
	actual := make(map[int]string)
	for val, key := range Swap(Map(input)) {
		actual[val] = key
	}
	expect := map[int]string{1: "1", 2: "2"}
	if !maps.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestTransform(t *testing.T) {
	transformFunc := func(v int) string {
		return fmt.Sprintf("%d", v)
	}

	actual := make([]string, 0, 3)
	for v := range Transform(SliceElem([]int{1, 2, 3}), transformFunc) {
		actual = append(actual, v)
	}
	expect := []string{"1", "2", "3"}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	actual = make([]string, 0, 2)
	i := 0
	for v := range Transform(SliceElem([]int{1, 2, 3}), transformFunc) {
		actual = append(actual, v)
		i++
		if i >= 2 {
			break
		}
	}
	expect = []string{"1", "2"}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestTransform2(t *testing.T) {
	transformFunc := func(k int, v int) (string, string) {
		return fmt.Sprintf("%d", k+10), fmt.Sprintf("%d", v+100)
	}

	actualV1 := make([]string, 0, 3)
	actualV2 := make([]string, 0, 3)
	for v1, v2 := range Transform2(Slice([]int{1, 2, 3}), transformFunc) {
		actualV1 = append(actualV1, v1)
		actualV2 = append(actualV2, v2)
	}
	expectV1 := []string{"10", "11", "12"}
	if !slices.Equal(expectV1, actualV1) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expectV1, actualV1))
	}
	expectV2 := []string{"101", "102", "103"}
	if !slices.Equal(expectV2, actualV2) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expectV2, actualV2))
	}

	actualV2 = make([]string, 0, 3)
	i := 0
	for _, v := range Transform2(Slice([]int{1, 2, 3}), transformFunc) {
		actualV2 = append(actualV2, v)
		i++
		if i >= 2 {
			break
		}
	}
	expectV2 = []string{"101", "102"}
	if !slices.Equal(expectV2, actualV2) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expectV2, actualV2))
	}
}

func TestTransform12(t *testing.T) {
	transformFunc := func(v int) (int, string) {
		return v + 10, fmt.Sprintf("%d", v)
	}

	actualV1 := make([]int, 0, 3)
	actualV2 := make([]string, 0, 3)
	for v1, v2 := range Transform12(SliceElem([]int{1, 2, 3}), transformFunc) {
		actualV1 = append(actualV1, v1)
		actualV2 = append(actualV2, v2)
	}
	expectV1 := []int{11, 12, 13}
	if !slices.Equal(expectV1, actualV1) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expectV1, actualV1))
	}
	expectV2 := []string{"1", "2", "3"}
	if !slices.Equal(expectV2, actualV2) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expectV2, actualV2))
	}

	actualV1 = make([]int, 0, 3)
	i := 0
	for v1, _ := range Transform12(SliceElem([]int{1, 2, 3}), transformFunc) {
		actualV1 = append(actualV1, v1)
		i++
		if i >= 2 {
			break
		}
	}
	expectV1 = []int{11, 12}
	if !slices.Equal(expectV1, actualV1) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expectV1, actualV1))
	}
}

func TestTransform21(t *testing.T) {
	transformFunc := func(k int, v int) string {
		return fmt.Sprintf("%d_%d", k, v)
	}

	actual := make([]string, 0, 3)
	for v := range Transform21(Slice([]int{1, 2, 3}), transformFunc) {
		actual = append(actual, v)
	}
	expect := []string{"0_1", "1_2", "2_3"}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	actual = make([]string, 0, 3)
	i := 0
	for v := range Transform21(Slice([]int{1, 2, 3}), transformFunc) {
		actual = append(actual, v)
		i++
		if i >= 2 {
			break
		}
	}
	expect = []string{"0_1", "1_2"}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}
