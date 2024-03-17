//go:build goexperiment.rangefunc

package goiter

import "iter"

// Filter returns an iterator that only yields the values of the input iterator that satisfy the predicate.
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

// Filter2 returns an iterator that only yields the key-values of the input iterator that satisfy the predicate.
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

// Distinct returns an iterator that only yields the distinct values of the input iterator.
// For example:
//
//	if the input iterator yields 1 2 3 3 2 1, Distinct function will yield 1 2 3.
func Distinct[V comparable](seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		yielded := map[any]bool{}

		next, stop := iter.Pull(seq)
		defer stop()
		for {
			v, ok := next()
			if !ok {
				return
			}
			if yielded[v] {
				continue
			}
			yielded[v] = true
			if !yield(v) {
				return
			}
		}
	}
}

// DistinctK returns an iterator that deduplicate the key-value pairs yielded by the input iterator according to the key
// For example:
//
//	if the input iterator yields ("john", 20) ("anne", 21) ("john", 22), DistinctK function will yield ("john", 20) ("anne", 21) because ("john", 22) has the same key as ("john", 20).
func DistinctK[K comparable, V any](seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		yielded := newDistinctor[K]()

		next, stop := iter.Pull2(seq)
		defer stop()
		for {
			k, v, ok := next()
			if !ok {
				return
			}
			if !yielded.mark(k) {
				continue
			}
			if !yield(k, v) {
				return
			}
		}
	}
}

// DistinctV is similar to DistinctK, but it deduplicates by the value instead of the key.
func DistinctV[K any, V comparable](seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		yielded := newDistinctor[V]()

		next, stop := iter.Pull2(seq)
		defer stop()
		for {
			k, v, ok := next()
			if !ok {
				return
			}
			if !yielded.mark(v) {
				continue
			}
			if !yield(k, v) {
				return
			}
		}
	}
}

// DistinctBy accepts a custom function to determine the deduplicate-key.
func DistinctBy[V any, DK comparable](seq iter.Seq[V], keySelector func(V) DK) iter.Seq[V] {
	return func(yield func(V) bool) {
		yielded := newDistinctor[DK]()

		next, stop := iter.Pull(seq)
		defer stop()
		for {
			v, ok := next()
			if !ok {
				return
			}
			if !yielded.mark(keySelector(v)) {
				continue
			}
			if !yield(v) {
				return
			}
		}
	}
}

// DistinctBy2 is an iter.Seq2 version of DistinctBy.
func DistinctBy2[K any, V any, DK comparable](seq iter.Seq2[K, V], keySelector func(K, V) DK) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		yielded := newDistinctor[DK]()

		next, stop := iter.Pull2(seq)
		defer stop()
		for {
			k, v, ok := next()
			if !ok {
				return
			}
			if !yielded.mark(keySelector(k, v)) {
				continue
			}
			if !yield(k, v) {
				return
			}
		}
	}
}

func newDistinctor[T comparable]() *distinctor[T] {
	return &distinctor[T]{
		dm: map[T]bool{},
	}
}

type distinctor[T comparable] struct {
	dm map[T]bool
}

func (d *distinctor[T]) mark(key T) bool {
	if _, ok := d.dm[key]; !ok {
		d.dm[key] = true
		return true
	}
	return false
}
