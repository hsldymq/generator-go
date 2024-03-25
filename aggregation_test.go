//go:build goexperiment.rangefunc

package goiter

import (
    "fmt"
    "testing"
)

func TestCount(t *testing.T) {
    input := []int{1, 2, 3, 4, 5}
    seq := Filter(SliceElem(input), func(v int) bool {
        return v%2 == 0
    })

    expect := 2
    actual := Count(seq)
    if actual != expect {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
}

func TestCount2(t *testing.T) {
    input := []int{1, 2, 3, 4, 5}
    seq := Filter2(Slice(input), func(idx int, v int) bool {
        return idx != 0
    })

    expect := 4
    actual := Count2(seq)
    if actual != expect {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
}

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
