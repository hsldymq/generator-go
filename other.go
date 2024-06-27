package goiter

import (
    "iter"
    "sync/atomic"
)

// Cache returns an iterator that caches the values of the input iterator.
func Cache[TIter SeqX[T], T any](it TIter) Iterator[T] {
    var cached []T
    var cacheFlag int32

    var dynIter iter.Seq[T]
    cachedIter := func(yield func(T) bool) {
        for _, v := range cached {
            if !yield(v) {
                return
            }
        }
    }

    originalIter := func(yield func(T) bool) {
        cTemp := make([]T, 0)
        next, stop := iter.Pull(iter.Seq[T](it))
        defer stop()
        for {
            v, ok := next()
            if !ok {
                break
            }
            if !yield(v) {
                return
            }
            cTemp = append(cTemp, v)
        }
        if atomic.CompareAndSwapInt32(&cacheFlag, 0, 1) {
            cached = cTemp
            dynIter = cachedIter
        }
    }
    dynIter = originalIter
    sIter := SeqSource(func() iter.Seq[T] {
        return dynIter
    })

    return func(yield func(T) bool) {
        for v := range sIter {
            if !yield(v) {
                return
            }
        }
    }
}

// Cache2 is iter.Seq2 version of Cache.
func Cache2[TIter Seq2X[T1, T2], T1 any, T2 any](it TIter) Iterator2[T1, T2] {
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
    sIter := Seq2Source(func() iter.Seq2[T1, T2] {
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
