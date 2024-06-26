package goiter

import (
    "fmt"
    "slices"
    "testing"
)

func TestReduce(t *testing.T) {
    foldFunc := func(a int, b int) int {
        return a + b
    }
    actual := Reduce(SliceElems([]int{1, 2, 3}), 0, foldFunc)
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
    for each := range Scan(SliceElems([]int{1, 2, 3, 4, 5}), 0, foldFunc) {
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
    for each := range Scan(SliceElems([]int{100, 200, 300, 400, 500}), st{}, accFunc) {
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

func TestCount(t *testing.T) {
    input := []int{1, 2, 3, 4, 5}
    it := Filter(SliceElems(input), func(v int) bool {
        return v%2 == 0
    })

    expect := 2
    actual := it.Count()
    if actual != expect {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
}

func TestCount2(t *testing.T) {
    input := []int{1, 2, 3, 4, 5}
    it := Filter2(Slice(input), func(idx int, v int) bool {
        return idx != 0
    })

    expect := 4
    actual := it.Count()
    if actual != expect {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
}
