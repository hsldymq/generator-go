//go:build goexperiment.rangefunc

package goiter

import (
	"iter"
	"sync"
	"sync/atomic"
)

func Once[T any](seq iter.Seq[T]) iter.Seq[T] {
	count := int64(0)
	return func(yield func(T) bool) {
		if atomic.AddInt64(&count, 1) > 1 {
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

func Once2[T1, T2 any](seq iter.Seq2[T1, T2]) iter.Seq2[T1, T2] {
	count := int64(0)
	return func(yield func(T1, T2) bool) {
		if atomic.AddInt64(&count, 1) > 1 {
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

func ContinuableOnce[T any](seq iter.Seq[T]) iter.Seq[T] {
	return continuable(Once(seq))
}

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
