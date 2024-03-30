//go:build goexperiment.rangefunc

package goiter

import (
    "iter"
    "sync"
    "sync/atomic"
)

// Once returns an iterator that can only be iterated over once;
// it cannot be reused after the iteration is complete or after breaking out of the loop. On subsequent attempts, it will not yield any values.
// you can not also iterate over it concurrently, only one goroutine can iterate over it.
func Once[TIter SeqX[T], T any](iterator TIter) Iterator[T] {
    flag := int32(0)
    return func(yield func(T) bool) {
        if !atomic.CompareAndSwapInt32(&flag, 0, 1) {
            return
        }

        next, stop := iter.Pull(iter.Seq[T](iterator))
        defer stop()

        for {
            v, ok := next()
            if !ok {
                break
            }
            if !yield(v) {
                break
            }
        }
    }
}

// Once2 is the Iterator2 version of Once.
func Once2[TIter Seq2X[T1, T2], T1, T2 any](iterator TIter) Iterator2[T1, T2] {
    flag := int32(0)
    return func(yield func(T1, T2) bool) {
        if !atomic.CompareAndSwapInt32(&flag, 0, 1) {
            return
        }

        next, stop := iter.Pull2(iter.Seq2[T1, T2](iterator))
        defer stop()

        for {
            v1, v2, ok := next()
            if !ok {
                break
            }
            if !yield(v1, v2) {
                break
            }
        }
    }
}

// FinishOnce unlike Once, FinishOnce can be iterated over multiple times until all values have been yields exactly once.
// so you can break out of the iteration midway, and then iterate continuously from where you left off.
// you can also iterate over it concurrently, FinishOnce will make sure that all values are yielded exactly once.
// once all values have been yielded, it will not yield any more values.
func FinishOnce[TIter SeqX[T], T any](iterator TIter) Iterator[T] {
    fetchLock := &sync.Mutex{}
    next, stop := iter.Pull(iter.Seq[T](Once(iterator)))
    stopFunc := sync.OnceFunc(stop)
    nextFunc := func() (T, bool) {
        fetchLock.Lock()
        defer fetchLock.Unlock()
        v, ok := next()
        if !ok {
            stopFunc()
        }
        return v, ok
    }
    return func(yield func(T) bool) {
        for {
            v, ok := nextFunc()
            if !ok {
                return
            }
            if !yield(v) {
                return
            }
        }
    }
}

// FinishOnce2 is the Iterator2 version of FinishOnce.
func FinishOnce2[TIter Seq2X[T1, T2], T1, T2 any](iterator TIter) Iterator2[T1, T2] {
    fetchLock := &sync.Mutex{}
    next, stop := iter.Pull2(iter.Seq2[T1, T2](Once2(iterator)))
    stopFunc := sync.OnceFunc(stop)
    nextFunc := func() (T1, T2, bool) {
        fetchLock.Lock()
        defer fetchLock.Unlock()
        v1, v2, ok := next()
        if !ok {
            stopFunc()
        }
        return v1, v2, ok
    }
    return func(yield func(T1, T2) bool) {
        for {
            v1, v2, ok := nextFunc()
            if !ok {
                return
            }
            if !yield(v1, v2) {
                return
            }
        }
    }
}
