[![Go Reference](https://pkg.go.dev/badge/github.com/hsldymq/goiter.svg)](https://pkg.go.dev/github.com/hsldymq/goiter)
[![Go Report Card](https://goreportcard.com/badge/github.com/hsldymq/goiter)](https://goreportcard.com/report/github.com/hsldymq/goiter)
[![codecov](https://codecov.io/gh/hsldymq/goiter/graph/badge.svg?token=1JE9U83U8K)](https://codecov.io/gh/hsldymq/goiter)
![GitHub License](https://img.shields.io/github/license/hsldymq/goiter?color=blue)
[![Test](https://github.com/hsldymq/goiter/actions/workflows/test.yml/badge.svg)](https://github.com/hsldymq/goiter/actions/workflows/test.yml)

[English](./README.md)

这个包针对range-over-func特性提供了一系列迭代器创建函数，简化使用者在各种常见场景下的迭代器创建逻辑, 例如转换、过滤、聚合、排序等。

# 为什么需要这个包?
尽管迭代器特性给我们提供了迭代任何数据结构的遍历, 但标准库中里提供的迭代器操作函数很有限的, 并且创建迭代器的逻辑并不是那么直观. 

通常,迭代器的创建看起来像是这样: 你提供一个创建函数, 这个函数返回另一个函数, 而返回的函数的参数也是一个函数. 这也是这个特性在社区中存在争议的原因之一.

这个包的目的就是作为一个基础工具包, 简化迭代器的创建逻辑, 减少手动创建迭代的场景, 并且提供一系列通用的迭代器操作函数, 给上层应用使用.

# 要求
* go 版本 >= 1.23.0

# 示例
### 示例 1: 遍历封装的集合
假设你需要让外部代码遍历结构体中的一个slice, 但不想暴露这个slice, 你可以使用 `goiter.Slice` 或 `goiter.SliceElem` 函数.

`goiter.Slice` 和 `goiter.SliceElem` 类似于标准库中的 `slices.All` 和 `slices.Values` 函数，但它们还提供了 `slices.Backward` 那样的反向遍历功能。你可以通过将 true 作为第二个可选参数来启用反向遍历。

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

// Students 方法返回一个迭代器，该迭代器遍历输出 School 中的每个Student对象，而无序直接暴露内部的slice。
func (s *School) Students() goiter.Iterator[*Student] {
    return goiter.SliceElems(s.students)
}

func PrintNames(school *School) {
    // 因此你可以像常规的slice遍历其中的元素.
    for student := range school.Students() {
        fmt.Println(student.Name)
    }
}
```

### 示例2: 序列生成
```go
package example2

import (
    "fmt"
    "github.com/hsldymq/goiter"
)

// 这个函数演示了goiter.Range 和 goiter.RangeStep函数的使用方法. 
// goiter.Range 和 goiter.RangeStep 提供了类似于 Python range 函数的功能.
// 不同的是，goiter.Range 和 goiter.RangeStep 生成的是闭区间的数字范围，而 Python 的 range 函数生成的是半开区间.
func RangeDemo() {
    // 所以这个循环会打印 0 1 2 3 4 5, 它等价于 python的 `range(0, 6)` 或 golang的 `for v := range 6`
    for v := range goiter.Range(0, 5) {
        fmt.Printf("%d ", v)
    }
    fmt.Println()

    // 这个循环会打印 3 2 1 0 -1 -2 -3
    for v := range goiter.Range(3, -3) {
        fmt.Printf("%d ", v)
    }
    fmt.Println()

    // RangeStep 函数允许你指定一个步长值作为第三个参数
    // 这个循环会打印 0 2 4 6 8 10
    for v := range goiter.RangeStep(0, 10, 2) {
        fmt.Printf("%d ", v)
    }
    fmt.Println()

    // 这个循环会打印 5 3 1 -1 -3 -5
    // 当以降序迭代时，你仍然需要提供一个正的步长值，因此你不需要根据迭代方向来调整步长的符号。
    // 这是与 Python 的 range 函数的另一个不同之处，后者在反向迭代时需要一个负的步长值。
    // 因此，如果你提供的步长为 0 或负数，RangeStep 将不会生成任何值。
    for v := range goiter.RangeStep(5, -5, 2) {
        fmt.Printf("%d ", v)
    }
    fmt.Println()
}

// 这个函数演示了goiter.Sequence函数的使用方法. 
// goiter.Sequence 是通用的序列生成器，你可以用它生成任何你想要的序列
// 下面是一个生成斐波那契数列的示例
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

    // 这个循环会打印斐波那契数列的前10个数字: 1 1 2 3 5 8 13 21 34 55
    for v := range goiter.Sequence(genFib(10)) {
        fmt.Printf("%d ", v)
    }
}
```

### 示例3: 数据转换
你还可以将一个迭代器与另一个迭代器链接到一起依次处理，于是你可以实现诸如数据转换之类的功能。

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

// Checkout 方法演示了如何使用 goiter Transform 函数在迭代中对数据进行转换
func (c *Cart) Checkout() goiter.Iterator2[string, float64] {
    // 这个转换器函数生成一个新的迭代器，它将源迭代器中输出的每个产品, 转换为对应的商品名称及金额
    return goiter.Transform12(goiter.SliceElems(c.products), func(p *Product) (string, float64) {
        return p.Name, float64(p.Quantity) * p.Price
    })
}

func PrintNamesAges(cart *Cart) {
    // 因此 Checkout 方法隐藏了Cart的内部细节，这样外部代码在遍历过程中只会拿到转换之后的数据
    for name, cost := range cart.Checkout() {
        fmt.Printf("The %s costs %.2f\n", name, cost)
    }
}
```

### 示例4: 过滤
```go
package example4

import (
    "fmt"
    "github.com/hsldymq/goiter"
)

// 这个示例演示了使用 goiter.Filter 函数过滤迭代器中的元素
func FilterDemo() {
    input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    iterator := goiter.SliceElems(input).Filter(func(v int) bool {
        return v % 2 == 0
    }) 
    // 这个循环会打印出 2 4 6 8 10
    for each := range iterator {
        fmt.Printf("%d ", each)
    }
}

// 这个示例演示了使用 goiter.Distinct 函数去重迭代器中的元素
func DistinctDemo() {
    input := []int{1, 2, 3, 3, 2, 1}
    // 这个循环会打印出 1 2 3
    for each := range goiter.Distinct(goiter.SliceElems(input)) {
        fmt.Printf("%d ", each)
    }
}
```

### 示例5: 排序
```go
package example5

import (
	"fmt"
	"github.com/hsldymq/goiter"
)

// 这个示例演示了如何使用 goiter.Order 函数对迭代器中的元素进行排序
// 注意: goiter.Order 系列函数在排序过程中会使用额外的内存来存储中间结果, 所以如果源迭代器会产生大量数据, 请谨慎使用.
func Demo() {
    // 这个循环打印 1 2 3 4
    for each := range goiter.Order(goiter.Items(1, 4, 3, 2)) {
        fmt.Printf("%d ", each)
    }

    // 传递 true 作为第二个参数以进行降序排序
    // 于是这个循环会打印 4 3 2 1
    for each := range goiter.Order(goiter.Items(1, 4, 3, 2), true) {
        fmt.Printf("%d ", each)
    }
}
```

# goiter 函数列表
以下是goiter所提供的函数.

如你所见, 有些函数提供了两个版本, 例如 `Filter` 和 `Filter2`, `Take` 和 `Take2`,  这些函数各自用于`iter.Seq`或`iter.Seq2`这两种版本的迭代器.

另外,有些函数的名称后缀有V1或V2, 它们针对iter.Seq2迭代器, V1是将操作作用于每一个迭代的2元组数据中的第一个元素, V2则是作用于第二个元素.

函数定义处的注释包含了简单的示例,以帮助你更好地理解如何使用这些函数.

### 数据转换
* `PickV1`
* `PickV2`
* `Swap`
* `Transform`
* `Transform2`
* `Transform12`
* `Transform21`

### 聚合
* `Count`
* `Count2`
* `Reduce`
* `Scan`

### 序列生成
* `Range`
* `RangeStep`
* `RangeTime`
* `Counter`
* `Sequence`
* `Sequence2`
* `Reverse`
* `Reverse2`

### 组合
* `Combine`
* `Zip`
* `ZipAs`
* `Concat`
* `Concat2`

### 过滤
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

### 排序
* `Order`
* `OrderBy`
* `Order2V1`
* `Order2V2`
* `Order2By`
* `StableOrderBy`
* `StableOrder2By`

### 不可重读迭代器
* `Once`
* `Once2`
* `FinishOnce`
* `FinishOnce2`

### 从数据源创建迭代器
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

# 链式调用
上面列出的很多函数返回的迭代器类型是 goiter.Iterator[T] (或者 goiter.Iterator2[T1, T2]). 

它们既是迭代器, 同时也拥有一些方法, 这样你可以以链式调用的方式将多个迭代器串起来.

下面两种调用方式生成的迭代器是等价的:
```go
// 非链式调用
iterator := goiter.Items(1, 2, 3, 4, 5, 6)
iterator = goiter.Filter(iterator, func(v int) bool {
    return v % 2 == 0
})

// 链式调用
iterator := goiter.Items(1, 2, 3, 4, 5, 6).Filter(func(v int) bool {
    return v % 2 == 0
})
```
