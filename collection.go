//go:build goexperiment.rangefunc

package goiter

import (
	"iter"
)

// Slice returns an iterator that allows you to traverse a slice in a forward or reverse direction.
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

// SliceElem only yields the elements of a slice.
func SliceElem[T any](s []T, backward ...bool) iter.Seq[T] {
	return PickV[int, T](Slice(s, backward...))
}

// Map returns an iterator that allows you to traverse a map.
func Map[K comparable, V any](m map[K]V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

// MapVal only yields the values of a map.
func MapVal[K comparable, V any](m map[K]V) iter.Seq[V] {
	return PickV[K, V](Map(m))
}

// MapKey only yields the keys of a map.
func MapKey[K comparable, V any](m map[K]V) iter.Seq[K] {
	return PickK[K, V](Map(m))
}

// Channel yields the values from a channel, it will stop when the channel is closed.
func Channel[V any](c <-chan V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range c {
			if !yield(v) {
				return
			}
		}
	}
}

// Empty returns an empty iterator.
func Empty[V any]() iter.Seq[V] {
	return func(yield func(V) bool) {
		return
	}
}

// Empty2 is iter.Seq2 version of Empty
func Empty2[K any, V any]() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		return
	}
}

// Concat returns an iterator that allows you to traverse multiple iterators in sequence.
func Concat[T any](seqs ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, seq := range seqs {
			func() {
				next, stop := iter.Pull(seq)
				defer stop()
				for {
					v, ok := next()
					if !ok {
						return
					}
					if !yield(v) {
						return
					}
				}
			}()
		}
	}
}

// Concat2 returns an iterator that allows you to traverse multiple iterators in sequence.
func Concat2[K any, V any](seqs ...iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, seq := range seqs {
			func() {
				next, stop := iter.Pull2(seq)
				defer stop()
				for {
					k, v, ok := next()
					if !ok {
						return
					}
					if !yield(k, v) {
						return
					}
				}
			}()
		}
	}
}

// Count counts the number of elements yielded by the input iterator.
func Count[V any](seq iter.Seq[V]) int {
	count := 0
	for _ = range seq {
		count++
	}
	return count
}

// Count2 counts the number of elements yielded by the input iterator.
func Count2[K any, V any](seq iter.Seq2[K, V]) int {
	count := 0
	for _, _ = range seq {
		count++
	}
	return count
}
