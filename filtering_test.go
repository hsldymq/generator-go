//go:build goexperiment.rangefunc

package goiter

import (
    "fmt"
    "maps"
    "slices"
    "testing"
)

func TestFilter(t *testing.T) {
    predicate := func(v int) bool {
        return v%2 == 0
    }
    actual := []int{}
    for v := range Range(0, 10).Filter(predicate) {
        actual = append(actual, v)
    }
    expect := []int{0, 2, 4, 6, 8}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    for _ = range Range(0, 10).Filter(predicate) {
        break
    }
}

func TestFilter2(t *testing.T) {
    predicate := func(name string, age int) bool {
        return name == "john"
    }
    input := map[string]int{"john": 20, "jane": 18}
    actual := map[string]int{}
    for k, v := range Map(input).Filter(predicate) {
        actual[k] = v
    }
    expect := map[string]int{"john": 20}
    if !maps.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    for _, _ = range Map(input).Filter(predicate) {
        break
    }
}

func TestOfType(t *testing.T) {
    cp1 := &Combined[int, int]{1, 1}
    cp2 := &Combined[int, int]{2, 2}
    input := []any{
        1,
        true,
        "hello",
        3.14,
        cp1,
        "world",
        nil,
        Combined[string, string]{"hello", "world"},
        cp2,
        nil,
        (*Combined[int, int])(nil),
        "nice",
    }

    // case 1
    actual := []Combined[string, string]{}
    for each := range OfType[Combined[string, string]](SliceElem(input)) {
        actual = append(actual, each)
    }
    expect := []Combined[string, string]{
        {"hello", "world"},
    }
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("test OfType failed, expect: %v, actual: %v", expect, actual))
    }

    // case 2
    actual2 := []*Combined[int, int]{}
    for each := range OfType[*Combined[int, int]](SliceElem(input)) {
        actual2 = append(actual2, each)
    }
    expect2 := []*Combined[int, int]{cp1, cp2, nil}
    if !slices.Equal(expect2, actual2) {
        t.Fatal(fmt.Sprintf("test OfType failed, expect: %v, actual: %v", expect2, actual2))
    }

    // case 3
    actual3 := []string{}
    for each := range OfType[string](SliceElem(input)) {
        actual3 = append(actual3, each)
        if each == "world" {
            break
        }
    }
    expect3 := []string{"hello", "world"}
    if !slices.Equal(expect3, actual3) {
        t.Fatal(fmt.Sprintf("test OfType failed, expect: %v, actual: %v", expect3, actual3))
    }
}

func TestTake(t *testing.T) {
    input := []int{1, 2, 3, 4, 5}

    actual := []int{}
    for v := range SliceElem(input).Take(2) {
        actual = append(actual, v)
    }
    expect := []int{1, 2}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    actual = []int{}
    for v := range SliceElem(input).Take(6) {
        actual = append(actual, v)
    }
    expect = []int{1, 2, 3, 4, 5}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    actual = []int{}
    for v := range SliceElem(input).Take(-1) {
        actual = append(actual, v)
    }
    expect = []int{}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    for _ = range SliceElem(input).Take(3) {
        break
    }
}

func TestTake2(t *testing.T) {
    type person struct {
        Name string
        Age  int
    }
    input := []person{
        {"alice", 20},
        {"bob", 21},
        {"eve", 22},
    }

    toNameAge := func(p person) (string, int) { return p.Name, p.Age }

    actual := []person{}
    for name, age := range Transform12(SliceElem(input), toNameAge).Take(2) {
        actual = append(actual, person{
            Name: name,
            Age:  age,
        })
    }
    expect := []person{
        {"alice", 20},
        {"bob", 21},
    }
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    actual = []person{}
    for name, age := range Transform12(SliceElem(input), toNameAge).Take(6) {
        actual = append(actual, person{
            Name: name,
            Age:  age,
        })
    }
    expect = []person{
        {"alice", 20},
        {"bob", 21},
        {"eve", 22},
    }
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    actual = []person{}
    for name, age := range Transform12(SliceElem(input), toNameAge).Take(0) {
        actual = append(actual, person{
            Name: name,
            Age:  age,
        })
    }
    expect = []person{}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    for _, _ = range Transform12(SliceElem(input), toNameAge).Take(2) {
        break
    }
}

func TestTakeLast(t *testing.T) {
    input := []int{1, 2, 3, 4, 5}

    // case 1
    actual := []int{}
    for v := range SliceElem(input).TakeLast(3) {
        actual = append(actual, v)
    }
    expect := []int{3, 4, 5}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("test TakeLast failed, expect: %v, actual: %v", expect, actual))
    }

    // case 2
    actual = []int{}
    for v := range SliceElem(input).TakeLast(7) {
        actual = append(actual, v)
    }
    expect = []int{1, 2, 3, 4, 5}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("test TakeLast failed, expect: %v, actual: %v", expect, actual))
    }

    // case 3
    actual = []int{}
    for v := range SliceElem(input).TakeLast(3) {
        actual = append(actual, v)
        if v == 4 {
            break
        }
    }
    expect = []int{3, 4}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("test TakeLast failed, expect: %v, actual: %v", expect, actual))
    }

    // case 4
    actual = []int{}
    for v := range SliceElem(input).TakeLast(0) {
        actual = append(actual, v)
    }
    expect = []int{}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("test TakeLast failed, expect: %v, actual: %v", expect, actual))
    }

    // case 5
    actual = []int{}
    for v := range SliceElem([]int{}).TakeLast(1) {
        actual = append(actual, v)
    }
    expect = []int{}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("test TakeLast failed, expect: %v, actual: %v", expect, actual))
    }
}

func TestTakeLast2(t *testing.T) {
    input := []int{1, 2, 3, 4, 5}

    // case 1
    actual := []Combined[int, int]{}
    for idx, v := range Slice(input).TakeLast(3) {
        actual = append(actual, Combined[int, int]{V1: idx, V2: v})
    }
    expect := []Combined[int, int]{
        {2, 3},
        {3, 4},
        {4, 5},
    }
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("test TakeLast2 failed, expect: %v, actual: %v", expect, actual))
    }

    // case 2
    actual = []Combined[int, int]{}
    for idx, v := range Slice(input).TakeLast(7) {
        actual = append(actual, Combined[int, int]{V1: idx, V2: v})
    }
    expect = []Combined[int, int]{
        {0, 1},
        {1, 2},
        {2, 3},
        {3, 4},
        {4, 5},
    }
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("test TakeLast2 failed, expect: %v, actual: %v", expect, actual))
    }

    // case 3
    actual = []Combined[int, int]{}
    for idx, v := range Slice(input).TakeLast(3) {
        actual = append(actual, Combined[int, int]{V1: idx, V2: v})
        if v == 4 {
            break
        }
    }
    expect = []Combined[int, int]{
        {2, 3},
        {3, 4},
    }
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("test TakeLast2 failed, expect: %v, actual: %v", expect, actual))
    }

    // case 4
    actual = []Combined[int, int]{}
    for idx, v := range Slice(input).TakeLast(0) {
        actual = append(actual, Combined[int, int]{V1: idx, V2: v})
    }
    expect = []Combined[int, int]{}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("test TakeLast failed, expect: %v, actual: %v", expect, actual))
    }

    // case 5
    actual = []Combined[int, int]{}
    for idx, v := range Slice([]int{}).TakeLast(1) {
        actual = append(actual, Combined[int, int]{V1: idx, V2: v})
    }
    expect = []Combined[int, int]{}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("test TakeLast failed, expect: %v, actual: %v", expect, actual))
    }
}

func TestSkip(t *testing.T) {
    input := []int{1, 2, 3, 4, 5}

    actual := []int{}
    for v := range SliceElem(input).Skip(2) {
        actual = append(actual, v)
    }
    expect := []int{3, 4, 5}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    actual = []int{}
    for v := range SliceElem(input).Skip(0) {
        actual = append(actual, v)
    }
    expect = []int{1, 2, 3, 4, 5}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    for _ = range SliceElem(input).Skip(1) {
        break
    }
}

func TestSkip2(t *testing.T) {
    type person struct {
        Name string
        Age  int
    }
    input := []person{
        {"alice", 20},
        {"bob", 21},
        {"eve", 22},
    }

    toNameAge := func(p person) (string, int) { return p.Name, p.Age }

    actual := []person{}
    for name, age := range Transform12(SliceElem(input), toNameAge).Skip(1) {
        actual = append(actual, person{
            Name: name,
            Age:  age,
        })
    }
    expect := []person{
        {"bob", 21},
        {"eve", 22},
    }
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    actual = []person{}
    for name, age := range Transform12(SliceElem(input), toNameAge).Skip(0) {
        actual = append(actual, person{
            Name: name,
            Age:  age,
        })
    }
    expect = []person{
        {"alice", 20},
        {"bob", 21},
        {"eve", 22},
    }
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    for _, _ = range Transform12(SliceElem(input), toNameAge).Skip(1) {
        break
    }
}

func TestSkipLast(t *testing.T) {
    input := []int{1, 2, 3, 4, 5}

    // case 1
    actual := []int{}
    for v := range SliceElem(input).SkipLast(3) {
        actual = append(actual, v)
    }
    expect := []int{1, 2}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("test SkipLast failed, expect: %v, actual: %v", expect, actual))
    }

    // case 2
    actual = []int{}
    for v := range SliceElem(input).SkipLast(7) {
        actual = append(actual, v)
    }
    expect = []int{}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("test SkipLast failed, expect: %v, actual: %v", expect, actual))
    }

    // case 3
    actual = []int{}
    for v := range SliceElem(input).SkipLast(2) {
        actual = append(actual, v)
        if v == 3 {
            break
        }
    }
    expect = []int{1, 2, 3}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("test SkipLast failed, expect: %v, actual: %v", expect, actual))
    }

    // case 4
    actual = []int{}
    for v := range SliceElem(input).SkipLast(0) {
        actual = append(actual, v)
    }
    expect = []int{1, 2, 3, 4, 5}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("test SkipLast failed, expect: %v, actual: %v", expect, actual))
    }
}

func TestSkipLast2(t *testing.T) {
    input := []int{1, 2, 3, 4, 5}

    // case 1
    actual := []Combined[int, int]{}
    for idx, v := range Slice(input).SkipLast(3) {
        actual = append(actual, Combined[int, int]{V1: idx, V2: v})
    }
    expect := []Combined[int, int]{
        {0, 1},
        {1, 2},
    }
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("test SkipLast2 failed, expect: %v, actual: %v", expect, actual))
    }

    // case 2
    actual = []Combined[int, int]{}
    for idx, v := range Slice(input).SkipLast(7) {
        actual = append(actual, Combined[int, int]{V1: idx, V2: v})
    }
    expect = []Combined[int, int]{}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("test SkipLast2 failed, expect: %v, actual: %v", expect, actual))
    }

    // case 3
    actual = []Combined[int, int]{}
    for idx, v := range Slice(input).SkipLast(2) {
        actual = append(actual, Combined[int, int]{V1: idx, V2: v})
        if v == 3 {
            break
        }
    }
    expect = []Combined[int, int]{
        {0, 1},
        {1, 2},
        {2, 3},
    }
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("test SkipLast2 failed, expect: %v, actual: %v", expect, actual))
    }

    // case 4
    actual = []Combined[int, int]{}
    for idx, v := range Slice(input).SkipLast(0) {
        actual = append(actual, Combined[int, int]{V1: idx, V2: v})
    }
    expect = []Combined[int, int]{
        {0, 1},
        {1, 2},
        {2, 3},
        {3, 4},
        {4, 5},
    }
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("test SkipLast failed, expect: %v, actual: %v", expect, actual))
    }
}

func TestDistinct(t *testing.T) {
    actual := []int{}
    for each := range Distinct(SliceElem([]int{1, 2, 3, 4, 4, 3, 2, 1})) {
        actual = append(actual, each)
    }
    expect := []int{1, 2, 3, 4}
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    //
    for _ = range Distinct(SliceElem([]int{1, 2, 3, 4, 4, 3, 2, 1})) {
        break
    }
}

func TestDistinctV1(t *testing.T) {
    type student struct {
        Name string
        Age  int
    }

    transFunc := func(s student) (string, int) {
        return s.Name, s.Age // name as key, age as value
    }
    input := []student{
        {"john", 20},
        {"jane", 18},
        {"john", 23}, // repeated name, so DistinctV1 will ignore this
    }
    actual := []student{}
    for name, age := range DistinctV1(Transform12(SliceElem(input), transFunc)) {
        actual = append(actual, student{name, age})
    }

    expect := []student{
        {"john", 20},
        {"jane", 18},
    }
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    //
    for _, _ = range DistinctV1(Transform12(SliceElem(input), transFunc)) {
        break
    }
}

func TestDistinctV2(t *testing.T) {
    type student struct {
        Name string
        Age  int
    }

    transFunc := func(s student) (string, int) {
        return s.Name, s.Age // name as key, age as value
    }
    input := []student{
        {"john", 20},
        {"jane", 18},
        {"alex", 20}, // alex has the same age as john, so DistinctV2 will ignore this
    }
    actual := []student{}
    for name, age := range DistinctV2(Transform12(SliceElem(input), transFunc)) {
        actual = append(actual, student{name, age})
    }

    expect := []student{
        {"john", 20},
        {"jane", 18},
    }
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    //
    for _, _ = range DistinctV2(Transform12(SliceElem(input), transFunc)) {
        break
    }
}

func TestDistinctBy(t *testing.T) {
    type student struct {
        Name string
        Age  int
    }

    input := []student{
        {"john", 20},
        {"jane", 18},
        {"john", 23},
    }
    transFunc := func(s student) string { return s.Name }
    actual := []student{}
    for each := range DistinctBy(SliceElem(input), transFunc) {
        actual = append(actual, each)
    }

    expect := []student{
        {"john", 20},
        {"jane", 18},
    }
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    //
    for _ = range DistinctBy(SliceElem(input), transFunc) {
        break
    }
}

func TestDistinctBy2(t *testing.T) {
    type student struct {
        Name string
        Age  int
    }

    input := []student{
        {"john", 20},
        {"jane", 18},
        {"alex", 20},
    }
    transFunc := func(s student) (int, student) { return s.Age, s }
    keySelector := func(age int, s student) int { return s.Age }
    actual := []student{}
    for _, each := range Distinct2By(Transform12(SliceElem(input), transFunc), keySelector) {
        actual = append(actual, each)
    }

    expect := []student{
        {"john", 20},
        {"jane", 18},
    }
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }

    //
    for _, _ = range Distinct2By(Transform12(SliceElem(input), transFunc), keySelector) {
        break
    }
}
