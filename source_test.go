//go:build goexperiment.rangefunc

package goiter

import (
    "fmt"
    "iter"
    "slices"
    "testing"
)

func TestSliceSource(t *testing.T) {
    input := []int{7, 8, 9}
    iterator := SliceSource(func() []int { return input })

    actual1 := make([]int, 0, 3)
    for _, v := range iterator {
        actual1 = append(actual1, v)
    }
    expect1 := []int{7, 8, 9}
    if !slices.Equal(expect1, actual1) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect1, actual1))
    }

    // replace the original slice with a new one, and the iterator should traverse the new slice instead of the old one.
    input = []int{1, 2, 3}
    actual2 := make([]int, 0, 3)
    for _, v := range iterator {
        actual2 = append(actual2, v)
    }
    expect2 := []int{1, 2, 3}
    if !slices.Equal(expect2, actual2) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect2, actual2))
    }

    actualIdx3 := make([]int, 0, 3)
    actualElem3 := make([]int, 0, 3)
    for idx, v := range SliceSource(func() []int { return input }, true) {
        if idx == 0 {
            break
        }
        actualIdx3 = append(actualIdx3, idx)
        actualElem3 = append(actualElem3, v)
    }
    expectIdx3 := []int{2, 1}
    if !slices.Equal(expectIdx3, actualIdx3) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expectIdx3, actualIdx3))
    }
    expectElem3 := []int{3, 2}
    if !slices.Equal(expectElem3, actualElem3) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expectElem3, actualElem3))
    }
}

func TestSliceSourceElems(t *testing.T) {
    input := []int{7, 8, 9}
    iterator := SliceSourceElems(func() []int { return input })

    actual1 := make([]int, 0, 3)
    for each := range iterator {
        actual1 = append(actual1, each)
    }
    expect1 := []int{7, 8, 9}
    if !slices.Equal(expect1, actual1) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect1, actual1))
    }

    // replace the original slice with a new one, and the iterator should traverse the new slice instead of the old one.
    input = []int{1, 2, 3}
    actual2 := make([]int, 0, 3)
    for each := range iterator {
        actual2 = append(actual2, each)
    }
    expect2 := []int{1, 2, 3}
    if !slices.Equal(expect2, actual2) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect2, actual2))
    }

    actual3 := make([]int, 0, 3)
    for v := range SliceSourceElems(func() []int { return input }, true) {
        if v == 1 {
            break
        }
        actual3 = append(actual3, v)
    }
    expect3 := []int{3, 2}
    if !slices.Equal(expect3, actual3) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect3, actual3))
    }
}

func TestMapSource(t *testing.T) {
    input := map[string]int{"foo": 1, "bar": 2}
    iterator := MapSource(func() map[string]int { return input })

    actual1 := make([]string, 0, 2)
    for k, _ := range iterator {
        actual1 = append(actual1, k)
    }
    slices.Sort(actual1)
    expect1 := []string{"bar", "foo"}
    if !slices.Equal(expect1, actual1) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect1, actual1))
    }

    // replace the original map with a new one, and the iterator should traverse the new slice instead of the old one.
    input = map[string]int{"goo": 1, "tar": 2}
    actual2 := make([]string, 0, 2)
    for k, _ := range iterator {
        actual2 = append(actual2, k)
    }
    slices.Sort(actual2)
    expect2 := []string{"goo", "tar"}
    if !slices.Equal(expect2, actual2) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect2, actual2))
    }

    for _, _ = range iterator {
        break
    }
}

func TestMapSourceVals(t *testing.T) {
    input := map[string]int{"foo": 1, "bar": 2}
    iterator := MapSourceVals(func() map[string]int { return input })

    actual1 := make([]int, 0, 2)
    for v := range Order(iterator) {
        actual1 = append(actual1, v)
    }
    expect1 := []int{1, 2}
    if !slices.Equal(expect1, actual1) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect1, actual1))
    }

    // replace the original map with a new one, and the iterator should traverse the new slice instead of the old one.
    input = map[string]int{"alice": 25, "bob": 20, "eve": 21}
    actual2 := make([]int, 0, 3)
    for v := range Order(iterator) {
        actual2 = append(actual2, v)
    }
    expect2 := []int{20, 21, 25}
    if !slices.Equal(expect2, actual2) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect2, actual2))
    }

    for _ = range iterator {
        break
    }
}

func TestMapSourceKeys(t *testing.T) {
    input := map[string]int{"foo": 1, "bar": 2}
    iterator := MapSourceKeys(func() map[string]int { return input })

    actual1 := make([]string, 0, 2)
    for v := range Order(iterator) {
        actual1 = append(actual1, v)
    }
    expect1 := []string{"bar", "foo"}
    if !slices.Equal(expect1, actual1) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect1, actual1))
    }

    // replace the original map with a new one, and the iterator should traverse the new slice instead of the old one.
    input = map[string]int{"eve": 21, "alice": 25, "bob": 20}
    actual2 := make([]string, 0, 3)
    for v := range Order(iterator) {
        actual2 = append(actual2, v)
    }
    expect2 := []string{"alice", "bob", "eve"}
    if !slices.Equal(expect2, actual2) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect2, actual2))
    }

    for _ = range iterator {
        break
    }
}

func TestChannel(t *testing.T) {
    ch := make(chan int, 3)
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)

    actual := make([]int, 0, 3)
    for v := range Chan(ch) {
        if v == 3 {
            break
        }
        actual = append(actual, v)
    }
    expect := []int{1, 2}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
}

func TestSeqSource(t *testing.T) {
    itFunc := func(yield func(int) bool) {
        yield(1)
        yield(2)
        yield(3)
    }
    iterator := SeqSource(func() iter.Seq[int] {
        return itFunc
    })

    actual := make([]int, 0, 3)
    for v := range iterator {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 3}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    itFunc = func(yield func(int) bool) {
        for _, each := range []int{4, 5, 6} {
            if !yield(each) {
                break
            }
        }
    }
    actual = make([]int, 0, 3)
    for v := range iterator {
        actual = append(actual, v)
        if v == 5 {
            break
        }
    }
    expect = []int{4, 5}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
}

func TestSeq2Source(t *testing.T) {
    type person struct {
        Name string
        Age  int
    }

    itFunc := func(yield func(string, int) bool) {
        yield("alice", 20)
        yield("bob", 21)
        yield("eve", 22)
    }
    iterator := Seq2Source(func() iter.Seq2[string, int] {
        return itFunc
    })

    actual := make([]person, 0, 3)
    for name, age := range iterator {
        actual = append(actual, person{name, age})
    }
    expect := []person{
        {"alice", 20},
        {"bob", 21},
        {"eve", 22},
    }
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    itFunc = func(yield func(string, int) bool) {
        people := []person{
            {"alice", 30},
            {"bob", 31},
            {"eve", 32},
        }
        for _, each := range people {
            if !yield(each.Name, each.Age) {
                break
            }
        }
    }
    actual = make([]person, 0, 3)
    for name, age := range iterator {
        actual = append(actual, person{name, age})
        if age == 31 {
            break
        }
    }
    expect = []person{
        {"alice", 30},
        {"bob", 31},
    }
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
}

func TestEmpty(t *testing.T) {
    i := 0
    for _ = range Empty[int]() {
        i++
    }

    if i != 0 {
        t.Fatal(fmt.Sprintf("expect: 0, actual: %d", i))
    }
}

func TestEmpty2(t *testing.T) {
    i := 0
    for _, _ = range Empty2[string, int]() {
        i++
    }

    if i != 0 {
        t.Fatal(fmt.Sprintf("expect: 0, actual: %d", i))
    }
}
