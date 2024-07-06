[![Go Reference](https://pkg.go.dev/badge/github.com/hsldymq/goiter.svg)](https://pkg.go.dev/github.com/hsldymq/goiter)
[![Go Report Card](https://goreportcard.com/badge/github.com/hsldymq/goiter)](https://goreportcard.com/report/github.com/hsldymq/goiter)
[![codecov](https://codecov.io/gh/hsldymq/goiter/graph/badge.svg?token=1JE9U83U8K)](https://codecov.io/gh/hsldymq/goiter)
![GitHub License](https://img.shields.io/github/license/hsldymq/goiter?color=blue)
[![Test](https://github.com/hsldymq/goiter/actions/workflows/test.yml/badge.svg)](https://github.com/hsldymq/goiter/actions/workflows/test.yml)

[中文](./README_zh-CN.md)

This package provides a series of iterator creation functions for the range-over-func feature, simplifying the iterator creation logic for users in various common scenarios, such as transformation, filtering, aggregation, sorting, and more.

# Why do we need this package?
Although the iterator feature provides us with the ability to iterate over any data structure, the iterator operation functions provided in the standard library are very limited, and the logic for creating iterators is not very intuitive.

Typically, creating an iterator looks like this: you provide a creation function, which returns another function, and the parameter of the returned function is also a function. This is one of the reasons why this feature is controversial in the community.

The purpose of this package is to serve as a foundational toolkit to simplify the logic for creating iterators, reduce scenarios where manual iteration creation is needed, and provide a set of common iterator operation functions for use in higher-level applications.

# Requirements
* go version >= 1.23.0

# Examples
### Example 1: Traversal of an encapsulated collection
Suppose you need to allow external code to traverse a slice within a struct, but do not want to expose the slice. In this case, you can use the goiter.Slice or goiter.SliceElem function.

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

// Students method returns an iterator that yields each student, rather than directly exposing the internal slice.
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

// This function demonstrates how to use the goiter.Range and goiter.RangeStep function.
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

    // RangeStep function allows you to specify a step value as the third parameter
    // This will print 0 2 4 6 8 10
    for v := range goiter.RangeStep(0, 10, 2) {
        fmt.Printf("%d ", v)
    }
    fmt.Println()

    // This will print 5 3 1 -1 -3 -5
    // When iterating in descending order, you still need to provide a positive step value, so you don't need to adjust the sign of the step based on the direction of the iteration.
    // This is another difference from Python's range function, which requires a negative step value when iterating in reverse.
    // So if you provide a step of 0 or a negative number, RangeStep will not yield any values, 
    for v := range goiter.RangeStep(5, -5, 2) {
        fmt.Printf("%d ", v)
    }
    fmt.Println()
}

// This function demonstrates how to use the goiter.Sequence function.
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

// Checkout demonstrates how to use the goiter Transform function to convert data during iteration.
func (c *Cart) Checkout() goiter.Iterator2[string, float64] {
    // This transformer function creates a new iterator. It converts each product from the source iterator into the corresponding product name and cost.
    return goiter.Transform12(goiter.SliceElems(c.products), func(p *Product) (string, float64) {
        return p.Name, float64(p.Quantity) * p.Price
    })
}

func PrintNamesAges(cart *Cart) {
    // Checkout hides the internal details of the cart. So external code can only access the transformed data during iteration.
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

// This example demonstrates the use of the goiter.Filter function to filter elements in an iterator.
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

// This example demonstrates the use of the goiter.Distinct function to remove duplicate elements from an iterator.
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

// This example demonstrates how to use the goiter.Order function to sort elements in an iterator.
// Note: The goiter.Order series functions use additional memory to store intermediate results during sorting. Therefore, if the source iterator produces a large amount of data, use these functions with caution.
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
* `PickV1`
* `PickV2`
* `Swap`
* `Transform`
* `Transform2`
* `Transform12`
* `Transform21`

### aggregation
* `Count`
* `Count2`
* `Reduce`
* `Scan`

### sequence
* `Range`
* `RangeStep`
* `RangeTime`
* `Counter`
* `Sequence`
* `Sequence2`
* `Reverse`
* `Reverse2`

### combining
* `Combine`
* `Zip`
* `ZipAs`
* `Concat`
* `Concat2`

### filtering
* `Filter`
* `Filter2`
* `OfType`
* `Take`
* `Take2`
* `TakeLast`
* `TakeLast2`
* `Skip`
* `Skip2`
* `SkipLast`
* `SkipLast2`
* `Distinct`
* `DistinctV1`
* `DistinctV2`
* `DistinctBy`
* `Distinct2By`

### ordering
* `Order`
* `OrderBy`
* `Order2V1`
* `Order2V2`
* `Order2By`
* `StableOrderBy`
* `StableOrder2By`

### unrepeatable iterator
* `Once`
* `Once2`
* `FinishOnce`
* `FinishOnce2`

### Creating iterators from sources
* `Items`
* `Slice`
* `SliceElems`
* `SliceSource`
* `SliceSourceElems`
* `Map`
* `MapVals`
* `MapKeys`
* `MapSource`
* `MapSourceVals`
* `MapSourceKeys`
* `Chan`
* `ChanSource`
* `SeqSource`
* `Seq2Source`
* `Empty`
* `Empty2`

# Method chaining
Many of the functions listed above return an iterator of type goiter.Iterator[T] (or goiter.Iterator2[T1, T2]).

They are iterators and also have methods that allow you to chain multiple operations together.

The following two methods of generating iterators are equivalent.
```go
// Non-method chaining
iterator := goiter.Items(1, 2, 3, 4, 5, 6)
iterator = goiter.Filter(iterator, func(v int) bool {
    return v % 2 == 0
})

// Method chaining
iterator := goiter.Items(1, 2, 3, 4, 5, 6).Filter(func(v int) bool {
    return v % 2 == 0
})
```
