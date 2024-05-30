//go:build goexperiment.rangefunc

package goiter

import (
    "fmt"
    "slices"
    "testing"
)

func TestIterator2_V1(t *testing.T) {
    input := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }
    actual := make([]string, 0, 3)
    for v := range Order(Map(input).V1()) {
        actual = append(actual, v)
    }
    expect := []string{"a", "b", "c"}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
}

func TestIterator2_V2(t *testing.T) {
    actual := make([]int, 0, 3)
    for idx := range Slice([]int{7, 8, 9}).V2() {
        actual = append(actual, idx)
    }
    expect := []int{7, 8, 9}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
}

func TestIterator2_Cache(t *testing.T) {
    count := 0

    iterator := Slice([]int{1, 2, 3, 4, 5, 6}).
        Filter(func(idx, v int) bool {
            count++
            return v%2 == 0
        }).Cache()
    actual := make([]int, 0, 3)
    for _, v := range iterator {
        actual = append(actual, v)
    }
    expect := []int{2, 4, 6}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
    if count != 6 {
        t.Fatal(fmt.Sprintf("expect: %d, actual: %d", 6, count))
    }
    for _, v := range iterator {
        if v == 6 {
            break
        }
    }
    if count != 6 {
        t.Fatal(fmt.Sprintf("expect: %d, actual: %d", 6, count))
    }

    iterator = Slice([]int{1, 2, 3, 4, 5, 6}).
        Filter(func(idx, v int) bool {
            count++
            return v%2 == 0
        }).Cache()
    for _, _ = range iterator {
        break
    }
}
