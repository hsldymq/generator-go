[![Go Reference](https://pkg.go.dev/badge/github.com/hsldymq/goiter.svg)](https://pkg.go.dev/github.com/hsldymq/goiter)
[![Go Report Card](https://goreportcard.com/badge/github.com/hsldymq/goiter)](https://goreportcard.com/report/github.com/hsldymq/goiter)
[![codecov](https://codecov.io/gh/hsldymq/goiter/graph/badge.svg?token=1JE9U83U8K)](https://codecov.io/gh/hsldymq/goiter)
![GitHub License](https://img.shields.io/github/license/hsldymq/goiter?color=blue)
[![Test](https://github.com/hsldymq/goiter/actions/workflows/test.yml/badge.svg)](https://github.com/hsldymq/goiter/actions/workflows/test.yml)

Go 1.22 introduced an experimental feature called [Rangefunc](https://go.dev/wiki/RangefuncExperiment), this is similar to the Generator in other languages (JavaScript, PHP, etc...). With this feature you can iterate data over any data structure.

The feature itself is simplistic designed, it does not provide many convenient functions for common use cases, such as sequence generation, data transformation, filtering, etc... 

So this package do the job.

To use this package, you must enable rangefunc feature.

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
func (s *School) Students() goiter.Iterator[*Student] {
    return goiter.SliceElem(s.students)
}

func PrintNames(school *School) {
    // iterate each student like a regular slice
    for student := range school.Students() {
        fmt.Println(student.Name)
    }
}
```

### Example 2: Sequence generation
```go
//go:build goexperiment.rangefunc

package example2

import (
    "fmt"
    "github.com/hsldymq/goiter"
)

// goiter.Range and goiter.RangeStep provide similar functionality to the Python's range function
func RangeDemo() {
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
    // When iterating in reverse, you still need to provide a positive step value, so you don't need to adjust the sign of the step based on the direction of the iteration.
    // If you provide a step of 0 or a negative number, RangeStep will not yield any values, this is different from the range function in Python.
    for v := range goiter.RangeStep(5, -5, 2) {
        fmt.Printf("%d ", v)
    }
    fmt.Println()
}

// goiter.Sequence is general purpose sequence generator, you can use it to generate any sequence you want 
// here is an example of generating Fibonacci sequence
func SequenceDemo() {
    genFib := func(n int) goiter.GeneratorFunc[int] {
        a, b := 0, 1
        return func() (int, bool) {
            if n <= 0 {
                return 0, false
            }
            n--
            a, b = b, a+b
            return a, true
        }
    }

    // this will print first 5 Fibonacci numbers: 1 1 2 3 5
    for v := range goiter.Sequence(genFib(5)) {
        fmt.Printf("%d ", v)
    }
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

func (s *School) Students() goiter.Iterator[*Student] {
    return goiter.SliceElem(s.students)
}

func PrintNamesAges(school *School) {
    // this iterator yields the age and name of each student
    iterator := goiter.Transform12(school.Students(), func(student *Student) (string, int) {
        return student.Name, student.Age
    })
    
    // so each round of iteration will yield a student's name and age, instead of a student struct
    for name, age := range iterator {
        fmt.Printf("name: %s, age: %d\n", name, age)
    }
}
```

### Example 4: Filtering
```go
//go:build goexperiment.rangefunc

package example4

import (
    "fmt"
    "github.com/hsldymq/goiter"
)

func FilterDemo() {
    input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    iterator := goiter.SliceElem(input).Filter(func(v int) bool {
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

### transformation
* `goiter.PickV1`
* `goiter.PickV2`
* `goiter.Swap`
* `goiter.Transform`
* `goiter.Transform2`
* `goiter.Transform12`
* `goiter.Transform21`

### aggregation
* `goiter.Count`
* `goiter.Count2`
* `goiter.Fold`
* `goiter.Scan`

### sequence
* `goiter.Range`
* `goiter.RangeStep`
* `goiter.RangeTime`
* `goiter.Counter`
* `goiter.Sequence`
* `goiter.Sequence2`
* `goiter.Reverse`
* `goiter.Reverse2`

### combining
* `goiter.Combine`
* `goiter.Zip`
* `goiter.ZipAs`
* `goiter.Concat`
* `goiter.Concat2`

### filtering
* `goiter.Filter`
* `goiter.Filter2`
* `goiter.OfType`
* `goiter.Take`
* `goiter.Take2`
* `goiter.TakeLast`
* `goiter.TakeLast2`
* `goiter.Skip`
* `goiter.Skip2`
* `goiter.SkipLast`
* `goiter.SkipLast2`
* `goiter.Distinct`
* `goiter.DistinctV1`
* `goiter.DistinctV2`
* `goiter.DistinctBy`
* `goiter.Distinct2By`

### ordering
* `goiter.Order`
* `goiter.OrderV1`
* `goiter.OrderV2`
* `goiter.OrderBy`
* `goiter.Order2By`
* `goiter.StableOrderBy`
* `goiter.StableOrder2By`

### unrepeatable iterator
* `goiter.Once`
* `goiter.Once2`
* `goiter.ContinuableOnce`
* `goiter.ContinuableOnce2`

### creating iterator
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
* `goiter.IterSource`
* `goiter.Iter2Source`
* `goiter.Empty`
* `goiter.Empty2`

### converting iterator
* `goiter.ToSlice`
* `goiter.ToMap`
* `goiter.ToMapAs`
* `goiter.ToMap2As`
