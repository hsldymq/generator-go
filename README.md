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

func (s *School) Students() iter.Seq2[int, *Student] {
	return goiter.Slice(s.students)
}

// in another package

func Handle(school *School) {
	for _, student := range school.Students() {
		fmt.Println(student.Name)
	}
}
```

### Example 2: Range function
`goiter.Range` and `goiter.RangeStep` provide similar functionality to the Python range function

```go
//go:build goexperiment.rangefunc

package example2

import (
	"fmt"
	"github.com/hsldymq/goiter"
)

func PrintInts() {
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

// in another package

func IterStudentInfo(seq iter.Seq[*Student]) iter.Seq2[string, int] {
	// This will return a new iterator that yields the age and name of each student
	return goiter.T12(seq, func(student *Student) (string, int) {
		return student.Name, student.Age 
	})
}

func Handle(school *School) {
	for name, age := range IterStudentInfo(school.Students()) {
		fmt.Printf("name: %s, age: %d\n", name, age)
	}
}
```

# List of goiter functions

### collection

* `goiter.Slice`
* `goiter.SliceElem`
* `goiter.SliceIdx`
* `goiter.Map`
* `goiter.MapVal`
* `goiter.MapKey`
* `goiter.Channel`
* `goiter.Concat`
* `goiter.Concat2`
* `goiter.Filter`
* `goiter.Filter2`
* `goiter.Count`
* `goiter.Count2`

### range
* `goiter.Range`
* `goiter.RangeStep`
* `goiter.RangeTime`

### transformation
* `goiter.PickK`
* `goiter.PickV`
* `goiter.SwapKV`
* `goiter.T1`
* `goiter.T2`
* `goiter.T12`
* `goiter.T21`
