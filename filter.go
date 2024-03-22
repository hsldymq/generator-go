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
func Filter2[T1 any, T2 any](seq iter.Seq2[T1, T2], predicate func(T1, T2) bool) iter.Seq2[T1, T2] {
	return func(yield func(T1, T2) bool) {
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
func Distinct[T comparable](seq iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
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

// DistinctV1 returns an iterator that deduplicate the 2-tuples provided by the input iterator according to the first element.
// For example:
//
//	if the input iterator yields ("john", 20) ("anne", 21) ("john", 22)
//	DistinctV1 function will yield ("john", 20) ("anne", 21) because ("john", 22) has the same key as ("john", 20).
func DistinctV1[T1 comparable, T2 any](seq iter.Seq2[T1, T2]) iter.Seq2[T1, T2] {
	return func(yield func(T1, T2) bool) {
		yielded := newDistinctor[T1]()

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

// DistinctV2 is similar to DistinctV1, but it deduplicates by the second element of the 2-tuple.
func DistinctV2[T1 any, T2 comparable](seq iter.Seq2[T1, T2]) iter.Seq2[T1, T2] {
	return func(yield func(T1, T2) bool) {
		yielded := newDistinctor[T2]()

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
func DistinctBy[T any, K comparable](seq iter.Seq[T], keySelector func(T) K) iter.Seq[T] {
	return func(yield func(T) bool) {
		yielded := newDistinctor[K]()

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
func DistinctBy2[T1 any, T2 any, K comparable](seq iter.Seq2[T1, T2], keySelector func(T1, T2) K) iter.Seq2[T1, T2] {
	return func(yield func(T1, T2) bool) {
		yielded := newDistinctor[K]()

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
