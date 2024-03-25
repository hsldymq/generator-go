//go:build goexperiment.rangefunc

package goiter

import (
    "iter"
    "sync"
    "sync/atomic"
)

// Once returns an iterator that can only be iterated over once;
// it cannot be reused after the iteration is complete or after breaking out of the loop. On subsequent attempts, it will not yield any values.
func Once[T any](seq iter.Seq[T]) iter.Seq[T] {
    flag := int32(0)
    return func(yield func(T) bool) {
        if !atomic.CompareAndSwapInt32(&flag, 0, 1) {
            return
        }

        next, stop := iter.Pull(seq)
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

// Once2 is the iter.Seq2 version of Once.
func Once2[T1, T2 any](seq iter.Seq2[T1, T2]) iter.Seq2[T1, T2] {
    flag := int32(0)
    return func(yield func(T1, T2) bool) {
        if !atomic.CompareAndSwapInt32(&flag, 0, 1) {
            return
        }

        next, stop := iter.Pull2(seq)
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

// ContinuableOnce is similar to the Once function.
// The difference is that if you break out of the iteration midway, it will continue to yield the remaining elements upon subsequent iterations until all values have been yielded.
func ContinuableOnce[T any](seq iter.Seq[T]) iter.Seq[T] {
    return continuable(Once(seq))
}

// ContinuableOnce2 is the iter.Seq2 version of ContinuableOnce.
func ContinuableOnce2[T1, T2 any](seq iter.Seq2[T1, T2]) iter.Seq2[T1, T2] {
    return continuable2(Once2(seq))
}

func continuable[T any](seq iter.Seq[T]) iter.Seq[T] {
    stopped := false
    var nextFunc func() (T, bool)
    var stopFunc func()

    pull := func() (func() (T, bool), func()) {
        if nextFunc == nil {
            stopped = false
            nextFunc, stopFunc = iter.Pull(seq)
            stopFunc = sync.OnceFunc(stopFunc)
        }
        return nextFunc, stopFunc
    }
    return func(yield func(T) bool) {
        next, stop := pull()
        defer func() {
            if stopped {
                stop()
                nextFunc = nil
                stopFunc = nil
            }
        }()

        for {
            v, ok := next()
            if !ok {
                stopped = true
                return
            }
            if !yield(v) {
                return
            }
        }
    }
}

func continuable2[T1, T2 any](seq iter.Seq2[T1, T2]) iter.Seq2[T1, T2] {
    stopped := false
    var nextFunc func() (T1, T2, bool)
    var stopFunc func()

    pull := func() (func() (T1, T2, bool), func()) {
        if nextFunc == nil {
            stopped = false
            nextFunc, stopFunc = iter.Pull2(seq)
            stopFunc = sync.OnceFunc(stopFunc)
        }
        return nextFunc, stopFunc
    }
    return func(yield func(T1, T2) bool) {
        next, stop := pull()
        defer func() {
            if stopped {
                stop()
                nextFunc = nil
                stopFunc = nil
            }
        }()

        for {
            v1, v2, ok := next()
            if !ok {
                stopped = true
                return
            }
            if !yield(v1, v2) {
                return
            }
        }
    }
}
