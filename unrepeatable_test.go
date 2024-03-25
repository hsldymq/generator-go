//go:build goexperiment.rangefunc

package goiter

import (
    "fmt"
    "slices"
    "testing"
)

func TestOnce(t *testing.T) {
    input := []int{1, 2, 3, 4, 5, 6}

    // case 1
    actual := make([]int, 0, 3)
    iterator := Once(SliceElem(input))
    for v := range iterator {
        actual = append(actual, v)
        if v == 3 {
            break
        }
    }
    for v := range iterator {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 3}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    // case 2
    iterator = Once(SliceElem(input))
    actual = make([]int, 0)
    for v := range iterator {
        actual = append(actual, v)
    }
    for v := range iterator {
        actual = append(actual, v)
    }
    expect = []int{1, 2, 3, 4, 5, 6}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
}

func TestOnce2(t *testing.T) {
    input := []int{1, 2, 3, 4, 5, 6}

    // case 1
    actual := make([]int, 0, 3)
    iterator := Once2(Slice(input))
    for idx, v := range iterator {
        actual = append(actual, v)
        if idx == 2 {
            break
        }
    }
    for _, v := range iterator {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 3}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    // case 2
    iterator = Once2(Slice(input))
    actual = make([]int, 0)
    for _, v := range iterator {
        actual = append(actual, v)
    }
    for _, v := range iterator {
        actual = append(actual, v)
    }
    expect = []int{1, 2, 3, 4, 5, 6}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
}

func TestContinuableOnce(t *testing.T) {
    input := []int{1, 2, 3, 4, 5, 6}

    // case 1
    actual := make([]int, 0)
    iterator := ContinuableOnce(SliceElem(input))
    for v := range iterator {
        actual = append(actual, v)
        if v == 3 {
            break
        }
    }
    for v := range iterator {
        actual = append(actual, v)
    }
    for v := range iterator {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 3, 4, 5, 6}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    // case 2
    iterator = ContinuableOnce(SliceElem(input))
    actual = make([]int, 0)
    for v := range iterator {
        actual = append(actual, v)
    }
    for v := range iterator {
        actual = append(actual, v)
    }
    expect = []int{1, 2, 3, 4, 5, 6}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
}

func TestContinuableOnce2(t *testing.T) {
    input := []int{1, 2, 3, 4, 5, 6}

    // case 1
    actual := make([]int, 0, 3)
    iterator := ContinuableOnce2(Slice(input))
    for idx, v := range iterator {
        actual = append(actual, v)
        if idx == 2 {
            break
        }
    }
    for _, v := range iterator {
        actual = append(actual, v)
    }
    for _, v := range iterator {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 3, 4, 5, 6}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    // case 2
    iterator = ContinuableOnce2(Slice(input))
    actual = make([]int, 0)
    for _, v := range iterator {
        actual = append(actual, v)
    }
    for _, v := range iterator {
        actual = append(actual, v)
    }
    expect = []int{1, 2, 3, 4, 5, 6}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
}

func TestContinuable(t *testing.T) {
    input := []int{1, 2, 3, 4, 5, 6}

    actual := make([]int, 0, 6)
    iterator := continuable(SliceElem(input))
    for v := range iterator {
        actual = append(actual, v)
        if v == 3 {
            break
        }
    }
    for v := range iterator {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 3, 4, 5, 6}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    actual = make([]int, 0, 12)
    for v := range iterator {
        actual = append(actual, v)
    }
    for v := range iterator {
        actual = append(actual, v)
    }
    expect = []int{1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 6}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
}

func TestContinuable2(t *testing.T) {
    input := []int{1, 2, 3, 4, 5, 6}

    actual := make([]int, 0, 6)
    iterator := continuable2(Slice(input))
    for idx, v := range iterator {
        actual = append(actual, v)
        if idx == 2 {
            break
        }
    }
    for _, v := range iterator {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 3, 4, 5, 6}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    actual = make([]int, 0, 12)
    for _, v := range iterator {
        actual = append(actual, v)
    }
    for _, v := range iterator {
        actual = append(actual, v)
    }
    expect = []int{1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 6}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
}
