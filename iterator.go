//go:build goexperiment.rangefunc

package goiter

import (
	"iter"
	"slices"
)

type Iterator[T any] iter.Seq[T]

func (it Iterator[T]) Seq() iter.Seq[T] {
	return iter.Seq[T](it)
}

func (it Iterator[T]) OrderBy(cmp func(T, T) int) Iterator[T] {
	return doOrderBy(it, cmp, slices.SortFunc[[]T, T])
}

func (it Iterator[T]) StableOrderBy(cmp func(T, T) int) Iterator[T] {
	return doOrderBy(it, cmp, slices.SortStableFunc[[]T, T])
}

func (it Iterator[T]) Count() int {
	count := 0
	for _ = range it {
		count++
	}
	return count
}
