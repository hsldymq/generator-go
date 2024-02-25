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

// SliceIdx only yields the indices of a slice.
func SliceIdx[T any](s []T, backward ...bool) iter.Seq[int] {
	return PickK[int, T](Slice(s, backward...))
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
func Channel[T any](c <-chan T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range c {
			if !yield(v) {
				return
			}
		}
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

// Filter returns an iterator that only yields the values of the input iterator that satisfy a predicate.
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

// Filter2 returns an iterator that only yields the key-values of the input iterator that satisfy a predicate.
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
