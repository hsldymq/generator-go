//go:build goexperiment.rangefunc

package goiter

import (
	"iter"
)

func Slice[T any](s []T, backward ...bool) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		if len(backward) == 0 || !backward[0] {
			for idx, elem := range s {
				if !yield(idx, elem) {
					return
				}
			}
		} else {
			for i := len(s) - 1; i >= 0; i-- {
				if !yield(i, s[i]) {
					return
				}
			}
		}
	}
}

func SliceElem[T any](s []T, backward ...bool) iter.Seq[T] {
	return PickV[int, T](Slice(s, backward...))
}

func SliceIdx[T any](s []T, backward ...bool) iter.Seq[int] {
	return PickK[int, T](Slice(s, backward...))
}

func Map[K comparable, V any](m map[K]V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

func MapVal[K comparable, V any](m map[K]V) iter.Seq[V] {
	return PickV[K, V](Map(m))
}

func MapKey[K comparable, V any](m map[K]V) iter.Seq[K] {
	return PickK[K, V](Map(m))
}

func Channel[T any](c <-chan T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range c {
			if !yield(v) {
				return
			}
		}
	}
}

func Filter[T any](seq iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for {
			v, ok := next()
			if !ok {
				return
			}
			if !predicate(v) {
				continue
			}
			if !yield(v) {
				return
			}
		}
	}
}

func Filter2[K any, V any](seq iter.Seq2[K, V], predicate func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		next, stop := iter.Pull2(seq)
		defer stop()
		for {
			k, v, ok := next()
			if !ok {
				return
			}
			if !predicate(k, v) {
				continue
			}
			if !yield(k, v) {
				return
			}
		}
	}
}