[![Go Reference](https://pkg.go.dev/badge/github.com/hsldymq/goiter.svg)](https://pkg.go.dev/github.com/hsldymq/goiter)
[![Go Report Card](https://goreportcard.com/badge/github.com/hsldymq/goiter)](https://goreportcard.com/report/github.com/hsldymq/goiter)
[![codecov](https://codecov.io/gh/hsldymq/goiter/graph/badge.svg?token=1JE9U83U8K)](https://codecov.io/gh/hsldymq/goiter)
![GitHub License](https://img.shields.io/github/license/hsldymq/goiter?color=blue)
[![Test](https://github.com/hsldymq/goiter/actions/workflows/test.yml/badge.svg)](https://github.com/hsldymq/goiter/actions/workflows/test.yml)

This package provides a set of functions for creating iterators using the range-over function feature.

With this feature, you can iterate over any data structure and perform operations on the data at each iteration, such as transformation, filtering, aggregation, and more.

However, the standard library lacks convenient functions for these use cases.

This package fills that gap.

# Requirements
* go version >= 1.23.0

# Examples
### Example 1: Traversal of an encapsulated collection
Suppose you want to provide traversal capability for a slice in a struct to outside, but do not want to expose the slice, you can use the `goiter.Slice` or `goiter.SliceElem` function.

`goiter.Slice` and `goiter.SliceElem` are similar to the `slices.All` and `slices.Values` functions in the standard library, but they also offer reverse traversal capabilities like slices.Backward. You can enable reverse traversal by passing true as the second optional parameter.

```go
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

// Students method uses the goiter to return an iterator that yields each student, rather than directly exposing the internal slice.
func (s *School) Students() goiter.Iterator[*Student] {
    return goiter.SliceElems(s.students)
}

func PrintNames(school *School) {
    // So you can iterate through each student like a regular slice.
    for student := range school.Students() {
        fmt.Println(student.Name)
    }
}
```

### Example 2: Sequence generation
```go
package example2

import (
    "fmt"
    "github.com/hsldymq/goiter"
)

// goiter.Range and goiter.RangeStep provide similar functionality to Python's range function.
// The difference is that goiter.Range and goiter.RangeStep generate a range of numbers as a closed interval, unlike Python's range function which generates a half-open interval
func RangeDemo() {
    // So this will print 0 1 2 3 4 5, it is equivalent to Python `range(0, 6)` or Golang `for v := range 6`
    for v := range goiter.Range(0, 5) {
        fmt.Printf("%d ", v)
    }
    fmt.Println()

    // This will print 3 2 1 0 -1 -2 -3
    for v := range goiter.Range(3, -3) {
        fmt.Printf("%d ", v)
    }
    fmt.Println()

    // This will print 0 2 4 6 8 10
    for v := range goiter.RangeStep(0, 10, 2) {
        fmt.Printf("%d ", v)
    }
    fmt.Println()

    // This will print 5 3 1 -1 -3 -5
    // When iterating in reverse, you still need to provide a positive step value, so you don't need to adjust the sign of the step based on the direction of the iteration.
    // This is another difference from Python's range function, which requires a negative step value when iterating in reverse.
    // So if you provide a step of 0 or a negative number, RangeStep will not yield any values, 
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

    // this will print first 10 Fibonacci numbers: 1 1 2 3 5 8 13 21 34 55
    for v := range goiter.Sequence(genFib(10)) {
        fmt.Printf("%d ", v)
    }
}
```

### Example 3: Transformation
You can chain an iterator to another iterator for chained processing, so you can implement functions such as data transformation

```go
package example3

import (
    "fmt"
    "github.com/hsldymq/goiter"
    "iter"
)

type Product struct {
    Name string
    Price float64
    Quantity int
}

type Cart struct {
    products []*Product
}

// The Checkout method demonstrates how to use the goiter Transform function to convert the data type in each iteration.
func (c *Cart) Checkout() goiter.Iterator2[string, float64] {
    // It returns an iterator that yields the name and the cost of each product in the cart.
    return goiter.Transform12(goiter.SliceElems(c.products), func(p *Product) (string, float64) {
        return p.Name, float64(p.Quantity) * p.Price
    })
}

func PrintNamesAges(cart *Cart) {
    // Checkout hides the internal details and provides only the necessary data during iteration.
    for name, cost := range cart.Checkout() {
        fmt.Printf("The %s costs %.2f\n", name, cost)
    }
}
```

### Example 4: Filtering
```go
package example4

import (
    "fmt"
    "github.com/hsldymq/goiter"
)

func FilterDemo() {
    input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    iterator := goiter.SliceElems(input).Filter(func(v int) bool {
        return v % 2 == 0
    }) 
    // This will print 2 4 6 8 10
    for each := range iterator {
        fmt.Printf("%d ", each)
    }
}

func DistinctDemo() {
    input := []int{1, 2, 3, 3, 2, 1}
    // This will print 1 2 3
    for each := range goiter.Distinct(goiter.SliceElems(input)) {
        fmt.Printf("%d ", each)
    }
}
```

### Example 5: Ordering
```go
package example5

import (
	"fmt"
	"github.com/hsldymq/goiter"
)

func Demo() {
    // This will print 1 2 3 4
    for each := range goiter.Order(goiter.Items(1, 4, 3, 2)) {
        fmt.Printf("%d ", each)
    }

    // pass true as the second argument to sort in descending order
    // So this will print 4 3 2 1
    for each := range goiter.Order(goiter.Items(1, 4, 3, 2), true) {
        fmt.Printf("%d ", each)
    }
}
```

# List of goiter functions
Below are the functions provided by goiter.

As you can see, some functions are provided in two versions, such as `Filter` and `Filter2`, `Take` and `Take2`. These functions are respectively used for iterators of the `iter.Seq` or `iter.Seq2` versions.

Additionally, some functions have the suffix V1 or V2 in their names. These functions are for iter.Seq2 iterators, with V1 applying operations to the first element of each 2-tuple in the iteration, and V2 applying operations to the second element.

The comments at the function definitions include simple examples to help you better understand how to use these functions.

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
* `goiter.Reduce`
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
* `goiter.OrderBy`
* `goiter.Order2V1`
* `goiter.Order2V2`
* `goiter.Order2By`
* `goiter.StableOrderBy`
* `goiter.StableOrder2By`

### unrepeatable iterator
* `goiter.Once`
* `goiter.Once2`
* `goiter.FinishOnce`
* `goiter.FinishOnce2`

### creating iterator
* `goiter.Items`
* `goiter.Slice`
* `goiter.SliceElems`
* `goiter.SliceSource`
* `goiter.SliceSourceElems`
* `goiter.Map`
* `goiter.MapVals`
* `goiter.MapKeys`
* `goiter.MapSource`
* `goiter.MapSourceVals`
* `goiter.MapSourceKeys`
* `goiter.Chan`
* `goiter.ChanSource`
* `goiter.SeqSource`
* `goiter.Seq2Source`
* `goiter.Empty`
* `goiter.Empty2`

# Iterator & Iterator2 Types
Many functions listed above return an iter.Seq-like function, which is a type of `goiter.Iterator[T]` (or `goiter.Iterator2[T1, T2]` for iter.Seq2).

These two types have their own methods, which you can use to chain multiple operations together. For example:
```go
// instead of:
iterator := goiter.Items(1, 2, 3, 4, 5, 6)
iterator = goiter.Filter(iterator, func(v int) bool {
    return v % 2 == 0
})

// you can do like this:
iterator := goiter.Items(1, 2, 3, 4, 5, 6).Filter(func(v int) bool {
    return v % 2 == 0
})
```
