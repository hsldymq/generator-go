[![Go Report Card](https://goreportcard.com/badge/github.com/hsldymq/goiter)](https://goreportcard.com/report/github.com/hsldymq/goiter)
[![Test](https://github.com/hsldymq/goiter/actions/workflows/test.yml/badge.svg)](https://github.com/hsldymq/goiter/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/hsldymq/goiter/graph/badge.svg?token=1JE9U83U8K)](https://codecov.io/gh/hsldymq/goiter)

Rangefunc is a feature introduced in golang 1.22, which is similar to the Generator in other languages (JavaScript, PHP, etc...), making the for statement have the opportunity to iterate over any data structure.

goiter provides some iterator generation functions for some common scenarios under this feature.

To use this package, you shoud enable rangefunc experiment feature.

# Requirements
* golang version >= 1.22.0
* enable rangefunc feature by building with GOEXPERIMENT=rangefunc

# Examples
### Example 1: Traversal of an encapsulated collection
Suppose you want to provide traversal capability for a slice in a struct to outside, but do not want to expose the slice, you can use the `goiter.Slice` or `goiter.SliceElem` function.

```go
//go:build goexperiment.rangefunc

package example1

import (
    "fmt"
    "github.com/hsldymq/goiter"
    "iter"
)

type Student struct {
    Name string
    Age  int
}

type School struct {
    students []*Student
}

// Students returns an iterator that yields each student, instead of exposing the slice of students directly
func (s *School) Students() iter.Seq[*Student] {
    return goiter.SliceElem(s.students)
}

func PrintNames(school *School) {
    // iterate each student like a regular slice
    for student := range school.Students() {
        fmt.Println(student.Name)
    }
}
```

### Example 2: Range function
`goiter.Range` and `goiter.RangeStep` provide similar functionality to the Python's range function

```go
//go:build goexperiment.rangefunc

package example2

import (
    "fmt"
    "github.com/hsldymq/goiter"
)

func Demo() {
    // This will print 0 1 2 3 4 5 6 7 8 9
    // It is equivalent to Python `range(0, 10)` or Golang `for v := range 10`
    for v := range goiter.Range(0, 10) {
        fmt.Printf("%d ", v)
    }
    fmt.Println()

    // This will print 5 4 3 2 1 0 -1 -2 -3 -4
    for v := range goiter.Range(5, -5) {
        fmt.Printf("%d ", v)
    }
    fmt.Println()

    // This will print 0 2 4 6 8
    for v := range goiter.RangeStep(0, 10, 2) {
        fmt.Printf("%d ", v)
    }
    fmt.Println()

    // This will print 5 3 1 -1 -3
    // When iterating in reverse, you still need to provide a positive step, so you don't need to adjust the sign of the step based on the direction of the iteration.
    // If you provide a step of 0 or a negative number, RangeStep will not iterate over any values.
    // This is different from the range function in Python.
    for v := range goiter.RangeStep(5, -5, 2) {
        fmt.Printf("%d ", v)
    }
    fmt.Println()
}
```

### Example 3: Transformation
You can chain an iterator to another iterator for chained processing, so you can implement functions such as data transformation

```go
//go:build goexperiment.rangefunc

package example3

import (
    "fmt"
    "github.com/hsldymq/goiter"
    "iter"
)

type Student struct {
    Name string
    Age  int
}

type School struct {
    students []*Student
}

func (s *School) Students() iter.Seq[*Student] {
    return goiter.SliceElem(s.students)
}

func PrintNamesAges(school *School) {
    // this iterator yields the age and name of each student
    iterator := goiter.T12(school.Students(), func(student *Student) (string, int) {
        return student.Name, student.Age
    })
    
    // so each round of iteration will yield a student's name and age, instead of a student struct
    for name, age := range iterator {
        fmt.Printf("name: %s, age: %d\n", name, age)
    }
}
```

### Example 4: Filter
```go
//go:build goexperiment.rangefunc

package example4

import (
	"fmt"
	"github.com/hsldymq/goiter"
)

func FilterDemo() {
    input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    iterator := goiter.Filter(goiter.SliceElem(input), func(v int) bool {
        return v % 2 == 0
    }) 
	// this will print 2 4 6 8 10
	for each := range iterator {
		fmt.Printf("%d ", each)
	}
}

func DistinctDemo() {
    input := []int{1, 2, 3, 3, 2, 1}
    // this will print 1 2 3
    for each := range goiter.Distinct(goiter.SliceElem(input)) {
        fmt.Printf("%d ", each)
    }
}
```

### Example 5: Ordering
```go
//go:build goexperiment.rangefunc

package example5

import (
	"fmt"
	"github.com/hsldymq/goiter"
)

func Demo() {
    input := []int{1, 4, 3, 2}
	// this will print 1 2 3 4
	for each := range goiter.Order(goiter.SliceElem(input)) {
		fmt.Printf("%d ", each)
	}

    // pass true as the second argument to sort in descending order
    // this will print 4 3 2 1
    for each := range goiter.Order(goiter.SliceElem(input), true) {
        fmt.Printf("%d ", each)
    }
}
```

# List of goiter functions

### collection

* `goiter.Slice`
* `goiter.SliceElem`
* `goiter.SliceSource`
* `goiter.SliceSourceElem`
* `goiter.Map`
* `goiter.MapVal`
* `goiter.MapKey`
* `goiter.MapSource`
* `goiter.MapSourceVal`
* `goiter.MapSourceKey`
* `goiter.Channel`
* `goiter.ChannelSource`
* `goiter.Empty`
* `goiter.Empty2`
* `goiter.Concat`
* `goiter.Concat2`
* `goiter.Count`
* `goiter.Count2`

## filter
* `goiter.Filter`
* `goiter.Filter2`
* `goiter.Distinct`
* `goiter.DistinctK`
* `goiter.DistinctV`
* `goiter.DistinctBy`
* `goiter.DistinctBy2`

### ordering
* `goiter.Order`
* `goiter.OrderK`
* `goiter.OrderV`
* `goiter.OrderBy`
* `goiter.OrderBy2`
* `goiter.StableOrderBy`
* `goiter.StableOrderBy2`

### range
* `goiter.Range`
* `goiter.RangeStep`
* `goiter.RangeTime`

### transformation
* `goiter.PickK`
* `goiter.PickV`
* `goiter.SwapKV`
* `goiter.CombineKV`
* `goiter.T1`
* `goiter.T2`
* `goiter.T12`
* `goiter.T21`
