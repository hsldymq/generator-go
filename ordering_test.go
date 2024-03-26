//go:build goexperiment.rangefunc

package goiter

import (
    "slices"
    "testing"
)

func TestOrder(t *testing.T) {
	// case 1: asc
	actual1 := make([]int, 0, 3)
	for v := range Order(SliceElem([]int{4, 1, 6})) {
		actual1 = append(actual1, v)
	}
	expect1 := []int{1, 4, 6}
	if !slices.Equal(expect1, actual1) {
		t.Fatal("expect:", expect1, "actual:", actual1)
	}

	// case 2: desc
	actual2 := make([]string, 0, 3)
	for v := range Order(SliceElem([]string{"bob", "alice", "eve"}), true) {
		actual2 = append(actual2, v)
	}
	expect2 := []string{"eve", "bob", "alice"}
	if !slices.Equal(expect2, actual2) {
		t.Fatal("expect:", expect2, "actual:", actual2)
	}

	// won't panic
	for _ = range Order(SliceElem([]int{1, 2, 3})) {
		break
	}
}

func TestOrderV1(t *testing.T) {
	input := map[string]int{
		"bob":   20,
		"eve":   18,
		"alice": 22,
	}

	// case 1: asc
	actual1E1 := []string{}
	actual1E2 := []int{}
	for v1, v2 := range OrderV1(Map(input)) {
		actual1E1 = append(actual1E1, v1)
		actual1E2 = append(actual1E2, v2)
	}
	expect1E1 := []string{"alice", "bob", "eve"}
	if !slices.Equal(expect1E1, actual1E1) {
		t.Fatal("expect:", expect1E1, "actual:", actual1E1)
	}
	expect1E2 := []int{22, 20, 18}
	if !slices.Equal(expect1E2, actual1E2) {
		t.Fatal("expect:", expect1E2, "actual:", actual1E2)
	}

	// case 2: desc
	actual2E1 := []string{}
	actual2E2 := []int{}
	for v1, v2 := range OrderV1(Map(input), true) {
		actual2E1 = append(actual2E1, v1)
		actual2E2 = append(actual2E2, v2)
	}
	expect2E1 := []string{"eve", "bob", "alice"}
	if !slices.Equal(expect2E1, actual2E1) {
		t.Fatal("expect:", expect2E1, "actual:", actual2E1)
	}
	expect2E2 := []int{18, 20, 22}
	if !slices.Equal(expect2E2, actual2E2) {
		t.Fatal("expect:", expect2E2, "actual:", actual2E2)
	}

	// won't panic
	for _, _ = range OrderV1(Map(input)) {
		break
	}
}

func TestOrderV2(t *testing.T) {
	input := map[string]int{
		"bob":   20,
		"eve":   18,
		"alice": 22,
	}

	// case 1: asc
	actual1E1 := []string{}
	actual1E2 := []int{}
	for v1, v2 := range OrderV2(Map(input)) {
		actual1E1 = append(actual1E1, v1)
		actual1E2 = append(actual1E2, v2)
	}
	expect1E1 := []string{"eve", "bob", "alice"}
	if !slices.Equal(expect1E1, actual1E1) {
		t.Fatal("expect:", expect1E1, "actual:", actual1E1)
	}
	expect1E2 := []int{18, 20, 22}
	if !slices.Equal(expect1E2, actual1E2) {
		t.Fatal("expect:", expect1E2, "actual:", actual1E2)
	}

	// case 2: desc
	actual2E1 := []string{}
	actual2E2 := []int{}
	for v1, v2 := range OrderV2(Map(input), true) {
		actual2E1 = append(actual2E1, v1)
		actual2E2 = append(actual2E2, v2)
	}
	expect2E1 := []string{"alice", "bob", "eve"}
	if !slices.Equal(expect2E1, actual2E1) {
		t.Fatal("expect:", expect2E1, "actual:", actual2E1)
	}
	expect2E2 := []int{22, 20, 18}
	if !slices.Equal(expect2E2, actual2E2) {
		t.Fatal("expect:", expect2E2, "actual:", actual2E2)
	}
}
