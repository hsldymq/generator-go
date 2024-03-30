package examples

import (
    "fmt"
    "github.com/hsldymq/goiter"
    "sync"
    "sync/atomic"
)

func OnceDemo() {
    data := []int{1, 2, 3, 4, 5, 6}

    // this is an example for Once function.
    // iterator returned by Once can be only used once.
    iterator := goiter.Once(goiter.SliceElem(data))

    // break the loop or finish the iteration will make the iterator unusable.
    for v := range iterator {
        // only the first 3 elements will be printed: 1 2 3
        fmt.Printf("%d ", v)
        if v == 3 {
            break
        }
    }
    fmt.Println()

    // so this won't print anything.
    for v := range iterator {
        fmt.Printf("%d ", v)
    }
    fmt.Println()
}

func FinishOnceDemo() {
    data := []int{1, 2, 3, 4, 5, 6}

    // this is an example for FinishOnce function.
    iterator := goiter.FinishOnce(goiter.SliceElem(data))

    // break the loop midway will not make the iterator unusable until all elements have been yielded.
    for v := range iterator {
        // only the first 3 elements will be printed: 1 2 3
        fmt.Printf("%d ", v)
        if v == 3 {
            break
        }
    }
    fmt.Println()

    // so this will print the remaining elements: 4 5 6
    for v := range iterator {
        fmt.Printf("%d ", v)
    }
    fmt.Println()

    // and this won't print anything because all elements have been yielded, iterator is unusable now.
    for v := range iterator {
        fmt.Printf("%d ", v)
    }
    fmt.Println()

    // you can also iterate over it concurrently.
    intSeqIterator := goiter.Range(int32(1), int32(10000)) // create an iterator yields 1 to 10000
    // iterator created by FinishOnce function allows you to iterate over it concurrently, and it will guarantee that all elements are yielded exactly once.
    iteratorConcurr := goiter.FinishOnce(intSeqIterator)
    // here we try to sum all elements concurrently.
    sum := int32(0)
    g := &sync.WaitGroup{}
    g.Add(3)
    go func() {
        for v := range iteratorConcurr {
            atomic.AddInt32(&sum, v)
        }
        g.Done()
    }()
    go func() {
        for v := range iteratorConcurr {
            atomic.AddInt32(&sum, v)
        }
        g.Done()
    }()
    go func() {
        for v := range iteratorConcurr {
            atomic.AddInt32(&sum, v)
        }
        g.Done()
    }()
    g.Wait()
    // so this should print the sum of 1 to 10000, which is 50005000
    fmt.Println(sum)
}
