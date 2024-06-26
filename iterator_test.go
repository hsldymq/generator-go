package goiter

import (
    "fmt"
    "slices"
    "testing"
)

func TestIterator_WithCounter(t *testing.T) {
    type rt struct {
        c int
        v string
    }
    actual := make([]rt, 0, 3)
    for c, v := range SliceElems([]string{"a", "b", "c"}).WithCounter(1) {
        actual = append(actual, rt{c, v})
    }
    expect := []rt{
        {1, "a"},
        {2, "b"},
        {3, "c"},
    }
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
}

func TestIterator_Cache(t *testing.T) {
    count := 0

    iterator := SliceElems([]int{1, 2, 3, 4, 5, 6}).
        Filter(func(v int) bool {
            count++
            return v%2 == 0
        }).Cache()
    actual := make([]int, 0, 3)
    for v := range iterator {
        actual = append(actual, v)
    }
    expect := []int{2, 4, 6}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
    if count != 6 {
        t.Fatal(fmt.Sprintf("expect: %d, actual: %d", 6, count))
    }
    for v := range iterator {
        if v == 6 {
            break
        }
    }
    if count != 6 {
        t.Fatal(fmt.Sprintf("expect: %d, actual: %d", 6, count))
    }

    iterator = SliceElems([]int{1, 2, 3, 4, 5, 6}).
        Filter(func(v int) bool {
            count++
            return v%2 == 0
        }).Cache()
    for _ = range iterator {
        break
    }
}
