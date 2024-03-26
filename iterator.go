//go:build goexperiment.rangefunc

package goiter

import (
	"iter"
)

type Iterator[T any] iter.Seq[T]

func (it Iterator[T]) Seq() iter.Seq[T] {
	return iter.Seq[T](it)
}

func (it Iterator[T]) OrderBy(cmp func(T, T) int) Iterator[T] {
	return OrderBy(it, cmp)
}

func (it Iterator[T]) StableOrderBy(cmp func(T, T) int) Iterator[T] {
	return StableOrderBy(it, cmp)
}

func (it Iterator[T]) Filter(predicate func(T) bool) Iterator[T] {
	return Filter(it, predicate)
}

func (it Iterator[T]) Concat(its ...Iterator[T]) Iterator[T] {
	return Concat(it, its...)
}

func (it Iterator[T]) Reverse() Iterator[T] {
	return Reverse(it)
}

func (it Iterator[T]) Count() int {
	return Count(it)
}

func (it Iterator[T]) ToSlice() []T {
	return ToSlice(it)
}
