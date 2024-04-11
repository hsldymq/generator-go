//go:build goexperiment.rangefunc

package goiter

import (
    "iter"
    "sync/atomic"
)

type Seq2X[T1, T2 any] interface {
    ~func(yield func(T1, T2) bool)
}

type Iterator2[T1, T2 any] iter.Seq2[T1, T2]

func (it Iterator2[T1, T2]) Seq() iter.Seq2[T1, T2] {
    return iter.Seq2[T1, T2](it)
}

func (it Iterator2[T1, T2]) OrderBy(cmp func(*Combined[T1, T2], *Combined[T1, T2]) int) Iterator2[T1, T2] {
    return Order2By(it, cmp)
}

func (it Iterator2[T1, T2]) StableOrderBy(cmp func(*Combined[T1, T2], *Combined[T1, T2]) int) Iterator2[T1, T2] {
    return StableOrder2By(it, cmp)
}

func (it Iterator2[T1, T2]) Filter(cmp func(T1, T2) bool) Iterator2[T1, T2] {
    return Filter2(it, cmp)
}

func (it Iterator2[T1, T2]) Take(n int) Iterator2[T1, T2] {
    return Take2(it, n)
}

func (it Iterator2[T1, T2]) TakeLast(n int) Iterator2[T1, T2] {
    return TakeLast2(it, n)
}

func (it Iterator2[T1, T2]) Skip(n int) Iterator2[T1, T2] {
    return Skip2(it, n)
}

func (it Iterator2[T1, T2]) SkipLast(n int) Iterator2[T1, T2] {
    return SkipLast2(it, n)
}

func (it Iterator2[T1, T2]) Combine() Iterator[*Combined[T1, T2]] {
    return Combine(it)
}

func (it Iterator2[T1, T2]) Concat(its ...Iterator2[T1, T2]) Iterator2[T1, T2] {
    return Concat2(it, its...)
}

func (it Iterator2[T1, T2]) Reverse() Iterator2[T1, T2] {
    return Reverse2(it)
}

func (it Iterator2[T1, T2]) Count() int {
    return Count2(it)
}

func (it Iterator2[T1, T2]) Through(f func(T1, T2) (T1, T2)) Iterator2[T1, T2] {
    return Transform2(it, f)
}

func (it Iterator2[T1, T2]) Cache() Iterator2[T1, T2] {
    var cached []*Combined[T1, T2]
    var cacheFlag int32

    var dynIter iter.Seq2[T1, T2]
    cachedIter := func(yield func(T1, T2) bool) {
        for _, v := range cached {
            if !yield(v.V1, v.V2) {
                return
            }
        }
    }

    originalIter := func(yield func(T1, T2) bool) {
        cTemp := make([]*Combined[T1, T2], 0)
        next, stop := iter.Pull2(iter.Seq2[T1, T2](it))
        defer stop()
        for {
            v1, v2, ok := next()
            if !ok {
                break
            }
            if !yield(v1, v2) {
                return
            }
            cTemp = append(cTemp, &Combined[T1, T2]{
                V1: v1,
                V2: v2,
            })
        }
        if atomic.CompareAndSwapInt32(&cacheFlag, 0, 1) {
            cached = cTemp
            dynIter = cachedIter
        }
    }
    dynIter = originalIter
    sIter := Iter2Source(func() iter.Seq2[T1, T2] {
        return dynIter
    })

    return func(yield func(T1, T2) bool) {
        for v1, v2 := range sIter {
            if !yield(v1, v2) {
                return
            }
        }
    }
}
