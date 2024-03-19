package goiter

import (
	"cmp"
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

func TestOrderK(t *testing.T) {
	input := map[string]int{
		"bob":   20,
		"eve":   18,
		"alice": 22,
	}

	// case 1: asc
	actual1K := []string{}
	actual1V := []int{}
	for k, v := range OrderK(Map(input)) {
		actual1K = append(actual1K, k)
		actual1V = append(actual1V, v)
	}
	expect1K := []string{"alice", "bob", "eve"}
	if !slices.Equal(expect1K, actual1K) {
		t.Fatal("expect:", expect1K, "actual:", actual1K)
	}
	expect1V := []int{22, 20, 18}
	if !slices.Equal(expect1V, actual1V) {
		t.Fatal("expect:", expect1V, "actual:", actual1V)
	}

	// case 2: desc
	actual2K := []string{}
	actual2V := []int{}
	for k, v := range OrderK(Map(input), true) {
		actual2K = append(actual2K, k)
		actual2V = append(actual2V, v)
	}
	expect2K := []string{"eve", "bob", "alice"}
	if !slices.Equal(expect2K, actual2K) {
		t.Fatal("expect:", expect2K, "actual:", actual2K)
	}
	expect2V := []int{18, 20, 22}
	if !slices.Equal(expect2V, actual2V) {
		t.Fatal("expect:", expect2V, "actual:", actual2V)
	}

	// won't panic
	for _, _ = range OrderK(Map(input)) {
		break
	}
}

func TestOrderV(t *testing.T) {
	input := map[string]int{
		"bob":   20,
		"eve":   18,
		"alice": 22,
	}

	// case 1: asc
	actual1K := []string{}
	actual1V := []int{}
	for k, v := range OrderV(Map(input)) {
		actual1K = append(actual1K, k)
		actual1V = append(actual1V, v)
	}
	expect1K := []string{"eve", "bob", "alice"}
	if !slices.Equal(expect1K, actual1K) {
		t.Fatal("expect:", expect1K, "actual:", actual1K)
	}
	expect1V := []int{18, 20, 22}
	if !slices.Equal(expect1V, actual1V) {
		t.Fatal("expect:", expect1V, "actual:", actual1V)
	}

	// case 2: desc
	actual2K := []string{}
	actual2V := []int{}
	for k, v := range OrderV(Map(input), true) {
		actual2K = append(actual2K, k)
		actual2V = append(actual2V, v)
	}
	expect2K := []string{"alice", "bob", "eve"}
	if !slices.Equal(expect2K, actual2K) {
		t.Fatal("expect:", expect2K, "actual:", actual2K)
	}
	expect2V := []int{22, 20, 18}
	if !slices.Equal(expect2V, actual2V) {
		t.Fatal("expect:", expect2V, "actual:", actual2V)
	}
}

func TestOrderBy(t *testing.T) {
	type person struct {
		name string
		age  int
	}
	input := []person{
		{"alice", 22},
		{"bob", 20},
		{"eve", 21},
	}
	actual := []person{}
	for each := range OrderBy(SliceElem(input), func(a, b person) int { return cmp.Compare(a.age, b.age) }) {
		actual = append(actual, each)
	}
	expect := []person{
		{"bob", 20},
		{"eve", 21},
		{"alice", 22},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal("expect:", expect, "actual:", actual)
	}
}

func TestOrderBy2(t *testing.T) {
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
	for k, v := range OrderBy2(Map(input), func(a, b *KV[string, int]) int { return cmp.Compare(a.V, b.V) }) {
		actual = append(actual, person{name: k, age: v})
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

func TestStableOrderBy(t *testing.T) {
	type person struct {
		name string
		age  int
	}
	input := []person{
		{"alice", 22},
		{"bob", 20},
		{"eve", 20},
	}
	actual := []person{}
	for each := range StableOrderBy(SliceElem(input), func(a, b person) int { return cmp.Compare(a.age, b.age) }) {
		actual = append(actual, each)
	}
	expect := []person{
		{"bob", 20},
		{"eve", 20},
		{"alice", 22},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal("expect:", expect, "actual:", actual)
	}
}

func TestStableOrderBy2(t *testing.T) {
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
	for _, v := range StableOrderBy2(Slice(input), func(a, b *KV[int, person]) int { return cmp.Compare(a.V.age, b.V.age) }) {
		actual = append(actual, v)
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
