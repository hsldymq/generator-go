//go:build goexperiment.rangefunc

package goiter

import (
    "fmt"
    "maps"
    "slices"
    "testing"
)

func TestPick1(t *testing.T) {
    actual := make([]int, 0, 3)
    for idx := range PickV1(Slice([]int{7, 8, 9})) {
        actual = append(actual, idx)
    }

    expect := []int{0, 1, 2}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
}

func TestPick2(t *testing.T) {
    actual := make([]int, 0, 3)
    for v := range PickV2(Slice([]int{7, 8, 9})) {
        actual = append(actual, v)
    }

    expect := []int{7, 8, 9}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
}

func TestSwap(t *testing.T) {
    input := map[string]int{"1": 1, "2": 2}
    actual := make(map[int]string)
    for val, key := range Swap(Map(input)) {
        actual[val] = key
    }
    expect := map[int]string{1: "1", 2: "2"}
    if !maps.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
}

func TestCombine(t *testing.T) {
    actual := make([]Combined[int, string], 0, 3)
    for v := range Combine(Slice([]string{"1", "2", "3"})) {
        actual = append(actual, *v)
    }
    expect := []Combined[int, string]{
        {V1: 0, V2: "1"},
        {V1: 1, V2: "2"},
        {V1: 2, V2: "3"},
    }
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
}

func TestTransform(t *testing.T) {
    transformFunc := func(v int) string {
        return fmt.Sprintf("%d", v)
    }

    actual := make([]string, 0, 3)
    for v := range Transform(SliceElem([]int{1, 2, 3}), transformFunc) {
        actual = append(actual, v)
    }
    expect := []string{"1", "2", "3"}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    actual = make([]string, 0, 2)
    i := 0
    for v := range Transform(SliceElem([]int{1, 2, 3}), transformFunc) {
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

func TestTransform2(t *testing.T) {
    transformFunc := func(k int, v int) (string, string) {
        return fmt.Sprintf("%d", k+10), fmt.Sprintf("%d", v+100)
    }

    actualV1 := make([]string, 0, 3)
    actualV2 := make([]string, 0, 3)
    for v1, v2 := range Transform2(Slice([]int{1, 2, 3}), transformFunc) {
        actualV1 = append(actualV1, v1)
        actualV2 = append(actualV2, v2)
    }
    expectV1 := []string{"10", "11", "12"}
    if !slices.Equal(expectV1, actualV1) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expectV1, actualV1))
    }
    expectV2 := []string{"101", "102", "103"}
    if !slices.Equal(expectV2, actualV2) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expectV2, actualV2))
    }

    actualV2 = make([]string, 0, 3)
    i := 0
    for _, v := range Transform2(Slice([]int{1, 2, 3}), transformFunc) {
        actualV2 = append(actualV2, v)
        i++
        if i >= 2 {
            break
        }
    }
    expectV2 = []string{"101", "102"}
    if !slices.Equal(expectV2, actualV2) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expectV2, actualV2))
    }
}

func TestTransform12(t *testing.T) {
    transformFunc := func(v int) (int, string) {
        return v + 10, fmt.Sprintf("%d", v)
    }

    actualV1 := make([]int, 0, 3)
    actualV2 := make([]string, 0, 3)
    for v1, v2 := range Transform12(SliceElem([]int{1, 2, 3}), transformFunc) {
        actualV1 = append(actualV1, v1)
        actualV2 = append(actualV2, v2)
    }
    expectV1 := []int{11, 12, 13}
    if !slices.Equal(expectV1, actualV1) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expectV1, actualV1))
    }
    expectV2 := []string{"1", "2", "3"}
    if !slices.Equal(expectV2, actualV2) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expectV2, actualV2))
    }

    actualV1 = make([]int, 0, 3)
    i := 0
    for v1, _ := range Transform12(SliceElem([]int{1, 2, 3}), transformFunc) {
        actualV1 = append(actualV1, v1)
        i++
        if i >= 2 {
            break
        }
    }
    expectV1 = []int{11, 12}
    if !slices.Equal(expectV1, actualV1) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expectV1, actualV1))
    }
}

func TestTransform21(t *testing.T) {
    transformFunc := func(k int, v int) string {
        return fmt.Sprintf("%d_%d", k, v)
    }

    actual := make([]string, 0, 3)
    for v := range Transform21(Slice([]int{1, 2, 3}), transformFunc) {
        actual = append(actual, v)
    }
    expect := []string{"0_1", "1_2", "2_3"}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    actual = make([]string, 0, 3)
    i := 0
    for v := range Transform21(Slice([]int{1, 2, 3}), transformFunc) {
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

func TestZip(t *testing.T) {
    // case 1
    seq1 := SliceElem([]string{"Alice", "Bob", "Eve"})
    seq2 := SliceElem([]int{20, 21, 22, 23}) // seq2 has one more element than seq1
    actual := make([]Combined[string, int], 0, 3)
    for v := range Zip(seq1, seq2) {
        actual = append(actual, *v)
    }
    expect := []Combined[string, int]{
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
    actual = make([]Combined[string, int], 0, 2)
    i := 0
    for v := range Zip(seq1, seq2) {
        actual = append(actual, *v)
        i++
        if i >= 2 {
            break
        }
    }
    expect = []Combined[string, int]{
        {V1: "Alice", V2: 20},
        {V1: "Bob", V2: 21},
    }
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
}

func TestZipAs(t *testing.T) {
    type person struct {
        Name string
        Age  int
    }
    transformer := func(zipped *ZippedE[string, int]) person {
        p := person{
            Name: zipped.V1,
            Age:  zipped.V2,
        }
        if !zipped.OK1 {
            p.Name = "?"
        }
        if !zipped.OK2 {
            p.Age = -1
        }
        return p
    }

    // case 1
    nameSeq := SliceElem([]string{"Alice", "Bob", "Eve"})
    ageSeq := SliceElem([]int{20, 21})
    zipSeq := ZipAs(nameSeq, ageSeq, transformer, true)
    actual := make([]person, 0, 3)
    for each := range zipSeq {
        actual = append(actual, each)
    }
    expect := []person{
        {Name: "Alice", Age: 20},
        {Name: "Bob", Age: 21},
        {Name: "Eve", Age: -1},
    }
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    // case 2
    nameSeq = SliceElem([]string{"Alice", "Bob", "Eve"})
    ageSeq = SliceElem([]int{20, 21, 22, 23})
    zipSeq = ZipAs(nameSeq, ageSeq, transformer, true)
    actual = make([]person, 0, 4)
    for each := range zipSeq {
        actual = append(actual, each)
    }
    expect = []person{
        {Name: "Alice", Age: 20},
        {Name: "Bob", Age: 21},
        {Name: "Eve", Age: 22},
        {Name: "?", Age: 23},
    }
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
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
    seq := func(yield func(string) bool) {
        yield("alice")
        yield("bob")
        yield("eve")
    }

    actual := ToMapBy(seq, func(name string) (string, int) {
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

func TestToMapBy2(t *testing.T) {
    seq := func(yield func(string, string) bool) {
        yield("Alice", "Paris")
        yield("Bob", "Shanghai")
        yield("Eve", "Bangkok")
    }

    actual := ToMapBy2(seq, func(name string, city string) (string, string) {
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
