//go:build goexperiment.rangefunc

package goiter

import (
	"fmt"
	"maps"
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

func TestSwapKV(t *testing.T) {
	input := map[string]int{"1": 1, "2": 2}
	actual := make(map[int]string)
	for k, v := range SwapKV(Map(input)) {
		actual[k] = v
	}
	expect := map[int]string{1: "1", 2: "2"}
	if !maps.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestCombineKV(t *testing.T) {
	input := map[string]int{"1": 1, "2": 2}
	actual := make([]KV[string, int], 0)
	for kvPair := range CombineKV(OrderV(Map(input))) {
		actual = append(actual, *kvPair)
	}
	expect := []KV[string, int]{
		{K: "1", V: 1},
		{K: "2", V: 2},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

}

func TestT11(t *testing.T) {
	transformFunc := func(v int) string {
		return fmt.Sprintf("%d", v)
	}

	actual := make([]string, 0, 3)
	for v := range T1(SliceElem([]int{1, 2, 3}), transformFunc) {
		actual = append(actual, v)
	}
	expect := []string{"1", "2", "3"}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	actual = make([]string, 0, 2)
	i := 0
	for v := range T1(SliceElem([]int{1, 2, 3}), transformFunc) {
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

func TestT12(t *testing.T) {
	transformFunc := func(v int) (int, string) {
		return v + 10, fmt.Sprintf("%d", v)
	}

	actualK := make([]int, 0, 3)
	actualV := make([]string, 0, 3)
	for k, v := range T12(SliceElem([]int{1, 2, 3}), transformFunc) {
		actualK = append(actualK, k)
		actualV = append(actualV, v)
	}
	expectK := []int{11, 12, 13}
	if !slices.Equal(expectK, actualK) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expectK, actualK))
	}
	expectV := []string{"1", "2", "3"}
	if !slices.Equal(expectV, actualV) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expectV, actualV))
	}

	actualK = make([]int, 0, 3)
	i := 0
	for k, _ := range T12(SliceElem([]int{1, 2, 3}), transformFunc) {
		actualK = append(actualK, k)
		i++
		if i >= 2 {
			break
		}
	}
	expectK = []int{11, 12}
	if !slices.Equal(expectK, actualK) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expectK, actualK))
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

	actual = make([]string, 0, 3)
	i := 0
	for v := range T21(Slice([]int{1, 2, 3}), transformFunc) {
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

func TestT22(t *testing.T) {
	transformFunc := func(k int, v int) (string, string) {
		return fmt.Sprintf("%d", k+10), fmt.Sprintf("%d", v+100)
	}

	actualK := make([]string, 0, 3)
	actualV := make([]string, 0, 3)
	for k, v := range T2(Slice([]int{1, 2, 3}), transformFunc) {
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

	actualV = make([]string, 0, 3)
	i := 0
	for _, v := range T2(Slice([]int{1, 2, 3}), transformFunc) {
		actualV = append(actualV, v)
		i++
		if i >= 2 {
			break
		}
	}
	expectV = []string{"101", "102"}
	if !slices.Equal(expectV, actualV) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expectV, actualV))
	}
}

func TestZip(t *testing.T) {
	// case 1
	seq1 := SliceElem([]string{"Alice", "Bob", "Eve"})
	seq2 := SliceElem([]int{20, 21, 22, 23}) // seq2 has one more element than seq1
	actual := make([]Zipped[string, int], 0, 3)
	for v := range Zip(seq1, seq2) {
		actual = append(actual, *v)
	}
	expect := []Zipped[string, int]{
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
	actual = make([]Zipped[string, int], 0, 2)
	i := 0
	for v := range Zip(seq1, seq2) {
		actual = append(actual, *v)
		i++
		if i >= 2 {
			break
		}
	}
	expect = []Zipped[string, int]{
		{V1: "Alice", V2: 20},
		{V1: "Bob", V2: 21},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestZipAs(t *testing.T) {
}

func TestToSlice(t *testing.T) {
	seq := func(yield func(int) bool) {
		yield(1)
		yield(2)
		yield(3)
	}

	actual := ToSlice(seq)
	expect := []int{1, 2, 3}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestToMap(t *testing.T) {
	seq := func(yield func(string, int) bool) {
		yield("alice", 20)
		yield("bob", 21)
		yield("eve", 22)
	}

	actual := ToMap(seq)
	expect := map[string]int{
		"alice": 20,
		"bob":   21,
		"eve":   22,
	}
	if !maps.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestToMapBy(t *testing.T) {
	seq := func(yield func(string, string) bool) {
		yield("Alice", "Paris")
		yield("Bob", "Shanghai")
		yield("Eve", "Bangkok")
	}

	actual := ToMapBy(seq, func(name string, city string) (string, string) {
		return name + "_" + city, string(name[0]) + "_" + string(city[0])
	})
	expect := map[string]string{
		"Alice_Paris":  "A_P",
		"Bob_Shanghai": "B_S",
		"Eve_Bangkok":  "E_B",
	}
	if !maps.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestToMapByV(t *testing.T) {
	seq := func(yield func(string) bool) {
		yield("alice")
		yield("bob")
		yield("eve")
	}

	actual := ToMapByV(seq, func(name string) (string, int) {
		return name, len(name)
	})
	expect := map[string]int{
		"alice": 5,
		"bob":   3,
		"eve":   3,
	}
	if !maps.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}
