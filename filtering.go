//go:build goexperiment.rangefunc

package goiter

import "iter"

// Filter returns an iterator that only yields the values of the input iterator that satisfy the predicate.
func Filter[TIter SeqX[T], T any](
	iterator TIter,
	predicate func(T) bool,
) Iterator[T] {
	return func(yield func(T) bool) {
		next, stop := iter.Pull(iter.Seq[T](iterator))
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

// Filter2 returns an iterator that only yields the 2-tuples of the input iterator that satisfy the predicate.
func Filter2[TIter Seq2X[T1, T2], T1 any, T2 any](
	iterator TIter,
	predicate func(T1, T2) bool,
) Iterator2[T1, T2] {
	return func(yield func(T1, T2) bool) {
		next, stop := iter.Pull2(iter.Seq2[T1, T2](iterator))
		defer stop()
		for {
			v1, v2, ok := next()
			if !ok {
				return
			}
			if !predicate(v1, v2) {
				continue
			}
			if !yield(v1, v2) {
				return
			}
		}
	}
}

// Take returns an iterator that yields the first n values of the input iterator.
func Take[TIter SeqX[T], T any](
	iterator TIter,
	n int,
) Iterator[T] {
	if n <= 0 {
		return Empty[T]()
	}

	return func(yield func(T) bool) {
		next, stop := iter.Pull(iter.Seq[T](iterator))
		defer stop()
		count := 0
		for {
			v, ok := next()
			if !ok {
				return
			}
			if !yield(v) {
				return
			}
			count++
			if count >= n {
				return
			}
		}
	}
}

// Take2 is the Iterator2 version of Take function.
func Take2[TIter Seq2X[T1, T2], T1, T2 any](
	iterator TIter,
	n int,
) Iterator2[T1, T2] {
	if n <= 0 {
		return Empty2[T1, T2]()
	}

	return func(yield func(T1, T2) bool) {
		next, stop := iter.Pull2(iter.Seq2[T1, T2](iterator))
		defer stop()
		count := 0
		for {
			v1, v2, ok := next()
			if !ok {
				return
			}
			if !yield(v1, v2) {
				return
			}
			count++
			if count >= n {
				return
			}
		}
	}
}

// Skip returns an iterator that suppress the first n values of the input iterator and yields the rest.
func Skip[TIter SeqX[T], T any](
	iterator TIter,
	n int,
) Iterator[T] {
	if n <= 0 {
		return Iterator[T](iterator)
	}

	return func(yield func(T) bool) {
		next, stop := iter.Pull(iter.Seq[T](iterator))
		defer stop()
		count := 0
		for {
			v, ok := next()
			if !ok {
				return
			}
			count++
			if count <= n {
				continue
			}
			if !yield(v) {
				return
			}
		}
	}
}

// Skip2 is the Iterator2 version of Skip function.
func Skip2[TIter Seq2X[T1, T2], T1, T2 any](
	iterator TIter,
	n int,
) Iterator2[T1, T2] {
	if n <= 0 {
		return Iterator2[T1, T2](iterator)
	}

	return func(yield func(T1, T2) bool) {
		next, stop := iter.Pull2(iter.Seq2[T1, T2](iterator))
		defer stop()
		count := 0
		for {
			v1, v2, ok := next()
			if !ok {
				return
			}
			count++
			if count <= n {
				continue
			}
			if !yield(v1, v2) {
				return
			}
		}
	}
}

// Distinct returns an iterator that only yields the distinct values of the input iterator.
// For example:
//
//	if the input iterator yields 1 2 3 3 2 1, Distinct function will yield 1 2 3.
//
// be careful, if this function is used on iterators that has massive amount of data, it might consume a lot of memory.
func Distinct[TIter SeqX[T], T comparable](iterator TIter) Iterator[T] {
	return func(yield func(T) bool) {
		yielded := map[any]bool{}

		next, stop := iter.Pull(iter.Seq[T](iterator))
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
//
// be careful, if this function is used on iterators that has massive amount of data, it might consume a lot of memory.
func DistinctV1[TIter Seq2X[T1, T2], T1 comparable, T2 any](iterator TIter) Iterator2[T1, T2] {
	return func(yield func(T1, T2) bool) {
		yielded := newDistinctor[T1]()

		next, stop := iter.Pull2(iter.Seq2[T1, T2](iterator))
		defer stop()
		for {
			v1, v2, ok := next()
			if !ok {
				return
			}
			if !yielded.mark(v1) {
				continue
			}
			if !yield(v1, v2) {
				return
			}
		}
	}
}

// DistinctV2 is similar to DistinctV1, but it deduplicates by the second element of the 2-tuple.
// be careful, if this function is used on iterators that has massive amount of data, it might consume a lot of memory.
func DistinctV2[TIter Seq2X[T1, T2], T1 any, T2 comparable](iterator TIter) Iterator2[T1, T2] {
	return func(yield func(T1, T2) bool) {
		yielded := newDistinctor[T2]()

		next, stop := iter.Pull2(iter.Seq2[T1, T2](iterator))
		defer stop()
		for {
			v1, v2, ok := next()
			if !ok {
				return
			}
			if !yielded.mark(v2) {
				continue
			}
			if !yield(v1, v2) {
				return
			}
		}
	}
}

// DistinctBy accepts a custom function to determine the deduplicate-key.
// be careful, if this function is used on iterators that has massive amount of data, it might consume a lot of memory.
func DistinctBy[TIter SeqX[T], T any, K comparable](
	iterator TIter,
	keySelector func(T) K,
) Iterator[T] {
	return func(yield func(T) bool) {
		yielded := newDistinctor[K]()

		next, stop := iter.Pull(iter.Seq[T](iterator))
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

// Distinct2By is an Iterator2 version of DistinctBy.
// be careful, if this function is used on iterators that has massive amount of data, it might consume a lot of memory.
func Distinct2By[TIter Seq2X[T1, T2], T1 any, T2 any, K comparable](
	iterator TIter,
	keySelector func(T1, T2) K,
) Iterator2[T1, T2] {
	return func(yield func(T1, T2) bool) {
		yielded := newDistinctor[K]()

		next, stop := iter.Pull2(iter.Seq2[T1, T2](iterator))
		defer stop()
		for {
			v1, v2, ok := next()
			if !ok {
				return
			}
			if !yielded.mark(keySelector(v1, v2)) {
				continue
			}
			if !yield(v1, v2) {
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
