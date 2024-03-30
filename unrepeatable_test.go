//go:build goexperiment.rangefunc

package goiter

import (
    "fmt"
    "slices"
    "sync"
    "testing"
    "time"
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

    // case 3
    iterator = Once(SliceElem(input))
    actual = make([]int, 0)
    actual2 := make([]int, 0)
    g := &sync.WaitGroup{}
    g.Add(2)
    go func() {
        for v := range iterator {
            time.Sleep(20 * time.Millisecond)
            actual = append(actual, v)
        }
        g.Done()
    }()
    go func() {
        time.Sleep(10 * time.Millisecond)
        for v := range iterator {
            actual2 = append(actual2, v)
        }
        g.Done()
    }()
    g.Wait()
    expect = []int{1, 2, 3, 4, 5, 6}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
    expect2 := []int{}
    if !slices.Equal(expect2, actual2) {
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

    // case 3
    iterator = Once2(Slice(input))
    actual = make([]int, 0)
    actual2 := make([]int, 0)
    g := &sync.WaitGroup{}
    g.Add(2)
    go func() {
        for _, v := range iterator {
            time.Sleep(20 * time.Millisecond)
            actual = append(actual, v)
        }
        g.Done()
    }()
    go func() {
        time.Sleep(10 * time.Millisecond)
        for _, v := range iterator {
            actual2 = append(actual2, v)
        }
        g.Done()
    }()
    g.Wait()
    expect = []int{1, 2, 3, 4, 5, 6}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
    expect2 := []int{}
    if !slices.Equal(expect2, actual2) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
}

func TestFinishOnce(t *testing.T) {
    input := []int{1, 2, 3, 4, 5, 6}

    // case 1
    actual := make([]int, 0)
    iterator := FinishOnce(SliceElem(input))
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
    iterator = FinishOnce(SliceElem(input))
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

    // case 3
    input = Range(1, 10000).ToSlice()
    iterator = FinishOnce(SliceElem(input))
    g := &sync.WaitGroup{}
    g.Add(3)
    actual1 := make([]int, 0)
    go func() {
        for v := range iterator {
            actual1 = append(actual1, v)
        }
        g.Done()
    }()
    actual2 := make([]int, 0)
    go func() {
        for v := range iterator {
            actual2 = append(actual2, v)
        }
        g.Done()
    }()
    actual3 := make([]int, 0)
    go func() {
        for v := range iterator {
            actual3 = append(actual3, v)
        }
        g.Done()
    }()
    g.Wait()
    actual = slices.Concat(actual1, actual2, actual3)
    slices.Sort(actual)
    if !slices.Equal(input, actual) {
        t.Fatal(fmt.Sprintf("\nexpect: %v\nactual: %v", input, actual))
    }
}

func TestFinishOnce2(t *testing.T) {
    input := []int{1, 2, 3, 4, 5, 6}

    // case 1
    actual := make([]int, 0, 3)
    iterator := FinishOnce2(Slice(input))
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
    iterator = FinishOnce2(Slice(input))
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

    // case 3: concurrent
    input = Range(1, 10001).ToSlice()
    iterator = FinishOnce2(Slice(input))
    g := &sync.WaitGroup{}
    g.Add(3)
    actual1 := make([]int, 0)
    go func() {
        for _, v := range iterator {
            actual1 = append(actual1, v)
        }
        g.Done()
    }()
    actual2 := make([]int, 0)
    go func() {
        for _, v := range iterator {
            actual2 = append(actual2, v)
        }
        g.Done()
    }()
    actual3 := make([]int, 0)
    go func() {
        for _, v := range iterator {
            actual3 = append(actual3, v)
        }
        g.Done()
    }()
    g.Wait()
    actual = slices.Concat(actual1, actual2, actual3)
    slices.Sort(actual)
    if !slices.Equal(input, actual) {
        t.Fatal(fmt.Sprintf("\nexpect: %v\nactual: %v", input, actual))
    }
}
